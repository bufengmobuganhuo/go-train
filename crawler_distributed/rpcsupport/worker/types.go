package worker

import (
	"errors"
	"fmt"
	"log"

	"mengyu.com/gotrain/crawler/engine"
	"mengyu.com/gotrain/crawler/zhenai/parser"
	"mengyu.com/gotrain/crawler_distributed/config"
)

// 用于序列化parser函数
type SerializedParser struct {
	// 函数名
	Name string
	// 函数入参
	Args interface{}
}

// 可以在网络上传递的request
type Request struct {
	Url    string
	Parser SerializedParser
}

// 可以在网络上传递的parse result
type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

// 将engine的request序列化
func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: parser,
	}, nil
}

// 根据函数名返回对应的函数
func deserializeParser(s SerializedParser) (engine.Parser, error) {
	switch s.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity, config.ParseCity), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParseProfile:
		if userName, ok := s.Args.(string); ok {
			return parser.NewProfileParser(userName), nil
		}
		return nil, fmt.Errorf("invalid arg: %v", s.Args)
	default:
		return nil, errors.New("unknown parser name")
	}
}

func DeserializeResult(r ParseResult) (engine.ParseResult, error) {
	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		request, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializing request: %v", req)
			continue
		}
		result.Requests = append(result.Requests, request)
	}
	return result, nil
}
