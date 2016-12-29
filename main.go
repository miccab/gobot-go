package main

import (
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot"
	"time"
	"gobot.io/x/gobot/platforms/raspi"
	"gobot.io/x/gobot/api"
)

func main() {
	master := gobot.NewMaster()
	server := api.NewAPI(master)
	server.Port = "3000"
	server.Start()
	adaptor := raspi.NewAdaptor()
	led := gpio.NewLedDriver(adaptor, "12")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{adaptor},
		[]gobot.Device{led},
		work,
	)
	println("INFO: remember to run using sudo !")
	master.AddRobot(robot)
	master.Start()
}
