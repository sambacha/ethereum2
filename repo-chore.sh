#!/usr/bin/env bash
git fetch --recurse-submodules -j10
git submodule sync --recursive && git submodule update --init --recursive --jobs=8
