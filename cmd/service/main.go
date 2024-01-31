package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/golobby/container/v3"

	pa "github.com/octoposprime/op-be-logging/internal/application/presentation/adapter"
	as "github.com/octoposprime/op-be-logging/internal/application/service"
	ds "github.com/octoposprime/op-be-logging/internal/domain/service"
	ia_ebus "github.com/octoposprime/op-be-logging/pkg/infrastructure/adapter/ebus"
	ia_repo "github.com/octoposprime/op-be-logging/pkg/infrastructure/adapter/repository"
	pc_grpc "github.com/octoposprime/op-be-logging/pkg/presentation/controller/grpc"
	pc_probe "github.com/octoposprime/op-be-logging/pkg/presentation/controller/probe"
	tconfig "github.com/octoposprime/op-be-shared/tool/config"
	tgorm "github.com/octoposprime/op-be-shared/tool/gorm"
	tredis "github.com/octoposprime/op-be-shared/tool/redis"
)

var internalConfig tconfig.InternalConfig
var dbConfig tconfig.DbConfig
var redisConfig tconfig.RedisConfig

func main() {
	internalConfig.ReadConfig()
	dbConfig.ReadConfig()
	redisConfig.ReadConfig()
	var err error

	fmt.Println("Starting Logger Service...")
	dbClient, err := tgorm.NewGormClient(tgorm.PostgresGormClient).Connect(dbConfig.PostgresDb.Host, dbConfig.PostgresDb.Port, dbConfig.PostgresDb.UserName, dbConfig.PostgresDb.Password, dbConfig.PostgresDb.Database)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to DB")

	redisClient := tredis.NewRedisClient(redisConfig.Redis.Host, redisConfig.Redis.Port, redisConfig.Redis.Password, redisConfig.Redis.Db)
	_, err = redisClient.Ping(context.TODO()).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to Redis")

	cont := container.New()

	//Domain Logger Service
	err = cont.Singleton(func() *ds.Service {
		return ds.NewService()
	})
	if err != nil {
		panic(err)
	}

	//Infrastructure Logger Db Repository Adapter
	err = cont.Singleton(func() ia_repo.DbAdapter {
		return ia_repo.NewDbAdapter(dbClient)
	})
	if err != nil {
		panic(err)
	}

	//Infrastructure Logger EBus Adapter
	err = cont.Singleton(func() ia_ebus.EBusAdapter {
		return ia_ebus.NewEBusAdapter(redisClient)
	})
	if err != nil {
		panic(err)
	}

	//Application Logger Service
	err = cont.Singleton(func(s *ds.Service, d ia_repo.DbAdapter, e ia_ebus.EBusAdapter) *as.Service {
		return as.NewService(s, &d, &e)
	})
	if err != nil {
		panic(err)
	}

	//Application Logger Query Adapter
	err = cont.Singleton(func(s *as.Service) pa.QueryAdapter {
		return pa.NewQueryAdapter(s)
	})
	if err != nil {
		panic(err)
	}

	//Application Logger Command Adapter
	err = cont.Singleton(func(s *as.Service) pa.CommandAdapter {
		return pa.NewCommandAdapter(s)
	})
	if err != nil {
		panic(err)
	}

	var queryHandler pa.QueryAdapter
	err = cont.Resolve(&queryHandler)
	if err != nil {
		panic(err)
	}

	var commandHandler pa.CommandAdapter
	err = cont.Resolve(&commandHandler)
	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}
	if !internalConfig.Local {
		wg.Add(1)
		go pc_probe.NewProbeAPI().Serve(internalConfig.Restapi.ProbePort)
	}
	wg.Add(1)
	go pc_grpc.NewGrpc(queryHandler, commandHandler).Serve(internalConfig.Grpc.LoggerPort)
	wg.Wait()

}
