package persist

import (
	"context"
	"log"

	"github.com/olivere/elastic/v7"
	"mengyu.com/gotrain/crawler/engine"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		// 在docker中必须关闭
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver got item #%d: %v", itemCount, item)
			itemCount++

			err := Save(client, index, item)
			if err != nil {
				log.Printf("Item Saver: error saving item %v:%v", item, err)
			}

		}
	}()
	return out, nil
}

func Save(client *elastic.Client, index string, item engine.Item) error {
	indexService := client.Index().Index(index).BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
