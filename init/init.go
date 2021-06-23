package init

import (
	"github.com/zs368/gin-example/internal/pkg/database"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB

func init() {
	if err := setupDBEngine(); err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}

func setupDBEngine() error {
	var err error
	Db, err = database.NewDB()
	if err != nil {
		return err
	}

	return nil
}
