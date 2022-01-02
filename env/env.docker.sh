#!/bin/bash

# host
export SUBNET=${SUBNET:-"172.18.1.0/24"}
export GATEWAY=${GATEWAY:-"172.18.1.1"}

# log envs
export LOG_PATH=${LOG_PATH:-"~/mud/logs"}

# resources
export BACK_CPU_LIMIT=
export BACK_MEM_LIMIT=
export BACK_CPU_RESERV=
export BACK_MEM_RESERV=
export FRONT_CPU_LIMIT=
export FRONT_MEM_LIMIT=
export FRONT_CPU_RESERV=
export FRONT_MEM_RESERV=

# backend envs
export BACK_CONTAINER_NAME=${BACK_CONTAINER_NAME:-"mud_back"}
export BACK_BASE_IMAGE=${BACK_IMAGE:-"golang"}
export BACK_BASE_TAG=${BACK_TAG:-"1.17-alpine3.14"}
export BACK_IMAGE=${BACK_IMAGE:-"mud_back"}
export BACK_TAG=${BACK_TAG:-"latest"}
export BACK_CMD=${BACK_CMD:-"mud"}
export BACK_TARGET=${BACK_TARGET:-"prod"}

export DB_HOST=${GATEWAY}
export DB_PORT=${DB_PORT:-"27017"}
export DB_USERNAME=${DB_USERNAME:-""}
export DB_PASSWORD=${DB_PASSWORD:-""}
export MUD_PEPPER=${MUD_PEPPER:-"_sudo_mud_PePPeP_"}

# frontend envs
export FRONT_CONTAINER_NAME=${FRONT_CONTAINER_NAME:-"mud_front"}
export FRONT_BASE_IMAGE=${FRONT_BASE_IMAGE:-"node"}
export FRONT_BASE_TAG=${FRONT_BASE_TAG:-"14.17.5-alpine3.14"}
export FRONT_IMAGE=${FRONT_IMAGE:-"mud_front"}
export FRONT_TAG=${FRONT_TAG:-"latest"}
export FRONT_CMD=${FRONT_CMD:-"npm start"}

export FRONT_PORT=${PORT:-"3000"}
export API_DOMAIN=${BACK_CONTAINER_NAME}
export API_PORT=${API_PORT:-"9011"}

# analyzer envs
export ANAL_CONTAINER_NAME=${ANAL_CONTAINER_NAME:-"mud_analyzer"}
export ANAL_HOST=${ANAL_CONTAINER_NAME}
export ANAL_PORT=${ANAL_PORT:-"9012"}
