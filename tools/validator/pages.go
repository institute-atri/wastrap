package validator

import (
	"github.com/institute-atri/ghttp"
	"github.com/institute-atri/glogger"
)

type Pages struct {
	Url   string
	Proxy string

	TOR         bool
	RandomAgent bool
}

// NewPages is a function that creates a new instance of the Pages struct with the specified parameters.
// It takes in the url (string) of the page, the proxy (string) to be used, a flag indicating whether to use TOR (bool),
// and a flag indicating whether to use a random User-Agent (bool).
// It returns a pointer to the new Pages instance.
func NewPages(url, proxy string, tor, randomAgent bool) *Pages {
	return &Pages{Url: url, Proxy: proxy, TOR: tor, RandomAgent: randomAgent}
}

func (p *Pages) Admin() Interesting {
	request := ghttp.NewHttp()

	err := request.SetURL(p.Url + "wp-admin/")
	glogger.ErrorHandling(err)

	if p.TOR {
		err := request.OnTor()
		glogger.ErrorHandling(err)
	} else if p.Proxy != "" {
		err := request.SetProxy(p.Proxy)
		glogger.ErrorHandling(err)
	}

	if p.RandomAgent {
		request.OnRandomUserAgent()
	}

	response, err := request.Do()
	glogger.ErrorHandling(err)

	entity := Interesting{
		BRaw:    response.BRaw,
		FoundBy: "Direct Access",
	}

	if response.StatusCode == 200 || response.StatusCode == 403 {
		entity.Confidence = 100
	}

	return entity
}

// Robots is a method of the Pages struct that sends an HTTP request to retrieve the "robots.txt" file from the specified URL.
// It uses the provided URL and appends "robots.txt" to it to form the complete URL for the request.
// If the TOR flag is set to true, it enables the TOR network for the request.
// If a proxy is specified, it sets the proxy for the request.
// If the RandomAgent flag is set to true, it sets a random User-Agent for the request.
// It then sends the HTTP request and retrieves the response.
// The response body is stored in the `BRaw` field of the returned Interesting struct.
// The `FoundBy` field is set to "Direct Access" to indicate that the file was directly accessed.
// If the response status code is 200 or 403, the `Confidence` field is set to 100 to indicate a successful retrieval.
// The method returns the Interesting struct.
// Example usage:
//
//	p := &Pages{
//	    Url:         "https://example.com/",
//	    Proxy:       "http://proxy.example.com",
//	    TOR:         true,
//	    RandomAgent: true,
//	}
//	result := p.Robots()
//	fmt.Println(result.BRaw, result.FoundBy, result.Confidence)
func (p *Pages) Robots() Interesting {
	request := ghttp.NewHttp()

	err := request.SetURL(p.Url + "robots.txt")
	glogger.ErrorHandling(err)

	if p.TOR {
		err := request.OnTor()
		glogger.ErrorHandling(err)
	} else if p.Proxy != "" {
		err := request.SetProxy(p.Proxy)
		glogger.ErrorHandling(err)
	}

	if p.RandomAgent {
		request.OnRandomUserAgent()
	}

	response, err := request.Do()
	glogger.ErrorHandling(err)

	entity := Interesting{
		BRaw:    response.BRaw,
		FoundBy: "Direct Access",
	}

	if response.StatusCode == 200 || response.StatusCode == 403 {
		entity.Confidence = 100
	}

	return entity
}

// Sitemap is a method of the Pages struct that sends an HTTP request to retrieve the "sitemap.xml" file from the specified URL.
// It uses the provided URL and appends "sitemap.xml" to it to form the complete URL for the request.
// If the TOR flag is set to true, it enables the TOR network for the request.
// If a proxy is specified, it sets the proxy for the request.
// If the RandomAgent flag is set to true, it sets a random User-Agent for the request.
// It then sends the HTTP request and retrieves the response.
// The response body is stored in the `BRaw` field of the returned Interesting struct.
// The `FoundBy` field is set to "Direct Access" to indicate that the file was directly accessed.
// If the response status code is 200 or 403, the `Confidence` field is set to 100 to indicate a successful retrieval.
// The method returns the Interesting struct.
//
// Example usage:
//
//	p := &Pages{
//	    Url:         "https://example.com/",
//	    Proxy:       "http://proxy.example.com",
//	    TOR:         true,
//	    RandomAgent: true,
//	}
//	result := p.Sitemap()
//	fmt.Println(result.BRaw, result.FoundBy, result.Confidence)
func (p *Pages) Sitemap() Interesting {
	request := ghttp.NewHttp()

	err := request.SetURL(p.Url + "sitemap.xml")
	glogger.ErrorHandling(err)

	if p.TOR {
		err := request.OnTor()
		glogger.ErrorHandling(err)
	} else if p.Proxy != "" {
		err := request.SetProxy(p.Proxy)
		glogger.ErrorHandling(err)
	}

	if p.RandomAgent {
		request.OnRandomUserAgent()
	}

	response, err := request.Do()
	glogger.ErrorHandling(err)

	entity := Interesting{
		BRaw:    response.BRaw,
		FoundBy: "Direct Access",
	}

	if response.StatusCode == 200 || response.StatusCode == 403 {
		entity.Confidence = 100
	}

	return entity
}

// Readme is a method of the Pages struct that sends an HTTP request to retrieve the "readme.html" file from the specified URL.
// It uses the provided URL and appends "readme.html" to it to form the complete URL for the request.
// If the TOR flag is set to true, it enables the TOR network for the request.
// If a proxy is specified, it sets the proxy for the request.
// If the RandomAgent flag is set to true, it sets a random User-Agent for the request.
// It then sends the HTTP request and retrieves the response.
// The response body is stored in the `BRaw` field of the returned Interesting struct.
// The `FoundBy` field is set to "Direct Access" to indicate that the file was directly accessed.
// If the response status code is 200 or 403, the `Confidence` field is set to 100 to indicate a successful retrieval.
// The method returns the Interesting struct.
func (p *Pages) Readme() Interesting {
	request := ghttp.NewHttp()

	err := request.SetURL(p.Url + "readme.html")
	glogger.ErrorHandling(err)

	if p.TOR {
		err := request.OnTor()
		glogger.ErrorHandling(err)
	} else if p.Proxy != "" {
		err := request.SetProxy(p.Proxy)
		glogger.ErrorHandling(err)
	}

	if p.RandomAgent {
		request.OnRandomUserAgent()
	}

	response, err := request.Do()
	glogger.ErrorHandling(err)

	entity := Interesting{
		BRaw:    response.BRaw,
		FoundBy: "Direct Access",
	}

	if response.StatusCode == 200 || response.StatusCode == 403 {
		entity.Confidence = 100
	}

	return entity
}
