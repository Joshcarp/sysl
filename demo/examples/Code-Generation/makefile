
# if this script doesn't run, try running `pip2 uninstall sysl` or
# `pip uninstall sysl` as make might be running an old sysl version

TMP = .tmp# Cache the server lib directory in tmp
SERVERLIB = /var/tmp
TRANSFORMS=$(wildcard .tmp/server-lib/codegen/transforms/*.sysl)
GRAMMAR=$(wildcard .tmp/server-lib/codegen/grammars/go.gen.g)
START=goFile
OUTPUT=output
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
	$(foreach file, $(TRANSFORMS), $(shell sysl codegen --transform $(file) --grammar ${GRAMMAR} --start ${START} --outdir=${OUTPUT} --app-name ${APPNAME} $(input)))
	gofmt -s -w output/*
	goimports -w output/*
	@echo --------------------------------------------------------

clean:
	@echo --------------------------------------------------------
	echo cleaning up $(TMP)
	rm -rf $(TMP)
	@echo --------------------------------------------------------

handler:
	mkdir implementation || true
	echo package implementation >> implementation/gencallback.go
	impl -dir simple 'g Callback' simple.GenCallback >> implementation/gencallback.go

	echo package implementation >> implementation/methods.go
	
	echo package implementation >> implementation/resthandler.go
	echo type RestHandler struct { >> implementation/resthandler.go
	echo Router             []handlerinitialiser.HandlerInitialiser >> implementation/resthandler.go
	echo ServerConfig       *config.CommonHTTPServerConfig >> implementation/resthandler.go
	echo MAdminServerConfig *config.CommonHTTPServerConfig >> implementation/resthandler.go
	echo MLibraryConfig     *config.LibraryConfig >> implementation/resthandler.go
	echo } >> implementation/resthandler.go
	impl -dir simple 'r Handler' simple.GenCallback >> implementation/resthandler.go



