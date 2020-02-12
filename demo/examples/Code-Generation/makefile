
# if this script doesn't run, try running `pip2 uninstall sysl` or
# `pip uninstall sysl` as make might be running an old sysl version

TMP = .tmp# Cache the server lib directory in tmp
SERVERLIB = /var/tmp
TRANSLOCATION = .tmp/server-lib/codegen/transforms
# TRANSFORMS =$(wildcard .tmp/server-lib/codegen/transforms/*.sysl)
TRANSFORMS= svc_error_types.sysl svc_handler.sysl svc_interface.sysl svc_router.sysl svc_types.sysl
GRAMMAR=$(wildcard .tmp/server-lib/codegen/grammars/go.gen.g)
START=goFile
OUTPUT=simple
APPNAME=Simple

# Always execute these with just `make`
.PHONY: setup clean gen
all: setup gen clean

# try to clone, then try to fetch and pull
setup:
	@echo --------------------------------------------------------
	@echo Syncing server-lib to $(SERVERLIB)
	@set -e
	@git clone https://github.service.anz/sysl/server-lib/ $(SERVERLIB)/server-lib || true  # Don't fail 
	@cd  $(SERVERLIB)/server-lib && git fetch && git pull && cd ../..
	@mkdir -p $(TMP)/server-lib/
	@echo Copying server-lib to $(TMP)
	@cp -rf $(SERVERLIB)/server-lib $(TMP)/
	@echo Fetching Success 
	@echo --------------------------------------------------------

# Generate files with internal git service
gen:
	@echo --------------------------------------------------------
	@echo Running sysl transforms ....
	@export GOPRIVATE="github.service.anz"
	@$(foreach file, $(TRANSFORMS), $(shell sysl codegen --transform $(TRANSLOCATION)/$(file) --grammar ${GRAMMAR} --start ${START} --outdir=${OUTPUT} --app-name ${APPNAME} $(input)))
	gofmt -s -w $(OUTPUT)/*
	goimports -w $(OUTPUT)/*
	@echo --------------------------------------------------------

clean:
	@echo --------------------------------------------------------
	echo cleaning up $(TMP)
	rm -rf $(TMP)
	@echo --------------------------------------------------------