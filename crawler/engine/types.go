package engine

type ParserFunc func(contents []byte, url string) ParseResult

type Request struct {
	// 要请求的URL
	Url string
	// 对应的解析器
	ParserFunc ParserFunc
}

type ParseResult struct {
	// 可能会产生新的request
	Requests []Request
	// 解析出来的结果，interface表示任何类型
	Items []Item
}

type Item struct {
	Url     string
	Id      string
	Payload interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}

type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
