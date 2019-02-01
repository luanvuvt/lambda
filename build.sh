#!/bin/bash

rm ./hello ./hello.zip > /dev/null 2>&1

GOOS=linux go build -o hello

zip hello.zip hello

if [[ -e ./hello.zip ]]; then
  echo "Uploading hello.zip"
  aws lambda update-function-code --function-name helloLambda --zip-file fileb://hello.zip
fi