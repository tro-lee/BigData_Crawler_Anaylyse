package manager

import (
	"regexp"
)

type PageParser struct {
	dataCh  []*PageData
	manager *Manager
}

func NewPageParser(manager *Manager) *PageParser {
	parse := new(PageParser)
	parse.dataCh = make([]*PageData, 0)
	parse.manager = manager
	return parse
}

func (p *PageParser) SendPageData(data *PageData) {
	p.dataCh = append(p.dataCh, data)
}

func (p *PageParser) Run() {
	result := make([]FilmData, 0)
	for _, data := range p.dataCh {
		result = append(result, p.parsePageData(data)...)
	}
	p.manager.ParserFinished(result)
}

func (p *PageParser) parsePageData(data *PageData) []FilmData {
	filmNameReg := regexp.MustCompile(`<img width="100" alt="(?s:(.*?))"`)
	filmNames := filmNameReg.FindAllStringSubmatch(data.data, -1)

	filmScoreReg := regexp.MustCompile(`<span class="rating_num" property="v:average">(.*)</span>`)
	filmScores := filmScoreReg.FindAllStringSubmatch(data.data, -1)

	filmScoreNumReg := regexp.MustCompile(`<span>(.*)人评价</span>`)
	filmScoreNum := filmScoreNumReg.FindAllStringSubmatch(data.data, -1)

	filmCommentReg := regexp.MustCompile(`<span class="inq">(.*)</span>`)
	filmComments := filmCommentReg.FindAllStringSubmatch(data.data, -1)

	filmContentReg := regexp.MustCompile(`(?s)<p class="">(.*?)</p>`)
	filmContents := filmContentReg.FindAllStringSubmatch(data.data, -1)

	films := make([]FilmData, len(filmNames))
	for i := 0; i < len(films); i++ {
		if len(filmComments) <= i {
			filmComments = append(filmComments, []string{"", ""})
		}

		films[i] = FilmData{
			Name:    filmNames[i][1],
			Score:   filmScores[i][1],
			People:  filmScoreNum[i][1],
			Comment: filmComments[i][1],
			Content: filmContents[i][1],
		}
	}

	return films
}
