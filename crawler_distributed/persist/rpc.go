package persist

import (
	"github.com/olivere/elastic/v7"
	"mengyu.com/gotrain/crawler/engine"
	"mengyu.com/gotrain/crawler/persist"
)

// 封装一个RPC服务
type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, item)
	if err == nil {
		*result = "ok"
	}
	return err
}
