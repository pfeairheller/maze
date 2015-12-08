#!/bin/sh

PWD=`pwd`
export MAZE_ROOT=$PWD
export MAZE_SOURCE=$MAZE_ROOT/src/github.com/MAZE
export GOPATH=$PWD:$PWD/vendor

export PATH=$PWD/bin:$PWD/vendor/bin:$PATH
