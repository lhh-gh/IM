package main

import (
	"fmt"
	"github/lhh-gh/IM/internal/common/dao/myMysql"
)

var dataSource = "root:123456@tcp(127.0.0.1:3306)/im_server_db?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	mysqldb := myMysql.MustNewMySQL(dataSource)
	if err := mysqldb.Migrate(); err != nil {
		fmt.Println("自动迁移失败")
	} else {
		fmt.Println("自动迁移成功")
	}
}
