#!/usr/bin/env bash
echo "Archiving..."
git archive -o "${PWD##*/}.zip" HEAD