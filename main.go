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
	"server/setting"
)

func main() {
	app := cmd.App()
	app.Run(fmt.Sprintf(":%d", setting.HttpPort))
}
