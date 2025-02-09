/*
 *  Powered By Luckykeeper <luckykeeper@luckykeeper.site | https://luckykeeper.site | https://github.com/luckykeeper>
 *  @CreateTime   : 2025-1-22 17:23
 *  @Author       : Luckykeeper
 *  @Contact      : https://github.com/luckykeeper | https://luckykeeper.site
 *  @Email        : luckykeeper@luckykeeper.site
 *  @project      : giteaCalendar
 */

package model

var GiteaCalendarConfig Config

type Config struct {
	DebugMode    bool   `json:"debugMode" yaml:"DebugMode"`
	LogLevel     uint32 `json:"logLevel" yaml:"LogLevel"`
	ApiPort      string `json:"apiPort" yaml:"ApiPort"`
	DatabaseType string `json:"dataBaseType" yaml:"DataBaseType"`
	DatabaseDsn  string `json:"databaseDsn" yaml:"DatabaseDsn"`

	Server   string `json:"server" yaml:"Server"`
	UserName string `json:"userName" yaml:"UserName"`
	ApiKey   string `json:"apiKey" yaml:"apiKey"`
}
