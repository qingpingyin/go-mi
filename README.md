

# go-mi 仿小米商城



<<<<<<< HEAD
##### 此项目为前后端分离的项目，已经部署至:
=======
##### 此项目为前后端分离的项目，已经部署至:Cancel changes
>>>>>>> c90620f6b5650b1ebc0997e5a844e374e43d124a

[http://101.132.127.139/#/index](http://101.132.127.139/#/index)，如需查看前端请前往[mi](https://github.com/qingpingyin/mi)

后续将完善 第三方登录(未完成)，支付(未完成)，手机短信验证码(目前使用测试代码)，Elasticsearch重写搜索系统（未完成）,后台管理系统(已完成)：

------

#### 项目依赖

- Gin
- Gorm
- MySql
- Redis
- jwt-go
- go-mail
- go-redis
- logrus
- viper
- 七牛云对象存储

------

#### 目录结构

```
│  go.mod //项目依赖
│  go.sum
│  main.go //程序入口
│  README.md
├─api //api请求入口
├─config //项目配置入口
├─log	//日志文件
├─middleware //中间件
├─models //数据模型
├─pkg //第三方包
│  ├─cache //go-redis
│  ├─email //go-mail
│  ├─jwt  //jwt-go
│  ├─logger //logrus
│  ├─setting //viper
│  ├─sms //验证码
│  └─validate //validator/v10
├─routers
│      router.go //路由入口
├─service  
└─utils //第三方工具包
    ├─common //通用工具
    ├─request //请求封装
    └─response //响应封装
```

#### 运行部署

​		本项目通过ngixn 部署， supervisor管理后台进程，具体实现可参考李文周老师：[Go项目部署](https://www.liwenzhou.com/posts/Go/deploy_go_app/)

1. ```
   git clone git@github.com:qingpingyin/go-mi.git
   ```

2. ```
   go mod tidy 
   ```

3. 修改config文件下config.yaml各项配置，初始化项目配置。

4. 由于使用的gorm的数据迁移，运行该项目之后导入mi.sql，进行数据的配置。

5. ```
   go build
   ```

