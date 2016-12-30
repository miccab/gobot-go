package experiments

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"time"
)

func LedBlinking(writer gpio.DigitalWriter) (devices []gobot.Device, work func())  {
	led := gpio.NewLedDriver(writer, "12")
	work = func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}
	devices = []gobot.Device{led}
	return
}

func TwoLedsBlinking(writer gpio.DigitalWriter) (devices []gobot.Device, work func())  {
	led1 := gpio.NewLedDriver(writer, "12")
	led2 := gpio.NewLedDriver(writer, "16")
	work = func() {
		led1.On()
		led2.Off()
		gobot.Every(1*time.Second, func() {
			led1.Toggle()
			led2.Toggle()
		})
	}
	devices = []gobot.Device{led1,led2}
	return
}


func TwoLedsBlinkingAndBuzz(writer gpio.DigitalWriter) (devices []gobot.Device, work func())  {
	ledGreen := gpio.NewLedDriver(writer, "12")
	ledRed := gpio.NewLedDriver(writer, "16")
	buzzer := gpio.NewBuzzerDriver(writer, "11")
	work = func() {
		ledRed.On()
		ledGreen.Off()
		buzzer.On()
		gobot.Every(1*time.Second, func() {
			ledGreen.Toggle()
			ledRed.Toggle()
			buzzer.Toggle()
		})
	}
	devices = []gobot.Device{ledGreen, ledRed}
	return
}
