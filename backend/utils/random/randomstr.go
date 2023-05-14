// 随机生成
// 随机字符串生成

package random

import (
	"math/rand"
)

// 字符串可能包含的字符
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

// 生成随机字符串
func RandStr(size int) string {
	// 创建一个存储size各单字节字符的bytes数据
	b := make([]rune, size)

	// 插入size次字符
	for i := range b {
		// 每次从letters中随机选取
		b[i] = letters[rand.Intn(len(letters))]
	}
	// 转化为string并返回
	return string(b)
}
