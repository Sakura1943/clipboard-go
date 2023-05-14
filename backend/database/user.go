// 数据库操作
// 用户数据库操作

package database

import (
	"backend/utils/permission"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

// 用户数据信息结构体
type User struct {
	// 用户ID，唯一值
	ID uint64 `json:"id" gorm:"id;primaryKey;unique;autoIncrement" xml:"id" form:"id"`
	// 用户名
	Name string `json:"name" gorm:"name;unique;not null" xml:"name" form:"name"`
	// 用户密码
	Password string `json:"password" gorm:"password;not null" xml:"password" form:"password"`
	// 用户权限
	Permission string `json:"permission" gorm:"permission;not null" xml:"permission" form:"permission"`
}

// 初始化用户数据库
func InitUserDatabase() (*gorm.DB, error) {
	// 判断data文件夹是否存在(data用于存储数据库文件)
	if _, err := os.Stat("data/"); err != nil {
		// 不存在则创建
		if err := os.MkdirAll("data/", os.ModePerm); err != nil {
			return nil, err
		}
	}

	// 打开数据库
	db, err := gorm.Open(sqlite.Open("data/user.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// 自动迁移数据库
	db.AutoMigrate(&User{})

	// 如果不存在admin用户，则初始化一个admin用户
	if !ExistsUser(db, "admin") {
		if err != nil {
			return nil, err
		}
		db.Create(&[]User{
			{
				ID:         1,
				Name:       "admin",
				Password:   "12345678",
				Permission: permission.Admin,
			},
		})
		// 打印初始化的信息
		log.Println("Init user `admin`, password `12345678`")
	}
	return db, err
}

// 创建用户
func CreateUser(db *gorm.DB, user *User) (*gorm.DB, bool) {
	// 判断数据的用户是否在数据库中，不存在则创建
	if !ExistsUser(db, user.Name) {
		db.Create(&user)
	} else {
		// 否则返回创建失败
		return nil, false
	}
	// 否则返回创建失败
	return nil, true
}

// 搜索用户
func SearchUser(db *gorm.DB, userName string) *User {
	return db.Model(&User{}).Where("name = ?", userName).First(&User{}).Statement.Dest.(*User)
}

// 更新用户
func UpdateUser(db *gorm.DB, oldName string, user *User) *User {
	return db.Model(&User{}).Where("name = ?", oldName).Updates(user).Statement.Dest.(*User)
}

// 判断用户是否存在
func ExistsUser(db *gorm.DB, userName string) bool {
	// 获取检索的用户名在数据库中的数据
	count := db.Model(&User{}).Where("name = ?", userName).Find(&User{}).RowsAffected
	// 不为0则存在
	return count != 0
}

// 通过ID判断用户是否存在
func ExistsUserById(db *gorm.DB, id uint64) bool {
	// 获取检索的用户名在数据库中的数据
	count := db.Model(&User{}).Where("id = ?", id).Find(&User{}).RowsAffected
	// 不为0则存在
	return count != 0
}

// 删除用户
func DeleteUser(db *gorm.DB, userName string) bool {
	db.Model(&User{}).Delete(&User{}, "name = ?", userName)
	// 删除结束判断数据是否还存在， 不存在则删除成功
	return !ExistsUser(db, userName)
}

// 获取所有用户的数据
func AllUsers(db *gorm.DB) *[]User {
	return db.Model(&[]User{}).Find(&[]User{}).Statement.Dest.(*[]User)
}
