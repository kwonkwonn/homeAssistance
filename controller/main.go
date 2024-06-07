package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/mochi-mqtt/server/v2/hooks/auth"
)

type authenticData struct{
	Name string
	age int
}

type moduleControlChannel struct{
	// authChan chan authenticData
	fanChan chan struct{}
	LEDChan chan struct{}
	// segChan chan int 

}
var controller moduleControlChannel

func main(){
	
	if(len(os.Args)<2){
		fmt.Println("need enough commands to start")
		panic("")
	}
	fmt.Println("program running....")
	command:=os.Args[1];
	switch command{
		case "run":
			run()
		default:
			panic("put valid command")
		
	}
	
}


func run(){
	wg:= sync.WaitGroup{}
	controller.fanChan=make(chan struct{})
	controller.LEDChan=make(chan struct{})

	AuthOption:= auth.Options{
		Ledger: &auth.Ledger{
		Auth: auth.AuthRules{ // Auth disallows all by default
			{Username: "kwon", Password: "123", Allow: true},
		},
		ACL: auth.ACLRules{ // ACL allows all by default
			},
		  },
		  }

	go MQTTInit(&AuthOption, controller)
	
	authentic:=make(chan authenticData)
	wg.Add(1)
	// _=InitContainer()
	go userAuthentication(authentic, &wg)
	fmt.Println("tag you NFC card for verification")
	var inputData string
loop:
	for {
		time.Sleep(1*time.Second)
		fmt.Println("looping")
		
		select{
		case data:= <-authentic:{
			wg.Add(1)
			fmt.Printf("%s %d",data.Name,data.age)
			
			go userAuthentication(authentic, &wg)}
		default:
		}
		
		if inputData=="exit"{
			break loop
		}
	}
	
	wg.Wait()

	fmt.Print("프로그램이 종료 되었습니다.")
}

func userAuthentication(channelForAuth chan authenticData,wg *sync.WaitGroup){
	//유저 NFC를 실행한다. 
	fmt.Println("nfc on")
	data,err:=NFCListener();
	if err!=nil{
		fmt.Print("error", err)
		data= authenticData{
			Name:"error",
			age:0,
		}
	}
	channelForAuth<-data
	defer wg.Done()
}