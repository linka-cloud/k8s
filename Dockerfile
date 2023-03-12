FROM golang

WORKDIR /tmp/dummy

COPY go.mod go.sum ./

RUN go install golang.org/x/tools/cmd/goimports && \
    go install k8s.io/code-generator/cmd/deepcopy-gen && \
    go install k8s.io/code-generator/cmd/go-to-protobuf && \
    go install github.com/gogo/protobuf/protoc-gen-gogo && \
    apt update && \
    apt install -y protobuf-compiler && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /go/src/k8s.io/api
RUN rm -rf /tmp/dummy && \
    mkdir -p /go/src/github.com/gogo/protobuf && \
    git clone --depth 1 https://github.com/gogo/protobuf /go/src/github.com/gogo/protobuf
