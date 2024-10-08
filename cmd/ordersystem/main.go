package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"

	graphqlhandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/andretefras/fullcycle-go-challenge-3-clean-architecture/configs"
	"github.com/andretefras/fullcycle-go-challenge-3-clean-architecture/internal/event/handler"
	"github.com/andretefras/fullcycle-go-challenge-3-clean-architecture/internal/infra/graph"
	"github.com/andretefras/fullcycle-go-challenge-3-clean-architecture/internal/infra/grpc/pb"
	"github.com/andretefras/fullcycle-go-challenge-3-clean-architecture/internal/infra/grpc/service"
	"github.com/andretefras/fullcycle-go-challenge-3-clean-architecture/internal/infra/web/webserver"
	"github.com/andretefras/fullcycle-go-challenge-3-clean-architecture/pkg/events"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rabbitMQChannel := getRabbitMQChannel()

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)
	listOrdersUseCase := NewListOrdersUseCase(db)

	httpServer := webserver.NewWebServer(configs.WebServerPort)
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)
	httpServer.AddHandler("/orders/list", webOrderHandler.List)
	httpServer.AddHandler("/orders/create", webOrderHandler.Create)
	fmt.Println("Starting web server on port", configs.WebServerPort)
	go httpServer.Start()

	grpcServer := grpc.NewServer()
	createOrderService := service.NewOrderService(*createOrderUseCase, *listOrdersUseCase)
	pb.RegisterOrderServiceServer(grpcServer, createOrderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go func() {
		err := grpcServer.Serve(lis)
		if err != nil {
			panic(err)
		}
	}()

	srv := graphqlhandler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase: *createOrderUseCase,
		ListOrdersUseCase:  *listOrdersUseCase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", configs.GraphQLServerPort)
	http.ListenAndServe(":"+configs.GraphQLServerPort, nil)
}

func getRabbitMQChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}
