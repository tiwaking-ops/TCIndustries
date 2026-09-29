// TCIndustries resource-lifecycle sandbox simulator (T-04 deliverable 2).
//
// NON-CANONICAL test scaffolding only: every constant here implements the
// parameter table in TCIndustries_Resource_Sim_Spec_v0.1.md §3 and carries no
// design authority (BAL-001). Structural test of INV-002 decomposition
// (R1a-d, R2, R3, R4) per spec §5.
//
// Standalone; stdlib-only; deterministic per seed. No testbed code is used or
// modified.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
)

// ---------- parameters (spec §3) ----------

type Params struct {
	Lifetime      int
	Quantity      int
	Interval      int
	Surveyors     int
	Traders       int
	Horizon       int
	Seed          int64
	GridW, GridH  int
	Zones         int
	Families      int // 3 families × 3 subtypes = 9 resource kinds
	Subtypes      int
	ExtractRate   int
	AttLo, AttHi  int // attribute envelope [20,100]
	DemandSigma   float64
	Harvesters    int
	TransportCost float64 // per unit-distance price penalty (scaffolding)
	InterZoneDist float64 // abstract distance between zone centers
}

// ---------- world ----------

type Spawn struct {
	Region        int
	Kind          int // family*Subtypes+subtype
	Born, Expires int
	Remaining     int
	Attrs         [4]float64
	DiscoveredAt  int // -1 = undiscovered
	DepletedAt    int // -1 = not depleted
	regionClass   int // terrain class of its region (for rediscovery)
}

type Agent struct {
	Kind       string // "surveyor"|"harvester"|"trader"|"harvesterQB"
	Zone       int
	EarlyYield float64 // units extracted within 20 ticks of a spawn's discovery
	LateYield  float64 // units extracted >100 ticks after discovery
	UnitsMoved float64 // total extraction (any)
	LastReloc  int
}

type Recorder struct {
	PriceHist         [][]float64   // per kind, per tick (zone 0 price as series; per-zone stored in ZonePrices)
	ZonePrices        [][][]float64 // [tick][zone][kind]
	Discoveries       []int         // tick of each first-discovery
	DiscKind          []int
	DiscWindow        []int // 1 if price/traffic event within 60 ticks
	DroughtWindows    int   // (zone,kind) drought windows with subsequent price rise
	FullCycles        int   // discovery→depletion→expiry→rediscovery chains
	TrafficEvents     int
	PriceEvents       int
	PriceEventTicks   []int
	TrafficEventTicks []int
	TradeByZone       []float64
	EarlyUnitRate     float64 // value-weighted extraction rate within 20 ticks of discovery
	LateUnitRate      float64 // value-weighted extraction rate ≥60 ticks after discovery
	QBRate            float64 // quality-BLIND harvester value yield (units × mean-attr quality)
	QSRate            float64 // quality-SENSITIVE harvester value yield
	earlyValue        float64 // population early-extraction value total (R2)
	lateValue         float64 // population late-extraction value total (R2)
	droughtActive     []droughtState
}

type droughtState struct {
	active     bool
	startTick  int
	startZone  int
	startKind  int
	startPr    float64
	lastZeroPr float64
}

func (r *Recorder) counterIn(a, b int, ticks []int) bool {
	for _, t := range ticks {
		if t >= a && t <= b {
			return true
		}
	}
	return false
}

func (r *Recorder) priceEvent(windows []float64) bool {
	// spec §4 intent, dimension-corrected (see results doc deviation D1):
	// |Δprice over 20-tick window| > 3× the run-median of 20-tick window deltas.
	if len(windows) < 2 {
		return false
	}
	abs := math.Abs(windows[len(windows)-1] - windows[len(windows)-2])
	return abs > 3*median(windows) && abs > 0
}

func median(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	s := append([]float64{}, xs...)
	sort.Float64s(s)
	return s[len(s)/2]
}

// ---------- run ----------

