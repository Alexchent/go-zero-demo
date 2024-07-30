## go-zero-demo

[English](README.md) | 简体中文



## 前言
记录一下使用go-zero开发的心得


### 技术栈

- k8s
- go-zero
- nginx网关
- filebeat
- kafka
- go-stash
- elasticsearch
- kibana
- prometheus
- grafana
- jaeger
- go-queue
- asynq
- asynqmon
- dtm
- docker
- docker-compose
- mysql
- redis

#### 启动user-api
```/bin/zsh
cd mall/user/api && go run user.go
```

#### 测试登录
```/bin/zsh
curl --location 'localhost:8888/user/login' \
--header 'Content-Type: application/json' \
--data '{"username":"434","password":"123456"}'
```