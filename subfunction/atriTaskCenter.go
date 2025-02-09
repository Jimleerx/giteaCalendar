/*
 *  Powered By Luckykeeper <luckykeeper@luckykeeper.site | https://luckykeeper.site | https://github.com/luckykeeper>
 *  @CreateTime   : 2025-2-8 20:05
 *  @Author       : Luckykeeper
 *  @Contact      : https://github.com/luckykeeper | https://luckykeeper.site
 *  @Email        : luckykeeper@luckykeeper.site
 *  @project      : giteaCalendar
 */

package subfunction

import (
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
)

var (
	AtriTaskCenter *cron.Cron
)

func AtriTaskInit() {
	logrus.Infoln(LogMarkIsChild(), "Initializing AtriTaskCenter During Early Start!")
	AtriTaskCenter = cron.New()
	AtriTaskCenter.Start()

	logrus.Infoln(LogMarkIsChild(), "[AtriTaskInit]添加定时获取 gitea 数据任务")
	err := AtriTaskCenter.AddFunc("0 */5 * * * ?", atriGetStatistics)
	if err != nil {
		logrus.Errorln(LogMarkIsChild(), "添加定时任务失败:", err)
	}

	//go func() {
	//	if model.GiteaCalendarConfig.DebugMode {
	//		atriGetStatistics()
	//	}
	//}()

	logrus.Infoln(LogMarkIsChild(), "Initializing AtriTaskCenter Succeeded!")
}
