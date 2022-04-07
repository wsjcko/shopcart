package main

import (
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	ratelimit4 "github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v4"
	opentracing4 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v4"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/opentracing/opentracing-go"
	"github.com/wsjcko/shopcart/common"
	"github.com/wsjcko/shopcart/domain/repository"
	"github.com/wsjcko/shopcart/domain/service"
	"github.com/wsjcko/shopcart/handler"
	pb "github.com/wsjcko/shopcart/protobuf/pb"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
)

var (
	MICRO_SERVICE_NAME   = "go.micro.service.shop.cart"
	MICRO_VERSION        = "latest"
	MICRO_ADDRESS        = "127.0.0.1:8087"
	MICRO_QPS            = 100
	MICRO_CONSUL_HOST    = "127.0.0.1"
	MICRO_CONSUL_PORT    = "8500"
	MICRO_CONSUL_ADDRESS = "127.0.0.1:8500"
)

func main() {
	//配置中心
	consulConfig, err := common.GetConsulConfig(MICRO_CONSUL_HOST, MICRO_CONSUL_PORT, "/micro/config")
	if err != nil {
		log.Fatal(err)
	}
	//注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			MICRO_CONSUL_ADDRESS,
		}
	})

	//链路追踪
	t, io, err := common.NewTracer(MICRO_SERVICE_NAME, "127.0.0.1:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//数据库设置
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//禁止副表 gorm默认使用复数映射，go代码的单数、复数struct形式都匹配到复数表中,开启后，将严格匹配，遵守单数形式
	db.SingularTable(true)

	// 设置服务
	srv := micro.NewService(
		micro.Name(MICRO_SERVICE_NAME),
		micro.Version(MICRO_VERSION),
		//这里设置地址和需要暴露的端口
		micro.Address(MICRO_ADDRESS),
		//添加注册中心
		micro.Registry(consulRegistry),
		//绑定链路追踪 服务端绑定handle 客户端绑定client
		micro.WrapHandler(opentracing4.NewHandlerWrapper(opentracing.GlobalTracer())),
		//添加限流
		micro.WrapHandler(ratelimit4.NewHandlerWrapper(MICRO_QPS)),
	)

	//初始化建表
	err = repository.NewCartRepository(db).InitTable()
	if err != nil {
		log.Fatal(err)
	}

	cartService := service.NewCartService(repository.NewCartRepository(db))

	// Initialise service
	srv.Init()

	// Register Handler
	pb.RegisterShopCartHandler(srv.Server(), &handler.ShopCart{CartService: cartService})

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
