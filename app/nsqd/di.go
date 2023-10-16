package nsqd

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"os/signal"
	"syscall"
	"vehicle-log-grpc-api/config"
	"vehicle-log-grpc-api/internal/common/utils"
	nsqLoginImpl "vehicle-log-grpc-api/internal/handlers/nsqlogin/impl"
	nsqLogOutImpl "vehicle-log-grpc-api/internal/handlers/nsqlogout/impl"
	"vehicle-log-grpc-api/internal/repository/db"
	repoImpl "vehicle-log-grpc-api/internal/repository/db/impl"
	"vehicle-log-grpc-api/internal/repository/redis"
	redisImpl "vehicle-log-grpc-api/internal/repository/redis/impl"
	logUcImpl "vehicle-log-grpc-api/internal/usecase/logs/impl"
)

type NSqd struct {
	repo  db.Repository
	redis redis.Redis
	cfg   *config.Config
}

func (n *NSqd) Start(types string) {

	configNqs := nsq.NewConfig()
	var consumerTemp *nsq.Consumer

	logUC := logUcImpl.New(n.repo, n.cfg, n.redis)

	switch types {
	case ListenLogIn:
		consumer, err := nsq.NewConsumer(TopicLogIn, ChannelMetrics, configNqs)
		if err != nil {
			log.Fatal(err)
		}
		consumerTemp = consumer
		consumer.AddHandler(nsqLoginImpl.New(logUC))

		err = consumer.ConnectToNSQLookupd(n.cfg.Nsq.Addr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("login state")
	case ListenLogOut:
		consumer, err := nsq.NewConsumer(TopicLogOut, ChannelMetrics, configNqs)
		if err != nil {
			log.Fatal(err)
		}
		consumerTemp = consumer
		consumer.AddHandler(nsqLogOutImpl.New(logUC))

		err = consumer.ConnectToNSQLookupd(n.cfg.Nsq.Addr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("logout state")
	default:

		log.Println("Selected types not found")
		return
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan

	// disconnect
	if consumerTemp != nil {
		consumerTemp.Stop()
	}

}

func New() NSqd {
	cfg := config.LoadDefault()
	sqlCon, err := utils.NewBunPostgresSQLCon(config.Env(cfg.Server.Mode), *cfg)
	if err != nil {
		panic(err)
	}
	repo := repoImpl.NewRepo(sqlCon)
	r := redisImpl.New(cfg)
	return NSqd{
		cfg:   cfg,
		repo:  repo,
		redis: r,
	}
}
