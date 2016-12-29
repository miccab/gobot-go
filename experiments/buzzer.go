package experiments

import (
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot"
	"time"
)

func GetBuzzerAndWork(writer gpio.DigitalWriter) (devices []gobot.Device, work func())  {
	buzzer := gpio.NewBuzzerDriver(writer, "11")
	work = func() {
		buzzer.On()
		time.Sleep(1*time.Second)
		buzzer.Off()
	}
	devices = []gobot.Device{buzzer}
	return
}

