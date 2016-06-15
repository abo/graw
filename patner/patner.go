package patner

import (
	"net/url"
	"os"

	"github.com/abo/patnsvc"
	httppatn "github.com/abo/patnsvc/http"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptrans "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

const defaultPatnSvcURL string = "https://api.lograw.com/v1/pattern"

// a pattern generator, client to pattern service
type patner struct {
	context.Context
	log.Logger
	generate endpoint.Endpoint
}

// Generate generate pattern by specific sample
func (c patner) Generate(lines []patnsvc.Line) ([]patnsvc.Pattern, error) {
	request := patnsvc.Request{Lines: lines}
	reply, err := c.generate(c.Context, request)
	if err != nil {
		c.Logger.Log(err)
		return make([]patnsvc.Pattern, 0), err
	}

	response := reply.(patnsvc.Response)
	return response.Patterns, nil
}

// NewPatner new a pattern generator
func NewPatner() patnsvc.Service {
	logger := log.NewLogfmtLogger(os.Stdout)
	p, _ := url.Parse(defaultPatnSvcURL)
	endpoint := httptrans.NewClient("POST", p, httppatn.RequestEncoder, httppatn.ResponseDecoder).Endpoint()
	ctx := context.Background()
	return patner{
		Context:  ctx,
		Logger:   logger,
		generate: endpoint,
	}
}
