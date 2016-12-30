package experiments

import (
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot"
)

func GetBuzzerAndWork(writer gpio.DigitalWriter) (devices []gobot.Device, work func())  {
	buzzer := gpio.NewBuzzerDriver(writer, "11")
	work = func() {
		playStarWars(buzzer)
	}
	devices = []gobot.Device{buzzer}
	return
}

func playStarWars(buzzer * gpio.BuzzerDriver) {
	buzzer.Tone(gpio.A4, 500)
	buzzer.Tone(gpio.A4, 500)
	buzzer.Tone(gpio.F4, 350)
	buzzer.Tone(gpio.C5, 150)

	buzzer.Tone(gpio.A4, 500)
	buzzer.Tone(gpio.F4, 350)
	buzzer.Tone(gpio.C5, 150)
	buzzer.Tone(gpio.A4, 1000)
	buzzer.Tone(gpio.E5, 500)
}

