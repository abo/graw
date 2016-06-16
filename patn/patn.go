package patn

import (
	"net/url"
	"os"
	"regexp"

	"github.com/abo/patnsvc"
	httppatn "github.com/abo/patnsvc/http"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptrans "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

const defaultPatnSvcURL string = "https://api.lograw.com/v1/pattern"

// a pattern generator, client to pattern service
type patnClient struct {
	context.Context
	log.Logger
	generate endpoint.Endpoint
}

// Generate rpc to patn server
func (c patnClient) Generate(lines []patnsvc.Line) ([]patnsvc.Pattern, error) {
	request := patnsvc.Request{Lines: lines}
	reply, err := c.generate(c.Context, request)
	if err != nil {
		c.Logger.Log(err)
		return make([]patnsvc.Pattern, 0), err
	}

	response := reply.(patnsvc.Response)
	return response.Patterns, nil
}

func newPatnClient(svcURL string) patnsvc.Service {
	logger := log.NewLogfmtLogger(os.Stdout)
	p, _ := url.Parse(svcURL)
	endpoint := httptrans.NewClient("POST", p, httppatn.RequestEncoder, httppatn.ResponseDecoder).Endpoint()
	ctx := context.Background()
	return patnClient{
		Context:  ctx,
		Logger:   logger,
		generate: endpoint,
	}
}

//Patner a pattern generator
type Patner struct {
	patnClient patnsvc.Service
}

// NewPatner create a new pattern generator
func NewPatner() Patner {
	return Patner{
		patnClient: newPatnClient(defaultPatnSvcURL),
	}
}

// Generate generate pattern to extract targets from raw
func (p *Patner) Generate(raw string, targets []string) ([]string, error) {
	lines := make([]patnsvc.Line, len(targets))
	for i, target := range targets {
		lines[i] = patnsvc.Line{
			Raw:      raw,
			Expected: target,
		}
	}
	ps, err := p.patnClient.Generate(lines)
	if err != nil {
		return nil, err
	}

	ret := make([]string, len(ps))
	for i, p := range ps {
		ret[i] = p.Expr
	}
	return ret, nil
}

// Extractor a field extractor
type Extractor struct {
	regexps []*regexp.Regexp
}

// NewExtractor create a new extractor
func NewExtractor(exprs []string) (Extractor, error) {
	regexps := make([]*regexp.Regexp, len(exprs))
	for i, expr := range exprs {
		re, err := regexp.Compile(expr)
		if err != nil {
			return Extractor{}, err
		}
		regexps[i] = re
	}
	return Extractor{
		regexps: regexps,
	}, nil
}

//Extract extract fields from raw
func (e *Extractor) Extract(raw string) []string {
	ret := make([]string, len(e.regexps))
	for i, re := range e.regexps {
		matches := re.FindStringSubmatch(raw)
		if len(matches) >= 2 {
			ret[i] = matches[1]
		}
	}
	return ret
}
