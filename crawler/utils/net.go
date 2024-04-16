package utils

import (
	"strconv"
)

const baseURL = `https://world.huanqiu.com/api/list?node="/e3pmh22ph/e3pmh2398","/e3pmh22ph/e3pmh26vv","/e3pmh22ph/e3pn6efsl","/e3pmh22ph/efp8fqe21"&offset=`

func GetPageUrl(page int) string {
	return baseURL + strconv.Itoa((page+1)*24) + `&limit=24`
}
