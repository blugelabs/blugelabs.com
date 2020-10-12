all: build

build: functions-with-index

public:
	hugo

site-index: public bluge_index_dir_exec
	@echo "building site index"
	bluge_index_dir/bluge_index_dir public site_index.bluge

bluge_index_dir_exec:
	@echo "installing bluge_index_dir"
	cd bluge_index_dir; go build

functions-with-index: site-index bluge_add_to_elf
	@echo "functions with index"
	./bluge_add_to_elf functions/site-search index site_index.bluge
	mv functions/site-search.withindex functions/site-search

functions:
	@echo "functions"
	mkdir -p functions
	cd funcsrc/site-search; go build -o ../../functions/site-search

bluge_add_to_elf: functions
	@echo "add to elf"
	GOBIN=`pwd` go get github.com/blugelabs/bluge_directory_elf/cmd/bluge_add_to_elf

clean:
	-rm -rf public functions site_index.bluge
