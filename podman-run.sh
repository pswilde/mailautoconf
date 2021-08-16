#!/usr/bin/env bash
podman run --name mailautoconf \
  --rm \
  -p "8010:8010" \
  -v ./config:/mailautoconf/config \
  pswilde/mailautoconf
