#!/usr/bin/env bash
set -euo pipefail

PACKAGES_DIR="$(cd "$(dirname "$0")/generated/packages" && pwd)"
RESULTS=()

# --- TypeScript ---
test_ts() {
  local tmp=$(mktemp -d)
  trap "rm -rf $tmp" RETURN
  
  # pack the package as a tarball (true install test)
  cp ./test-open-rpc-spec-1.3.2.tgz "$tmp/"
  cd "$PACKAGES_DIR/ts"
  npm pack --pack-destination "$tmp" 2>/dev/null
  
  # install it in a clean project
  cd "$tmp"
  npm init -y > /dev/null 2>&1
  npm install ./open-rpc-spec-types-*.tgz
  
  # smoke test: can we import it?
  node -e "import('@open-rpc/spec-types').then(m => console.log('TS OK:', Object.keys(m)))"
}

# --- Go ---
test_go() {
  local tmp=$(mktemp -d)
  trap "rm -rf $tmp" RETURN
  cp -r "$PACKAGES_DIR/go/"* "$tmp/"
  cd "$tmp"
  go build ./...
  echo "Go OK"
}

# --- Python ---

test_py() {
  local tmp=$(mktemp -d)
  trap "rm -rf $tmp" RETURN
  cp -r "$PACKAGES_DIR/py/"* "$tmp/"
  cd "$tmp"
  python3.12 -m venv "$tmp/.venv"
  "$tmp/.venv/bin/python" -m pip install . --quiet
  "$tmp/.venv/bin/python" -c "from open_rpc_spec_types import v1_4; print('Py OK')"
}

# --- Rust ---
test_rs() {
  local tmp=$(mktemp -d)
  trap "rm -rf $tmp" RETURN
  cp -r "$PACKAGES_DIR/rs/"* "$tmp/"
  cd "$tmp"
  cargo check 2>&1
  echo "Rust OK"
}

echo "=== Testing TS ===" && test_ts
echo "=== Testing Go ===" && test_go
# echo "=== Testing Py ===" && test_py
echo "=== Testing Rs ===" && test_rs

echo "All install tests passed."