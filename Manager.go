package main

import (
	"sync"

	"github.com/golang/glog"
)

type Manager struct {
	startPage   uint
	endPage     uint
	pageGetters []*PageGetter
	pageParse   []*PageParse
}

func NewManager(start, end uint) *Manager {
	manager := &Manager{
		startPage:   start,
		endPage:     end,
		pageGetters: make([]*PageGetter, end-start+1),
		pageParse:   make([]*PageParse, end-start+1),
	}

	for i := range manager.pageGetters {
		manager.pageGetters[i] = NewPageGetter(getPageUrl(start+uint(i)), start+uint(i), manager)
	}

	for i := range manager.pageParse {
		manager.pageParse[i] = NewPageParse()
	}

	return manager
}

func (m Manager) Run() {
	waitGroup := sync.WaitGroup{}

	waitGroup.Add(len(m.pageParse))
	for _, parse := range m.pageParse {
		go func(parse *PageParse) {
			defer waitGroup.Done()
			parse.Run()
		}(parse)
	}

	waitGroup.Add(len(m.pageGetters))
	for _, getter := range m.pageGetters {
		go func(getter *PageGetter) {
			defer waitGroup.Done()
			getter.Run()
		}(getter)
	}

	waitGroup.Wait()
}

func (m *Manager) SendPageData(data *PageData) {

	index := data.index - m.startPage
	if index < 0 || index > (m.endPage-m.startPage) {
		glog.Fatalf("send page data error, index %d less than zero or over page size %d", index, m.endPage-m.endPage)
	}
	parse := m.pageParse[index]
	if parse == nil {
		glog.Fatalf("get page parse error, parse is null")
	}

	parse.SendPageData(data)
}
