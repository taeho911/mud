#!/bin/bash

# log envs
export BUILD_LOG=${BUILD_LOG:-"~/mud/logs/mud.local.build.log"}
export BACK_LOG=${BACK_LOG:-"~/mud/logs/mud.local.back.log"}
export FRONT_LOG=${FRONT_LOG:-"~/mud/logs/mud.local.front.log"}

# backend envs
export BACK_CONTAINER_NAME=${BACK_CONTAINER_NAME:-"mud_back"}
export BACK_BASE_IMAGE=${BACK_IMAGE:-"golang"}
export BACK_BASE_TAG=${BACK_TAG:-"1.17-alpine3.14"}
export BACK_IMAGE=${BACK_IMAGE:-"mud_back"}
export BACK_TAG=${BACK_TAG:-"latest"}

export DB_HOST=${DB_HOST:-"localhost"}
export DB_PORT=${DB_PORT:-"27017"}
export DB_USERNAME=${DB_USERNAME:-"taeho"}
export DB_PASSWORD=${DB_PASSWORD:-""}
export MUD_PEPPER=${MUD_PEPPER:-"_sudo_mud_PePPeP_"}

# frontend envs
export FRONT_CONTAINER_NAME=${FRONT_CONTAINER_NAME:-"mud_front"}
export FRONT_BASE_IMAGE=${FRONT_BASE_IMAGE:-"node"}
export FRONT_BASE_TAG=${FRONT_BASE_TAG:-"14.17.5-alpine3.14"}
export FRONT_IMAGE=${FRONT_IMAGE:-"mud_front"}
export FRONT_TAG=${FRONT_TAG:-"latest"}

export FRONT_PORT=${PORT:-"3000"}
export API_DOMAIN=${BACK_CONTAINER_NAME}
export API_PORT=${API_PORT:-"9011"}

# analyzer envs
export ANAL_CONTAINER_NAME=${ANAL_CONTAINER_NAME:-"mud_analyzer"}
export ANAL_HOST=${ANAL_CONTAINER_NAME}
export ANAL_PORT-${ANAL_PORT:-"9012"}
