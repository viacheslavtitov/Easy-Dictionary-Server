#!/usr/bin/env bash
set -euo pipefail

PORT="${1:-8080}"

echo "Looking for processes listening on port ${PORT}..."

if PIDS=$(lsof -n -i tcp:"${PORT}" -sTCP:LISTEN -t 2>/dev/null); then
  if [[ -z "${PIDS}" ]]; then
    echo "No listeners on port ${PORT}."
    exit 0
  fi
else
  PIDS=$(netstat -anv | awk -v p=".${PORT}" '
    /LISTEN/ && $4 ~ p { print $9 }' 2>/dev/null || true)
  if [[ -z "${PIDS}" ]]; then
    echo "No listeners on port ${PORT}."
    exit 0
  fi
fi

echo "Found PID(s): ${PIDS}"

for pid in ${PIDS}; do
  echo "Sending SIGTERM to PID ${pid}..."
  kill "${pid}" 2>/dev/null || true
done

sleep 1

STILL_ALIVE=()
for pid in ${PIDS}; do
  if kill -0 "${pid}" 2>/dev/null; then
    STILL_ALIVE+=("${pid}")
  fi
done

if [[ ${#STILL_ALIVE[@]} -gt 0 ]]; then
  echo "Forcing kill (SIGKILL) for: ${STILL_ALIVE[*]}"
  for pid in "${STILL_ALIVE[@]}"; do
    kill -9 "${pid}" 2>/dev/null || true
  done
else
  echo "All processes terminated gracefully."
fi

echo "Done."
