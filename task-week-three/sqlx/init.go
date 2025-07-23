package sqlx

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

/*
*
题目1：使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
题目2：实现类型安全映射
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*/
/**
create table employees(
    id bigint primary key auto_increment,
    name varchar(255),
    department varchar(255),
    salary float8
);

create table books(
    id bigint primary key auto_increment,
    title varchar(255),
    author varchar(255),
    price float8
);

insert into employees(name,department,salary)values ('张三','技术部',1500.34), ('张一','技术部',20000), ('张二','技术部',1500.34);

insert into books(title, author, price) values ('go语言','李思',58.04),('go语言从入门到精通','李思',58.04),('java语言','李思',100);
*/
var db *sqlx.DB

func init() {
	var connectionString string = "root:123456@tcp(127.0.0.1:3307)/gorm_example?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = sqlx.Connect("mysql", connectionString)
	if err != nil {
		fmt.Println("连接数据库异常", err)
		return
	}
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(10)
	return
}

type Employee struct {
	Id         int64
	Name       string
	Department string
	Salary     string
}

type Book struct {
	Id     int64
	Title  string
	Author string
	Price  string
}

func GetEmployeeByDepartment() {

	query_sql := "select * from employees where department=?"
	var employees []Employee
	err := db.Select(&employees, query_sql, "技术部")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(employees)

}

func GetMaxSalaryEmployee() {

	query_sql := "select * from employees where salary = (select max(salary) from employees)"
	var e Employee
	err := db.Get(&e, query_sql)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(e)

}

func GetBookByPrice() {
	query_sql := "select * from books where price >= ?"
	var books []Book
	err := db.Select(&books, query_sql, 50)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(books)
}
