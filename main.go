package main

import
(
	"fmt"
	"flag"
)
var (
	runTime = flag.Int("runtime", 60, "Defines running time in second")
	maxConcurrecy = flag.Int("max-concurrency", 0, "Defines how many process will be used")
	parameters = flag.String("parameters", "", "Defines parameter for the request")
)

func main(){

    flag.Parse()

    fmt.Println(*runTime)
    fmt.Println(*maxConcurrecy)
    fmt.Println(*parameters)
}