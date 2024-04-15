package utils

import (
	"strconv"
)

var PageUrlBase = "https://movie.douban.com/top250?start=%s&filter="

func GetPageUrl(page int) string {
	return "https://movie.douban.com/top250?start=" + strconv.Itoa((page-1)*25) + "&filter="
}
