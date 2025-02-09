/*
 *  Powered By Luckykeeper <luckykeeper@luckykeeper.site | https://luckykeeper.site | https://github.com/luckykeeper>
 *  @CreateTime   : 2025-1-22 17:23
 *  @Author       : Luckykeeper
 *  @Contact      : https://github.com/luckykeeper | https://luckykeeper.site
 *  @Email        : luckykeeper@luckykeeper.site
 *  @project      : giteaCalendar
 */

package router

import (
	"giteaCalendar/handler"
	"giteaCalendar/model"
	"giteaCalendar/subfunction"
	assets "giteaCalendar/wwwroot"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	GiteaCalendarServer *fiber.App
)

func StartApiServer() {
	GiteaCalendarServer = fiber.New(fiber.Config{
		ServerHeader:          "giteaCalendar By Luckykeeper<https://github.com/luckykeeper>",
		StrictRouting:         false,
		CaseSensitive:         false,
		UnescapePath:          true,
		BodyLimit:             512 * 1024 * 1024,
		ReadBufferSize:        4 * 1024,
		WriteBufferSize:       4 * 1024,
		AppName:               "giteaCalendar " + model.AppVersion,
		EnableIPValidation:    true,
		DisableStartupMessage: false,
		EnablePrintRoutes:     model.GiteaCalendarConfig.DebugMode,
		Prefork:               !model.GiteaCalendarConfig.DebugMode,
	})

	GiteaCalendarServer.Use(etag.New())

	GiteaCalendarServer.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	GiteaCalendarServer.Use(helmet.New(helmet.Config{
		XSSProtection: "1; mode=block",
	}))

	GiteaCalendarServer.Get("/metrics", monitor.New(monitor.Config{Title: "服务器状态监控 - giteaCalendar", Refresh: 1 * time.Second}))

	GiteaCalendarServer.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(assets.WebappFiles),
		PathPrefix: "webapp",
	}))

	checkPointRouter := GiteaCalendarServer.Group("/checkPoint")
	checkPointRouter.Get("/ready", handler.ServiceReady)
	checkPointRouter.Get("/liveness", handler.ServiceReady)

	apiRouter := GiteaCalendarServer.Group("/api")
	apiRouter.Get("/Calendar", handler.Calendar)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		logrus.Warnln(subfunction.LogMarkIsChild(), "Shutting Down giteaCalendar Gracefully With Max Lifetime 15 Seconds...")
		err := GiteaCalendarServer.ShutdownWithTimeout(15 * time.Second)
		if err != nil {
			logrus.Errorln(subfunction.LogMarkIsChild(), "giteaCalendar Gracefully Shutting Down ==> ", err)
		} else {
			logrus.Infoln(subfunction.LogMarkIsChild(), "giteaCalendar Has Gracefully Shutdown")
		}
	}()

	err := GiteaCalendarServer.Listen(":" + model.GiteaCalendarConfig.ApiPort)
	if err != nil {
		logrus.Fatalln(subfunction.LogMarkIsChild(), "giteaCalendar ListenAndServe @:", model.GiteaCalendarConfig.ApiPort, " err:", err)
	}
	logrus.Infoln(subfunction.LogMarkIsChild(), "Has Completed!")
}
