package persist_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/olivere/elastic/v7"
	"mengyu.com/gotrain/crawler/model"
	"mengyu.com/gotrain/crawler/persist"
)

func TestSave(t *testing.T) {
	expected := model.Profile{
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
	}
	id, err := persist.Save(expected)
	if err != nil {
		panic(err)
	}
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	resp, err := client.Get().Index("dating_profile").Id(id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	var actual model.Profile
	err = json.Unmarshal([]byte(resp.Source), &actual)
	if err != nil {
		panic(err)
	}
	if actual != expected {
		t.Errorf("got %v; but expect %v", actual, expected)
	}
}
