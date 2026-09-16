#!/bin/sh
# Phase 10 standalone verification: fresh DB, fast-cycle, phase10test +
# phase10report. Exit 0 = all observable proxies PASS; exit 2 = one or more
# FAIL-with-telemetry (an accepted exit form per proposal §2.1 — read the
# telemetry before judging); other codes are setup failures.
#
# Uses a dedicated port (P10_HTTP_ADDR / P10_SERVER_URL) so it can run
# alongside any other local server on 8080 — coexistence hardening (B6).
set -u
P10_HTTP_ADDR="${P10_HTTP_ADDR:-:8090}"
P10_SERVER_URL="${P10_SERVER_URL:-http://localhost:8090}"
export P10_HTTP_ADDR P10_SERVER_URL

go build -o /tmp/srv10 ./cmd/server || exit 1
go build -o /tmp/p10 ./cmd/phase10test || exit 1
go build -o /tmp/p10rep ./cmd/phase10report || exit 1
rm -f /tmp/p10.db /tmp/p10.db-shm /tmp/p10.db-wal
DB_PATH=/tmp/p10.db HTTP_ADDR="$P10_HTTP_ADDR" WS_ADDR="$P10_HTTP_ADDR" \
  TESTBED_FAST_CYCLE=1 /tmp/srv10 > /tmp/srv10.log 2>&1 &
SRV=$!
sleep 12
P10_SERVER_URL="$P10_SERVER_URL" /tmp/p10
CODE=$?
kill $SRV 2>/dev/null
# B3/B4: the balance report reads the DB after the server has shut down
# (no concurrent-writer risk); it appends the B4 section to the same file.
cd "$(dirname "$0")" && /tmp/p10rep /tmp/p10.db
echo "PHASE10_EXIT:$CODE"
