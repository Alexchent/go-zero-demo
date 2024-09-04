.PHONY: doc
doc:
	goctl api doc --dir .

.PHONY: protoc
protoc:
	goctl rpc protoc policy.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=.

.PHONY: model-ddl
model-ddl:
	goctl model mysql ddl --src model/*.sql --dir model --cache

.PHONY: model-source
model-source:
	goctl model mysql datasource -t policy --url 'root:root@tcp(127.0.0.1:3306)/demo' --dir model

#生成 greet.proto 文件
.PHONY: proto
proto:
	goctl rpc -o polic.proto

USER=$(shell echo $$USER)
GITHASH?=$(shell git rev-parse --short HEAD)

.PHONY: img
img:
	docker build -t ${USER}/policy:${GITHASH} .

# 使用 .env 文件中的变量
include .env
.PHONY: test
test:
	echo ${DB_DATABASE}