package collector

import (
	"fmt"
	"sync"

	"github.com/gocolly/colly"
)

func search(maxDepth, parallelism int, userAgent, searchString, site string, async bool, wg *sync.WaitGroup) {
	c := colly.NewCollector(
		colly.MaxDepth(maxDepth),
		colly.Async(async),
	)
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: parallelism})
	c.UserAgent = userAgent
	c.OnHTML("a[href]", func(element *colly.HTMLElement) {
		fmt.Printf("a:%#v \n", element.Attr("href"))
		//if ok := strings.Contains(element.Attr("content"), searchString); ok {
		//	// todo
		//}
		element.Request.Visit(element.Attr("href"))
	})
	c.Visit(site)
	wg.Done()
}

func SearchWeb() {
	guideWebsites := []string{
		"https://www.2345.com/",
		"https://www.hao123.com/",
		"https://hao.360.cn/",
		"https://www.114la.com/",
	}
	searchString := "网址导航"
	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36 Edge/16.16299"
	maxDepth := 2
	async := false
	parallelism := 2
	wg := &sync.WaitGroup{}
	wg.Add(len(guideWebsites))
	//wg.Add(1)

	for _, site := range guideWebsites {
		go search(maxDepth, parallelism, userAgent, searchString, site, async, wg)
		//search(maxDepth,parallelism,userAgent,searchString,site,async,wg)
	}
	wg.Wait()
}
