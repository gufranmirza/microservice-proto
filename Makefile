all: generate-proto

generate-proto:
	protoc -I=./proto/v1/product/ --go_out="${GOPATH}/src/" ./proto/v1/product/*.proto 
	protoc -I=./proto/v1/product/ --doc_out=markdown,doc.md:./proto/v1/product ./proto/v1/product/*.proto 