#!/bin/bash

pushd ../..

docker build -f build/package/Dockerfile -t csv-to-json .