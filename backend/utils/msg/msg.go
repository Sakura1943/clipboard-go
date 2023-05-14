// 响应信息
// 请求响应信息

package msg

import (
	"encoding/json"
	"log"
	"net/http"
)

// 响应的信息的结构体
type Message struct {
	// http响应码
	Code uint16 `json:"code" xml:"code" form:"code"`
	// 响应类型
	Type string `json:"type,omitempty" xml:"type,omitempty" form:"type,omitempty"`
	// 成功响应的信息
	Message string `json:"message,omitempty" xml:"message,omitempty" form:"message,omitempty"`
	// 错误响应的信息
	Error string `json:"error,omitempty" xml:"error,omitempty" form:"error,omitempty"`
	// 响应额外的说明信息
	Extra interface{} `json:"extra,omitempty" xml:"extra,omitempty" form:"extra,omitempty"`
}

// 消息响应类型
const (
	// 用户已存在
	UserExists = "user_exists"
	// 已登录
	Logined = "logined"
	// 登录失败
	LoginError = "login_error"
	// 更新失败
	UpdateError = "update_error"
	// 已更新
	Updated = "updated"
	// token携带的信息与数据库中的信息不相等
	TokenInfoNotEqual = "token_info_not_equal"
	// 密码错误
	PasswordWrong = "password_wrong"
	// 已注册
	Registered = "registered"
	// 更改生成token
	GeneratedToken = "generated_token"
	/// 生成token错误
	GenerateTokenError = "generate_token_error"
	// 验证失败(未验证成功)，操作不被允许
	Unauthorized = "unauthorized"
	// 获取token失败
	GetTokenError = "get_token_error"
	// form表单信息解析失败
	FormDataParseError = "form_data_parse_error"
	// 初始化数据库失败
	InitDatabaseError = "init_database_error"
	// 获取用户信息失败
	GetUserInfoError = "get_user_info_error"
	// 删除用户失败
	DeleteUserError = "delete_user_error"
	// 已删除用户
	DeletedUser = "deleted_user"
	// 权限不足
	PermissionDenied = "permission_denied"
	// 用户不存在
	UserNotExixts = "user_not_exists"
	// 操作不被允许
	OperationNotAllowed = "operation_not_allowed"
	// 字段不存在
	FieldNotExists = "filed_not_exists"
	// 上传错误
	UploadError = "upload_error"
	// 不是文本
	IsNotText = "is_not_text"
	// 创建文件夹失败
	CreateDirError = "create_dir_error"
	// 保存文件失败
	SaveFileError = "save_file_error"
	// 读取文件失败
	ReadFileError = "read_file_error"
	// 文件已存在
	FileExists = "file_exists"
	// 保存成功
	SaveSuccess = "save_sucess"
	// 空值
	Empty = "empty"
	// 已删除数据
	DeletedData = "deleted_data"
	// 删除数据失败
	DeleteDataError = "delete_data_error"
	// 数据不存在
	DataNotExists = "data_not_exists"
	// 查询失败
	SearchError = "search_error"
	// 遍历不存在
	VarNotExists = "var_not_exists"
	// 位置权限
	UnknownPermission = "unknown_permission"
	// 不允许改变admin用户信息
	NotAllowedToChangeAdminInfo = "not_allowed_to_change_admin_info"
	// 创建用户失败
	CreatingUserFailed = "creating_user_failed"
)

// 初始化响应信息
func Default() Message {
	return Message{
		Code: http.StatusOK,
	}
}

// 自定义http状态码
func (m *Message) WithCode(code uint16) {
	m.Code = code
}

// 自定义响应的消息类型
func (m *Message) WithType(Type string) {
	m.Type = Type
}

// 自定义成功响应的信息
func (m *Message) WithMessage(msg string) {
	m.Message = msg
}

// 自定义错误响应的信息
func (m *Message) WithError(err string) {
	m.Error = err
}

// 自定义额外的说明信息
func (m *Message) WithExtra(extra interface{}) {
	m.Extra = extra
}

// 打印响应的信息
func PrintlnMessage(message *Message) error {
	// 序列化为json字符串
	jsonData, err := json.Marshal(message)
	// 打印json字符串
	log.Println(string(jsonData))
	return err
}