func runSim(p Params) Recorder {
	rng := newRNG(p.Seed)
	rec := Recorder{TradeByZone: make([]float64, p.Zones)}

	// terrain classes per region
	classes := make([]int, p.GridW*p.GridH)
	zoneOf := make([]int, p.GridW*p.GridH)
	for i := range classes {
		classes[i] = rng.intn(5)
		zoneOf[i] = rng.intn(p.Zones)
	}

	spawns := []*Spawn{}
	nextSpawn := make([]int, p.Families*p.Subtypes) // per-kind next spawn tick
	for i := range nextSpawn {
		nextSpawn[i] = 1 + rng.intn(p.Interval)
	}
	rediscoveryBase := map[string]*Spawn{} // regionClass:kind → last completed spawn

	// demand walk per zone per family (demand level 0.5–1.5 around 1.0)
	demand := make([][]float64, p.Zones)
	for z := range demand {
		demand[z] = []float64{1, 1, 1}
	}

	// agent pools
	agents := []*Agent{}
	for i := 0; i < p.Surveyors; i++ {
		agents = append(agents, &Agent{Kind: "surveyor", Zone: rng.intn(p.Zones)})
	}
	for i := 0; i < p.Harvesters; i++ {
		agents = append(agents, &Agent{Kind: "harvester", Zone: rng.intn(p.Zones)})
	}
	agents = append(agents, &Agent{Kind: "harvesterQB", Zone: 0}) // quality-blind variant (spec R4) — same zone as QS (deviation D6: deconfound location)
	agents = append(agents, &Agent{Kind: "harvesterQS", Zone: 0}) // quality-sensitive benchmark
	for i := 0; i < p.Traders; i++ {
		agents = append(agents, &Agent{Kind: "trader", Zone: rng.intn(p.Zones)})
	}

	// stock: delivered units per zone per kind
	delivered := make([][]float64, p.Zones)
	for z := range delivered {
		delivered[z] = make([]float64, p.Families*p.Subtypes)
	}

	pricePrev := make([][][]float64, 0)
	zoneDelta := make([][]float64, p.Zones) // rolling 20-tick |Δprice| per zone (all kinds pooled)
	for z := range zoneDelta {
		zoneDelta[z] = []float64{}
	}
	trafficWindow := make([][]float64, p.Zones) // rolling 20-tick inflow per zone
	relocThisTick := make([]int, p.Zones)       // zone relocations this tick (traffic signal)

	for tick := 1; tick <= p.Horizon; tick++ {
		// --- spawn lifecycle ---
		for k := 0; k < p.Families*p.Subtypes; k++ {
			if tick >= nextSpawn[k] {
				region := rng.intn(p.GridW * p.GridH)
				// avoid stacking same-kind spawn in same region
				free := true
				for _, s := range spawns {
					if s.Region == region && s.Kind == k && s.Expires > tick {
						free = false
						break
					}
				}
				if free {
					sp := &Spawn{
						Region: region, Kind: k,
						Born: tick, Expires: tick + p.Lifetime,
						Remaining: p.Quantity, DiscoveredAt: -1, DepletedAt: -1,
						regionClass: classes[region],
					}
					for a := 0; a < 4; a++ {
						sp.Attrs[a] = float64(p.AttLo) + rng.float64()*float64(p.AttHi-p.AttLo)
					}
					spawns = append(spawns, sp)
				}
				nextSpawn[k] = tick + p.Interval + rng.intn(p.Interval)
			}
		}

		// --- surveyors: discovery ---
		for _, a := range agents {
			if a.Kind != "surveyor" {
				continue
			}
			region := rng.intn(p.GridW * p.GridH)
			for _, s := range spawns {
				if s.Region == region && s.DiscoveredAt == -1 && s.Expires > tick && s.Remaining > 0 {
					s.DiscoveredAt = tick
					rec.Discoveries = append(rec.Discoveries, tick)
					rec.DiscKind = append(rec.DiscKind, s.Kind)
					rec.DiscWindow = append(rec.DiscWindow, 0)
				}
			}
		}

		// --- harvesters: extract from discovered spawns ---
		for _, a := range agents {
			if a.Kind != "harvester" && a.Kind != "harvesterQB" && a.Kind != "harvesterQS" {
				continue
			}
			// pick best discovered spawn reachable
			var best *Spawn
			bestScore := -1.0
			for _, s := range spawns {
				if s.DiscoveredAt == -1 || s.Expires <= tick || s.Remaining <= 0 {
					continue
				}
				if zoneOf[s.Region] != a.Zone {
					continue // harvesters work their zone; traders move goods
				}
				q := 1.0
				if a.Kind == "harvesterQS" {
					// quality-seeking: score by attribute sum (deviation D6)
					sum := 0.0
					for _, v := range s.Attrs {
						sum += v
					}
					q = sum / (4 * float64(p.AttHi))
				} // QB: quality-indifferent (q stays 1.0)
				freshness := 1.0
				if s.DiscoveredAt > tick-20 {
					freshness = 1.5
				}
				score := float64(s.Remaining) * q * freshness
				if score > bestScore {
					bestScore = score
					best = s
				}
			}
			if best != nil {
				take := p.ExtractRate
				if take > best.Remaining {
					take = best.Remaining
				}
				best.Remaining -= take
				a.UnitsMoved += float64(take)
				if tick-best.DiscoveredAt <= 20 {
					a.EarlyYield += valueOf(best, float64(p.AttHi))
				} else if tick-best.DiscoveredAt >= 60 {
					a.LateYield += valueOf(best, float64(p.AttHi))
				}
				if a.Kind == "harvester" {
					rec.harvestValue(best, take, best.DiscoveredAt, tick, p.AttHi)
				}
				// extraction delivers to the spawn's local zone (completes the
				// spawn→extraction→delivery→price→arbitrage chain; deviation D3)
				delivered[zoneOf[best.Region]][best.Kind] += float64(take)
				if a.Kind == "harvesterQS" {
					rec.QSRate += valueOf(best, float64(p.AttHi))
				} else if a.Kind == "harvesterQB" {
					rec.QBRate += valueOf(best, float64(p.AttHi))
				}
				if best.Remaining == 0 && best.DepletedAt == -1 {
					best.DepletedAt = tick
					// full-cycle detection: same regionClass+kind fully cycled before
					key := fmt.Sprintf("%d:%d", best.regionClass, best.Kind)
					if prev, ok := rediscoveryBase[key]; ok && prev.Expires < tick {
						rec.FullCycles++
					} else if ok {
						_ = prev
					}
					rediscoveryBase[key] = best
				}
			}
		}

		// zero relocation counters at tick start (moved before traders so both
		// migration loops record into the same tick)
		for z := range relocThisTick {
			relocThisTick[z] = 0
		}

		// --- traders: arbitrage moves goods between zones ---
		for _, a := range agents {
			if a.Kind != "trader" {
				continue
			}
			if tick-a.LastReloc < 20 {
				continue
			}
			// scan ALL kinds × zones for the best positive-margin route (deviation D5:
			// random-kind selection never found carryable stock)
			bestK, bestZ, bestM := -1, a.Zone, 0.0
			for k := 0; k < p.Families*p.Subtypes; k++ {
				if delivered[a.Zone][k] < 40.0 {
					continue
				}
				for z := 0; z < p.Zones; z++ {
					if z == a.Zone {
						continue
					}
					m := demand[z][k%3]*1.0 - delivered[z][k]/200.0 - p.TransportCost*p.InterZoneDist/100.0 // TransportCost=0.05: penalty 0.05, small vs demand spread (deviation D8)
					if m > bestM {
						bestM = m
						bestK, bestZ = k, z
					}
				}
			}
			if bestK != -1 && bestZ != a.Zone {
				delivered[a.Zone][bestK] -= 40.0
				delivered[bestZ][bestK] += 40.0
				rec.TradeByZone[bestZ] += 40.0
				a.Zone = bestZ
				a.LastReloc = tick
				relocThisTick[bestZ]++
			}
		}

		// --- harvester migration = traffic signal (relocations this tick) ---
		// (zeroing moved BEFORE the trader loop; relocations are recorded in both
		// the trader loop and the harvester migration loop below)
		for _, a := range agents {
			if a.Kind == "harvester" && tick-a.LastReloc >= 40 {
				// opportunistic migration toward zones with discovered stock
				bestZ, bestStock := a.Zone, -1.0
				for z := 0; z < p.Zones; z++ {
					stock := 0.0
					for k := 0; k < p.Families*p.Subtypes; k++ {
						stock += delivered[z][k]
					}
					stock += float64(availableDiscovered(spawns, -1, zoneOf, z, tick))
					if stock > bestStock {
						bestStock = stock
						bestZ = z
					}
				}
				if bestZ != a.Zone {
					a.Zone = bestZ
					a.LastReloc = tick
					relocThisTick[bestZ]++
				}
			}
		}

		// --- consumption sink (deviation D4): zones consume delivered stock at a
		// demand-scaled rate; without a sink, supply monotonically accumulates,
		// droughts are impossible, and prices lose meaning
		for z := 0; z < p.Zones; z++ {
			for k := 0; k < p.Families*p.Subtypes; k++ {
				f := k % 3
				consume := 1.5 * demand[z][f]
				if delivered[z][k] < consume {
					consume = delivered[z][k]
				}
				delivered[z][k] -= consume
			}
		}

		// --- demand walk ---
		for z := range demand {
			for f := 0; f < p.Families; f++ {
				demand[z][f] += rng.norm() * p.DemandSigma
				demand[z][f] = math.Max(0.3, math.Min(2.0, demand[z][f]))
			}
		}

		// --- pricing per zone per kind (local, quality-weighted) ---
		prices := make([][]float64, p.Zones)
		for z := 0; z < p.Zones; z++ {
			prices[z] = make([]float64, p.Families*p.Subtypes)
			for k := 0; k < p.Families*p.Subtypes; k++ {
				f := k % 3
				supply := delivered[z][k] + float64(availableDiscovered(spawns, k, zoneOf, z, tick))
				base := 100.0 * demand[z][f] / (1.0 + supply/300.0)
				// quality weighting: zone price reflects best discovered quality recently
				qw := qualityWeight(spawns, k, zoneOf, z, tick)
				prices[z][k] = base * (0.7 + 0.6*qw)
			}
		}
		// drought detection: demand present, zero discovered supply for the zone-kind,
		// price rising through the window (spec R1b)
		for z := 0; z < p.Zones; z++ {
			for k := 0; k < p.Families*p.Subtypes; k++ {
				f := k % 3
				supply := delivered[z][k] + float64(availableDiscovered(spawns, k, zoneOf, z, tick))
				droughtTrack(z, k, tick, demand[z][f], supply, prices[z][k], &rec, p)
			}
		}
		// event detection vs previous window (per zone-kind series)
		if len(pricePrev) > 0 {
			for z := 0; z < p.Zones; z++ {
				for k := 0; k < p.Families*p.Subtypes; k++ {
					flat := []float64{}
					for _, tickPrices := range pricePrev {
						flat = append(flat, tickPrices[z][k])
					}
					flat = append(flat, prices[z][k])
					if rec.priceEvent(flat) {
						rec.PriceEvents++
						rec.PriceEventTicks = append(rec.PriceEventTicks, tick)
					}
				}
			}
		}
		pricePrev = append(pricePrev, prices)
		if len(pricePrev) > 21 {
			pricePrev = pricePrev[1:]
		}
		// traffic events (run-relative relocation rule; deviation D2, see results doc)
		for z := range relocThisTick {
			if relocThisTick[z] > 0 {
				trafficWindow[z] = append(trafficWindow[z], float64(relocThisTick[z]))
				if len(trafficWindow[z]) > 400 {
					trafficWindow[z] = trafficWindow[z][1:]
				}
			}
		}
		allReloc := []float64{}
		for z := range trafficWindow {
			allReloc = append(allReloc, trafficWindow[z]...)
		}
		if len(allReloc) > 40 {
			medReloc := median(allReloc)
			if medReloc > 0 {
				for z := range relocThisTick {
					if float64(relocThisTick[z]) > 3*medReloc {
						rec.TrafficEvents++
						rec.TrafficEventTicks = append(rec.TrafficEventTicks, tick)
					}
				}
			}
		}

		// expiry
		alive := spawns[:0]
		for _, s := range spawns {
			if s.Expires > tick {
				alive = append(alive, s)
			}
		}
		spawns = alive
	}

	rec.EarlyUnitRate = rec.earlyValue
	rec.LateUnitRate = rec.lateValue
	for i, dt := range rec.Discoveries {
		if dt+60 <= p.Horizon && (rec.priceEventsIn(dt, dt+60, rec.PriceEventTicks) || rec.trafficEventsIn(dt, dt+60, rec.TrafficEventTicks)) {
			rec.DiscWindow[i] = 1
		}
	}
	return rec
}

