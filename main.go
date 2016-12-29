package main

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/raspi"
	"gobot.io/x/gobot/api"
	"github.com/miccab/gobot-go/led"
	"flag"
)

func main() {
	master := gobot.NewMaster()
	server := api.NewAPI(master)
	server.Port = "3000"
	server.Start()
	adaptor := raspi.NewAdaptor()

	var experiment string
	flag.StringVar(&experiment, "experiment", "UNkNOWN", "type of experiment")
	flag.Parse()

	var devices []gobot.Device
	var work func()

	switch experiment {
	case "1led":
		devices, work = led.GetLedAndWork(adaptor)
	case "2leds":
		devices, work = led.Get2LedsAndWork(adaptor)
	default:
		panic("unknown experiment")
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{adaptor},
		devices,
		work,
	)
	println("INFO: remember to run using sudo !")
	master.AddRobot(robot)
	master.Start()
}
