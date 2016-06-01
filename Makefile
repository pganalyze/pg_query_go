root_dir := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

LIB_PG_QUERY_TAG = 9.5-1.3.0

TMPDIR = $(root_dir)/tmp
LIBDIR = $(TMPDIR)/libpg_query
LIBDIRGZ = $(TMPDIR)/libpg_query-$(LIB_PG_QUERY_TAG).tar.gz

default: test

$(LIBDIR): $(LIBDIRGZ)
	mkdir -p $(LIBDIR)
	cd $(TMPDIR); tar -xzf $(LIBDIRGZ) -C $(LIBDIR) --strip-components=1
	cd $(LIBDIR); make build

$(LIBDIRGZ):
	mkdir -p $(TMPDIR)
	curl -o $(LIBDIRGZ) https://codeload.github.com/lfittl/libpg_query/tar.gz/$(LIB_PG_QUERY_TAG)

build: $(LIBDIR)

clean:
	-@ $(RM) -r $(TMPDIR)

test: build
	go get github.com/kr/pretty
	go test -v ./ ./nodes

.PHONY: default build clean test
