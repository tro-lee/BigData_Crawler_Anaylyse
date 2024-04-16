package manager

type PageData struct {
	data  string
	index int
}

type News struct {
	Aid            string `json:"aid"`
	Title          string `json:"title"`
	Summary        string `json:"summary"`
	AddlType       string `json:"addltype"`
	ExtDisplayTime string `json:"ext_displaytime"`
	ExtDeferTime   string `json:"ext_defertime"`
	CTime          string `json:"ctime"`
	XTime          string `json:"xtime"`
	Cover          string `json:"cover"`
	Host           string `json:"host"`
	ExtSerious     string `json:"ext-serious"`
	ExtWeight      string `json:"ext-weight"`
}

type Data struct {
	List []News `json:"list"`
}
