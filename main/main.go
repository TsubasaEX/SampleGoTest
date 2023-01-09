package main

import (
	mygreeting "github.com/TsubasaEX/SampleGoModPkg/greeting"
	myshow "github.com/TsubasaEX/SampleGoModPkg/testshow/show"
	"github.com/TsubasaEX/SampleGoModTest/greeting"
)

func main() {
	greeting.Say("Hi World")
	greeting.SayWithColor("Hi World with Color")
	mygreeting.SayWithHello("Zach")
	myshow.ShowWithHello("123")
	// f.Println("123")
}
