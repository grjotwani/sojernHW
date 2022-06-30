package server

import (
	"context"
	"net/http"
	"os"
)

// Process listens on a requests channel and responds to requests
func Process(requestsCh chan RequestResponse, cancelCtx context.Context) {
	var statusCode int
	var body string
	for {
		select {
		case req := <-requestsCh:
			info, err := os.Stat(req.Request.Filepath)
			statusCode = http.StatusServiceUnavailable
			if err == nil && !info.IsDir() {
				body = "OK"
				statusCode = http.StatusOK
			}

			req.ResponseCh <- Response{
				StatusCode: statusCode,
				Body:       body,
				Request:    req.Request,
			}
		case <-cancelCtx.Done():
			return
		}
	}
}
