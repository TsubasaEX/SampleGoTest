package main

import (
	"fmt"
	"net/http"

	mygreetingV2 "github.com/TsubasaEX/SampleGoModPkg/v2/greeting"
	myshowV2 "github.com/TsubasaEX/SampleGoModPkg/v2/testshow/show"

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

	mygreetingV2.SayWithHelloV2("Zach")
	myshowV2.ShowWithHelloV2("123")

	result := calc.AddTwo(10)
	fmt.Println(result)

	http.ListenAndServe(":1123", appRouter())
}
