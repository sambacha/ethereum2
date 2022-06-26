#!/usr/bin/env bash

[ -z "$1" -o $(echo "$1" | tr '/' '\n' | wc -l) != 2 ] &&
  {
    echo "Usage: $(basename "$0") <vendor>/<repo> [destdir] # 'destdir' defaults to 'vendor/repo'"
    exit 1
  }
REPO="$1"

DEST="vendor/${REPO#*/}"
[ -n "$2" ] && DEST="$2"