package sqlx

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
// ● 要求 ：
//   ○ 定义一个 Book 结构体，包含与 books 表对应的字段。
//   ○ 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。

// Book 结构体，对应books表的字段
type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func Sqlx2TestDemo() {

	// 连接数据库
	db, err := sqlx.Connect("mysql", "username:password@tcp(localhost:3306)/database_name?parseTime=true")
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	defer db.Close()

	// 测试数据库连接
	err = db.Ping()
	if err != nil {
		log.Fatalf("数据库连接测试失败: %v", err)
	}

	fmt.Println("数据库连接成功")

	// 执行复杂查询：查询价格大于50元的书籍
	books, err := queryExpensiveBooks(db, 50.0)
	if err != nil {
		log.Fatalf("查询失败: %v", err)
	}

	// 输出结果
	fmt.Printf("\n价格大于50元的书籍 (%d本):\n", len(books))
	for i, book := range books {
		fmt.Printf("%d. ID: %d, 书名: 《%s》, 作者: %s, 价格: ￥%.2f\n",
			i+1, book.ID, book.Title, book.Author, book.Price)
	}

}

// queryExpensiveBooks 查询价格大于指定值的书籍
func queryExpensiveBooks(db *sqlx.DB, minPrice float64) ([]Book, error) {
	var books []Book

	// 使用命名参数进行类型安全的查询
	query := `
		SELECT id, title, author, price 
		FROM books 
		WHERE price > ? 
		ORDER BY price DESC
	`

	// 执行查询并将结果映射到Book结构体切片中
	err := db.Select(&books, query, minPrice)
	if err != nil {
		return nil, fmt.Errorf("查询执行失败: %v", err)
	}

	return books, nil
}

// 其他可能的查询函数示例

// QueryBooksByAuthor 按作者查询书籍
func QueryBooksByAuthor(db *sqlx.DB, author string) ([]Book, error) {
	var books []Book

	query := `
		SELECT id, title, author, price 
		FROM books 
		WHERE author = ? 
		ORDER BY title
	`

	err := db.Select(&books, query, author)
	if err != nil {
		return nil, fmt.Errorf("按作者查询失败: %v", err)
	}

	return books, nil
}

// QueryBooksInPriceRange 查询价格区间的书籍
func QueryBooksInPriceRange(db *sqlx.DB, minPrice, maxPrice float64) ([]Book, error) {
	var books []Book

	query := `
		SELECT id, title, author, price 
		FROM books 
		WHERE price BETWEEN ? AND ? 
		ORDER BY price ASC
	`

	err := db.Select(&books, query, minPrice, maxPrice)
	if err != nil {
		return nil, fmt.Errorf("价格区间查询失败: %v", err)
	}

	return books, nil
}
