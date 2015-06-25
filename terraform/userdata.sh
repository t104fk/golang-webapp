#!/bin/bash

docker run --name ecs-agent -d \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v /var/log/ecs/:/log -p 127.0.0.1:51678:51678 \
  -v /var/lib/ecs/data:/data \
  -e ECS_LOGFILE=/log/ecs-agent.log \
  -e ECS_LOGLEVEL=info \
  -e ECS_DATADIR=/data \
  -e ECS_CLUSTER=golang-webapp \
  amazon/amazon-ecs-agent:v1.2.0
