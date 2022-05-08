package engine

type ParserFunc func(contents []byte, url string) ParseResult

// 代表一个parser函数
type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

type Request struct {
	// 要请求的URL
	Url string
	// 对应的解析器
	Parser Parser
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

type NilParser struct {
}

func (NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

type FuncParser struct {
	// parser函数
	parser ParserFunc
	// 函数名
	name string
}

// FuncParser实现了Parser接口，所以这里的返回可以作为Request的参数
func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}
