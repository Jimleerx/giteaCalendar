/*
 *  Powered By Luckykeeper <luckykeeper@luckykeeper.site | https://luckykeeper.site | https://github.com/luckykeeper>
 *  @CreateTime   : 2025-1-23 10:45
 *  @Author       : Luckykeeper
 *  @Contact      : https://github.com/luckykeeper | https://luckykeeper.site
 *  @Email        : luckykeeper@luckykeeper.site
 *  @project      : giteaCalendar
 */

package model

import "time"

type GiteaCalendar struct {
	Id          uint64    `xorm:"pk autoincr 'Id' comment('ID')"`
	Date        time.Time `xorm:"'Date' comment('日期')"`
	Contributes uint64    `xorm:"'Contributes' comment('贡献数')"`
}
