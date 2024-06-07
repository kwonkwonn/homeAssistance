package main

import (
	"fmt"
	"os/exec"
	"sync"
	"time"
)


func StartListening(){

}

func ModuleController(name string, wg *sync.WaitGroup, channel chan struct{}){
	cmd:= exec.Command("python3", "../controlNods/"+name+".py")
		err:=cmd.Start()
		if err!=nil{
		panic( err)		
	}
	go func(){
		for {
			time.Sleep(1*time.Second)
			select {
			case <-channel:
				return
			default:
			}
		}	
	}()
	cmd.Process.Kill()
	wg.Done()
}

func NFCListener()(data authenticData, error error){
	// cmd:= exec.Command("python3", "../controlNods/NFCReader.py")
	// cmd.Run()
	time.Sleep(5*time.Second)
	fmt.Println("121")

	datadata:=authenticData{ Name:"택용", age:25}
	return datadata,nil
}