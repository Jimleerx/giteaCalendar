/*
 *  Powered By Luckykeeper <luckykeeper@luckykeeper.site | https://luckykeeper.site | https://github.com/luckykeeper>
 *  @CreateTime   : 2025-2-8 21:25
 *  @Author       : Luckykeeper
 *  @Contact      : https://github.com/luckykeeper | https://luckykeeper.site
 *  @Email        : luckykeeper@luckykeeper.site
 *  @project      : giteaCalendar
 */

package model

type CheckPoint struct {
	StatusCode    int    `json:"statusCode"`
	AppName       string `json:"appName"`
	Version       string `json:"version"`
	ServerMessage string `json:"serverMessage"`
}
