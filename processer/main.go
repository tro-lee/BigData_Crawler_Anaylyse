package main

import (
	"encoding/json"
	"log"
	"regexp"

	"github.com/go-ego/gse"
)

var (
	seg gse.Segmenter
)

func main() {
	path, err := GetFlagPath()
	if err != nil {
		log.Fatalln(err)
	}

	content, err := ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	news, err := JsonToNews(content)
	if err != nil {
		log.Fatalln(err)
	}

	// 清洗数据

	clearNews, _ := GetClearData(news)
	clear(news)

	jsonData, _ := json.Marshal(clearNews)
	err = JsonToFile(jsonData, "./result/clear_news.json")
	if err != nil {
		log.Fatalln(err)
	}

	// 开始分词
	loadDict()

	titleResult := make(map[string]int, len(clearNews))
	contentResult := make(map[string]int, len(clearNews))

	for i := range clearNews {
		words := segCut(clearNews[i].Title)

		for _, word := range words {
			reg := regexp.MustCompile("^[0-9a-zA-Z[:space:]，：（）《》〈〉“；—、`”.,;_+='？！]+$")
			if !reg.MatchString(word) {
				titleResult[word]++
			}
		}

		contents := segCut(clearNews[i].Summary)
		for _, content := range contents {
			reg := regexp.MustCompile("^[0-9a-zA-Z[:space:]，：（）《》〈〉“；—、`”.,;_+='？！]+$")

			if !reg.MatchString(content) {
				contentResult[content]++
			}
		}
	}

	titleJson, _ := json.Marshal(titleResult)
	JsonToFile(titleJson, "./result/seg_title_result.json")

	contentJson, _ := json.Marshal(contentResult)
	JsonToFile(contentJson, "./result/seg_content_result.json")

	// 提取国家关键字
	countryRegex := getCountryRegex()
	countryResult := make(map[string]int, len(clearNews))

	for _, value := range clearNews {
		if countryRegex.MatchString(value.Title) {
			title := countryRegex.Find([]byte(value.Title))
			countryResult[string(title)]++
		}
	}
	countryJson, _ := json.Marshal(countryResult)
	JsonToFile(countryJson, "./result/country_result.json")
}

func GetClearData(news []News) ([]ClearNews, error) {
	result := make([]ClearNews, len(news))
	for i := range news {
		result[i] = ClearNews{
			Title:   news[i].Title,
			Summary: news[i].Summary,
			CTime:   news[i].CTime,
		}
	}
	return result, nil
}

func JsonToNews(content string) ([]News, error) {
	var news []News
	err := json.Unmarshal([]byte(content), &news)
	if err != nil {
		return nil, err
	}
	return news, nil
}

func loadDict() {
	seg.LoadDict()
}

func segCut(content string) []string {
	return seg.Slice(content)
}
