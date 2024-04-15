package manager

import (
	"douban/utils"
)

type Manager struct {
	pageGetters []*PageGetter
	PageParser  *PageParser
	RawResults  []FilmData
	Result      chan []ResultData
}

func New() *Manager {
	manager := &Manager{
		pageGetters: make([]*PageGetter, 10),
		PageParser:  nil,
		Result:      make(chan []ResultData),
	}

	for i := range manager.pageGetters {
		manager.pageGetters[i] = NewPageGetter(utils.GetPageUrl(i), i, manager)
	}
	manager.PageParser = NewPageParser(manager)

	return manager
}

func (m Manager) Run() {
	for _, getter := range m.pageGetters {
		go func(getter *PageGetter) {
			getter.Run()
		}(getter)
	}

	m.PageParser.Run()
}

func (m *Manager) GetterFininshed(data *PageData) {
	m.PageParser.SendPageData(data)
}

func (m *Manager) ParserFinished(data []FilmData) {
	m.RawResults = append(m.RawResults, data...)

	if checkGetterAndPerserFinished(m) {
		m.final()
	}
}

func checkGetterAndPerserFinished(manager *Manager) bool {
	if len(manager.PageParser.dataCh) > 0 {
		return false
	}

	for _, getter := range manager.pageGetters {
		if !getter.isFinish {
			return false
		}
	}

	return true
}

func (m *Manager) final() {
	realResult := m.processRawData()

	// m.PageParser.Close()
	m.Result <- realResult
	close(m.Result)
}

func (m *Manager) processRawData() []ResultData {
	realResut := []ResultData{}

	for _, data := range m.RawResults {
		content := utils.ProcessContent(data.Content)
		result := ResultData{
			Name:     data.Name,
			Score:    data.Score,
			People:   data.People,
			Comment:  data.Comment,
			Director: content.Director,
			Country:  content.Country,
			Year:     content.Year,
			Genre:    content.Genre,
		}
		realResut = append(realResut, result)
	}
	return realResut
}
