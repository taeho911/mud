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
    docker build --target test -t ${BACK_IMAGE}_test:${BACK_TAG} ${ROOTDIR}/back
    docker run -it --rm ${BACK_IMAGE}_test:${BACK_TAG}
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
