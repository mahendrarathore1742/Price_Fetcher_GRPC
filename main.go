package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/mahendrarathore1742/priceFetcher/client"
	"github.com/mahendrarathore1742/priceFetcher/proto"
)

func main() {

	// client := client.New("http://localhost:3000")

	// price, err := client.FetcherPrice(context.Background(), "ET")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%+v\n", price)
	var (
		jsonAddr = flag.String("json", ":3000", "listen address of the json transport")
		grpcAddr = flag.String("grpc", ":4000", "listen address of the grpc transport")
		svc      = loggingService{&priceService{}}
		ctx      = context.Background()
	)

	flag.Parse()

	grpcClient, err := client.NewGRPCClinet(":4000")

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		time.Sleep(3 * time.Second)
		resp, err := grpcClient.FetchPrice(ctx, &proto.PriceRequest{Ticker: "BTC"})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", resp)
	}()

	go makeGRPCServerAndRun(*grpcAddr, svc)

	jsonServer := NewJsonAPiServer(*jsonAddr, svc)
	jsonServer.Run()

}
