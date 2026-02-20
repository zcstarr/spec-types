#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")"

echo "=== Running install tests in Docker ==="

docker compose -f docker-compose.test.yml up \
  --remove-orphans \
  --force-recreate 2>&1

FAILED=$(docker compose -f docker-compose.test.yml ps -a --format '{{.Name}} {{.ExitCode}}' | grep -v ' 0$' || true)

docker compose -f docker-compose.test.yml down --volumes --remove-orphans 2>/dev/null

if [ -n "$FAILED" ]; then
  echo ""
  echo "=== FAILED ==="
  echo "$FAILED"
  exit 1
else
  echo ""
  echo "=== All install tests passed ==="
fi
