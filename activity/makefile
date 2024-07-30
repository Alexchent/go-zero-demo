.PHONY: doc
doc:
	goctl api doc --dir .

.PHONY: go
go:
	goctl api go --dir . --api *.api

.PHONY: model-ddl
model-ddl:
	goctl model mysql ddl --src model/mysql/*.sql --dir model

.PHONY: model-source
model-source:
	goctl model mysql datasource -t wechat_users --url 'root:root@tcp(127.0.0.1:3306)/demo' --dir model

# 生成docker并运行
#docker:
#	goctl docker --go *.go --exe activity \
#	&& docker build -t activity:v1 . \
# 	&& docker run --rm -it -p 8887:8888 activity:v1



