
# if this script doesn't run, try running `pip2 uninstall sysl` or
# `pip uninstall sysl` as make might be running an old sysl version

MODEL=model/1_project.sysl
TRANSFORMS=$(wildcard server-lib/codegen/transforms/*.sysl)
GRAMMAR=$(wildcard server-lib/codegen/grammars/go.gen.g)
START=goFile
TMP=.tmp
OUTPUT=output
APPNAME=Simple

# Always execute these with just `make`
.PHONY: download clean gen
all: setup this

# try to clone, then try to fetch and pull
setup:
	set -e
	git clone https://github.service.anz/sysl/server-lib/ .tmp/server-lib || true  # Don't fail 
	cd  .tmp/server-lib  && git fetch && git pull && cd ../..

# Generate files with internal git service
gen: */*
		export GOPRIVATE="github.service.anz"
		$(foreach file, $(TRANSFORMS), $(shell sysl codegen --transform $(file) --grammar ${GRAMMAR} --start ${START}  --outdir=output --app-name ${APPNAME} ${MODEL}))
		go fmt ${OUTPUT}/*.go
		goimports -w ${OUTPUT}/*.go

clean:
	rm -rf .tmp