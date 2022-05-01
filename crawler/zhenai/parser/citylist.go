package parser

import (
	"regexp"

	"mengyu.com/gotrain/crawler/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[A-Za-z0-9]+)" [^>]*>([^<]+)</a>`

// 解析城市列表
func ParseCityList(contents []byte) engine.ParseResult {
	// ^>: 表示不是'>'
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		// 解析出的城市
		result.Items = append(result.Items, "City "+string(m[2]))
		// 城市对应的URL，生成新的request
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			// 生成的新URL对应的parser
			ParserFunc: ParseCity,
		})
	}
	return result
}
