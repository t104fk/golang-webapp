#!/bin/bash

aws ecs create-cluster --cluster-name "golang-webapp" --profile takasing
aws ecs list-clusters --profile takasing

aws ecs register-task-definition --cli-input-json file:///Users/a12681/go/src/golang-webapp/ecs/golang_webapp_task.json --profile takasing
aws ecs list-tasks --profile takasing

aws ecs create-service --cluster golang-webapp --service-name ecs-golang-webapp --task-definition golang-webapp --desired-count 1 --profile takasing
