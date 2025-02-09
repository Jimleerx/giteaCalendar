/*
 *  Powered By Luckykeeper <luckykeeper@luckykeeper.site | https://luckykeeper.site | https://github.com/luckykeeper>
 *  @CreateTime   : 2025-1-22 17:23
 *  @Author       : Luckykeeper
 *  @Contact      : https://github.com/luckykeeper | https://luckykeeper.site
 *  @Email        : luckykeeper@luckykeeper.site
 *  @project      : giteaCalendar
 */

package subfunction

import "github.com/gofiber/fiber/v2"

func LogMarkIsChild() string {
	if !fiber.IsChild() {
		return "[giteaCalendarSelf]"
	} else {
		return "[giteaCalendarFork]"
	}
}
