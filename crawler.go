

package main

import (
"encoding/json"
"fmt"
"net/url"
"time"

"github.com/gocolly/colly"
"github.com/gocolly/colly/debug"
)

type T struct {
	Result []struct {
		SrcShortUrl string
		TitleUrl    string
		ImgUrl      string
		VideoUrl    string
		Duration    string
		Restitle    string
		Site_logo   string
		Nsclick_v   string
	}
}

func main() {

	t3 := time.Now()
	keyword := url.QueryEscape("台风")
	url1 := fmt.Sprintf("http://app.video.baidu.com/app?word=%s&pn=0&rn=50&order=1", keyword)

	// Instantiate default collector
	c := colly.NewCollector()

	// Attach a debugger to the collector
	c.SetDebugger(&debug.LogDebugger{})

	// Limit the number of threads started by colly to two
	// when visiting links which domains' matches "*httpbin.*" glob
	c.Limit(&colly.LimitRule{

		Parallelism: 10,
		//Delay:      5 * time.Second,
	})

	// Start scraping in four threads on https://httpbin.org/delay/2
	for i := 0; i < 1000; i++ {
		go c.Visit(fmt.Sprintf("http://app.video.baidu.com/app?word=%s&pn=%d&rn=50&order=1", keyword, i*50))
	}
	// Start scraping on https://httpbin.org/delay/2
	c.Visit(url1)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnResponse(func(r *colly.Response) {

		var t T
		json.Unmarshal(r.Body, &t)

		fmt.Println(t)
	})
	// Wait until threads are finished
	c.Wait()
	delt := time.Since(t3)
	fmt.Println("爬虫结束,总共耗时: ", delt)
}

