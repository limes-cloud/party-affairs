package migrate

import (
	"github.com/limes-cloud/kratosx"
	gte "github.com/limes-cloud/kratosx/library/db/gormtranserror"

	"party-affairs/config"
	"party-affairs/internal/model"
)

func IsInit(ctx kratosx.Context) bool {
	db := ctx.DB().Migrator()
	return db.HasTable(model.Banner{}) &&
		db.HasTable(model.News{}) &&
		db.HasTable(model.NewsClassify{}) &&
		db.HasTable(model.Resource{}) &&
		db.HasTable(model.ResourceClassify{}) &&
		db.HasTable(model.Task{}) &&
		db.HasTable(model.TaskValue{})
}

func Init(ctx kratosx.Context, _ *config.Config) {
	db := ctx.DB()
	_ = db.Set("gorm:table_options", "COMMENT='轮播' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(model.Banner{})
	_ = db.Set("gorm:table_options", "COMMENT='新闻分类' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(model.NewsClassify{})
	_ = db.Set("gorm:table_options", "COMMENT='新闻' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(model.News{})
	_ = db.Set("gorm:table_options", "COMMENT='资料分类' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(model.ResourceClassify{})
	_ = db.Set("gorm:table_options", "COMMENT='资料' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(model.Resource{})
	_ = db.Set("gorm:table_options", "COMMENT='任务' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(model.Task{})
	_ = db.Set("gorm:table_options", "COMMENT='任务值' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(model.TaskValue{})

	// 重新载入gorm错误插件
	_ = gte.NewGlobalGormErrorPlugin().Initialize(ctx.DB())
}
