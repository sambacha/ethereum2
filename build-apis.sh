#!usr/bin/env bash

cd mev-builder-specs
git submodule sync --recursive && git submodule update --init --recursive
redocly bundle builder-oapi.yaml > bundle.yml
cd ..
cp lib/builder-specs/mev-boost-api.yml $PWD
#cp mev-builder-specs/bundle.yml $PWD/mev-builder-api.yml