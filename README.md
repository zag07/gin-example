



### 目录结构


```text
├── assets
│   ├── static
│   └── template
├── cmd
├── configs
├── init
├── docs
├── internal
│   ├── app
│   │   ├── controllers
│   │   ├── models
│   │   └── services
│   ├── pkg
│   │   └── ...
│   └── routes
│       └── api.go
├── pkg
│   ├── config
│   ├── database
│   ├── logger
│   └── routing 
├── storage
│   └── logs 
├── test
├── main.go
├── go.mod
├── go.sum
├── ...
```



### Swaager

| 注解     | 描述                                                         |
| -------- | ------------------------------------------------------------ |
| @Summary | 摘要                                                         |
| @Produce | API 可以产生的 MIME 类型的列表，MIME 类型你可以简单的理解为响应类型，例如：json、xml、html 等等 |
| @Param   | 参数格式，从左到右分别为：参数名、入参类型、数据类型、是否必填、注释 |
| @Success | 响应成功，从左到右分别为：状态码、参数类型、数据类型、注释   |
| @Failure | 响应失败，从左到右分别为：状态码、参数类型、数据类型、注释   |
| @Router  | 路由，从左到右分别为：路由地址，HTTP 方法                    |
