package parser_test

import (
	"io/ioutil"
	"testing"

	"mengyu.com/gotrain/crawler/fetcher"
	"mengyu.com/gotrain/crawler/zhenai/parser"
)

func TestParseProfile(t *testing.T) {
	contents, err := fetcher.Fetch("http://album.zhenai.com/u/108906739")
	err = ioutil.WriteFile("profile_test_data.html", contents, 0644)

	if err != nil {
		panic(err)
	}

	result := parser.ParseProfile(contents)
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v", result.Items)
	}

	// profile := result.Items[0].(model.Profile)

	// expected := model.Profile{
	// 	Age:        34,
	// 	Height:     162,
	// 	Weight:     57,
	// 	Income:     "3001-5000元",
	// 	Gender:     "女",
	// 	Name:       "安静的雪",
	// 	Xinzuo:     "牡羊座",
	// 	Occupation: "人事/行政",
	// 	Marriage:   "离异",
	// 	House:      "已购房",
	// 	Hokou:      "山东菏泽",
	// 	Education:  "大学本科",
	// 	Car:        "未购车",
	// }
}