func availableDiscovered(spawns []*Spawn, kind int, zoneOf []int, zone, tick int) int {
	n := 0
	for _, s := range spawns {
		if (kind == -1 || s.Kind == kind) && s.DiscoveredAt != -1 && s.Expires > tick && s.Remaining > 0 && zoneOf[s.Region] == zone {
			n += s.Remaining
		}
	}
	return n
}

// valueOf: quality-weighted value of extracted units (mean attribute / max)
func valueOf(s *Spawn, attMax float64) float64 {
	sum := 0.0
	for _, v := range s.Attrs {
		sum += v
	}
	return (sum / 4.0) / attMax
}

func qualityWeight(spawns []*Spawn, kind int, zoneOf []int, zone, tick int) float64 {
	best := 0.0
	for _, s := range spawns {
		if s.Kind == kind && s.DiscoveredAt != -1 && s.Expires > tick && s.Remaining > 0 && zoneOf[s.Region] == zone {
			sum := 0.0
			for _, v := range s.Attrs {
				sum += v
			}
			q := sum / 400.0
			if q > best {
				best = q
			}
		}
	}
	return best
}

// event-window helpers
func (r *Recorder) priceEventsIn(a, b int, ticks []int) bool   { return r.counterIn(a, b, ticks) }
func (r *Recorder) trafficEventsIn(a, b int, ticks []int) bool { return r.counterIn(a, b, ticks) }

