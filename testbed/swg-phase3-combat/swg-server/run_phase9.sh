#!/bin/sh
# Phase 9 standalone verification: fresh DB, fast-cycle, phase9test only.
set -u
go build -o /tmp/srv ./cmd/server || exit 1
go build -o /tmp/p9 ./cmd/phase9test || exit 1
rm -f /tmp/p9.db /tmp/p9.db-shm /tmp/p9.db-wal
DB_PATH=/tmp/p9.db HTTP_ADDR=:8080 TESTBED_FAST_CYCLE=1 /tmp/srv > /tmp/srv.log 2>&1 &
SRV=$!
sleep 12
/tmp/p9
CODE=$?
kill $SRV 2>/dev/null
echo "PHASE9_EXIT:$CODE"
