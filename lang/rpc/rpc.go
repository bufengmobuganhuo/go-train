package rpcdemo

import "errors"

// 定义一个RPC服务
type DemoService struct{}

type Args struct {
	A, B int
}

// 返回结果必须是指针
func (DemoService) Div(args Args, result *float64) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}
	*result = float64(args.A) / float64(args.B)
	return nil
}
