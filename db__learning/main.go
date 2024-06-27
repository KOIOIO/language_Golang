package main

import (
	"fmt"
	"os"
)

func desktop() {
	fmt.Print(
		"学生管理系统\n" +
			"1.添加学生\n" +
			"2.删除学生\n" +
			"3.修改学生\n" +
			"4.查询学生\n" +
			"5.展示学生\n" +
			"6.退出系统\n",
	)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("初始化失败，数据库连接失败 error=%v\n", err)
		return
	} else {
		jia()
		fmt.Printf("初始化成功\n")
	}
	for {
		desktop()
		var input string
		fmt.Scanf("%s", &input)
		switch input {
		case "1":
			insert()
		case "2":
			deleteRow()
		case "3":
			update()
		case "4":
			query()
		case "5":
			all_query()
		case "6":
			os.Exit(1314)
		default:
			fmt.Println("gun")
			fmt.Println(add_data(1, 2))
			os.Exit(250)
			break
		}
	}
}
