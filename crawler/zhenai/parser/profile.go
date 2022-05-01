package parser

import (
	"regexp"
	"strconv"

	"mengyu.com/gotrain/crawler/engine"
	"mengyu.com/gotrain/crawler/model"
)

var ageRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">43岁</div>`)

var heightRe = regexp.MustCompile(
	`<td><span class="label">身高：</span>(\d+)CM</td>`)
var incomeRe = regexp.MustCompile(
	`<td><span class="label">月收入：</span>([^<]+)</td>`)
var weightRe = regexp.MustCompile(
	`<td><span class="label">体重：</span><span field="">(\d+)KG</span></td>`)
var genderRe = regexp.MustCompile(
	`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var xinzuoRe = regexp.MustCompile(
	`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var marriageRe = regexp.MustCompile(
	`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(
	`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(
	`<td><span class="label">职业：</span><span field="">([^<]+)</span></td>`)
var hokouRe = regexp.MustCompile(
	`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(
	`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(
	`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
var guessRe = regexp.MustCompile(
	`<a class="exp-user-name"[^>]*href="(.*album\.zhenai\.com/u/[\d]+)">([^<]+)</a>`)
var idUrlRe = regexp.MustCompile(
	`.*album\.zhenai\.com/u/([\d]+)`)

// 解析某个具体用户的信息
func ParseProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}

	if age, err := strconv.Atoi(extractString(contents, ageRe)); err != nil {
		profile.Age = age
	}
	if height, err := strconv.Atoi(extractString(contents, heightRe)); err != nil {
		profile.Height = height
	}
	if weight, err := strconv.Atoi(extractString(contents, weightRe)); err != nil {
		profile.Weight = weight
	}
	profile.Income = extractString(
		contents, incomeRe)
	profile.Gender = extractString(
		contents, genderRe)
	profile.Car = extractString(
		contents, carRe)
	profile.Education = extractString(
		contents, educationRe)
	profile.Hokou = extractString(
		contents, hokouRe)
	profile.House = extractString(
		contents, houseRe)
	profile.Marriage = extractString(
		contents, marriageRe)
	profile.Occupation = extractString(
		contents, occupationRe)
	profile.Xinzuo = extractString(
		contents, xinzuoRe)

	result := engine.ParseResult {
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
