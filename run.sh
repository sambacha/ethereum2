
(cd lib/builder-specs; )
git submodule sync --recursive && git submodule update --init --recursive --jobs=8
redocly bundle builder-oapi.yaml > mev-boost-api.yml