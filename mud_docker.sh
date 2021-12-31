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
    source $ROOTDIR/env/env.docker.sh
}

build_back() {
    cd $ROOTDIR/back
    echo "--- BACK ${NOW} ---" >> $ROOTDIR/$BUILD_LOG
    go install 2>&1 >> $ROOTDIR/$BUILD_LOG &
}

build_front() {
    cd $ROOTDIR/front
    npm_install
    echo "--- FRONT ${NOW} ---" >> $ROOTDIR/$BUILD_LOG
    npm run build 2>&1 >> $ROOTDIR/$BUILD_LOG &
}

start_back() {
    mkdir -p $ROOTDIR/logs
    echo "--- START ${NOW} ---" >> $ROOTDIR/$BACK_LOG
    mud 2>&1 >> $ROOTDIR/$BACK_LOG &
}

start_front() {
    cd $ROOTDIR/front
    npm_install
    mkdir -p $ROOTDIR/logs
    echo "--- START ${NOW} ---" >> $ROOTDIR/$BACK_LOG
    npm start 2>&1 >> $ROOTDIR/$FRONT_LOG &
}

kill_back() {
    kill -9 `ps | grep 'mud$' | awk '{print $1}'` 2>&1 > /dev/null
}

kill_front() {
    kill -9 `ps | grep 'node$' | awk '{print $1}'` 2>&1 > /dev/null
}

# ---------------------------------------------------
# Logic
# ---------------------------------------------------
setenv

case $CMD in

"build")
    build_back && build_front
    ;;

"start")
    start_back && start_front
    ;;

"kill")
    kill_back
    kill_front
    ;;

*)
    kill_back
    kill_front
    build_back
    build_front
    start_back
    start_front
    ;;

esac
