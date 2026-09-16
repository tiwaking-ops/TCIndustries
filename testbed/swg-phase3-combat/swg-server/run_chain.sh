#!/bin/sh
# Full-chain regression (Phase 10 B7): phases 1–9 suites re-run green on
# fresh DBs, then the Phase 10 suite. Each phase gets a fresh DB and its own
# server lifecycle; exit codes are accumulated and reported per phase.
# PHASES="3 9" filters to a subset (default: 1–9 then 10).
#
# Port note: phase 1–9 clients hardcode http://localhost:8080 (pre-existing;
# flagged for a follow-up env-override pass), so phases 1–9 REQUIRE 8080
# free — same precondition as run_phase9.sh. The Phase 10 suite runs on its
# own dedicated port (8090) via P10_SERVER_URL.
set -u
cd "$(dirname "$0")"
PHASES="${PHASES:-1 2 3 4 5 6 7 8 9 10}"
FAILS=""
RESULT=""

for P in $PHASES; do
  echo "=== PHASE $P: build + fresh DB + run (port 8080) ==="
  # Port hygiene: a previous phase's server must be gone before this one
  # binds (MSYS `kill` on a background job does not reap the exe; the next
  # boot then dies on bind and the client reports "Server not reachable").
  for OLD in $(netstat -ano | grep ":8080" | grep LISTENING | awk '{print $NF}' | sort -u); do
    taskkill //F //PID "$OLD" 2>/dev/null
  done
  sleep 2
  go build -o /tmp/srv ./cmd/server || { echo "PHASE $P: server build FAILED"; FAILS="$FAILS $P"; RESULT="$RESULT p$P=build"; continue; }
  go build -o /tmp/p "./cmd/phase${P}test" || { echo "PHASE $P: client build FAILED"; FAILS="$FAILS $P"; RESULT="$RESULT p$P=build"; continue; }
  rm -f /tmp/chain.db /tmp/chain.db-shm /tmp/chain.db-wal
  DB_PATH=/tmp/chain.db HTTP_ADDR=":8080" WS_ADDR=":8080" TESTBED_FAST_CYCLE=1 \
    /tmp/srv > /tmp/chain_srv.log 2>&1 &
  SRV=$!
  sleep 12
  /tmp/p
  CODE=$?
  kill $SRV 2>/dev/null
  # Reap hard: taskkill the srv.exe image this phase spawned so the next
  # phase's bind cannot collide (Windows/MSYS background-job reaping gap).
  for OLD in $(netstat -ano | grep ":8080" | grep LISTENING | awk '{print $NF}' | sort -u); do
    taskkill //F //PID "$OLD" 2>/dev/null
  done
  sleep 2
  RESULT="$RESULT p$P=$CODE"
  if [ "$CODE" != "0" ]; then FAILS="$FAILS $P"; fi
done

# Phase 10: dedicated port, env-overridable client (B2/B6 design).
if echo " $PHASES " | grep -q " 10 "; then
  echo "=== PHASE 10: fresh DB + run (port 8090) ==="
  # Phase 10 clients default to 8080 constants unless redirected; the
  # runner exports its dedicated port for them.
  P10_SERVER_URL="${P10_SERVER_URL:-http://localhost:8090}"
  export P10_SERVER_URL
  sh ./run_phase10.sh
  CODE=$?
  RESULT="$RESULT p10=$CODE"
  if [ "$CODE" != "0" ]; then FAILS="$FAILS 10"; fi
fi

echo
echo "=== CHAIN RESULTS:$RESULT ==="
if [ -n "$FAILS" ]; then
  echo "CHAIN:FAILED (phases:$FAILS)"
  exit 1
fi
echo "CHAIN:ALL-GREEN"
