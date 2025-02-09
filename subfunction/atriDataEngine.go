/*
 *  Powered By Luckykeeper <luckykeeper@luckykeeper.site | https://luckykeeper.site | https://github.com/luckykeeper>
 *  @CreateTime   : 2025-1-22 17:54
 *  @Author       : Luckykeeper
 *  @Contact      : https://github.com/luckykeeper | https://luckykeeper.site
 *  @Email        : luckykeeper@luckykeeper.site
 *  @project      : giteaCalendar
 */

package subfunction

import (
	"giteaCalendar/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	_ "modernc.org/sqlite"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/caches"
)

var (
	AtriDataEngine *xorm.Engine
)

func InitializeDatabase() {

	var err error
	AtriDataEngine, err = xorm.NewEngine(model.GiteaCalendarConfig.DatabaseType, model.GiteaCalendarConfig.DatabaseDsn)
	if err != nil {
		logrus.Fatalln(LogMarkIsChild(), "数据库初始化失败：", err)
		return
	}

	AtriDataEngine.Dialect().SetQuotePolicy(1)

	cacher := caches.NewLRUCacher(caches.NewMemoryStore(), 1000)
	AtriDataEngine.SetDefaultCacher(cacher)

	AtriDataEngine.ShowSQL(model.GiteaCalendarConfig.DebugMode)
	//AtriDataEngine.ShowSQL(true)

	err = AtriDataEngine.Ping()
	if err != nil {
		logrus.Fatalln(LogMarkIsChild(), "数据库连接失败！检查数据库信息是否正确，程序返回原因为：", err)
	} else {
		AtriDataEngine.SetConnMaxLifetime(time.Second * 60)
		AtriDataEngine.SetMaxOpenConns(50)
		AtriDataEngine.SetMaxIdleConns(3)
	}
	if !fiber.IsChild() {
		syncDataTable()
	}
}

func syncDataTable() {
	logrus.Infoln(LogMarkIsChild(), "同步表结构...")
	logrus.Infoln(LogMarkIsChild(), "同步表结构：GiteaCalendar")
	err := AtriDataEngine.Table("GitLabCalendar").Sync(new(model.GiteaCalendar))
	if err != nil {
		logrus.Errorln(LogMarkIsChild(), "同步表结构失败(model.GiteaCalendar)，原因为：", err)
		return
	}
	logrus.Infoln(LogMarkIsChild(), "同步表结构完成")
}
