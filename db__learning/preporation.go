package main

//导入github上的第三方数据库驱动包
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 定义与表结构相同字段的结构体
type student struct {
	password string
	name     string
	age      int
	ID       int
}

// 定义一个数据库的链接池
var database *sql.DB

// 初始化数据库
func initDB() error {
	//连接数据库
	dsn := "root:wwy040609@tcp(127.0.0.1:3306)/mydb1" //“用户名:密码@tcp(端口)/数据库名”
	//连接数据库
	var err error                          //链接失败后返回的错误
	database, err = sql.Open("mysql", dsn) //链接成功返回连接池，err返回为nil，如果链接失败err不是nil
	if err != nil {                        //如果链接失败打印失败原因//注意这次链接不校验密码和用户
		fmt.Printf("Open %s failed,err:%v\n", dsn, err)
		return err
	}
	err = database.Ping() //进行校验的链接，如果密码或用户错误err不为nil
	if err != nil {       //链接错误并打印原因
		fmt.Printf("Open %s failed,err:%v\n", dsn, err)
		return err
	}
	return nil
}

// 查询一条数据
func query() {
	var ID int //输入的接收值
	fmt.Println("请输入要查找的学生")
	fmt.Scanf("%d", &ID)
	sqlStr := "select * from stu where id=?" //查询的sql语句
	rowObj := database.QueryRow(sqlStr, ID)  //执行aql语句
	var stu student
	rowObj.Scan(&stu.name, &stu.password, &stu.age, &stu.ID) //用一个变量承接返回出的结果
	if stu.ID == 0 {                                         //如果ID==0
		fmt.Println("查找失败")
		return
	}
	fmt.Println(stu)
}

func all_query() {
	sqlStr := "select * from stu where id >= 0" //要执行的sql语句
	rowObj, err := database.Query(sqlStr)       //执行sql语句,执行失败err不等于nil
	if err != nil {
		fmt.Printf("Query %s failed,err:%v\n", sqlStr, err) //打印失败原因
		return
	}
	defer rowObj.Close() //defer语句最后关闭
	for rowObj.Next() {  //rowObj的一个方法，直到数据库中没有下一个元素
		var stu student
		err := rowObj.Scan(&stu.name, &stu.password, &stu.age, &stu.ID) //如果没有接收到返回的值，err不是nil
		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err) //打印失败的原因
			return
		}
		fmt.Println(stu)
	}
}

func insert() {
	var name string
	var age int
	var Id int
	var password string
	fmt.Println("请一次输入年龄，密码，ID号，姓名")
	fmt.Scanf("%d", &age)
	if age < 18 || age > 65 { //对年两进行判断
		fmt.Println("年龄输入错误，请从新输入")
		return
	}
	fmt.Scanf("%s", &password)
	fmt.Scanf("%d", &Id)
	fmt.Scanf("%s", &name)
	//对已经存在的数据进行查重
	sqlStr1 := "select * from stu where id=?"
	rowObj := database.QueryRow(sqlStr1, Id)
	var stu1 student
	rowObj.Scan(&stu1.name, &stu1.password, &stu1.age, &stu1.ID)
	if stu1.ID != 0 {
		fmt.Println("此账号已经存在，请从新输入")
		return
	}

	//1.先写sql语句
	sqlStr := "insert into stu(password,name,age,ID) values(?,?,?,?)"
	//2.exec
	ret, err := database.Exec(sqlStr, password, name, age, Id)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	//如果是插入数据的操作，能拿到插入数据的ID值
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
	}
	fmt.Printf("insert success, id:%d\n", id)
}

// 更新操作
// 与insert操作类似
func update() {

	var name string
	var age int
	var Id int
	var password string

	var nextID int
	fmt.Println("请输入要更改的ID")
	fmt.Scanf("%d", &nextID)

	fmt.Println("请输入你要更新后的值")
	fmt.Scanf("%s", &name)
	fmt.Scanf("%d", &age)
	if age < 18 || age > 65 {
		fmt.Println("年龄输入错误,请从新输入年龄")
		return
	}
	fmt.Scanf("%d", &Id)
	fmt.Scanf("%s", &password)

	sqlStr1 := "select * from stu where id=?"
	rowObj := database.QueryRow(sqlStr1, Id)

	var stu1 student
	rowObj.Scan(&stu1.name, &stu1.password, &stu1.age, &stu1.ID)
	if stu1.ID != 0 {
		fmt.Println("此账号已经存在，请从新输入")
		return
	}

	sqlStr := "update stu set name=?,age=?,ID=?,password=? where ID=?"
	ret, err := database.Exec(sqlStr, name, age, Id, password, nextID)
	if err != nil {
		fmt.Printf("update failed,err:%v\n", err)
		return
	}
	id, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return
	}
	fmt.Printf("update success, id:%d\n", id)
}

func deleteRow() {
	sqlStr := "delete from stu where ID=?"
	var ID int
	fmt.Println("请输入你要删除的ID")
	fmt.Scanf("%d", &ID)
	result, err := database.Exec(sqlStr, ID)

	if err != nil {
		fmt.Printf("delete failed,err:%v\n", err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("failed to get rows affected, err:%v\n", err)
		return
	}

	if rowsAffected == 0 {
		fmt.Println("No rows were deleted")
		return
	} else {
		fmt.Println("delete success, ", rowsAffected, " rows were deleted")
		return
	}
}
