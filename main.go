package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	userPb "github.com/yannyy/istio-user/proto"
	"google.golang.org/grpc"
)

const (
	address     = "user:80"
	defaultName = "world"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/hello/:name", sayHello)
	r.Run()
}

func sayHello(ctx *gin.Context) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	name := ctx.Param("name")

	c := userPb.NewUserClient(conn)
	ctx1, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	r, err := c.SayHello(ctx1, &userPb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	greeting := r.GetGreeting()
	ctx.JSON(200, gin.H{
		"message": greeting,
	})
}
