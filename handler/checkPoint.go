/*
 *  Powered By Luckykeeper <luckykeeper@luckykeeper.site | https://luckykeeper.site | https://github.com/luckykeeper>
 *  @CreateTime   : 2025-2-8 21:25
 *  @Author       : Luckykeeper
 *  @Contact      : https://github.com/luckykeeper | https://luckykeeper.site
 *  @Email        : luckykeeper@luckykeeper.site
 *  @project      : giteaCalendar
 */

package handler

import (
	"giteaCalendar/model"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func ServiceReady(ctx *fiber.Ctx) error {
	var serviceStatus model.CheckPoint
	serviceStatus.AppName = "giteaCalendar - Powered By Luckykeeper <https://github.com/luckykeeper | https://luckykeeper.site>"
	serviceStatus.Version = model.AppVersion
	serviceStatus.ServerMessage = model.StartUpString
	serviceStatus.StatusCode = http.StatusOK
	return ctx.Status(http.StatusOK).JSON(serviceStatus)
}
