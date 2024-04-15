package manager

import (
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
	result, err := g.getPageData()
	if err != nil {
		log.Fatal(err)
	}

	g.finsihed(&PageData{data: result, index: g.index})
}

func (g *PageGetter) getPageData() (string, error) {
	request, err := http.NewRequest("GET", g.url, nil)
	if err != nil {
		return "", err
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.88 Safari/537.36")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	buff := make([]byte, 1024)
	var result string

	for {
		n, err1 := resp.Body.Read(buff)
		if n <= 0 {
			break
		}

		if err1 != nil && err1 != io.EOF {
			err = err1
			break
		}

		result += string(buff[:n])
	}

	if err != nil {
		return "", err
	}

	return result, nil
}

func (g *PageGetter) finsihed(page *PageData) {
	g.manager.GetterFininshed(page)
	g.isFinish = true
}
