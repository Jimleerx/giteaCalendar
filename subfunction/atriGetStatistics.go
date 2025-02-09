/*
 *  Powered By Luckykeeper <luckykeeper@luckykeeper.site | https://luckykeeper.site | https://github.com/luckykeeper>
 *  @CreateTime   : 2025-2-8 20:11
 *  @Author       : Luckykeeper
 *  @Contact      : https://github.com/luckykeeper | https://luckykeeper.site
 *  @Email        : luckykeeper@luckykeeper.site
 *  @project      : giteaCalendar
 */

package subfunction

import (
	"code.gitea.io/sdk/gitea"
	"encoding/json"
	"giteaCalendar/model"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"time"
)

func atriGetStatistics() {
	logrus.Infoln(LogMarkIsChild(), "定时任务开始，获取 Gitea 相关信息")

	atriClient, err := gitea.NewClient(model.GiteaCalendarConfig.Server, gitea.SetToken(model.GiteaCalendarConfig.ApiKey))
	if err != nil {
		logrus.Errorln(LogMarkIsChild(), "构造 Gitea 客户端失败，错误信息:", err)
		return
	}
	user, resp, err := atriClient.GetUserInfo(model.GiteaCalendarConfig.UserName)
	if err != nil {
		logrus.Errorln(LogMarkIsChild(), "获取用户信息失败，错误信息:", err)
		return
	}
	logrus.Infoln(LogMarkIsChild(), "HTTP CODE:", resp.StatusCode, "/", resp.Status)
	if resp.StatusCode != 200 {
		logrus.Errorln(LogMarkIsChild(), "服务器返回了非 200 的状态码，等待下次尝试运行")
		return
	}
	logrus.Infoln(LogMarkIsChild(), "用户名:", user.UserName)
	logrus.Infoln(LogMarkIsChild(), "用户 ID :", user.ID)

	atriHttpClient := resty.New()

	hcResp, hcErr := atriHttpClient.R().
		SetHeaders(map[string]string{"Authorization": "token " + model.GiteaCalendarConfig.ApiKey}).
		Get(model.GiteaCalendarConfig.Server + "/api/v1/users/" + user.UserName + "/heatmap")
	if hcErr != nil {
		logrus.Errorln(LogMarkIsChild(), "获取用户提交信息失败:", err)
		return
	}

	logrus.Infoln(LogMarkIsChild(), "HTTP CODE:", resp.StatusCode, "/", resp.Status)
	if hcResp.StatusCode() != 200 {
		logrus.Errorln(LogMarkIsChild(), "服务器返回了非 200 的状态码，等待下次尝试运行")
		return
	}
	var serverReturn []model.GiteaHeatMap
	if err := json.Unmarshal(hcResp.Body(), &serverReturn); err != nil {
		logrus.Errorln(LogMarkIsChild(), "解析服务器返回的 json 异常:", err)
		return
	}
	logrus.Debugln(LogMarkIsChild(), "贡献 heatmap data:", serverReturn)

	var calendarDataList []model.GiteaCalendar
	for _, heatMap := range serverReturn {
		var (
			calendarData model.GiteaCalendar
			findMark     = false
		)
		calendarData.Contributes = uint64(heatMap.Contributions)
		date := time.Unix(int64(heatMap.Timestamp), 8).UTC().Format(time.DateOnly)
		calendarData.Date, err = time.Parse(time.DateOnly, date)
		if err != nil {
			logrus.Errorln(LogMarkIsChild(), "转换 heatmap 结果失败:", err)
			return
		}
		logrus.Debugln(LogMarkIsChild(), "日期:", calendarData.Date)
		logrus.Debugln(LogMarkIsChild(), "Contributes:", calendarData.Contributes)

		for seq, thisCalendar := range calendarDataList {
			if calendarData.Date.Equal(thisCalendar.Date) {
				findMark = true
				calendarDataList[seq].Contributes += calendarData.Contributes
			}
		}

		if !findMark {
			calendarDataList = append(calendarDataList, calendarData)
		}

	}
	logrus.Infoln(LogMarkIsChild(), "将入库数据（{ID}/日期/Contributes）:", calendarDataList)
	for i := 0; i < len(calendarDataList); i++ {
		hasResult, err := AtriDataEngine.Table("GitLabCalendar").Where("Date = ?", calendarDataList[i].Date.Format(time.DateTime)).Exist(&model.GiteaCalendar{})
		if err != nil {
			logrus.Errorln(LogMarkIsChild(), "判断数据库是否存在应覆盖数据失败:", err)
			return
		}
		if hasResult {
			_, err = AtriDataEngine.Table("GitLabCalendar").Cols("Contributes").Where("Date = ?", calendarDataList[i].Date.Format(time.DateTime)).Update(&calendarDataList[i])
			if err != nil {
				logrus.Errorln(LogMarkIsChild(), "覆盖 Contributes 数据失败:", err)
				return
			}
		} else {
			_, err = AtriDataEngine.Table("GitLabCalendar").Insert(&calendarDataList[i])
			if err != nil {
				logrus.Errorln(LogMarkIsChild(), "写入 Contributes 数据失败:", err)
				return
			}
		}
	}
	logrus.Infoln(LogMarkIsChild(), "本阶段任务成功完成")
}
