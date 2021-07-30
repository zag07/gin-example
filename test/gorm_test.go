// https://github.com/go-gorm/gorm/tree/master/tests

package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/zs368/gin-example/internal/app/models"
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

	// result := db.Debug().Create(&article)	// INSERT INTO `article` (`created_at`,`updated_at`,`deleted_at`,`title`,`desc`,`cover_image_url`,`content`,`status`,`created_by`,`updated_by`) VALUES ('2021-07-29 15:59:13.005','2021-07-29 15:59:13.005',NULL,'123','456','','789',0,'','')

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
	/*result := db.Debug().Clauses(clause.OnConflict{
		// DoNothing: true,	// ->  ON DUPLICATE KEY UPDATE `id`=`id`  (os: 无变化)

		// Columns:   []clause.Column{{Name: "id"}},
		// DoUpdates: clause.Assignments(map[string]interface{}{"title": "测试重复ID"}),	// 在`id`冲突时，将列更新为默认值

		// Columns:   []clause.Column{{Name: "id"}},
		// DoUpdates: clause.AssignmentColumns([]string{"Title", "Content"}),	// 在`id`冲突时，将列更新为新值

		UpdateAll: true, // 在冲突时，更新除主键以外的所有列到新值
	}).Model(&models.Article{}).Create(map[string]interface{}{
		"ID":      23,
		"Title":   "测试重复ID2",
		"Content": "gogogo",
	})*/

	/*if result.Error != nil {
		fmt.Println(result.Error)
	}
	if result.RowsAffected > 0 {
		fmt.Printf("影响记录的条数为：%d\n", result.RowsAffected)
	}*/
}

func TestQuery(t *testing.T) {
	// var article models.Article
	// var articles []models.Article

	// db.Debug().First(&article)	// SELECT * FROM `article` ORDER BY `article`.`id` LIMIT 1

	// db.Debug().Take(&article)	// SELECT * FROM `article` LIMIT 1

	// db.Debug().Last(&article)	// SELECT * FROM `article` ORDER BY `article`.`id` DESC LIMIT 1

	// First, Last方法将按主键排序查找第一/最后一条记录，只有在用struct查询或提供model value时才有效，如果当前model没有定义主键，将按第一个字段排序
	// Take 可以配合 Table 方法使用，不需要提供model
	// result := map[string]interface{}{"ID": ""}

	// db.Debug().Table("article").Take(&result)

	// fmt.Printf("%+v\n", result)

	// db.Debug().Find(&article, 5)	// SELECT * FROM `article` WHERE `article`.`id` = 5

	// db.Debug().Find(&article, "id = ?", "5")	// SELECT * FROM `article` WHERE id = '5'

	// db.Debug().Where("id = ?", 5).Find(&article) // SELECT * FROM `article` WHERE id = 5

	// db.Debug().Where("id in ?", []int{5, 6}).Find(&articles)

	// 注意: 使用 struct 作为条件查询时，GORM 只会查询非零值字段，查询条件为0的话可以参考下面第二条写法
	// db.Debug().Where(&models.Article{Title: "zs_test", Status: 0}).Find(&article)	// SELECT * FROM `article` WHERE `article`.`title` = 'zs_test'

	// db.Debug().Where(&models.Article{Title: "zs_test"}, "title", "status").Find(&article) // SELECT * FROM `article` WHERE `article`.`title` = 'zs_test' AND `article`.`status` = 0

	// db.Debug().Where(map[string]interface{}{"title": "zs_test", "status": 0}).Find(&article) // SELECT * FROM `article` WHERE `status` = 0 AND `title` = 'zs_test'

	// 內联条件 (os: 其实就是把 where 放在里面)
	// db.Debug().Find(&article, map[string]interface{}{"title": "zs_test"})

	// Not 条件  		.Not()
	// Or 条件 			.Or()
	// 选择特定字段		.Select()
	// Order			.Order()
	// Limit & Offset	.Limit().Offset()
	// Group & Having	.Group().Having()
	// Distinct			.Distinct()

	// Joins (os: 这种方式有点啰嗦)
	// db.Debug().Model(&models.Article{}).Select("article.title").
	// 	Joins("left join article_tag b on article.id = b.article_id").Find(&article, "b.id is not null")

	// fmt.Printf("%+v\n", article)
}

func TestAdvancedQuery(t *testing.T) {
	var article models.Article

	// Smart Select Fields
	/*result := struct {
		ID    string
		Title string
	}{}

	db.Debug().Model(&models.Article{}).Limit(1).Find(&result)

	fmt.Printf("%+v\n", result)*/

	// Locking (FOR UPDATE)
	/*db.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).Find(&article)	// SELECT * FROM `article` WHERE `article`.`deleted_at` IS NULL FOR UPDATE

	db.Debug().Clauses(clause.Locking{
		Strength: "SHARE",
		Table: clause.Table{Name: clause.CurrentTable},
	}).Find(&article)*/ // SELECT * FROM `article` WHERE `article`.`deleted_at` IS NULL AND `article`.`id` = 5 FOR SHARE OF `article`

	// SubQuery
	// db.Debug().Where("id > (?)", db.Table("article").Select("AVG(id)")).Find(&article)

	// From SubQuery
	// db.Debug().Table("(?) as a", db.Model(models.Article{})).
	// 	Select("title").Find(&article, 5) // SELECT `title` FROM (SELECT * FROM `article`) as a WHERE `a`.`id` = 5

	// Group Conditions
	/*db.Debug().Where(
		db.Where("").Or(""),
	).Or(
		db.Where("").Where(""),
	).Find(&article)*/

	// Named Argument
	// Find To Map
	// FirstOrInit
	// FirstOrCreate
	// Optimizer/Index Hints

	// Iteration (os: 只有一条sql)
	/*rows, _ := db.Debug().Model(&models.Article{}).Where("title = ?", "123").Rows()
	defer rows.Close()

	for rows.Next() {
		db.ScanRows(rows, &article)
		fmt.Printf("%+v\n", article)
	}*/

	// FindInBatches (os: 这个好像只支持传入model)
	/*var results []models.Article
	// var results []map[string]interface{}	// panic: reflect: call of reflect.Value.Field on map Value
	_ = db.Debug().Model(&models.Article{}).Where("title = ?", "123").
		Select("id", "title", "created_at").
		FindInBatches(&results, 3, func(tx *gorm.DB, batch int) error {
			fmt.Printf("results: %+v\n", results)

			fmt.Printf("tx.RowsAffected: %v\n", tx.RowsAffected)

			fmt.Printf("batch: %v\n", batch)

			return nil
		})*/

	// Pluck
	/*var titles []string
	db.Debug().Model(&models.Article{}).Pluck("title", &titles)
	fmt.Printf("%+v\n", titles)*/

	var mulcoulmns []map[string]interface{}
	db.Debug().Model(&models.Article{}).Select("id", "title").Find(&mulcoulmns)
	fmt.Printf("%v\n", mulcoulmns)

	// Scopes: allows you to specify commonly-used queries which can be referenced as method calls

	// Count: get matched records count

	fmt.Printf("%+v\n", article)
}

func TestUpdate(t *testing.T) {

}
