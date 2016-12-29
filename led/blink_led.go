package led

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"time"
)

func GetDevicesAndWork(writer gpio.DigitalWriter) (devices []gobot.Device, work func())  {
	led := gpio.NewLedDriver(writer, "12")
	work = func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}
	devices = []gobot.Device{led}
	return
}
