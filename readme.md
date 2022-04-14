# bluebell

bluebell是一个基于gin框架和Vue框架搭建的前后端分离的web项目。

## 项目清单

1. 后端：后端使用golang和gin作为编程语言和框架开发了Restful API，返回状态码遵循HTTP语义。
2. 分布式ID生成
3. JWT认证
4. zap日志库
5. Viper配置管理
6. 令牌桶限流
7. Go语言操作MySQL
8. Go语言操作Redis

## 压力测试

Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.25ms    1.75ms  42.89ms   91.77%
    Req/Sec    13.47k     2.59k   37.19k    75.29%
  Latency Distribution
     50%    0.95ms
     75%    1.56ms
     90%    2.70ms
     99%    8.58ms
  3216645 requests in 30.03s, 481.62MB read
Requests/sec: 107130.17
Transfer/sec:     16.04MB

## 下载及运行

### 下载

```git
git clone https://github.com/Xinky777/bluebell.git
```

### 运行

```go
go build

./bulebell
```





