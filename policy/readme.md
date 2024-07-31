# 策略服务

## 测试
需要先安装工具grpcurl
```
brew install grpcurl
```

1. 添加一条规则
```/bin/zsh
grpcurl -plaintext -H "package: TD1201" -d '{
  "cate": "pay_mode",
  "attr": "2231",
  "rule": "ReadTime >= 1200 || ListenTime >= 1200"
}' localhost:12345 policy.Policy.Add
```

2. 测试用户可访问的资源
```/bin/zsh
grpcurl -plaintext -d '{
  "cate": "pay_mode",
  "user": {
    "Balance": "90",
    "AppId": "TD1201"
  }
}' localhost:12345 policy.Policy.Check
```