/**
 * @Author: Tomonori
 * @Date: 2019/7/27 14:42
 * @File: IMemoriseService
 * @Desc:
 */
package service

import "server/models"

type IMemoriseService interface {
	// 记忆学习
	Add(memory models.Memorise) map[string]interface{}
	// 回复
	Reply(memory models.Memorise) (int, map[string]interface{})
	// 忘记
	Forget(answer string) bool
	// 状态
	Status() int
}
