package manager

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type PageGetter struct {
	url      string
	index    int
	manager  *Manager
	isFinish bool
}

func NewPageGetter(url string, index int, manager *Manager) (getter *PageGetter) {
	getter = new(PageGetter)
	getter.url = url
	getter.index = index
	getter.manager = manager
	getter.isFinish = false
	return
}

func (g *PageGetter) Run() {
	pageString, err := g.getPageData()
	if err != nil {
		log.Fatal(err)
		g.isFinish = true
	}

	result := getNews(pageString)
	g.manager.GetterFininshed(result)
	g.isFinish = true
}

func (g *PageGetter) getPageData() (string, error) {
	request, err := http.NewRequest("GET", g.url, nil)
	if err != nil {
		return "", err
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.88 Safari/537.36")
	request.Header.Set("Cookie", "HMF_CI=0ad043adfa5571dd32e411cb93623d299c25486c1ab0b40a4324c29e7bfc1064c18bc60a2432bff8563dd4bd27e923918834ccc6574e9561abfd21a26ebcae73dd; HBB_HC=05220b8b05a8fc8646abfcaf0b967b94ac587294a452d3386af02affa2bd8605a04f3a4d6a31d83903c234c8227d968678; HMY_JC=4386ca00d6c50ba59eb6d1be319b4c326947c076aa3ad0c8fcf3ce62fdf4e18e5e,; country_code=CN")

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	result := string(bodyBytes)

	if len(result) > 0 {
		return result, nil
	} else {
		return "", fmt.Errorf("No match found")
	}
}

func getNews(data string) []News {
	var response Data
	err := json.Unmarshal([]byte(data), &response)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return []News{}
	}
	return response.List
}
