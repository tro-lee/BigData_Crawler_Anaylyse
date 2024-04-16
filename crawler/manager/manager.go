package manager

import (
	"bigdata/utils"
	"fmt"
)

type Manager struct {
	pageGetters []*PageGetter
	RawResults  chan []News
	Result      chan []News
}

func New() *Manager {
	manager := &Manager{
		pageGetters: make([]*PageGetter, 415),
		RawResults:  make(chan []News),
		Result:      make(chan []News),
	}

	for i := range manager.pageGetters {
		manager.pageGetters[i] = NewPageGetter(utils.GetPageUrl(i), i, manager)
	}

	return manager
}

func (m Manager) Run() {
	for _, getter := range m.pageGetters {
		go getter.Run()
	}

	realResult := []News{}
	for {
		result := <-m.RawResults
		realResult = append(realResult, result...)
		fmt.Println("Got result", len(realResult))
		if m.checkFinish() {
			break
		}
	}

	m.Result <- realResult
}

func (m *Manager) checkFinish() bool {
	for _, getter := range m.pageGetters {
		if !getter.isFinish {
			return false
		}
	}
	return true
}

func (m *Manager) GetterFininshed(data []News) {
	m.RawResults <- data
}
