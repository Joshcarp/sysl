
# if this script doesn't run, try running `pip2 uninstall sysl` or
# `pip uninstall sysl` as make might be running an old sysl version

TMP = .tmp# Cache the server lib directory in tmp
SERVERLIB = /var/tmp
TRANSLOCATION = .tmp/server-lib/codegen/transforms
TRANSFORMS= svc_error_types.sysl svc_handler.sysl svc_interface.sysl svc_router.sysl svc_types.sysl

GRAMMAR=$(wildcard .tmp/server-lib/codegen/grammars/go.gen.g)
START=goFile
app=Simple

# Always execute these with just `make`
.PHONY: setup clean gen
all: setup gen format tmp

# try to clone, then try to fetch and pull
setup:
	@echo Syncing server-lib to $(SERVERLIB)
	@set -e
	@git clone https://github.service.anz/sysl/server-lib/ $(SERVERLIB)/server-lib || true  # Don't fail 
	@cd  $(SERVERLIB)/server-lib && git fetch && git checkout tags/v0.1.8
	@mkdir -p $(TMP)/server-lib/
	@echo Copying server-lib to $(TMP)
	@cp -rf $(SERVERLIB)/server-lib $(TMP)/
	@echo Fetching Success 
	@mkdir -p ${app}
	

# Generate files with internal git service
gen:
	@echo Running sysl transforms ....
	@$(foreach file, $(TRANSFORMS), $(shell sysl codegen --transform $(TRANSLOCATION)/$(file) --grammar ${GRAMMAR} --start ${START} --outdir=${app} --app-name ${app} $(input)))
	

format:
	gofmt -s -w ${app}/*
	goimports -w ${app}/*

# Remove the tmp directory after 
tmp:
	rm -rf $(TMP)

# Remove the generated files
clean:
	rm -rf ${app}
	