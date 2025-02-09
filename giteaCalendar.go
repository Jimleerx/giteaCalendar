/*
 *  Powered By Luckykeeper <luckykeeper@luckykeeper.site | https://luckykeeper.site | https://github.com/luckykeeper>
 *  @CreateTime   : 2025-1-22 17:23
 *  @Author       : Luckykeeper
 *  @Contact      : https://github.com/luckykeeper | https://luckykeeper.site
 *  @Email        : luckykeeper@luckykeeper.site
 *  @project      : giteaCalendar
 */

package main

import (
	"giteaCalendar/model"
	"giteaCalendar/router"
	"giteaCalendar/subfunction"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

func init() {
	logrus.SetLevel(logrus.Level(5))
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:   false,
		TimestampFormat: time.RFC3339,
		ForceColors:     true,
		FullTimestamp:   true,
	})
	logrus.SetOutput(os.Stdout)
	writers := []io.Writer{
		os.Stdout}
	stdoutWriter := io.MultiWriter(writers...)
	logrus.SetOutput(stdoutWriter)

	logrus.Infoln(subfunction.LogMarkIsChild(), "Project giteaCalendar Start Initializing...")
	if !fiber.IsChild() {
		model.StartUpMessage()
	}
	subfunction.ReadConfig()
	if model.GiteaCalendarConfig.DebugMode {
		logrus.Warnln("Start As Debug Mode! NOT RECOMMEND IN PRODUCTION ENVIRONMENT!!!")
	}
	logrus.SetLevel(logrus.Level(model.GiteaCalendarConfig.LogLevel))
	subfunction.InitializeDatabase()
	if !fiber.IsChild() {
		subfunction.AtriTaskInit()
	}
}

func main() {
	router.StartApiServer()
}
