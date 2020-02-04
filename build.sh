#!/usr/bin/env bash
docker build --no-cache -t dashbase/heplify:latest -f docker/heplify/Dockerfile .

docker push dashbase/heplify:latest
