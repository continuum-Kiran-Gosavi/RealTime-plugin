package main

import (
	"fmt"

	"github.com/ContinuumLLC/SamplePlugin/src/communication"
)

func main() {

	pluginListener := &communication.RTSListener{}
	err := pluginListener.SendMessage("Hello from plugin!")
	if err != nil {
		fmt.Println("error :%v", err)
	}
	fmt.Printf("successs")

}
