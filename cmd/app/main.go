package main

import (
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/MorselShogiew/UsersManager/packages/color"
	pb "github.com/MorselShogiew/UsersManager/proto/user"
	"github.com/MorselShogiew/UsersManager/repos"
	"google.golang.org/grpc"
)

func main() {
	conf := NewConfig()
	var wg sync.WaitGroup

	db := &DBManager{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		db.Connect(&conf.psql)
	}()

	kafka := &KafkaProducer{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		kafka.Connect(&conf.kafka)
	}()

	redis := &RedisManager{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		redis.Connect(&conf.redis)
	}()

	wg.Wait()

	if db.ok {
		log.Println("Connection to the DB: " + color.Green + "done" + color.Reset)
		defer db.Close()
	} else {
		log.Fatalf(color.Red+"ERROR"+color.Reset+": Failed DB connection (%s:%d)", conf.psql.host, conf.psql.port)
	}

	if kafka.ok {
		log.Println("Connection to the Kafka: " + color.Green + "done" + color.Reset)
		defer kafka.Close()
	} else {
		log.Fatalf(color.Red+"ERROR"+color.Reset+": Failed Kafka connection (%s)", conf.kafka.server)
	}

	if redis.ok {
		log.Println("Connection to the Redis: " + color.Green + "done" + color.Reset)
		defer redis.Close()
	} else {
		log.Fatalf(color.Red+"ERROR"+color.Reset+": Failed Redis connection (%s)", conf.redis.address)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", conf.grpc.port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserRepoServer(grpcServer, repos.NewUserDB(db, redis, kafka))
	log.Printf("Start gRPC server %s", lis.Addr().String())
	grpcServer.Serve(lis)
}
