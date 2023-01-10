package main

import (
	"fmt"
	"net/http"

	mygreeting "github.com/TsubasaEX/SampleGoModPkg/greeting"
	myshow "github.com/TsubasaEX/SampleGoModPkg/testshow/show"
	"github.com/TsubasaEX/SampleGoModTest/calc"
	"github.com/TsubasaEX/SampleGoModTest/greeting"
)

func handleSlothfulMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Stay slothful!"}`))
}

func appRouter() http.Handler {
	rt := http.NewServeMux()
	rt.HandleFunc("/sloth", handleSlothfulMessage)
	return rt
}

func main() {
	greeting.Say("Hi World")
	greeting.SayWithColor("Hi World with Color")
	mygreeting.SayWithHello("Zach")
	myshow.ShowWithHello("123")

	result := calc.AddTwo(10)
	fmt.Println(result)

	http.ListenAndServe(":1123", appRouter())
}
