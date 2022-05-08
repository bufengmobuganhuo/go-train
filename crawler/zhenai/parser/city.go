package parser

import (
	"regexp"

	"mengyu.com/gotrain/crawler/engine"
)

var (
	profileRe = regexp.MustCompile(
		`<a href="(.*album\.zhenai\.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(
		`href="(.*www\.zhenai\.com/zhenghun/[^"]+)"`)
)

// 解析具体城市页，取出一个城市中，每个具体用户的链接
func ParseCity(contents []byte, _ string) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			// 函数式编程，使函数兼容
			Parser: NewProfileParser(string(m[2])),
		})
	}

	// 查找城市下一页
	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
	}

	return result
}
