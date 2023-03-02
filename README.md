# go-zero-demo



## 启动user-api
`cd mall/user/api && go run user.api`

测试登录
```zsh
curl --location 'localhost:8888/user/login' \
--header 'Content-Type: application/json' \
--data '{"username":"434","password":"123456"}'
```