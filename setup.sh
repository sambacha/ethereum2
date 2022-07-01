#!/usr/bin/env bash


python3 --version
echo "Creating Python3 Virtual Env..."
python3 -m venv ~/venv
source ~/venv/bin/activate

cd module_api/concensus-specs
python3 setup.py install
