#!/usr/bin/env bash
gradlew --version
cd module_clients/teku
./gardlew build
cd -
cp -rf module_clients/teku/ethereum/spec/build/resources/main/tech/pegasys/teku/spec/config/presets/mainnet output/