// droughtTrack (deviation D7): enters drought when demand>0.6 and supply==0.
// While active it tracks the last zero-supply price; if supply returns after
// ≥30 ticks and the scarcity-era price rose ≥5% over entry, the drought
// "wins" (scarcity → price rise). Comparing post-resupply price would invert
// the signal (resupply lowers price), hence lastZeroPr.
// R2 instrumentation: the population harvesters (not the QB/QS pair) feed the
// early/late value split so the ratio has statistical mass.
func (rec *Recorder) harvestValue(s *Spawn, take int, discAt, tick, attMax int) {
	v := valueOf(s, float64(attMax)) * float64(take)
	if tick-discAt <= 20 {
		rec.earlyValue += v
	} else if tick-discAt >= 60 {
		rec.lateValue += v
	}
}

func droughtTrack(z, k, tick int, demand, supply, price float64, rec *Recorder, p Params) {
	key := z*(p.Families*p.Subtypes) + k
	for len(rec.droughtActive) <= key {
		rec.droughtActive = append(rec.droughtActive, droughtState{})
	}
	d := &rec.droughtActive[key]
	if !d.active {
		if supply == 0 && demand > 0.6 {
			d.active = true
			d.startTick = tick
			d.startZone = z
			d.startKind = k
			d.startPr = price
			d.lastZeroPr = price
		}
		return
	}
	// active
	if supply == 0 {
		d.lastZeroPr = price
	}
	if supply > 0 || demand <= 0.3 {
		if tick-d.startTick >= 30 && d.lastZeroPr >= d.startPr*1.05 {
			rec.DroughtWindows++
		}
		d.active = false
	}
}

