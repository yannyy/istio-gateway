# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
PACKAGE=github.com/yannyy/istio-gateway
PROTOC=protoc
BINARY_NAME=istio-gateway
DOCKER=docker
GITVER=`git rev-parse --short HEAD`

all: push
build: test
	GOOS=linux GOARCH=amd64 $(GOBUILD) 
test:
	$(GOTEST) -v ./...
docker: build 
	docker build -f Dockerfile -t $(BINARY_NAME):$(GITVER) .
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run: grpc
	$(GOBUILD) $(PACKAGE)	
	./$(BINARY_NAME)
