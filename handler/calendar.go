/*
 *  Powered By Luckykeeper <luckykeeper@luckykeeper.site | https://luckykeeper.site | https://github.com/luckykeeper>
 *  @CreateTime   : 2025-2-8 21:23
 *  @Author       : Luckykeeper
 *  @Contact      : https://github.com/luckykeeper | https://luckykeeper.site
 *  @Email        : luckykeeper@luckykeeper.site
 *  @project      : giteaCalendar
 */

package handler

import (
	"fmt"
	"giteaCalendar/model"
	"giteaCalendar/subfunction"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Calendar(ctx *fiber.Ctx) error {
	var (
		calenderList = make([]model.GiteaCalendar, 0)
		result       string
	)

	err := subfunction.AtriDataEngine.Table("GitLabCalendar").Find(&calenderList)
	if err != nil {
		logrus.Errorln("从数据库获取贡献数据异常:", err)
		return fiber.ErrInternalServerError
	}
	for seq, calender := range calenderList {
		if seq == 0 {
			result += "{"
		}
		result += "\"" + calender.Date.Format(time.DateOnly) + "\":" + fmt.Sprint(calender.Contributes)
		if seq < len(calenderList)-1 {
			result += ","
		}
		if seq == len(calenderList)-1 {
			result += "}"
		}
	}
	return ctx.Status(http.StatusOK).SendString(result)
}
