#!/bin/bash

# log envs
export BUILD_LOG=${BUILD_LOG:-"logs/mud.local.build.log"}
export BACK_LOG=${BACK_LOG:-"logs/mud.local.back.log"}
export FRONT_LOG=${FRONT_LOG:-"logs/mud.local.front.log"}

# backend envs
export DB_HOST=${DB_HOST:-"localhost"}
export DB_PORT=${DB_PORT:-"27017"}
export DB_USERNAME=${DB_USERNAME:-"taeho"}
export DB_PASSWORD=${DB_PASSWORD:-""}
export MUD_PEPPER=${MUD_PEPPER:-"_sudo_mud_PePPeP_"}

# frontend envs
export PORT=${PORT:-"80"}
export API_DOMAIN=${API_DOMAIN:-"localhost"}
export API_PORT=${API_PORT:-"8080"}