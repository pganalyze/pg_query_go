.PHONY: default build clean test update_source

default: test

build:
	go build

test: build
	go get github.com/kr/pretty
	go test -v ./ ./nodes

# --- Below only needed for releasing new versions

LIB_PG_QUERY_TAG = 9.5-1.3.0

root_dir := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
TMPDIR = $(root_dir)/tmp
LIBDIR = $(TMPDIR)/libpg_query
LIBDIRGZ = $(TMPDIR)/libpg_query-$(LIB_PG_QUERY_TAG).tar.gz

$(LIBDIR): $(LIBDIRGZ)
	mkdir -p $(LIBDIR)
	cd $(TMPDIR); tar -xzf $(LIBDIRGZ) -C $(LIBDIR) --strip-components=1

$(LIBDIRGZ):
	mkdir -p $(TMPDIR)
	curl -o $(LIBDIRGZ) https://codeload.github.com/lfittl/libpg_query/tar.gz/$(LIB_PG_QUERY_TAG)

update_source: $(LIBDIR)
	rm -f parser/*.{c,h}
	rm -fr parser/include
	# Reduce everything down to one directory
	cp -a $(LIBDIR)/src/* parser/
	mv parser/postgres/* parser/
	rmdir parser/postgres
	cp -a $(LIBDIR)/pg_query.h parser/include
	# Make sure every .c file in the top-level directory is its own translation unit
	mv parser/*{_conds,_defs,_helper,scan}.c parser/include
	# Other support files
	rm -fr testdata
	cp -a $(LIBDIR)/testdata testdata
	# Update nodes directory
	ruby scripts/generate_nodes.rb

clean:
	-@ $(RM) -r $(TMPDIR)
