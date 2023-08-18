// main
package main

import (
	"fmt"
	"log"
	"net/url"
	_ "strings"

	"github.com/gocolly/colly"
)

func f1() {

	baseURL := "https://www.jobkorea.co.kr"
	searchKeyword := "인도네시아"
	searchURL := baseURL + "/Search/?stext=" + searchKeyword

	c := colly.NewCollector(
		colly.AllowedDomains("www.jobkorea.co.kr"),
	)

	var sResultNodes, sResultPagination string

	c.OnHTML("div.list-default", func(e *colly.HTMLElement) {
		e.ForEach(".list-post", func(i int, item *colly.HTMLElement) {
			sResultNodes += item.Text
			sResultNodes += "\n"
		})
	})

	c.OnHTML("div.tplPagination.newVer.wide", func(e *colly.HTMLElement) {
		sResultPagination = e.Text
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c.Visit(searchURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Nodes:")
	fmt.Println(sResultNodes)
	fmt.Println("Pagination:")
	fmt.Println(sResultPagination)
}

func f2() {

	baseURL := "https://www.jobkorea.co.kr"
	searchKeyword := "인도네시아"
	searchURL := baseURL + "/Search/?stext=" + searchKeyword

	c := colly.NewCollector(
		colly.AllowedDomains("www.jobkorea.co.kr"),
	)

	var sResult string

	c.OnHTML("div.list-default", func(e *colly.HTMLElement) {
		e.ForEach(".list-post", func(i int, item *colly.HTMLElement) {
			surl := item.Attr("data-gavirturl")
			sinfo := item.Attr("data-gainfo")

			node1 := item.DOM.Find(".post-list-corp")
			node2 := item.DOM.Find(".post-list-info")

			sResult += surl + "\r\n" + sinfo + "\r\n" + node1.Text() + "\r\n" + node2.Text() + "\r\n\r\n"
		})
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c.Visit(searchURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sResult)
}

func f3() {

	baseURL := "https://www.jobkorea.co.kr"
	searchKeyword := "인도네시아"
	searchURL := baseURL + "/Search/?stext=" + url.QueryEscape(searchKeyword)

	c := colly.NewCollector() // colly.New() 대신 colly.NewCollector() 사용

	var sResult string

	c.OnHTML("div.tplPagination.newVer.wide", func(e *colly.HTMLElement) {
		e.ForEach("li a[href]", func(_ int, item *colly.HTMLElement) {
			href := item.Attr("href")
			sResult += href + "\r\n"
		})
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c.Visit(searchURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sResult)
}

func main() {
	f1()
	//f2()
	//f3()
}
