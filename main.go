package main

import (
	"databaseAPI/goMySql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// POSTリクエストを処理する
func postJsonHandler(rw http.ResponseWriter, req *http.Request) {
	// リクエストの設定
	rw.Header().Set("Content-Type", "application/json")

	// メソッドの確認
	if req.Method != "POST" {
		fmt.Fprint(rw, "Method Not POST.")
		return
	}

	// データの取得
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprint(rw, err.Error())
		fmt.Println(err.Error())
		return
	}

	// データの代入
	var input goMySql.Log
	err = json.Unmarshal(body, &input)
	if err != nil {
		fmt.Fprint(rw, err.Error())
		fmt.Println(err.Error())
		return
	}

	// 受け取ったデータの表示
	fmt.Printf("%#v\n", input)

	// クライアントへのレスポンス
	fmt.Fprint(rw, "success post data!")

	// 受け取ったデータをDBに書き込む
	goMySql.SqlWrite(input)
}

// GETリクエストを処理する
func getJsonHandler(rw http.ResponseWriter, req *http.Request) {
	// リクエストの設定
	rw.Header().Set("Content-Type", "application/json")

	// DBからデータを取得
	users := goMySql.SqlRead()

	// 構造体をJSONに変換
	response, err := json.Marshal(users)
	if err != nil {
		fmt.Fprint(rw, err.Error())
		return
	}

	// クライアントにデータを投げる
	fmt.Fprint(rw, string(response))
}

func main() {
	// ハンドラの設定
	http.HandleFunc("/post", postJsonHandler)
	http.HandleFunc("/get", getJsonHandler)
	http.ListenAndServe(":9999", nil)
}
