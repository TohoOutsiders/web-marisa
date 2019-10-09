/**
 * @Author: Tomonori
 * @Date: 2019/7/27 14:21
 * @File: IDb
 * @Desc:
 */
package datasource

import "github.com/jinzhu/gorm"

type IDb interface {
	DB() *gorm.DB
	Connect() error
}
