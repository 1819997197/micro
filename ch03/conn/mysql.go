package conn

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"micro/ch03/config"
)

var SqlDB *gorm.DB

func InitMysql(c *config.MysqlConfig) error {
	var err error
	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", c.UserName, c.Password, c.ServerName, c.ServerPort, c.DbName, c.Charset)
	SqlDB, err = gorm.Open("mysql", args)
	if err != nil {
		fmt.Println("sql.Open err:", err)
		return err
	}
	SqlDB.LogMode(true) //打开调试模式

	return nil
}
