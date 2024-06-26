package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func InitDB() (*sql.DB, error) { // 修改函数簽名以返回错误
	// 设置資料库连接字符串
	dsn := "host=db user=user dbname=mydatabase password=password port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// 连接到資料库
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// 测试資料库连接
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	fmt.Println("Connected to the database successfully!")
	return db, nil
}
