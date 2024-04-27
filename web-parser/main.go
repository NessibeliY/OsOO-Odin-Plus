package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	url := "https://hypeauditor.com/top-instagram-all-russia/"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("error fetching URL:", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("error reading response body:", err)
	}

	rowPattern := `<div class="row__top"`

	rowRe := regexp.MustCompile(rowPattern)

	indexes := rowRe.FindAllStringIndex(string(body), -1)

	rows := make([][]byte, 0, 50)
	matches := make([][]string, 50)

	for _, index := range indexes {
		rows = append(rows, body[index[0]:index[0]+3583])
	}

	for i, row := range rows {
		rankPattern := `<span data-v-b11c405a>(\d+)</span>`
		rank := FindRegex(rankPattern, row)

		contributorPattern := `<div class="contributor__name-content" data-v-c5a99f5a>([\s\S]*?)<\/div><i class="fas verified"`
		contributor := FindRegex(contributorPattern, row)

		categoryPattern := `<div class="row-cell category" data-v-b11c405a>([\s\S]*?)<\/div><div class="row-cell subscribers"`
		category := FindRegex(categoryPattern, row)
		subcategoryPattern := `<div class="tag__content ellipsis" data-v-677f0c75>([\s\S]*?)<\/div><!---->`
		subcategories := FindAllRegexes(subcategoryPattern, category)
		for len(subcategories) < 3 {
			subcategories = append(subcategories, "")
		}

		followersPattern := `<div class="row-cell subscribers" data-v-b11c405a>([\s\S]*?)<\/div><div class="row-cell audience"`
		followers := FindRegex(followersPattern, row)

		countryPattern := `<div class="row-cell audience" data-v-b11c405a data-v-452bf4ed>([\s\S]*?)<\/div><div class="row-cell authentic"`
		country := FindRegex(countryPattern, row)

		authPattern := `<div class="row-cell authentic" data-v-b11c405a data-v-452bf4ed>([\s\S]*?)<\/div><div class="row-cell engagement"`
		auth := FindRegex(authPattern, row)

		avgPattern := `<div class="row-cell engagement" data-v-b11c405a data-v-452bf4ed>([\s\S]*?)<\/div><div class="row-cell share margin-auto"`
		avg := FindRegex(avgPattern, row)

		matches[i] = append(matches[i], rank, contributor)
		matches[i] = append(matches[i], subcategories...)
		matches[i] = append(matches[i], followers, country, auth, avg)

	}
	f, err := os.Create("file.csv")
	if err != nil {
		log.Fatal("error creating file:", err)
	}
	defer f.Close()

	_, err = f.WriteString("Rank,Influencer,Category1,Category2,Category3,Followers,Country,Eng.(Auth.),Eng.(Avg.)" + "\n")
	if err != nil {
		log.Fatal("error writing to file:", err)
	}

	for _, match := range matches {
		_, err := f.WriteString(strings.Join(match, ",") + "\n")
		if err != nil {
			log.Fatal("error writing to file:", err)
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func FindRegex(pattern string, row []byte) string {
	rankRe := regexp.MustCompile(pattern)
	rankBytes := rankRe.FindSubmatch(row)
	rank := string(rankBytes[1])

	return rank
}

func FindAllRegexes(pattern string, category string) []string {
	rankRe := regexp.MustCompile(pattern)
	rankArr := rankRe.FindAllStringSubmatch(category, -1)
	ans := []string{}
	for _, rank := range rankArr {
		ans = append(ans, strings.ReplaceAll(rank[1], "&amp;", "&"))
	}

	return ans
}
