package engine

type Request struct {
	// 要请求的URL
	Url string
	// 对应的解析器
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	// 可能会产生新的request
	Requests []Request
	// 解析出来的结果，interface表示任何类型
	Items []interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
