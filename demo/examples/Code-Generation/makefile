
# if this script doesn't run, try running `pip2 uninstall sysl` or
# `pip uninstall sysl` as make might be running an old sysl version


MODEL=model/1_project.sysl
TRANSFORMS=$(wildcard transforms/*.sysl)
GRAMMAR=grammars/go.gen.g
START=goFile

# TODOS_MODEL=model/1_project.sysl
# TODOS_OUT=../gen/1_project
OUTPUT=output
APPNAME=Simple

this: */*
		which sysl

		$(foreach file, $(TRANSFORMS), $(shell sysl codegen --transform $(file) --grammar ${GRAMMAR} --start ${START}  --outdir=output --app-name ${APPNAME} ${MODEL}))
		# go fmt ${TODOS_OUT}/*.go
		# goimports -w ${TODOS_OUT}/*.go
		

