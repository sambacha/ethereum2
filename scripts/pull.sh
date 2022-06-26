#!/usr/bin/env bash
set -o errexit

git pull && git submodule sync --recursive && git submodule update --init --recursive

echo "Submodule pull finished"
sleep 1
exit 0