package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main(){
	var url = "https://www.cnyes.com/usstock"
	var news []string
	var urls []string

	c := colly.NewCollector()

	c.OnHTML("div.news-list  ul  li ",func(e *colly.HTMLElement){
		// fmt.Print(e.Text,e.ChildAttr("a","href"))
		linksStr:=e.Text
		if len(linksStr)!=0{
			linksStr=linksStr[5:]
			news=append(news,linksStr)
		}

		if len(e.ChildAttr("a","href"))!=0{
			urls=append(urls,e.ChildAttr("a","href"))
		}
	})

	c.Visit(url)

	helper(news,urls)
}

func helper(news []string,urls []string){
	if len(news)>0 && len(news)==len(urls){
		for i:=0;i<len(news);i++{
			fmt.Println(news[i],urls[i])
		}
	}
}