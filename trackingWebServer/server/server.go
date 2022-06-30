package server

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const filePath = "/tmp/ok"

type Response struct {
	Body       string
	StatusCode int
	Request    Request
}

type Request struct {
	Filepath string
}

// RequestResponse encapsulates a request and response channel
type RequestResponse struct {
	Request    Request
	ResponseCh chan Response
}

// Server handles requests
type Server struct {
	ReqResCh chan RequestResponse
	InfoLog  *log.Logger
	ErrLog   *log.Logger
}

// HandlePing sends a request on a request channel and waits for response
func (p *Server) HandlePing(w http.ResponseWriter, r *http.Request) {
	resCh := make(chan Response)
	p.ReqResCh <- RequestResponse{ResponseCh: resCh, Request: Request{Filepath: filePath}}
	response := <-resCh

	w.WriteHeader(response.StatusCode)

	if response.StatusCode == http.StatusOK {
		io.WriteString(w, response.Body)
	}

}

// HandleImg responds with a constant response
func (p *Server) HandleImg(w http.ResponseWriter, r *http.Request) {
	// log request
	log := []string{r.RequestURI, r.Host}

	if reqHeadersBytes, err := json.Marshal(r.Header); err != nil {
		p.ErrLog.Printf("could not Marshal Req Headers, got err : %s", err)
	} else {
		log = append(log, string(reqHeadersBytes))
	}
	p.InfoLog.Println(log)

	// return response
	base64pixelImg := "R0lGODlhAQABAIAAAP///wAAACwAAAAAAQABAAACAkQBADs="
	w.Header().Set("Content-Type", "image/gif")
	decoded, err := base64.StdEncoding.DecodeString(base64pixelImg)
	if err != nil {
		p.ErrLog.Printf("could not perform decodeString, got err : %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	io.WriteString(w, string(decoded))
}
