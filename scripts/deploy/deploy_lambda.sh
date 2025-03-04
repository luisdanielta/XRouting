#!/bin/bash

rm -rf package lambda_package.zip
mkdir package
cd ..
pip install \
  --platform manylinux2014_x86_64 \
  --target deploy/package \
  --implementation cp \
  --python-version 3.12 \
  --only-binary=:all: --upgrade \
  -r requirements.txt
cp -r adapters core ports utils main.py deploy/package/
cd deploy/package
zip -r ../lambda_package.zip .