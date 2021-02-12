#!/usr/bin/env bash

while true; do

  curl -X POST \
    -d"@payload/${1}.json" \
    -H'Content-Type:application/cloudevents+json' \
    http://localhost/demo/products

  sleep 0.5

done
