#!/bin/bash

# ---------------------------------------------------
# Environment Variables
# ---------------------------------------------------
CMD=$1
ROOTDIR=`pwd`
NOW=`date +"%Y-%m-%d %H:%M:%S"`

# ---------------------------------------------------
# Functions
# ---------------------------------------------------
setenv() {
    source ${ROOTDIR}/env/env.docker.sh
}

back_test() {
    export BACK_TARGET=test
    export BACK_IMAGE=${BACK_IMAGE}_test
    docker-compose build backend
    docker run -it --rm ${BACK_IMAGE}:${BACK_TAG}
}

build() {
    docker-compose build
}

up() {
    docker-compose up
}

down() {
    docker-compose down
}

# ---------------------------------------------------
# Logic
# ---------------------------------------------------
setenv

case $CMD in

"test")
    back_test
    ;;

"build")
    build
    ;;

"up")
    up
    ;;

"down")
    down
    ;;

*)
    ;;

esac
