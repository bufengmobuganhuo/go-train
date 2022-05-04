package persist

import (
	"context"
	"log"

	"github.com/olivere/elastic/v7"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver got item #%d: %v", itemCount, item)
			itemCount++

			_, err := Save(item)
			if err != nil {
				log.Print("Item Saver: error saving item %v: %v", item, err)
			}

		}
	}()
	return out
}

func Save(item interface{}) (string, error) {
	client, err := elastic.NewClient(
		// 在docker中必须关闭
		elastic.SetSniff(false))
	if err != nil {
		return "", nil
	}
	resp, err := client.Index().Index("dating_profile").BodyJson(item).Do(context.Background())
	if err != nil {
		return "", err
	}
	return resp.Id, nil
}
