package main

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/raspi"
	"gobot.io/x/gobot/api"
	"github.com/miccab/gobot-go/led"
)

func main() {
	master := gobot.NewMaster()
	server := api.NewAPI(master)
	server.Port = "3000"
	server.Start()
	adaptor := raspi.NewAdaptor()

	devices, work := led.GetDevicesAndWork(adaptor)

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{adaptor},
		devices,
		work,
	)
	println("INFO: remember to run using sudo !")
	master.AddRobot(robot)
	master.Start()
}
