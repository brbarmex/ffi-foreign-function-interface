PROTO_DIR=./cloudsdk/protos
OUT_DIR=./cloudsdk/protos

JAVA_OUT=$(OUT_DIR)/java
CSHARP_OUT=$(OUT_DIR)/csharp
PYTHON_OUT=$(OUT_DIR)/python
GO_OUT=$(OUT_DIR)/go
GO_LOCAL_OUT=$(OUT_DIR)

export LD_LIBRARY_PATH=/Users/wakuba/workspace/fii-poc/ffi-foreign-function-interface/java
export LD_LIBRARY_PATH=.:$LD_LIBRARY_PATH


all: generate

build:
	cd ./cloudsdk && go build -o cloudsdk.so -buildmode=c-shared main.go

generate:
	mkdir -p $(JAVA_OUT) $(CSHARP_OUT) $(PYTHON_OUT) $(GO_OUT)
	protoc \
		--go_out=$(GO_OUT) --go_opt=paths=source_relative \
		--go_out=$(GO_LOCAL_OUT) --go_opt=paths=source_relative \
		--java_out=$(JAVA_OUT) \
		--csharp_out=$(CSHARP_OUT) \
		--python_out=$(PYTHON_OUT) \
		$(wildcard $(PROTO_DIR)/*.proto)

clean:
	rm -rf $(GO_OUT)/*.pb.go $(JAVA_OUT)/*.java $(CSHARP_OUT)/*.cs $(PYTHON_OUT)/*.py
