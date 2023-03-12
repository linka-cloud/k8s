IMAGE := k8s-api-builder

ORIGINAL_MODULE := k8s.io/api
MODULE := go.linka.cloud/k8s

API_DIRS := $(shell find . -name 'types.go'|xargs dirname|sed 's|./||'|sort)
API_GROUPS := $(addprefix $(ORIGINAL_MODULE)/,$(API_DIRS))
API_GROUP_LIST := $(shell echo $(API_GROUPS)|sed 's| |,|g')

gen-scheme:
	@hack/gen-scheme > scheme.go

# replace all struct fields with pointers
# regenerate deep copy
# regenerate proto
unstructured-local:
	@find . -name 'types.go' -type f -exec sed -i -E 's|^\t(\w+) (\w+) (`json:.+)|\t\1 *\2 \3|g' {} \;
	@find . -name 'types.go' -type f -exec sed -i 's|*ResourceList|ResourceList|g' {} \;
	@echo "generating deepcopy for k8s.io/api"
	@cd /go/src
	@deepcopy-gen \
		--input-dirs "$(API_GROUP_LIST)" \
		--output-file-base zz_generated.deepcopy \
		--go-header-file hack/boilerplate.go.txt \
		--add_dir_header \
		--bounding-dirs "$(API_GROUP_LIST)" \
		--logtostderr
	@echo "generating proto for k8s.io/api"
	@go-to-protobuf \
		--go-header-file hack/boilerplate.go.txt \
		--proto-import $$GOPATH/src/github.com/gogo/protobuf/protobuf \
		--packages="$(API_GROUP_LIST)"
	@./hack/gen-scheme > scheme.go
	@find . -name '*.go' -type f -exec sed -i 's|"$(ORIGINAL_MODULE)"|"$(MODULE)"|g' {} \;
	@find . -type f -exec sed -i 's|"$(ORIGINAL_MODULE)/|"$(MODULE)/|g' {} \;
	@find . -maxdepth 1 -name '*.go' -type f -exec sed -i 's|package api|package k8s|g' {} \;
	@find . -name '*.proto' -type f -exec sed -i 's|package k8s.io.api.|package go.linka.cloud.k8s.|g' {} \;
	@sed -i 's|module $(ORIGINAL_MODULE)|module $(MODULE)|g' go.mod
	@find . -name 'generated.pb.go' -exec sed -i 's|"k8s\.io\.api\.|"go.linka.cloud.k8s.|g' {} \;


unstructured: tidy docker-image
	@docker run --rm -i -t -v $(PWD):/go/src/$(ORIGINAL_MODULE) -w /go/src/$(ORIGINAL_MODULE) $(IMAGE) \
		make unstructured-local

show-api-groups:
	@echo $(API_GROUP_LIST)

docker-image:
	@docker build -t $(IMAGE) .

tidy:
	@go mod tidy
