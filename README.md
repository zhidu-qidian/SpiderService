# SpiderService

为爬虫提供服务,避免创建大量到数据库的连接(**目前只为段子抓取提供存储接口**)

## Getting Started

### API

参见 API.md

### Prerequisites

`golang 1.8.3`

### Installing

Clone this project and execute: 

```
go build
./SpiderService
```

## Deployment

`go build` 后得到一个可执行文件 `SpiderService`

`./SpiderService` 或使用 `supervisord`

*config.json 文件必须和 SpiderService 可执行文件在同一个目录*

## Built With

* [elastic.v3](https://github.com/olivere/elastic) - elasticsearch golang driver
* [gin](https://github.com/gin-gonic/gin) - The web framework used
* [grpc](https://google.golang.org/grpc) - A high performance, open-source universal RPC framework

## Authors

* **[Sven Lee](https://github.com/lixianyang)** - *python工程师*

## License

智度科技股份有限公司
