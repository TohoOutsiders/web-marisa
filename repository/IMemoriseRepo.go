/**
 * @Author: Tomonori
 * @Date: 2019/7/27 14:36
 * @File: IMemoriseRepo
 * @Desc:
 */
package repository

import "server/models"

type IMemoriseRepo interface {
	// 插入记忆
	AddMemory(data map[string]interface{}) bool
	// 读取所有记忆
	FetchAllMemory() (memorise []models.Memorise)
	// 读取一条记忆
	FetchMemory(answer string) (memorise models.Memorise)
	// 删除一条记忆
	DeleteMemoryByAnswer(answer string) bool
}
