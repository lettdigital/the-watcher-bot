#!/usr/bin/env bash
set -e

$(aws ecr get-login --no-include-email)
#export ENVIROMENT_DEPLOY=$1
#export ECR_IMAGE_URL=868884350453.dkr.ecr.us-east-1.amazonaws.com/team-data-capture/the-watcher-bot:$1

docker-compose -f build.yml build
docker-compose -f build.yml push

docker push 868884350453.dkr.ecr.us-east-1.amazonaws.com/team-data-capture/the-watcher-bot:latest
docker rmi $(docker image ls -q 868884350453.dkr.ecr.us-east-1.amazonaws.com/team-data-capture/the-watcher-bot:latest)