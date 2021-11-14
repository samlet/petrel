package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/olivere/elastic/v7"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	indexName = "subject"
	servers   = []string{"http://localhost:9200/"}
	client, _ = elastic.NewClient(elastic.SetURL(servers...))
)

type Subject struct {
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	Genres []string `json:"genres"`
}

func fetch(url string) *html.Node {
	log.Println("Fetch Url", url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Http get err:", err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("Http status code:", resp.StatusCode)
	}
	defer resp.Body.Close()
	doc, err := htmlquery.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func parseUrls(url string, ch chan bool) {
	doc := fetch(url)
	ctx := context.Background()
	nodes := htmlquery.Find(doc, `//ol[@class="grid_view"]/li`)

	subjects := []Subject{}

	for _, node := range nodes {
		url := htmlquery.FindOne(node, `.//div[@class="hd"]/a/@href`)
		title := htmlquery.FindOne(node, `.//span[@class="title"]/text()`)
		genre := htmlquery.Find(node, `.//div[@class="bd"]/p/text()`) // 包含了演员、导演、类型等内容的文本Node

		id, _ := strconv.Atoi(strings.Split(htmlquery.InnerText(url), "/")[4]) // 把ID转成数字
		genreStr := strings.Split(htmlquery.InnerText(genre[1]), "/")[2]       // 选择第二个Node，用`/`分割后选择第三部分就是类型了
		subject := Subject{id, htmlquery.InnerText(title),
			strings.Split(strings.TrimSpace(genreStr), " ")} // 拼一个结构体对象
		subjects = append(subjects, subject) // 使用append放到subjects slice里面
	}

	bulkRequest := client.Bulk() // 借用之前的批量操作的代码
	for _, subject := range subjects {
		doc := elastic.NewBulkIndexRequest().Index(indexName).Id(strconv.Itoa(subject.ID)).Doc(subject)
		bulkRequest = bulkRequest.Add(doc)
	}

	response, err := bulkRequest.Do(ctx) // 每页的25个文档一次性写入到es
	if err != nil {
		//panic(err)
		fmt.Println("fail: ", err.Error())
	}
	failed := response.Failed()
	l := len(failed)
	if l > 0 {
		fmt.Printf("Error(%d)", l, response.Errors)
	}

	log.Println("Finished Url", url)

	ch <- true
}

func doubanCrawler() {
	start := time.Now()
	ch := make(chan bool)

	for i := 0; i < 10; i++ {
		go parseUrls("https://movie.douban.com/top250?start="+strconv.Itoa(25*i), ch)
	}

	for i := 0; i < 10; i++ {
		<-ch
	}

	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func PrintQuery(src interface{}) {
	fmt.Println("*****")
	data, err := json.MarshalIndent(src, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func queryMovies() {
	query := elastic.NewTermQuery("genres", "动画")
	src, err := query.Source()
	if err != nil {
		panic(err)
	}
	PrintQuery(src)

	boolQuery := elastic.NewBoolQuery()
	boolQuery = boolQuery.Must(elastic.NewTermQuery("genres", "剧情"))
	boolQuery = boolQuery.Filter(elastic.NewTermQuery("id", 1))
	src, err = boolQuery.Source()
	if err != nil {
		panic(err)
	}
	PrintQuery(src)

	rangeQuery := elastic.NewRangeQuery("born").
		Gte("2012/01/01").
		Lte("now").
		Format("yyyy/MM/dd")
	src, err = rangeQuery.Source()
	PrintQuery(src)
}
