package main

import (
	"context"
	"flag"
	"github.com/grjotwani/sojernHW/trackingWebServer/server"
	"log"
	"net/http"
	"os"
)

var requestBufferSize = flag.Int("requestBufferSize", 100, "buffer size for requests")

func main() {
	flag.Parse()
	infoLog := log.New(os.Stdout, "", 0)
	errLog := log.New(os.Stderr, "", 0)
	requestsCh := make(chan server.RequestResponse, *requestBufferSize)

	cancelCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go server.Process(requestsCh, cancelCtx)

	server := &server.Server{ReqResCh: requestsCh, InfoLog: infoLog, ErrLog: errLog}

	http.HandleFunc("/ping", server.HandlePing)
	http.HandleFunc("/img", server.HandleImg)

	http.ListenAndServe(":8090", nil)
}
