package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"github.com/pelletier/go-toml"
	"github.com/use-iris/configs"
	"github.com/use-iris/libs"
	"log"
	"net/http"
)
var dec = mahonia.NewDecoder("gbk")

func main(){
	res, err := http.Get("http://www.chinastor.com/xinpian/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".active_list .textcont").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		var detail_url, _ = s.Find("a").Attr("href")
		getDetail(detail_url)

	})
}

func getDetail(detail_url string){
	res, err := http.Get(detail_url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var art_title = dec.ConvertString(doc.Find(".art-title h1").Text())
	var create_time = doc.Find(".source").Find("span").First().Text()
	var source = dec.ConvertString(doc.Find(".source").Find("span").Last().Text())
	var brief = dec.ConvertString(doc.Find(".art-desc").Text())
	//var content = dec.ConvertString(doc.Find(".article").Text())
	fmt.Println(art_title)
	fmt.Println(create_time)
	fmt.Println(source)
	fmt.Println(brief)

}
func init(){
	//后面做接参数获取配置项路径
	var configPath = flag.String("c","github.com/use-iris/configs/config.toml","defaut config")
	var err error
	configs.ConfigTree,err = toml.LoadFile(*configPath)
	if err != nil{
		fmt.Println(err)
		return
	}
	libs.InitAllDB()
}