go install go-micro.dev/v4/cmd/micro@master

micro new service github.com/wsjcko/shopcart

mkdir -p domain/{model,repository,service} 
mkdir -p protobuf/{pb,pbserver} 
mkdir -p proto/{pb,pbserver}
mkdir common

go mod edit --module=github.com/wsjcko/shopcart
go mod edit --go=1.17  

gorm 有个根据创建表sql 生成model  : gormt

清除mod下载的包
go clean -modcache


### consul 微服务注册中心和配置中心
docker search --filter is-official=true --filter stars=3 consul
docker pull consul

## 生产环境要注意数据落盘  -v /data/consul:/data/consul
docker run -d -p 8500:8500 consul:latest 

### 注册中心
"github.com/asim/go-micro/plugins/registry/consul/v4"

### 配置中心
"github.com/asim/go-micro/plugins/config/source/consul/v4"

### consul数据库配置
http://127.0.0.1:8500/ui/dc1/kv/create

key: micro/config/mysql

{
  "host":"127.0.0.1",
  "user":"root",
  "pwd":"123456",
  "database":"shopdb",
  "port":3306
}


### 链路追踪jaeger 耶格 
[官方文档](https://www.jaegertracing.io/docs/1.32/)

docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 9411:9411 \
  jaegertracing/all-in-one:latest

  http://127.0.0.1:16686/search


### 绑定链路追踪 服务端绑定handle 客户端绑定client
micro.WrapClient(opentracing4.NewClientWrapper(opentracing.GlobalTracer()))
micro.WrapClient(opentracing4.NewClientWrapper(opentracing.GlobalTracer()))

opentracing4 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v4"

### 创建链路追踪实例
"github.com/opentracing/opentracing-go"
"github.com/uber/jaeger-client-go"
"github.com/uber/jaeger-client-go/config"


### 集流量控制、熔断、容错，负载均衡等hystrix-go
docker search hystrix
docker pull mlabouardy/hystrix-dashboard
docker run --name hystrix-dashboard -d -p 9002:9002 mlabouardy/hystrix-dashboard:latest


### 购物车加入熔断（客户端），限流（服务端），负载均衡（客户端）

github.com/asim/go-micro/go-plugins/wrapper/ratelimiter/uber/v4 限流










  监控告警Prometheus、日志接入ELK，到最后k8s部署