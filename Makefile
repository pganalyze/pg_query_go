root_dir := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

TMPDIR = $(root_dir)/tmp
LIBDIR = $(TMPDIR)/libpg_query-master
LIBDIRGZ = $(TMPDIR)/libpg_query-master.tar.gz

default: test

$(LIBDIR): $(LIBDIRGZ)
	cd $(TMPDIR); tar -xf $(LIBDIRGZ)
	cd $(LIBDIR); make

$(LIBDIRGZ):
	mkdir -p $(TMPDIR)
	curl -o $(LIBDIRGZ) https://codeload.github.com/lfittl/libpg_query/tar.gz/master

build: $(LIBDIR)

clean:
	-@ $(RM) -r $(TMPDIR)

test: build
	go test -v

.PHONY: default build clean test
