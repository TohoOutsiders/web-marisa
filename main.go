/**
 * @Author: Tomonori
 * @Date: 2019/6/18 14:30
 * @File: main
 * @Desc:
 */
package main

import (
	"fmt"
	"server/cmd"
	"server/common/setting"
)

func main() {
	app := cmd.App()
	app.Run(fmt.Sprintf(":%s", setting.Config.Server.Port))
}
