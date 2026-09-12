// ============================================================================
// SWG Pre-CU Core Data Types
// These map 1:1 to future C++ structs. Keep them POD-like and serializable.
// ============================================================================

// --- Resource Stat IDs -----------------------------------------------------
// Pre-CU uses 9 core attributes, all 0-1000.
export type StatId =
  | "OQ" // Overall Quality - weighted into EVERY experimental prop
  | "PE" // Potential Energy
  | "UT" // Unit Toughness
  | "CR" // Cold Resistance
  | "CD" // Conductivity
  | "DR" // Decay Resistance
  | "HR" // Heat Resistance
  | "MA" // Malleability
  | "SR"; // Shock Resistance

export const ALL_STATS: StatId[] = ["OQ", "PE", "UT", "CR", "CD", "DR", "HR", "MA", "SR"];

// Stats that a resource class actually generates; irrelevant stats are 0.
export type StatMap = Readonly<Record<StatId, number>>;

// --- Resource Class Hierarchy ----------------------------------------------
// Resources are organized in a tree. A schematic slot asking for "Ferrous
// Metal" accepts Iron, Steel, etc., but not Aluminum (Non-Ferrous).
export interface ResourceClass {
  readonly id: string;          // e.g. "mineral.ferrous.iron"
  readonly name: string;        // "Iron"
  readonly parent?: string;     // "mineral.ferrous"
  readonly relevantStats: ReadonlyArray<StatId>; // subset of ALL_STATS
}

// --- Concrete Resource Spawn -----------------------------------------------
// A unique vein/pool active on a planet at a given time. Despawns.
export interface Resource {
  readonly id: string;          // uuid for this spawn
  readonly name: string;        // e.g. "Vatrisian Iron" (random two-part name)
  readonly classId: string;     // FK to ResourceClass.id
  readonly planetId: PlanetId;
  readonly stats: StatMap;      // sampled at spawn, immutable for life of spawn
  readonly availableUnits: number; // how much is left in this vein
}

// --- Planets ---------------------------------------------------------------
export type PlanetId =
  | "tatooine"
  | "naboo"
  | "corellia"
  | "dantooine"
  | "endor"
  | "dathomir"
  | "yavin4"
  | "rori"
  | "talus"
  | "lok";

// --- Inventory Stacks ------------------------------------------------------
// Resources are stacked, identical to a spawn reference.
export interface ResourceStack {
  readonly resourceId: string;
  readonly units: number;
}

// --- Schematics ------------------------------------------------------------
// A schematic slot demands some units of a resource that MUST be a subclass
// (or exact match) of `requiredClassId`.
export interface SchematicSlot {
  readonly id: string;          // "main_metal", "fuel"
  readonly label: string;       // UI label
  readonly requiredClassId: string;
  readonly requiredUnits: number;
  readonly statWeights: Readonly<Partial<Record<StatId, number>>>;
  // How much this slot contributes to each experimental property.
  // Keys are experimental prop ids on the schematic. Values are 0..1 weights.
  readonly experimentalContribution: Readonly<Record<string, number>>;
}

// An experimental property on the finished item (e.g. "Min Damage").
export interface ExperimentalProperty {
  readonly id: string;
  readonly label: string;
  // Each stat's weight in the weighted-average pool. Weights need not sum to 1;
  // OQ is almost always the dominant co-stat.
  readonly statWeights: Readonly<Partial<Record<StatId, number>>>;
  readonly minValue: number;    // worst-case value (all crit fail)
  readonly maxValue: number;    // theoretical max with perfect resources + all amazing
  readonly unit?: string;       // "dmg", "%"
}

export interface Schematic {
  readonly id: string;
  readonly name: string;
  readonly description: string;
  readonly profession: string;  // Artisan, Weaponsmith, Architect...
  readonly slots: ReadonlyArray<SchematicSlot>;
  readonly experimentalProps: ReadonlyArray<ExperimentalProperty>;
  // Experimentation points budget allocated to the player during the phase.
  readonly experimentPoints: number;
  // Assembly complexity modifies the base success/failure thresholds.
  readonly complexity: number;
}

// --- Crafting Session State ------------------------------------------------
// These are mutable objects during an active craft. They map to transient
// server-side state in the real game.
export interface SlotAssignment {
  readonly slotId: string;
  readonly resourceId: string | null; // null if unfilled
  readonly unitsUsed: number;
}

export type ExperimentOutcome = "amazing" | "great" | "success" | "failure" | "critical";

export interface ExperimentLogEntry {
  readonly attempt: number;
  readonly pointsSpent: number;
  readonly propertyId: string;
  readonly outcome: ExperimentOutcome;
  readonly boxDelta: number;    // net change in progress boxes
  readonly dieRoll: number;     // 1..100 for inspection
}

// Each experimental property has 0..N "boxes" filled by experimentation.
// Final value is interpolated from min..max using (boxes / maxBoxes) combined
// with the resource cap.
export interface ExperimentProgress {
  readonly propertyId: string;
  boxes: number;               // mutable during session
  readonly maxBoxes: number;
}

// A finished, concrete item instance after crafting completes.
export interface CraftedItem {
  readonly id: string;
  readonly schematicId: string;
  readonly serialNumber: number;
  readonly resourcesUsed: ReadonlyArray<{ readonly resourceId: string; readonly units: number }>;
  readonly finalPropertyValues: Readonly<Record<string, number>>;
  readonly overallQuality: number; // 0..1 scalar, convenience stat
  readonly experimentLog: ReadonlyArray<ExperimentLogEntry>;
}
