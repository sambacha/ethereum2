#!usr/bin/env bash
(cd execution-apis; yarn install)
npm run build:spec && npm run generate-clients
ls build/docs/gatsby/src
cd -
cp execution-apis/build/docs/gatsby/src/openrpc.json $PWD
./node_modules/.bin/openrpc-cli validate openrpc.json

