package main

import "fmt"

type PowerDrawer interface {
	Draw()
	GetName() string
}

type Socket struct {
	Name string
}

func (this Socket) Power(device PowerDrawer) {
	fmt.Println(this.Name, "is powering up", device.GetName())
	device.Draw()
}

type Mixer struct {
	Name string
}

func (this Mixer) Draw() {
	fmt.Println(this.Name, "is drawing power")
}
func (this Mixer) GetName() string {
	return this.Name
}

type Kettle struct {
	Name string
}

func (this Kettle) Draw() {
	fmt.Println(this.Name, "is drawing power")
}

func (this Kettle) GetName() string {
	return this.Name
}

func main() {
	kettle, socket, mixer := Kettle{"Kettle"}, Socket{"Socket"}, Mixer{"Mixer"}

	socket.Power(kettle)
	socket.Power(mixer)

}
