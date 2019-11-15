package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Post struct {
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
	Title     string `json:"title"`
	Url       string `json:"url"`
}

func main() {
	// コマンドラインオプションを取得する
	var (
		user = flag.String("user", "", "取得したい投稿を持つユーザーのid")
		tag  = flag.String("tag", "", "取得したい投稿のタグ")
	)

	flag.Parse()

	isUserEmpty := *user == ""
	isTagEmpty := *tag == ""

	// ユーザーもタグも指定されていなかったら終了する
	if isUserEmpty && isTagEmpty {
		log.Fatal("少なくともユーザーかタグのどちらかは指定してください")
	}

	url := "https://qiita.com/api/v2/items?query="

	if !isUserEmpty {
		url += ("user%3A" + *user)
	}
	if !isTagEmpty {
		if !isUserEmpty {
			url += "+"
		}
		url += ("tag%3A" + *tag)
	}

	res, error := http.Get(url)

	if error != nil {
		log.Fatal(error)
	}

	defer res.Body.Close()

	bytes, _ := ioutil.ReadAll(res.Body)

	var posts []Post
	if err := json.Unmarshal(bytes, &posts); err != nil {
		log.Fatal(err)
	}
	// デコードしたデータを表示
	for _, p := range posts {
		fmt.Printf("Title: %s\nURL: %s\nCreatedAt: %s\nUpdateAt: %s\n\n", p.Title, p.Url, p.CreatedAt, p.UpdatedAt)
	}
}
