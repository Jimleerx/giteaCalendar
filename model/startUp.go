/*
 *  Powered By Luckykeeper <luckykeeper@luckykeeper.site | https://luckykeeper.site | https://github.com/luckykeeper>
 *  @CreateTime   : 2025-1-22 17:23
 *  @Author       : Luckykeeper
 *  @Contact      : https://github.com/luckykeeper | https://luckykeeper.site
 *  @Email        : luckykeeper@luckykeeper.site
 *  @project      : giteaCalendar
 */

package model

import (
	"fmt"
)

func StartUpMessage() {
	var count []rune
	fmt.Println(GiteaCalendarATRILogo)
	fmt.Printf("\t\t\t\t\t\t\t%v\n", StartUpString)
	fmt.Printf("\t\t\t\t\t\t\t=========>Version:%v\n", AppVersion)
	for _, i := range GiteaCalendarATRILogo {
		count = append(count, i)
	}
	for _, i := range StartUpString {
		count = append(count, i)
	}
	fmt.Println("Powered By Luckykeeper <luckykeeper@luckykeeper.site | https://luckykeeper.site | https://github.com/luckykeeper>")
}
