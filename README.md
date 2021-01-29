# webGo
### IRIS学习

#### IRIS安装

```
参考链接
https://docs.iris-go.com/iris/getting-started/installation
```

#### 数据库安装

```
参考链接
https://gobook.io/read/gitea.com/xorm/manual-en-US/
```

#### 数据库结构

```
user表结构
+-----------+--------------+------+-----+---------+----------------+
| Field     | Type         | Null | Key | Default | Extra          |
+-----------+--------------+------+-----+---------+----------------+
| id        | int(10)      | NO   | PRI | NULL    | auto_increment |
| user_name | varchar(255) | NO   |     | NULL    |                |
| pwd       | varchar(255) | NO   |     | NULL    |                |
+-----------+--------------+------+-----+---------+----------------+
```

#### 目录文件说明

|--config
<br>
|&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|-------config.go 读取配置文件
<br>
|&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|-------config.toml 配置文件文档
<br>
|--database
<br>
|&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|-------database.go 数据库处理
<br>
|--utils
<br>
|&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|-------result.go 结果格式化
<br>
|--model
<br>
|&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|-------user.go 用户登录结构体
<br>
|--control
<br>
|&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|-------Test1.go 测试
<br>
|&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|-------user.go 用户登录相关
<br>
|--service
<br>
|&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|-------user.go 用户登录相关
<br>





