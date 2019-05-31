package test

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shopspring/decimal"
	"github.com/tietang/dbx"
	"net/http"
	_ "net/http/pprof"
	"time"
)

var db *dbx.Database

func init() {

	settings := dbx.Settings{
		DriverName: "mysql",
		User:       "root",
		Password:   "111111",
		//Password:   "123456",
		//Host: "192.168.232.175:3306",
		Host:            "172.16.1.248:3306",
		Database:        "po0",
		MaxOpenConns:    10,
		MaxIdleConns:    5,
		ConnMaxLifetime: time.Minute * 30,
		LoggingEnabled:  false,
		Options: map[string]string{
			"charset":              "utf8",
			"parseTime":            "true",
			"autocommit":           "true",
			"allowNativePasswords": "true",
		},
	}
	var err error
	db, err = dbx.Open(settings)
	if err != nil {
		fmt.Println(err)
	}
	db.SetLogging(false)
	db.RegisterTable(&GoodsSigned{}, "goods")
	db.RegisterTable(&GoodsSigned2{}, "red_envelope_goods3")
	db.RegisterTable(&GoodsUnsigned{}, "goods_unsigned")
	pprof()
}

//运行pprof分析器
func pprof() {
	go func() {
		fmt.Println(http.ListenAndServe(":16060", nil))
	}()
}

//事务锁方案
var query = "select * from goods  where envelope_no=? for update"
var update = "update goods  set remain_amount=?,remain_quantity=? where envelope_no=? "

func UpdateForLock(g *Goods) {
	//通过db.tx函数构建事务锁代码块
	err := db.Tx(func(runner *dbx.TxRunner) error {
		//第一步：锁定需要修改的资源，也就是需要修改的数据行
		//编写事务锁查询语句，使用for update子句来锁定资源

		out := &GoodsSigned{}
		_, err := runner.Get(out, query, g.EnvelopeNo)
		if err != nil {
			return err
		}
		//第二部：计算剩余金额和剩余数量
		subAmount := decimal.NewFromFloat(0.01)
		remainAmount := out.RemainAmount.Sub(subAmount)
		remainQuantity := out.RemainQuantity - 1
		//第三步：执行更新

		_, row, err := runner.Execute(update, remainAmount, remainQuantity, g.EnvelopeNo)
		if err != nil {
			return err
		}
		if row < 1 {
			return errors.New("库存扣减失败")
		}
		return nil
	})
	//BenchmarkConcurrentUpdateForLock-4   	    1000	  51071139 ns/op
	//BenchmarkConcurrentUpdateForLock-4   	    1000	  45212671 ns/op

	if err != nil {
		fmt.Println(err)
	}
}

//数据库无符号类型+直接更新方案
func UpdateForUnsigned(g *Goods) {
	update := "update goods_unsigned " +
		"set remain_amount=remain_amount-?, " +
		"remain_quantity=remain_quantity-? " +
		"where envelope_no=?"
	_, row, err := db.Execute(update, 0.01, 1, g.EnvelopeNo)
	//db.Execute(update, 0.01, 1, g.EnvelopeNo)
	//BenchmarkConcurrentUpdateForUnsigned-4   	  500000	    140751 ns/op

	if err != nil {
		fmt.Println(err)
	}
	if row < 1 {
		fmt.Println("扣减失败")
	}
}

//乐观锁方案
func UpdateForOptimistic(g *Goods) {
	update := "update goods " +
		"set remain_amount=remain_amount-?, " +
		" remain_quantity=remain_quantity-? " +
		" where envelope_no=? " +
		" and remain_amount>=? " +
		" and remain_quantity>=? "
	_, row, err := db.Execute(update, 0.01, 1, g.EnvelopeNo, 0.01, 1)
	if err != nil {
		fmt.Println(err)
	}
	if row < 1 {
		fmt.Println("扣减失败")
	}
}

//乐观锁+无符号字段双保险方案
func UpdateForOptimisticAndUnsigned(g *Goods) {
	update := "update goods_unsigned " +
		"set remain_amount=remain_amount-?, " +
		" remain_quantity=remain_quantity-? " +
		" where envelope_no=? " +
		" and remain_amount>=? " +
		" and remain_quantity>=? "
	_, row, err := db.Execute(update, 0.01, 1, g.EnvelopeNo, 0.01, 1)
	if err != nil {
		fmt.Println(err)
	}
	if row < 1 {
		fmt.Println("扣减失败")
	}
}
