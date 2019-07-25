#!/bin/bash

GOOS=linux go build -a --ldflags '-linkmode external -extldflags "-static"' . ;
docker build -t rodriguesdossantosvincent/loginsrv . ;
docker login -u="$DOCKER_USERNAME" -p="$DOCKER_PASSWORD" ;
docker push rodriguesdossantosvincent/loginsrv ;
