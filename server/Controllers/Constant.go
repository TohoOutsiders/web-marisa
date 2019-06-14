/**
 * @Author: Tomonori
 * @Date: 2019/6/14 11:07
 * @File: Constant
 * @Desc:
 */
package Controllers

type ModelAndView struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
