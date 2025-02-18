package def

type User struct {
	UserId uint32 `json:"userid"`
}

// 用户登陆时输入的信息
type Login struct {
	Email    string `form:"email,required" json:"email,required"`
	Password string `form:"password,required" json:"password,required"`
}

// 用户注册时输入的信息
type Register struct {
	Email    string `form:"email,required" json:"email,required"`
	Username string `form:"username,required" json:"username,required"`
	Password string `form:"password,required" json:"password,required"`
}
