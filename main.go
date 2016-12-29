package main

import (
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot"
	"time"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
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
	println("INFO: run using sudo !")
	robot.Start()
}
