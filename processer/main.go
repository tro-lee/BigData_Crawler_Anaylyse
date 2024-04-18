package main

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
	"unicode/utf8"

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

	successChan := make(chan int)

	fmt.Println("开始加载词典")

	// 开始分词
	loadDict()

	fmt.Println("开始加工")

	// 提取关键字
	go segWord(clearNews, successChan)

	// 提取国家关键字
	go countryResult(clearNews, successChan)

	// 提取国家词汇
	go countryResultPro(clearNews, successChan)

	successedNum := 0
	for {
		if successedNum >= 3 {
			break
		}

		num := <-successChan
		fmt.Println(num)
		successedNum++
	}
	fmt.Println("完成加工")
}

// 得到分词文件
func segWord(clearNews []ClearNews, successChan chan int) {
	titleResult := make(map[string]int, len(clearNews))
	contentResult := make(map[string]int, len(clearNews))

	for i := range clearNews {
		words := segCut(clearNews[i].Title)

		for _, word := range words {
			reg := regexp.MustCompile("^[0-9a-zA-Z[:space:]，：（）《》〈〉“；—、`”.,;_+='？！]+$")

			if utf8.RuneCountInString(word) <= 1 {
				continue
			}

			if !reg.MatchString(word) {
				titleResult[word]++
			}
		}

		contents := segCut(clearNews[i].Summary)
		for _, content := range contents {
			reg := regexp.MustCompile("^[0-9a-zA-Z[:space:]，：（）《》〈〉“；—、`”.,;_+='？！]+$")

			if utf8.RuneCountInString(content) <= 1 {
				continue
			}

			if !reg.MatchString(content) {
				contentResult[content]++
			}
		}
	}

	titleJson, _ := json.Marshal(titleResult)
	JsonToFile(titleJson, "./result/seg_title_result.json")
	fmt.Println("新闻标题分词完成")

	contentJson, _ := json.Marshal(contentResult)
	JsonToFile(contentJson, "./result/seg_content_result.json")
	fmt.Println("新闻内容分词完成")

	successChan <- 1
}

// 得到关键字提取
func countryResult(clearNews []ClearNews, successChan chan int) {
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
	fmt.Println("国家关键字提取完成")

	successChan <- 2
}

// 提取词汇
func countryResultPro(clearNews []ClearNews, successChan chan int) {
	countryRegex := getCountryRegex()

	// 提取国家词汇
	countryResultPro := make(map[string]map[string]int, len(clearNews))
	for _, value := range clearNews {
		if countryRegex.MatchString(value.Title) {
			title := countryRegex.Find([]byte(value.Title))
			if _, ok := countryResultPro[string(title)]; !ok {
				countryResultPro[string(title)] = make(map[string]int)
			}

			words := segCut(value.Title)
			for _, word := range words {
				if strings.EqualFold(word, string(title)) {
					continue
				}
				if utf8.RuneCountInString(word) <= 1 {
					continue
				}

				reg := regexp.MustCompile("^[0-9a-zA-Z[:space:]，：（）《》〈〉“；—、`”.,;_+='？！]+$")
				if !reg.MatchString(word) {
					countryResultPro[string(title)][word]++
				}
			}
		}
	}

	countryJsonPro, _ := json.Marshal(countryResultPro)
	JsonToFile(countryJsonPro, "./result/country_result_pro.json")
	fmt.Println("国家词汇提取完成")

	// 判断哪个词汇，国家占比最多
	countryMaxResult := make(map[string]map[string]int, len(clearNews))

	for key, value := range countryResultPro {
		max := 0
		maxKey := ""
		for k, v := range value {
			// 如果是国家名，直接跳过
			if countryRegex.MatchString(k) {
				continue
			}
			if utf8.RuneCountInString(k) <= 1 {
				continue
			}

			if v > max {
				max = v
				maxKey = k
			}
		}
		countryMaxResult[key] = map[string]int{maxKey: max}
	}
	countryMaxJson, _ := json.Marshal(countryMaxResult)
	JsonToFile(countryMaxJson, "./result/country_max_result.json")
	fmt.Println("国家词汇Pro提取完成")

	successChan <- 3
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
