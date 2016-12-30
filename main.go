package main

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/raspi"
	"gobot.io/x/gobot/api"
	"github.com/miccab/gobot-go/experiments"
	"flag"
	"os"
	"strings"
)

const (
	LED_1 = "led1"
	LED_2 = "led2"
	BUZZ  = "buzz"
	LED_AND_BUZZ = "ledAndBuzz"
)

func main() {
	adaptor := raspi.NewAdaptor()

	var experiment string
	flag.StringVar(&experiment, "e", "", "type of experiment: " + strings.Join([]string{LED_1,LED_2,BUZZ,LED_AND_BUZZ}, ","))
	flag.Parse()

	var devices []gobot.Device
	var work func()

	switch experiment {
	case LED_1:
		devices, work = experiments.LedBlinking(adaptor)
	case LED_2:
		devices, work = experiments.TwoLedsBlinking(adaptor)
	case BUZZ:
		devices, work = experiments.Buzz(adaptor)
	case LED_AND_BUZZ:
		devices, work = experiments.TwoLedsBlinkingAndBuzz(adaptor)
	default:
		flag.Usage()
		os.Exit(0)
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{adaptor},
		devices,
		work,
	)
	println("INFO: remember to run using sudo !")
	master := gobot.NewMaster()
	server := api.NewAPI(master)
	server.Port = "3000"
	server.Start()
	master.AddRobot(robot)
	master.Start()
}
