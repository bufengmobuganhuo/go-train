package parser_test

import (
	"io/ioutil"
	"testing"

	"mengyu.com/gotrain/crawler/zhenai/parser"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	result := parser.ParseCityList(contents, "")

	const resultSize = 470
	var expectedUrls = []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	var expectedCities = []string{
		"City 阿坝",
		"City 阿克苏",
		"City 阿拉善盟",
	}
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Requests))
	}
	if len(result.Items) != resultSize {
		t.Errorf("result should have %d items; but had %d", resultSize, len(result.Items))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s", i, url, result.Requests[i].Url)
		}
	}
	for i, city := range expectedCities {
		if result.Items[i].Payload.(string) != city {
			t.Errorf("expected city #%d: %s; but was %s", i, city, result.Items[i].Payload.(string))
		}
	}
	// verify result

}
