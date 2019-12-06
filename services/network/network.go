package network

import (
	"github.com/go-resty/resty"
	"log"
)

var client *resty.Client

func init() {
	client = resty.New()
}

func GetRequest(url string, headers map[string]string, queryParams map[string]string, callback func (res *resty.Response, err error)) {
	go func() { res, err := client.R().
		SetQueryParams(queryParams).
		EnableTrace().
		SetHeaders(headers).
		Get(url)
		log.Printf("Get request Trace Info: %v", res.Request.TraceInfo())
		if res != nil {
			log.Printf("Response: %s", res.String())
		} 
		if err != nil {
			log.Printf("Error: %s", err.Error())
		}
		callback(res, err)
	}()
}

func PostRequest(url string, headers map[string]string, body interface{}, callback func(res *resty.Response, err error)) {
	go func() { res, err := client.R(). 
		SetBody(body). 
		EnableTrace(). 
		SetHeaders(headers). 
		Post(url)
		log.Printf("Post request Trace Info: %v", res.Request.TraceInfo())
		if res != nil {
			log.Printf("Response: %s", res.String())
		} 
		if err != nil {
			log.Printf("Error: %s", err.Error())
		}
		callback(res, err)
	}()
}