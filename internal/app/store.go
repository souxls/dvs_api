package app

import (
	"github.com/jinzhu/gorm"
	"github.com/souxls/dvs_api/internal/app/config"
	igorm "github.com/souxls/dvs_api/internal/app/model/impl/gorm"
	"go.uber.org/dig"
)

// InitStore 初始化存储
func InitStore(container *dig.Container) (func(), error) {
	var storeCall func()
	cfg := config.Global()

	db, err := initGorm()
	if err != nil {
		return nil, err
	}

	storeCall = func() {
		db.Close()
	}

	igorm.SetTablePrefix(cfg.Gorm.TablePrefix)

	if cfg.Gorm.EnableAutoMigrate {
		err = igorm.AutoMigrate(db)
		if err != nil {
			return nil, err
		}

		// 注入DB
		_ = container.Provide(func() *gorm.DB {
			return db
		})

		_ = igorm.Inject(container)
	}

	return storeCall, nil
}

// initGorm 实例化gorm存储
func initGorm() (*gorm.DB, error) {
	cfg := config.Global()

	var dsn string
	dsn = cfg.MySQL.DSN()

	return igorm.NewDB(&igorm.Config{
		Debug:        cfg.Gorm.Debug,
		DBType:       cfg.Gorm.DBType,
		DSN:          dsn,
		MaxIdleConns: cfg.Gorm.MaxIdleConns,
		MaxLifetime:  cfg.Gorm.MaxLifetime,
		MaxOpenConns: cfg.Gorm.MaxOpenConns,
	})
}
