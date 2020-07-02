package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	sql.Register("myMysql", &mysql.MySQLDriver{})
	// 连接数据库, driverName数据库驱动名称，dataSourceName数据库连接信息
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	// 关闭数据库
	defer db.Close()
	// 检查与数据库的连接是否仍有效，如果需要会创建连接
	fmt.Println(db.Ping())
	// 设置与数据库建立连接的最大数目
	// 如果n大于0且小于最大闲置连接数，会将最大闲置连接数减小到匹配最大开启连接数的限制
	// 如果n <= 0，不会限制最大开启连接数，默认为0（无限制）
	db.SetMaxOpenConns(0)
	// 设置连接池中的最大闲置连接数
	// 如果n大于最大开启连接数，则新的最大闲置连接数会减小到匹配最大开启连接数的限制
	// 如果n <= 0，不会保留闲置连接
	db.SetMaxIdleConns(0)
	// 设置连接最大生存时间
	// 过期的连接可能会在重用之前被缓慢地关闭
	// 如果d <= 0，连接将永久重用
	db.SetConnMaxLifetime(0)
	// 查询
	query(db)
	// 执行一次命令，不返回任何结果
	exec(db)
	// 事务
	transaction(db)
}

func query(db *sql.DB) {
	// 查询
	// 等同于 db.QueryContext(context.Background(), "select nickname from users where id=?", 1)
	rows, err := db.Query("SELECT nickname FROM users WHERE id=?", 1)
	if err != nil {
		log.Fatal(err)
	}
	res := make([]string, 0)
	// 循环取值
	for rows.Next() {
		var nickname string
		// 赋值，rows中的列的数量必须与值的数量相同
		if err := rows.Scan(&nickname); err != nil {
			log.Fatal(err)
		}
		res = append(res, nickname)
	}
	// 关闭查询
	if err := rows.Close(); err != nil {
		log.Fatal(err)
	}
	// 返回扫描赋值时的最后个错误
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

func exec(db *sql.DB) {
	// 执行一次命令（包括查询、删除、更新、插入等），不返回任何执行结果。参数args表示query中的占位参数
	// 等同于 db.ExecContext(context.Background(), `select nickname from users where id = 1`)
	res, err := db.Exec(`SELECT nickname FROM users WHERE id = 1`)
	if err != nil {
		log.Fatal(err)
	}
	// 返回一个数据库生成的回应命令的整数
	// 当插入新行时，一般来自一个"自增ID"
	// 不是所有的数据库都支持该功能，该状态的语法也各有不同
	fmt.Println(res.LastInsertId())
	// 返回被update、insert或delete命令影响的行数
	// 不是所有的数据库都支持该功能
	fmt.Println(res.RowsAffected())
}

func transaction(db *sql.DB) {
	// 开启事务
	// 等同于 db.BeginTx(context.Background(), nil)
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	// 使用提供的参数执行准备好的查询状态，返回Rows类型查询结果
	// 其它Query, Exec等用法参见db相关方法
	row := tx.QueryRow("SELECT nickname FROM users WHERE id = ?", 1)
	var name string
	if err := row.Scan(&name); err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
	// 提交事务并关闭
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
	// 回滚事务并关闭
	//tx.Rollback()
}
