package persist_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/olivere/elastic/v7"
	"mengyu.com/gotrain/crawler/engine"
	"mengyu.com/gotrain/crawler/model"
	"mengyu.com/gotrain/crawler/persist"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url: "http://album.zhenai.com/u/108906739",
		Id:  "108906739",
		Payload: model.Profile{
			Age:        34,
			Height:     162,
			Weight:     57,
			Income:     "3001-5000元",
			Gender:     "女",
			Name:       "安静的雪",
			Xinzuo:     "牡羊座",
			Occupation: "人事/行政",
			Marriage:   "离异",
			House:      "已购房",
			Hokou:      "山东菏泽",
			Education:  "大学本科",
			Car:        "未购车",
		},
	}
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	const index = "dating_test"
	err = persist.Save(client, index, expected)
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().Index(index).Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	var actual engine.Item
	err = json.Unmarshal([]byte(resp.Source), &actual)
	if err != nil {
		panic(err)
	}
	// 由于Item的Playload是一个interface，所以json反序列化时不知道反序列化成什么，所以这里需要手动反序列化
	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile
	if actual != expected {
		t.Errorf("got %v; but expect %v", actual, expected)
	}
}
