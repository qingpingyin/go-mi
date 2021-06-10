package models

import (
	"MI/pkg/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var (
	Db *gorm.DB
	err error
)

func Setup(){

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.MysqlConf.User,
		setting.MysqlConf.Pwd,
		setting.MysqlConf.Host,
		setting.MysqlConf.Port,
		setting.MysqlConf.Db,
	)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   setting.MysqlConf.Prefix, // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true,                     // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})
	if err != nil {
		log.Printf("Failed to connect database:%v", err)
	}
	sqlDB, _ := Db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(setting.MysqlConf.MaxIdle)
	sqlDB.SetMaxOpenConns(setting.MysqlConf.MaxActive)
	//开启自动迁移
	AutoMigrate()
}
func AutoMigrate(){
	Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT 'device'").AutoMigrate(&Device{})
	Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT 'trace'").AutoMigrate(&Trace{})
	Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT 'users'").AutoMigrate(&Users{})
	Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT 'carousel'").AutoMigrate(&Carousel{})
	Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT 'categories'").AutoMigrate(&Categories{})
	Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT 'product'").AutoMigrate(&Product{})
	Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT 'product_img'").AutoMigrate(&ProductImg{})
}

// 通用分页获取偏移量
func GetOffset(page, pageSize int) int {
	if page <= 1 {
		return 0
	}
	return (page - 1) * pageSize
}
// 设置条件
func MultiWhere(where ...interface{}) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(where[0], where[1:]...)
	}
}