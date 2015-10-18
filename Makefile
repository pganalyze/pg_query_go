root_dir := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

TMPDIR = $(root_dir)/tmp
LIBDIR = $(TMPDIR)/libpg_query-master
LIBDIRGZ = $(TMPDIR)/libpg_query-master.tar.gz

build:
	mkdir -p $(TMPDIR)
	curl -o $(LIBDIRGZ) https://codeload.github.com/lfittl/libpg_query/tar.gz/master
	cd $(TMPDIR); tar -xf $(LIBDIRGZ)
	cd $(LIBDIR); make

clean:
	-@ $(RM) -r $(TMPDIR)

.PHONY: build clean
