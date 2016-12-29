package experiments

import (
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot"
)

func GetBuzzerAndWork(writer gpio.DigitalWriter) (devices []gobot.Device, work func())  {
	buzzer := gpio.NewBuzzerDriver(writer, "11")
	work = func() {
		buzzer.Tone(gpio.C5, gpio.Whole)
		buzzer.Tone(gpio.D5, gpio.Half)
		buzzer.Tone(gpio.E5, gpio.Quarter)
		buzzer.Tone(gpio.F5, gpio.Eighth)
	}
	devices = []gobot.Device{buzzer}
	return
}