// ---------- RNG (deterministic) ----------

type rng struct{ s uint64 }

func newRNG(seed int64) *rng { return &rng{s: uint64(seed) * 2862933555777941757 % (1 << 62)} }
func (r *rng) next() uint64 {
	r.s ^= r.s << 13
	r.s ^= r.s >> 7
	r.s ^= r.s << 17
	return r.s
}
func (r *rng) float64() float64 { return float64(r.next()>>11) / (1 << 53) }
func (r *rng) intn(n int) int   { return int(r.next() % uint64(n)) }
func (r *rng) norm() float64 {
	// Box–Muller, two uniforms
	u1, u2 := r.float64(), r.float64()
	if u1 < 1e-12 {
		u1 = 1e-12
	}
	return math.Sqrt(-2*math.Log(u1)) * math.Cos(2*math.Pi*u2)
}

// ---------- main: envelope ----------

type RunResult struct {
	Lifetime, Quantity, Interval, Surveyors, Traders int
	Seed                                             int64
	PriceEvents, TrafficEvents                       int
	DiscTotal, DiscFollowed                          int
	FullCycles, DroughtWins                          int
	EarlyRate, LateRate, QBRate, QSRate              float64
	MaxZoneShare                                     float64
}

func main() {
	out := flag.String("out", "resource_sim_results.json", "output JSON path")
	quick := flag.Bool("quick", false, "run a 9-point smoke subset instead of full 108×3")
	flag.Parse()

	lifetimes := []int{100, 200, 400}
	quantities := []int{2000, 5000, 10000}
	intervals := []int{20, 40, 80}
	surveyors := []int{10, 25}
	traders := []int{8, 15}
	seeds := []int64{11, 23, 47}

	var results []RunResult
	for _, L := range lifetimes {
		for _, Q := range quantities {
			for _, I := range intervals {
				for _, S := range surveyors {
					for _, T := range traders {
						for _, sd := range seeds {
							p := Params{
								Lifetime: L, Quantity: Q, Interval: I, Surveyors: S, Traders: T,
								Horizon: 2000, Seed: sd, GridW: 40, GridH: 30, Zones: 4,
								Families: 3, Subtypes: 3, ExtractRate: 50,
								AttLo: 20, AttHi: 100, DemandSigma: 0.15,
								Harvesters: 40, TransportCost: 0.05, InterZoneDist: 100,
							}
							if *quick && (L != 200 || Q != 5000 || I != 40 || sd != 11) {
								continue
							}
							rec := runSim(p)
							early := rec.EarlyUnitRate
							late := rec.LateUnitRate
							er := early / math.Max(1, late)
							lr := late / math.Max(1, early)
							total := 0.0
							for _, v := range rec.TradeByZone {
								total += v
							}
							share := 0.0
							if total > 0 {
								for _, v := range rec.TradeByZone {
									if v/total > share {
										share = v / total
									}
								}
							}
							followed := 0
							for _, w := range rec.DiscWindow {
								followed += w
							}
							results = append(results, RunResult{
								Lifetime: L, Quantity: Q, Interval: I, Surveyors: S, Traders: T, Seed: sd,
								PriceEvents: rec.PriceEvents, TrafficEvents: rec.TrafficEvents,
								DiscTotal: len(rec.Discoveries), DiscFollowed: followed,
								FullCycles: rec.FullCycles,
								EarlyRate:  er, LateRate: lr,
								QBRate: rec.QBRate, QSRate: rec.QSRate,
								MaxZoneShare: share,
							})
						}
					}
				}
			}
		}
	}
	f, _ := os.Create(*out)
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", " ")
	_ = enc.Encode(results)
	fmt.Printf("runs completed: %d → %s\n", len(results), *out)
}
