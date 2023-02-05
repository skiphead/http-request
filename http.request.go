package request

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"
)

type Params struct {
	Ctx       context.Context
	Timeout   int
	Method    string
	URL       string `json:"url"`
	Body      io.Reader
	Transport *http.Transport
}

func (params *Params) Request() []byte {

	ctx, cancel := context.WithTimeout(params.Ctx, time.Duration(params.Timeout)*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, params.Method, params.URL, params.Body)
	defer ctx.Done()
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{Transport: params.Transport}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	byteBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {

		}
	}(res.Body)

	return byteBody
}
