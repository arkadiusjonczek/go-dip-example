package main

func main() {
	lamp := NewLamp()
	coffeeMaschine := NewCoffeeMaschine()

	switchWithLampConnected := NewSwitch(lamp)
	switchWithCoffeeMaschineConnected := NewSwitch(coffeeMaschine)

	switchWithLampConnected.On()
	switchWithLampConnected.Off()

	switchWithCoffeeMaschineConnected.On()
	switchWithCoffeeMaschineConnected.Off()
}
