#!/bin/sh
# Full-chain regression (Phase 10 B7): phases 1-9 suites re-run green on
# fresh DBs, then the Phase 10 suite. Each phase gets a fresh DB and its own
# server lifecycle; exit codes are accumulated and reported per phase.
# PHASES="3 9" filters to a subset (default: 1-9 then 10).
#
# Port note: phase 1-9 clients now honor the SERVER_URL env override (port-
# override pass, Option A), so the chain runs on a dedicated port with NO
# dependency on 8080 being free. The Phase 10 suite runs on its own dedicated
# port (8090) via P10_SERVER_URL.
set -u
cd "$(dirname "$0")"
CHAIN_PORT="${CHAIN_PORT:-8091}"
CHAIN_URL="http://localhost:${CHAIN_PORT#*:}"
PHASES="${PHASES:-1 2 3 4 5 6 7 8 9 10}"
FAILS=""
RESULT=""

for P in $PHASES; do
  if [ "$P" = "10" ]; then
    continue
  fi
  echo "=== PHASE $P: build + fresh DB + run (port $CHAIN_PORT) ==="
  # Port hygiene: a previous phase's server must be gone before this one
  # binds.
  for OLD in $(netstat -ano | grep ":$CHAIN_PORT" | grep LISTENING | awk '{print $NF}' | sort -u); do
    taskkill //F //PID "$OLD" 2>/dev/null
  done
  sleep 2
  go build -o /tmp/srv ./cmd/server || { echo "PHASE $P: server build FAILED"; FAILS="$FAILS $P"; RESULT="$RESULT p$P=build"; continue; }
  go build -o /tmp/p "./cmd/phase${P}test" || { echo "PHASE $P: client build FAILED"; FAILS="$FAILS $P"; RESULT="$RESULT p$P=build"; continue; }
  rm -f /tmp/chain.db /tmp/chain.db-shm /tmp/chain.db-wal
  DB_PATH=/tmp/chain.db HTTP_ADDR=":$CHAIN_PORT" WS_ADDR=":$CHAIN_PORT" TESTBED_FAST_CYCLE=1 \
    /tmp/srv > /tmp/chain_srv.log 2>&1 &
  SRV=$!
  sleep 12
  SERVER_URL="$CHAIN_URL" /tmp/p
  CODE=$?
  kill $SRV 2>/dev/null
  for OLD in $(netstat -ano | grep ":$CHAIN_PORT" | grep LISTENING | awk '{print $NF}' | sort -u); do
    taskkill //F //PID "$OLD" 2>/dev/null
  done
  sleep 2
  RESULT="$RESULT p$P=$CODE"
  if [ "$CODE" != "0" ]; then FAILS="$FAILS $P"; fi
done

# Phase 10: dedicated port, env-overridable client (B2/B6 design).
if echo " $PHASES " | grep -q " 10 "; then
  echo "=== PHASE 10: fresh DB + run (port 8090) ==="
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