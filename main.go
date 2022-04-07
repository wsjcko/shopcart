package main

import (
	"fmt"

	"github.com/asim/go-micro/plugins/registry/consul/v4"
	ratelimit4 "github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v4"
	opentracing4 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v4"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/opentracing/opentracing-go"
	cli2 "github.com/urfave/cli/v2"
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
	MICRO_SERVICE_NAME = "go.micro.service.shop.cart"
	MICRO_VERSION      = "latest"
	MICRO_ADDRESS      = "127.0.0.1:8087"
	MICRO_QPS          = 100
	DOCKER_HOST        = "127.0.0.1"
)

func SetDockerHost(host string) {
	DOCKER_HOST = host
}

func main() {
	function := micro.NewFunction(
		micro.Flags(
			&cli2.StringFlag{ //micro 多个选项 --ip
				Name:  "ip",
				Usage: "docker Host IP(ubuntu)",
				Value: "0.0.0.0",
			},
		),
	)

	function.Init(
		micro.Action(func(c *cli2.Context) error {
			ipstr := c.Value("ip").(string)
			if len(ipstr) > 0 { //后续校验IP
				fmt.Println("docker Host IP(ubuntu)1111", ipstr)
			}
			SetDockerHost(ipstr)
			return nil
		}),
	)

	fmt.Println("DOCKER_HOST ", DOCKER_HOST)

	var (
		MICRO_CONSUL_HOST    = DOCKER_HOST
		MICRO_CONSUL_PORT    = "8500"
		MICRO_CONSUL_ADDRESS = DOCKER_HOST + ":8500"
		MICRO_JAEGER_ADDRESS = DOCKER_HOST + ":6831"
	)
	fmt.Println("MICRO_CONSUL_ADDRESS ", MICRO_CONSUL_ADDRESS)
	fmt.Println("MICRO_JAEGER_ADDRESS ", MICRO_JAEGER_ADDRESS)
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
	t, io, err := common.NewTracer(MICRO_SERVICE_NAME, MICRO_JAEGER_ADDRESS)
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//数据库设置
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@tcp("+mysqlInfo.Host+":"+mysqlInfo.Port+")/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
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
	// err = repository.NewCartRepository(db).InitTable()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	cartService := service.NewCartService(repository.NewCartRepository(db))

	// Initialise service
	srv.Init(
		micro.BeforeStart(func() error {
			log.Info("Log BeforeStart")
			return nil
		}),
		micro.AfterStart(func() error {
			log.Info("Log AfterStart")
			return nil
		}),
		micro.BeforeStop(func() error {
			log.Info("Log BeforeStop")
			return nil
		}),
		micro.AfterStop(func() error {
			log.Info("Log AfterStop")
			return nil
		}),
	)

	// Register Handler
	pb.RegisterShopCartHandler(srv.Server(), &handler.ShopCart{CartService: cartService})

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
