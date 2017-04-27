package models

import (
	"database/sql"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

/*
* 初始化数据库
*包括穿件数据库，表，以及插入部分数据
 */
func InitDB() {
	createdb()
	Connect()
	orm.RunSyncdb("default", false, true)

}

/**
* 创建数据库
 */
func createdb() {
	var sqlstring string
	dns, db_name := getConfig(0)
	sqlstring = fmt.Sprintf("CREATE DATABASE  if not exists `%s` CHARSET utf8 COLLATE utf8_general_ci", db_name)
	db, err := sql.Open("mysql", dns)
	if err != nil {
		panic(err.Error())
	}
	r, err := db.Exec(sqlstring)
	if err != nil {
		beego.Error("err is :", err.Error(), "and r is:", r)
	} else {
		beego.Info("Database: ", db_name, " created succes")
	}
	defer db.Close()
}

func Connect() {
	dns, _ := getConfig(1)
	beego.Info("数据库is %s", dns)
	err := orm.RegisterDataBase("default", "mysql", dns)
	if err != nil {
		beego.Error("数据库连接失败")
	} else {
		beego.Info("数据库连接sucess ")
	}
}

/*
* 获取配置
	flag ==1 表示 只链接
	==0 创建 加链接
*/
func getConfig(flag int) (string, string) {
	var dns string
	db_host := beego.AppConfig.String("mysqlurls")
	db_port := beego.AppConfig.String("mysqlport")
	db_user := beego.AppConfig.String("mysqluser")
	db_pass := beego.AppConfig.String("mysqlpass")
	db_name := beego.AppConfig.String("mysqldb")
	if flag == 1 {
		// fmt.Println("链接数据库")
		orm.RegisterDriver("mysql", orm.DRMySQL)
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local", db_user, db_pass, db_host, db_port, db_name)
	} else {
		// fmt.Println("创建数据库")
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)?charset=utf8", db_user, db_pass, db_host, db_port)
	}
	return dns, db_name
}
