#!/bin/bash

rm data/out/output.json

docker run --rm -v $(pwd)/data/in:/in:ro -v $(pwd)/data/out:/out csv-to-json /in/input.csv /out/output.json

cat data/out/output.json