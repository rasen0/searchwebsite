package collector

import (
	"bytes"
	"io/ioutil"
	"strings"
	"sync"
	"time"
	"unicode/utf8"

	"com.rasen/common/structlog"
	"com.rasen/scraper/config"
	"com.rasen/scraper/database"
	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func saveDataMap(hrefMap *sync.Map, client *mongo.Client) {
	collection := client.Database("collectData").Collection("webHref")
	database.SaveDataMapToMgo(hrefMap, collection)
}

func loopRun(hrefMap *sync.Map, client *mongo.Client, succInterval time.Duration, failInterval time.Duration,
	fn func(hrefMap *sync.Map, client *mongo.Client)) {
	defer func() {
		if err := recover(); err != nil {
			structlog.Logger.WithFields(logrus.Fields{"err": err}).Error("loopRun function fail")
			//fmt.Fprintln(os.Stderr,"loopRun function fail")
			loopRun(hrefMap, client, succInterval, failInterval, fn)
		}
	}()
	time.Sleep(failInterval)
	for {
		fn(hrefMap, client)
		time.Sleep(succInterval)
	}
}

func search(maxDepth, parallelism int, userAgent, searchString, site string, async bool, wg *sync.WaitGroup, hrefMap *sync.Map) {
	c := colly.NewCollector(
		colly.MaxDepth(maxDepth),
		colly.Async(async),
	)
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: parallelism})
	c.UserAgent = userAgent
	c.OnHTML("a[href]", func(element *colly.HTMLElement) {
		href := element.Attr("href")
		if !strings.Contains(href, "http") {
			return
		}
		title := element.Attr("title")
		if title == "" {
			lastDot := strings.LastIndex(href, ".")
			lastSlash := strings.LastIndex(href, "/")
			if lastDot < lastSlash+1 {
				return
			}
			title = href[lastSlash+1 : lastDot]
		}
		if !utf8.ValidString(title) {
			reader := transform.NewReader(bytes.NewReader([]byte(title)), simplifiedchinese.GBK.NewDecoder())
			b, err := ioutil.ReadAll(reader)
			if err != nil {
				structlog.Logger.WithFields(logrus.Fields{"err": err}).Error("convert to utf8 fail")
			} else {
				title = string(b)
			}
		}
		hrefMap.Store(title, href)
		element.Request.Visit(element.Attr("href"))
		//fmt.Printf("a:%#v, title:%v \n", href, title)
	})
	c.Visit(site)
	wg.Done()
}

func SearchWeb(config *config.DefaultConfig, client *mongo.Client) {
	//guideWebsites := []string{
	//  "www.sohu.com/",
	//	"https://www.qq.com/",
	//	"https://www.163.com/",
	//	"https://www.csdn.net/",
	//	"https://www.cnblogs.com/",
	//	"https://www.sina.com.cn/",
	//	"http://www.people.com.cn/"
	//}
	//userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36 Edge/16.16299"
	hrefMap := &sync.Map{}

	searchString := "网址导航"
	maxDepth := 2
	async := false
	parallelism := 2
	wg := &sync.WaitGroup{}
	wg.Add(len(config.GuideWeb))
	//wg.Add(1)

	// 把收集到的数据存入mongodb
	go loopRun(hrefMap, client, 500*time.Millisecond, 1000*time.Millisecond, saveDataMap)

	for _, site := range config.GuideWeb {
		go search(maxDepth, parallelism, config.UserAgent, searchString, site, async, wg, hrefMap)
		//search(maxDepth,parallelism,userAgent,searchString,site,async,wg)
	}
	wg.Wait()

}
