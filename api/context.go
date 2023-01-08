package api

import (
	"log"
	"net/http"
)

type ApiContext struct {
	Request   *http.Request
	method    string
	uri       string
	Endpoints string
	Operation string
	Username  string
	Password  string
	Tenant    string
	Token     string
}

func New(endpoint string, tenantid string, option string) *ApiContext {
	context := &ApiContext{}
	context.SetUri(endpoint, tenantid, option)
	return context
}

func (ctx *ApiContext) Run() (*http.Response, error) {
	ctx.Request.Header.Add("Accept", "application/json")
	ctx.Request.Header.Add("X-Auth-Token", ctx.Token)
	client := &http.Client{}
	response, err := client.Do(ctx.Request)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}

func (ctx *ApiContext) SetQueryOptions(query map[string]string) {
	for key, value := range query {
		q := ctx.Request.URL.Query()
		q.Add(key, value)
		ctx.Request.URL.RawQuery = q.Encode()
	}
}
