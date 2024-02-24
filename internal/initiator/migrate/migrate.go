package migrate

import (
	"github.com/limes-cloud/kratosx"
	gte "github.com/limes-cloud/kratosx/library/db/gormtranserror"

	"party-affairs/config"
	"party-affairs/internal/biz"
)

func IsInit(ctx kratosx.Context) bool {
	db := ctx.DB().Migrator()
	return db.HasTable(biz.Banner{}) &&
		db.HasTable(biz.NewsContent{}) &&
		db.HasTable(biz.NewsClassify{}) &&
		db.HasTable(biz.ResourceContent{}) &&
		db.HasTable(biz.ResourceClassify{})
	// &&
	//	db.HasTable(biz.Task{}) &&
	//	db.HasTable(biz.TaskValue{})
}

func Init(ctx kratosx.Context, _ *config.Config) {
	db := ctx.DB()
	_ = db.Set("gorm:table_options", "COMMENT='通知' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.Notice{})
	_ = db.Set("gorm:table_options", "COMMENT='通知用户' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.NoticeUser{})

	_ = db.Set("gorm:table_options", "COMMENT='轮播' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.Banner{})
	_ = db.Set("gorm:table_options", "COMMENT='新闻分类' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.NewsClassify{})
	_ = db.Set("gorm:table_options", "COMMENT='新闻内容' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.NewsContent{})
	_ = db.Set("gorm:table_options", "COMMENT='新闻评论' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.NewsComment{})
	_ = db.Set("gorm:table_options", "COMMENT='资料分类' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.ResourceClassify{})
	_ = db.Set("gorm:table_options", "COMMENT='资料' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.ResourceContent{})
	_ = db.Set("gorm:table_options", "COMMENT='任务' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.Task{})
	_ = db.Set("gorm:table_options", "COMMENT='任务值' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.TaskValue{})
	_ = db.Set("gorm:table_options", "COMMENT='视频分类' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.VideoClassify{})
	_ = db.Set("gorm:table_options", "COMMENT='视频内容' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.VideoContent{})
	_ = db.Set("gorm:table_options", "COMMENT='用户视频进度' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.UserVideoProcess{})

	// 重新载入gorm错误插件
	_ = gte.NewGlobalGormErrorPlugin().Initialize(ctx.DB())
}
