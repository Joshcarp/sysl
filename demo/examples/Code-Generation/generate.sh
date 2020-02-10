#!/bin/sh

set -e -x

MODEL="model/1_project.sysl"
TRANSFORMS="transforms"
GRAMMAR="grammars/go.gen.g"
START="goFile"

TODOS_MODEL="model/1_project.sysl"
TODOS_OUT="../gen/1_project"
OUTPUT="output"
APPNAME=Simple
declare -a allxforms=(svc_client.sysl)

for xform in "${allxforms[@]}"
do
    sysl codegen  --transform $TRANSFORMS/$xform --grammar $GRAMMAR --start $START  --outdir=output --app-name $APPNAME $MODEL
done

go fmt $TODOS_OUT/*.go
goimports -w $TODOS_OUT/*.go