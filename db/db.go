package db

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

/*
	host: "127.0.0.1"
	port: "3306"
	dbname: "test"
	user: "root"
	password: "123456"
	charset : "utf8"

	exmple: db.DbInit("root", "123456", "127.0.0.1", "3306", "test", "utf8")
*/
func DbInit(
	user,
	pw,
	host,
	port,
	dbname,
	charset string) (*sql.DB, error) {

	DB_Driver := user + ":" + pw + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=" + charset
	var err error
	db, err = sql.Open("mysql", DB_Driver)
	if err != nil {
		return nil, err
	} else {
		db.SetMaxOpenConns(100)
		// 闲置连接数
		db.SetMaxIdleConns(20)
		// 最大连接周期
		db.SetConnMaxLifetime(100 * time.Second)
		return db, nil
	}
}

/*
input:
	"select * from user where id < ?"  3

	exmple: data := db.Select("select * from user where id = ?", 1);

output:
	[0] =>{"id":1, "name": "小明"}
	[1] =>{"id":2, "name": "小李"}

	exmple: fmt.Printf(data[0]["name"]);
*/
func Select(strSql string, arg ...interface{}) []map[string]string {
	rows, err := db.Query(strSql, arg...)
	if err != nil {
		fmt.Printf("读取数据失败:" + err.Error())
		return nil
	}
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	//获取表信息
	colTypes, colNames, length := getRowInfo(rows)
	//创建列对应变量
	values := make([]interface{}, 0)
	for i := 0; i < length; i++ {
		values = append(values, reflect.New(colTypes[i]).Interface())
	}
	//返回的结果
	res := make([]map[string]string, 0)
	for rows.Next() {
		rows.Scan(values...)
		row := make(map[string]string)
		for i := 0; i < length; i++ {
			if colTypes[i].Name() == "RawBytes" {
				row[colNames[i]] = string(*values[i].(*sql.RawBytes))
			} else if colTypes[i].Name() == "int64" {
				row[colNames[i]] = strconv.FormatInt(*values[i].(*int64), 10)
			}
		}
		res = append(res, row)
	}
	return res
}

/*
input:
	INSERT INTO test.`user`(name) VALUES('?');  "miz"

	exmple: 	db.Insert("INSERT INTO test.`user`(name) VALUES('miz');")
*/
func Insert(strSql string, arg ...interface{}) {
	rows, err := db.Query(strSql, arg...)
	rows.Close()
	if err != nil {
		fmt.Println("插入数据失败:" + err.Error())
	} else {
		fmt.Println("插入数据成功：")
	}
}

func getRowInfo(rows *sql.Rows) ([]reflect.Type, []string, int) {
	//获取表列类型
	sqlTypes, _ := rows.ColumnTypes()
	length := len(sqlTypes)
	columnTypes := []reflect.Type{}
	for i := 0; i < length; i++ {
		columnTypes = append(columnTypes, sqlTypes[i].ScanType())
	}
	//获取列名
	colNames, _ := rows.Columns()
	return columnTypes, colNames, length
}
