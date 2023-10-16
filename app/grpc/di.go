package grpc

import (
	"google.golang.org/grpc"
	"log"
	"net"
	protoPb "vehicle-log-grpc-api/common/grpc/proto"
	"vehicle-log-grpc-api/config"
	"vehicle-log-grpc-api/internal/common/utils"
	logsHandlerImpl "vehicle-log-grpc-api/internal/handlers/logs/impl"
	townHandlerImpl "vehicle-log-grpc-api/internal/handlers/towns/impl"
	typesHandlerImpl "vehicle-log-grpc-api/internal/handlers/types/impl"
	vehicleHandlerImpl "vehicle-log-grpc-api/internal/handlers/vehicle/impl"
	"vehicle-log-grpc-api/internal/repository/db"
	repoImpl "vehicle-log-grpc-api/internal/repository/db/impl"
	"vehicle-log-grpc-api/internal/repository/redis"
	redisImpl "vehicle-log-grpc-api/internal/repository/redis/impl"
	logsUCImpl "vehicle-log-grpc-api/internal/usecase/logs/impl"
	townUCImpl "vehicle-log-grpc-api/internal/usecase/towns/impl"
	typesUseCaseImpl "vehicle-log-grpc-api/internal/usecase/types/impl"
	vehilceUCImpl "vehicle-log-grpc-api/internal/usecase/vehicle/impl"
)

type Grpc struct {
	repo  db.Repository
	redis redis.Redis
	cfg   *config.Config
}

func (s *Grpc) registerServer(srv *grpc.Server) {
	// Declare UseCase
	vehicleUseCase := vehilceUCImpl.New(s.repo, s.cfg, s.redis)
	townUseCase := townUCImpl.New(s.repo, s.cfg, s.redis)
	typesUseCase := typesUseCaseImpl.New(s.repo, s.cfg, s.redis)
	logUseCase := logsUCImpl.New(s.repo, s.cfg, s.redis)

	// Declare Server
	vehicleServer := vehicleHandlerImpl.New(vehicleUseCase)
	townServer := townHandlerImpl.New(townUseCase)
	typesServer := typesHandlerImpl.New(typesUseCase)
	logsServer := logsHandlerImpl.New(logUseCase)

	protoPb.RegisterVehiclesServer(srv, vehicleServer)
	protoPb.RegisterTownsServer(srv, townServer)
	protoPb.RegisterTypesServer(srv, typesServer)
	protoPb.RegisterLogsServer(srv, logsServer)
}

func (s *Grpc) Start() {
	srv := grpc.NewServer()

	if !config.Env(s.cfg.Server.Mode).IsProd() {
		grpc.EnableTracing = true
	}

	s.registerServer(srv)

	l, err := net.Listen("tcp", s.cfg.Server.Port)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", s.cfg.Server.Port, err)
	}

	log.Printf("grpc serve at potrt :%s", s.cfg.Server.Port)
	err = srv.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}

func NewGrpc() Grpc {
	cfg := config.LoadDefault()
	sqlCon, err := utils.NewBunPostgresSQLCon(config.Env(cfg.Server.Mode), *cfg)
	if err != nil {
		panic(err)
	}
	repo := repoImpl.NewRepo(sqlCon)
	r := redisImpl.New(cfg)
	return Grpc{
		cfg:   cfg,
		repo:  repo,
		redis: r,
	}
}
