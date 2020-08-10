.PHONY: default build test benchmark update_source clean

default: test

build:
	go build

test: build
	go get github.com/kr/pretty
	go test -v ./ ./nodes

benchmark:
	go build -a
	go test -test.bench=. -test.run=XXX -test.benchtime 10s -test.benchmem -test.cpu=4
	#go test -c -o benchmark
	#GODEBUG=schedtrace=100 ./benchmark -test.bench=BenchmarkRawParseCreateTableParallel -test.run=XXX -test.benchtime 20s -test.benchmem -test.cpu=16

# --- Below only needed for releasing new versions

LIB_PG_QUERY_TAG = 12-latest-develop-protobuf

root_dir := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
LIB_TMPDIR = $(root_dir)/tmp
LIBDIR = $(LIB_TMPDIR)/libpg_query
LIBDIRGZ = $(TMPDIR)/libpg_query-$(LIB_PG_QUERY_TAG).tar.gz

$(LIBDIR): $(LIBDIRGZ)
	mkdir -p $(LIBDIR)
	cd $(LIB_TMPDIR); tar -xzf $(LIBDIRGZ) -C $(LIBDIR) --strip-components=1

$(LIBDIRGZ):
	mkdir -p $(LIB_TMPDIR)
	curl -o $(LIBDIRGZ) https://codeload.github.com/lfittl/libpg_query/tar.gz/$(LIB_PG_QUERY_TAG)

IGNORED_PROTOBUF_FILES := $(wildcard $(LIBDIR)/protobuf/protobuf-v3.12.0/google/protobuf/*test.cc) \
$(wildcard $(LIBDIR)/protobuf/protobuf-v3.12.0/google/protobuf/*test.pb.cc) \
$(wildcard $(LIBDIR)/protobuf/protobuf-v3.12.0/google/protobuf/test_*.cc) \
$(wildcard $(LIBDIR)/protobuf/protobuf-v3.12.0/google/protobuf/unittest_*.cc) \
$(LIBDIR)/protobuf/protobuf-v3.12.0/google/protobuf/map_lite_test_util.cc

PROTOBUF_FILES := $(filter-out $(IGNORED_PROTOBUF_FILES), $(wildcard $(LIBDIR)/protobuf/protobuf-v3.12.0/google/protobuf/*.cc)) \
$(filter-out $(wildcard $(LIBDIR)/protobuf/protobuf-v3.12.0/google/protobuf/stubs/*test.cc), $(wildcard $(LIBDIR)/protobuf/protobuf-v3.12.0/google/protobuf/stubs/*.cc)) \
$(filter-out $(wildcard $(LIBDIR)/protobuf/protobuf-v3.12.0/google/protobuf/io/*test.cc), $(wildcard $(LIBDIR)/protobuf/protobuf-v3.12.0/google/protobuf/io/*.cc)) \
$(filter-out $(wildcard $(LIBDIR)/protobuf/protobuf-v3.12.0/google/protobuf/util/*test.cc), $(wildcard $(LIBDIR)/protobuf/protobuf-v3.12.0/google/protobuf/util/*.cc)) \
$(filter-out $(wildcard $(LIBDIR)/protobuf/protobuf-v3.12.0/google/protobuf/util/internal/*test.cc), $(wildcard $(LIBDIR)/protobuf/protobuf-v3.12.0/google/protobuf/util/internal/*.cc))

update_source: $(LIBDIR)
	rm -f parser/*.{c,cc,h}
	rm -fr parser/include
	rm -fr parser/protobuf-v*
	# Reduce everything down to one directory
	cp -a $(LIBDIR)/src/* parser/
	mv parser/postgres/* parser/
	rmdir parser/postgres
	cp -a $(LIBDIR)/pg_query.h parser/include
	# Make sure every .c and .cc file in the top-level directory is its own translation unit
	mv parser/*{_conds,_defs,_helper}.c parser/include
	mv parser/*{_conds,_defs}.cc parser/include
	# Protobuf definitions
	cp -a $(LIBDIR)/protobuf/parse_tree.proto nodes/
	mkdir -p $(PWD)/bin
	GOBIN=$(PWD)/bin go install github.com/golang/protobuf/protoc-gen-go
	PATH=$(PWD)/bin:$(PATH) protoc --go_out=. nodes/parse_tree.proto
	# Protobuf library code
	mkdir -p parser/include/protobuf
	cp -a $(LIBDIR)/protobuf/*.h parser/include/protobuf
	mkdir -p parser/include/protobuf-c
	cp -a $(LIBDIR)/protobuf-c/*.h parser/include/protobuf-c
	cp -a $(LIBDIR)/protobuf-c/*.h parser/include
	cp -a $(PROTOBUF_FILES) parser/
	cp -a $(LIBDIR)/protobuf/protobuf-v3.12.0/google parser/include
	cp -a $(LIBDIR)/protobuf/*.cc parser/
	cp -a $(LIBDIR)/protobuf/*.c parser/
	cp -a $(LIBDIR)/protobuf-c/*.c parser/
	# Other support files
	rm -fr testdata
	cp -a $(LIBDIR)/testdata testdata
	# Update nodes directory
	#ruby scripts/generate_nodes.rb

clean:
	-@ $(RM) -r $(LIB_TMPDIR)
