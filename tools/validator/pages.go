package validator

import (
	"github.com/institute-atri/glogger"
	"github.com/institute-atri/gnet"
)

type pages struct {
	Url   string
	Proxy string

	TOR         bool
	RandomAgent bool
}

// The `NewPages` function is responsible for instantiating a set of other functions with the purpose of checking the existence and accessibility of common pages on websites.
func NewPages(url, proxy string, tor, randomAgent bool) *pages {
	return &pages{Url: url, Proxy: proxy, TOR: tor, RandomAgent: randomAgent}
}

func (p *pages) Admin() Interesting {
	var request = gnet.NewHttp()

	request.SetURL(p.Url + "wp-admin/")

	if p.TOR {
		request.OnTor()
	} else if p.Proxy != "" {
		request.SetProxy(p.Proxy)
	}

	if p.RandomAgent {
		request.OnRandomUserAgent()
	}

	response, err := request.Do()

	if err != nil {
		glogger.Fatal(err)
	}

	var entity = Interesting{BRaw: response.BRaw, FoundBy: "Direct Access"}

	if response.StatusCode == 200 || response.StatusCode == 403 {
		entity.Confidence = 100
	}

	return entity
}

func (p *pages) Robots() Interesting {
	var request = gnet.NewHttp()

	request.SetURL(p.Url + "robots.txt")

	if p.TOR {
		request.OnTor()
	} else if p.Proxy != "" {
		request.SetProxy(p.Proxy)
	}

	if p.RandomAgent {
		request.OnRandomUserAgent()
	}

	response, err := request.Do()

	if err != nil {
		glogger.Fatal(err)
	}

	var entity = Interesting{BRaw: response.BRaw, FoundBy: "Direct Access"}

	if response.StatusCode == 200 || response.StatusCode == 403 {
		entity.Confidence = 100
	}

	return entity
}

func (p *pages) Sitemap() Interesting {
	var request = gnet.NewHttp()

	request.SetURL(p.Url + "sitemap.xml")

	if p.TOR {
		request.OnTor()
	} else if p.Proxy != "" {
		request.SetProxy(p.Proxy)
	}

	if p.RandomAgent {
		request.OnRandomUserAgent()
	}

	response, err := request.Do()

	if err != nil {
		glogger.Fatal(err)
	}

	var entity = Interesting{BRaw: response.BRaw, FoundBy: "Direct Access"}

	if response.StatusCode == 200 || response.StatusCode == 403 {
		entity.Confidence = 100
	}

	return entity
}

func (p *pages) Readme() Interesting {
	var request = gnet.NewHttp()

	request.SetURL(p.Url + "readme.html")

	if p.TOR {
		request.OnTor()
	} else if p.Proxy != "" {
		request.SetProxy(p.Proxy)
	}

	if p.RandomAgent {
		request.OnRandomUserAgent()
	}

	response, err := request.Do()

	if err != nil {
		glogger.Fatal(err)
	}

	var entity = Interesting{BRaw: response.BRaw, FoundBy: "Direct Access"}

	if response.StatusCode == 200 || response.StatusCode == 403 {
		entity.Confidence = 100
	}

	return entity
}
