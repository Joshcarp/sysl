MODEL=model/1_project.sysl
TRANSFORMS=transforms/*
GRAMMAR=grammars/go.gen.g
START=goFile

TODOS_MODEL=model/1_project.sysl
TODOS_OUT=../gen/1_project
OUTPUT=output
APPNAME=Simple

# declare -a allxforms=(*.sysl)

this: */*
		$(foreach file, ${TRANSFORMS}, sysl codegen  --transform ${TRANSFORMS} --grammar ${GRAMMAR} --start ${START}  --outdir=output --app-name ${APPNAME} ${MODEL})
		go fmt $TODOS_OUT/*.go
		goimports -w $TODOS_OUT/*.go
		

