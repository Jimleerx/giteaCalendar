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
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func Calendar(ctx *fiber.Ctx) error {
	var (
		calenderList = make([]model.GiteaCalendar, 0)
		result       string
		level        string
	)

	err := subfunction.AtriDataEngine.Table("GitLabCalendar").Find(&calenderList)
	if err != nil {
		logrus.Errorln("从数据库获取贡献数据异常:", err)
		return fiber.ErrInternalServerError
	}
	result = "{\"total\": {\"year\": " + fmt.Sprint(len(calenderList)) + "}, \"contributions\": ["
	for seq, calender := range calenderList {
		if calender.Contributes <= 0 || calender.Contributes%5 == 0 {
			level = "5"
		} else {
			level = fmt.Sprint(calender.Contributes % 5)
		}
		result += "{\"date\":\"" + calender.Date.Format(time.DateOnly) + "\", \"count\":" + fmt.Sprint(calender.Contributes) + ", \"level\":" + level + "}"
		if seq < len(calenderList)-1 {
			result += ","
		}
	}
	result += "]}"
	return ctx.Status(http.StatusOK).SendString(result)
}
