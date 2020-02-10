#!/bin/sh
set -e ;\
if [ ! -d ".tmp/server-lib" ]; then \
    git clone https://github.service.anz/sysl/server-lib/ .tmp/server-lib \;
fi
cd  .tmp/server-lib  && git fetch && git pull && cd ../..
