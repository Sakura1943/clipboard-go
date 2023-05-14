// 数据库操作
// 剪切板数据库操作

package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

// 剪切板数据结构体
type Document struct {
	// ID
	ID uint64 `json:"id" gorm:"id;primaryKey;unique;autoIncrement" binding:"required" xml:"id" form:"id"`
	// 文件路径
	Path string `json:"path" gorm:"path;unique;not null" xml:"path" form:"path"`
	// 文件文本
	Text string `json:"text" gorm:"text;not null" xml:"text" form:"text"`
	// 文件的语言
	Lang string `json:"lang" gorm:"lang;not null" xml:"lang" form:"lang"`
	// 所属用户名
	UserName string `json:"user_name" gorm:"user_name;not null" xml:"user_name" form:"user_name"`
}

// 初始化剪切板数据
func InitDocumentDatabase() (*gorm.DB, error) {
	// 判断data文件夹是否存在(data用于存储数据库文件)
	if _, err := os.Stat("data/"); err != nil {
		// 不存在则创建
		err := os.MkdirAll("data/", os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	// 打开数据库
	db, err := gorm.Open(sqlite.Open("data/document.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 自动迁移数据库
	db.AutoMigrate(&Document{})
	return db, err
}

// 增加数据到数据库中
func CreateData(db *gorm.DB, data *Document) (*gorm.DB, bool) {
	// 判断数据的文本及其路径信息是否在数据库中，不存在则创建
	if !ExistsFile(db, data.Path) {
		db.Create(&data)
	} else {
		// 否则返回创建失败
		return nil, false
	}
	// 否则返回创建失败
	return nil, true
}

// 是否存在对应的路径和文本在数据库中
func ExistsFile(db *gorm.DB, filePath string) bool {
	// 获取在数据库中符合筛选条件的数据
	count := db.Model(&Document{}).Where("path = ?", filePath).Find(&Document{}).RowsAffected
	// 如果等于0则不存在
	return count != 0
}

// 获取所有数据的信息
func AllDatas(db *gorm.DB) *[]Document {
	return db.Model(&[]Document{}).Find(&[]Document{}).Statement.Dest.(*[]Document)
}

// 删除数据
func DeleteData(db *gorm.DB, filePath string) bool {
	db.Model(&Document{}).Delete(&Document{}, "path = ?", filePath)
	// 删除结束判断数据是否还存在， 不存在则删除成功
	return !ExistsFile(db, filePath)
}

// 查询数据
func SearchData(db *gorm.DB, filePath string) (*Document, bool) {
	// 判断数据的文本及其路径信息是否在数据库中，不存在则查询失败
	if !ExistsFile(db, filePath) {
		return nil, false
	}
	// 否则查询成功
	return db.Model(&Document{}).Where("path = ?", filePath).First(&Document{}).Statement.Dest.(*Document), true
}
