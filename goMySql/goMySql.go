package goMySql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// 登録されているユーザ情報を格納する構造体
type Authentication struct {
	Name  string
	ID    string
	Image string
}

// LogをDBに書き込む際に用いる構造体
type Log struct {
	Name   string
	Number string
	Result bool
	Status string
}

// DBに書き込む
func SqlWrite(user Log) {
	// DBに接続　
	db, err := sql.Open("mysql", "root:secret@/rails-app_development")
	if err != nil {
		return
	}
	defer db.Close()

	// insert文
	stmt, err := db.Prepare("INSERT INTO logs VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return
	}

	// 行を取得
	rows, err := db.Query("SELECT COUNT(id) FROM logs")
	if err != nil {
		return
	}
	defer rows.Close()

	var count int
	for rows.Next() {
		rows.Scan(&count)
	}
	fmt.Println(count)

	// データをDBに挿入
	const layout = "2006-01-02 15:04:05"
	tm := time.Now().Format(layout)
	fmt.Println(tm)
	// ID UserName UserNumber CheckTime Result Status Create Update
	_, err = stmt.Exec(count+1, user.Name, user.Number, tm, user.Result, user.Status, tm, tm)
	if err != nil {
		return
	}
	defer stmt.Close()
}

// DBから読み込む
func SqlRead() ([]Authentication) {
	// 接続
	db, err := sql.Open("mysql", "root:secret@/rails-app_development")
	if err != nil {
		return nil
	}
	defer db.Close()

	// Select
	rows, err := db.Query("SELECT user_name, user_number, face_image FROM authentications")
	if err != nil {
		return nil
	}
	defer rows.Close()

	// 構造体へマッピング
	var (
		users []Authentication
		user  Authentication
	)
	for rows.Next() {
		rows.Scan(&user.Name, &user.ID, &user.Image)
		users = append(users, user)
	}

	return users
}
