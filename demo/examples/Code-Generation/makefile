
# if this script doesn't run, try running `pip2 uninstall sysl` or
# `pip uninstall sysl` as make might be running an old sysl version

TMP = .tmp# Cache the server lib directory in tmp

TRANSFORMS=$(wildcard .tmp/server-lib/codegen/transforms/*.sysl)
GRAMMAR=$(wildcard .tmp/server-lib/codegen/grammars/go.gen.g)
START=goFile
OUTPUT=output
APPNAME=Simple

# Always execute these with just `make`
.PHONY: download clean gen
all: setup this

# try to clone, then try to fetch and pull
setup:
	@echo --------------------------------------------------------
	@echo Syncing server-lib ....
	set -e
	git clone https://github.service.anz/sysl/server-lib/ $(TMP)/server-lib || true  # Don't fail 
	cd  $(TMP) && git fetch && git pull && cd ../.. && \
	cp -rf $(TMP) $(TMP)
	echo Fetching Success 
	@echo --------------------------------------------------------

# Generate files with internal git service
gen:
	@echo --------------------------------------------------------
	@echo Running sysl transforms ....
	export GOPRIVATE="github.service.anz"
	$(foreach file, $(TRANSFORMS), $(shell sysl codegen --transform $(file) --grammar ${GRAMMAR} --start ${START} --outdir=output --app-name ${APPNAME} $(input)))
	gofmt -s -w .
	goimports -w .
	@echo --------------------------------------------------------

clean:
	rm -rf $(TMP)