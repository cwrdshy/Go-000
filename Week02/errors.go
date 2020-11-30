package main

import (
	"database/sql"
	"errors"
	"fmt"
)

var errSqlNoRows = errors.New("sql row not found")

func dao() (string, error) {
	var name string
	row := db.Con().QueryRow("SELECT name FROM users WHERE id = 1;")
	err := row.Scan(&name)
	if err != nil && err == sql.ErrNoRows {
		// 如果报错直接返回
		return "", errSqlNoRows
	}
	return row, nil
}

func service() (int, string) {
	data, err := dao()
	// 根据dao层返回的错误进行判定，给main函数调用
	// 屏蔽dao错误，返回状态码和数据
	// 思考：
	// 1.这样做可以屏蔽dao底层技术实现，增加重用性
	// 2.对上层api调用层友好，上层api先根据状态码判断
	// 再决定是否要取值
	if errors.Is(err, errSqlNoRows) {
		return 404, ""
	}
	return 200, data
}

func main() {
	switch code, data := service(); code {
	case 200:
		fmt.Println(data)
	case 404:
		fmt.Println("not found")
	default:
		fmt.Println("...")
	}
}
