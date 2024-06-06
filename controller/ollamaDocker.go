package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type Conv struct{
	signal [2]int
	message string
}

func InitContainer()(cli *client.Client){
	
	var option container.StartOptions
	cli, err:= client.NewClientWithOpts(client.WithVersion("1.43"))
	if err!=nil{
		panic(err)
	}
	err = cli.ContainerStart(context.Background(), "1a8bd0ac705c", option)
	if err!=nil{
		panic(err)
	}
	return cli
}

func Communicate(message chan Conv,  inputMessage string){
	fmt.Println(inputMessage)
	newMessage:= Conv{
		signal:[2]int{0,0},
		message:"how are you doing",
	}
	message <-newMessage
}