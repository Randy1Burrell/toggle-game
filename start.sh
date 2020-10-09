#!/bin/bash

docker stop go-docker
docker rm go-docker
docker build -t go-docker -f dev.dockerfile . | grep l
docker run -it --rm -v `pwd`:/app --name go-docker go-docker
