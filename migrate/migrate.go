package main

import (
	"fmt"
	"go-rest-api-app/db"
	"go-rest-api-app/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")//defer は、関数が終了するときに実行される処理を予約するキーワード
	defer db.CloseDB(dbConn)//複数の defer がある場合、逆順で実行
	dbConn.AutoMigrate(&model.User{}, &model.Task{})//データベースにテーブルを自動生成
}

// データベースに接続し、
// モデルに基づいてデータベースのテーブルを作成し、
// 接続を解放する。