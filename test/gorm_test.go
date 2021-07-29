// https://github.com/go-gorm/gorm/tree/master/tests

package test

import (
	"fmt"
	"github.com/zs368/gin-example/internal/app/models"
	"gorm.io/gorm/clause"
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=%s",
		"root", "root", "127.0.0.1:33061", "gin_example", "utf8mb4", true, "Local")
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("init.setDatabase err: %v", err)
	}
}

func TestCreate(t *testing.T) {
	// article := models.Article{Title: "123", Desc: "456", Content: "789"}

	// result := db.Debug().Create(&article)	// INSERT INTO `article` (`created_at`,`updated_at`,`deleted_at`,`title`,`desc`,`cover_image_url`,`content`,`state`,`created_by`,`updated_by`) VALUES ('2021-07-29 15:59:13.005','2021-07-29 15:59:13.005',NULL,'123','456','','789',0,'','')

	// result := db.Debug().Select("Title", "Desc").Create(&article)	// INSERT INTO `article` (`created_at`,`updated_at`,`title`,`desc`) VALUES ('2021-07-29 16:03:29.589','2021-07-29 16:03:29.589','123','456')

	// result := db.Debug().Session(&gorm.Session{SkipHooks: true}).Create(&article)

	// var articles = []models.Article{{Title: "999"}, {Title: "888"}}

	// result := db.Debug().Create(&articles)

	// result := db.Debug().CreateInBatches(&articles, 1)

	// 注意: map 不会触发钩子
	// result := db.Debug().Model(&models.Article{}).Create(map[string]interface{}{"Title": "zs_test"})	// INSERT INTO `article` (`title`) VALUES ('zs_test')

	// sql 表达式
	// 1、map[string]interface{}  2、自定义数据类型 (os: 感觉就是把 sql 封装一下)
	// result := db.Debug().Model(&models.Article{}).Create(map[string]interface{}{
	// 	"Title": "zs_test2",
	// 	"Desc":  clause.Expr{SQL: "LEFT(?, ?)", Vars: []interface{}{"hahahahaha", "6"}},
	// })	// INSERT INTO `article` (`desc`,`title`) VALUES (LEFT('hahahahaha', '6'),'zs_test2')

	// 测试冲突
	// DoNothing: true  ->  ON DUPLICATE KEY UPDATE `id`=`id`  (os: 无变化 FIXME 感觉这儿有点不合理)

	result := db.Debug().Clauses(clause.OnConflict{
		// DoNothing: true,

		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"title": "测试重复ID"}),
	}).Model(&models.Article{}).Create(map[string]interface{}{
		"ID": 23,
		"Title": "测试重复ID",
	})

	if result.Error != nil {
		fmt.Println(result.Error)
	}
	if result.RowsAffected > 0 {
		fmt.Printf("影响记录的条数为：%d\n", result.RowsAffected)
	}
}
