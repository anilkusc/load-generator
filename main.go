package main

import
(
	"fmt"
	"flag"
	"strings"
	"log"
	"os"
)
var (
	url = flag.String("url", "", "(Required)Defines url of the service")
	method = flag.String("method", "GET", "Defines method of the request")
	runTime = flag.Int("runtime", 60, "Defines running time in second")
	requestCount = flag.Int("request-count", 100, "Defines count of requests")
	maxConcurrecy = flag.Int("max-concurrency", 1, "Defines how many process will be used")
	parameters = flag.String("parameters", "", "Defines parameter for the request")
	getHeaders = flag.String("headers","", "Defines headers of the requests.You should seperate headers with comma(,)")
	sem = make(chan int, *maxConcurrecy)
) 

func main(){
	flag.Parse()
	if *url == "" {
		log.Println("Required flags are not set")
		flag.PrintDefaults()
		os.Exit(1)
	}
	for i := 0; i < *requestCount+1; i++ {
		sem <- 0
		go fmt.Println(LoadGenerator())
		<-sem
	}
			
}

func LoadGenerator() (int,string){
	headers := strings.Split(*getHeaders,",")
	return MakeRequest(*url,*method,headers)
}