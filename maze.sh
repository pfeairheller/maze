#!/bin/sh

PWD=`pwd`
export G3PO_ROOT=$PWD
export G3PO_SOURCE=$G3PO_ROOT/src/github.com/g3po
export GOPATH=$PWD:$PWD/vendor

export PATH=$PWD/bin:$PWD/vendor/bin:$PATH
export SLACK_BOT_TOKEN=xoxb-10899465207-JNXxFicSLMl7GkPO0PXSDRji
