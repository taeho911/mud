#!/bin/bash

# ---------------------------------------------------
# Environment Variables
# ---------------------------------------------------
CMD=$1
ROOTDIR=`pwd`

# ---------------------------------------------------
# Functions
# ---------------------------------------------------
setenv() {
    source $ROOTDIR/env/env.local.sh
}

npm_install() {
    [ -d "node_modules" ] || npm install
}

build_back() {
    cd $ROOTDIR/back
    go install 2>&1 >> $ROOTDIR/$BUILD_LOG &
}

build_front() {
    cd $ROOTDIR/front
    npm_install
    npm run build 2>&1 >> $ROOTDIR/$BUILD_LOG &
}

start_back() {
    mkdir -p $ROOTDIR/logs
    echo "--- START ${date} ---" >> $ROOTDIR/$BACK_LOG
    mud 2>&1 >> $ROOTDIR/$BACK_LOG &
}

start_front() {
    cd $ROOTDIR/front
    npm_install
    mkdir -p $ROOTDIR/logs
    echo "--- START ${date} ---" >> $ROOTDIR/$BACK_LOG
    npm start 2>&1 >> $ROOTDIR/$FRONT_LOG &
}

open_log_back() {
    code $ROOTDIR/$BACK_LOG
}

open_log_front() {
    code $ROOTDIR/$FRONT_LOG
}

kill_back() {
    kill `ps | grep 'mud$' | awk '{print $1}'`
}

kill_front() {
    kill `ps | grep 'node$' | awk '{print $1}'`
}

# ---------------------------------------------------
# Logic
# ---------------------------------------------------
setenv

case $CMD in

"build")
    build_back
    build_front
    ;;

"start")
    kill_back
    kill_front
    start_back
    start_front
    ;;

"log")
    open_log_back
    open_log_front
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
