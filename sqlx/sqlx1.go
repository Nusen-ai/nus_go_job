package sqlx

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // 或其他数据库驱动
	"github.com/jmoiron/sqlx"
)

// 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
// ● 要求 ：
//   ○ 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
//   ○ 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。

// Employee 结构体映射数据库表
type Employee struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     int    `db:"salary"`
}

func Sqlx1TestDemo() {
	// 连接数据库（这里以MySQL为例）
	db, err := sqlx.Connect("mysql", "username:password@tcp(localhost:3306)/database_name")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 查询1：查询技术部所有员工
	technicalEmployees, err := getEmployeesByDepartment(db, "技术部")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("技术部员工信息：")
	for _, emp := range technicalEmployees {
		fmt.Printf("ID: %d, Name: %s, Department: %s, Salary: %d\n",
			emp.ID, emp.Name, emp.Department, emp.Salary)
	}

	// 查询2：查询工资最高的员工
	highestPaidEmployee, err := getHighestPaidEmployee(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n工资最高的员工：\nID: %d, Name: %s, Department: %s, Salary: %d\n",
		highestPaidEmployee.ID, highestPaidEmployee.Name,
		highestPaidEmployee.Department, highestPaidEmployee.Salary)
}

// getEmployeesByDepartment 查询指定部门的所有员工
func getEmployeesByDepartment(db *sqlx.DB, department string) ([]Employee, error) {
	var employees []Employee

	query := "SELECT id, name, department, salary FROM employees WHERE department = ?"
	err := db.Select(&employees, query, department)
	if err != nil {
		return nil, err
	}

	return employees, nil
}

// getHighestPaidEmployee 查询工资最高的员工
func getHighestPaidEmployee(db *sqlx.DB) (*Employee, error) {
	var employee Employee

	query := "SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1"
	err := db.Get(&employee, query)
	if err != nil {
		return nil, err
	}

	return &employee, nil
}
