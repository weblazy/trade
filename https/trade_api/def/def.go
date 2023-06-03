package def

type GetUserInfoRequest struct {
	Uid int `json:"uid" binding:"required,min=1,max=100000"` // 用户ID

}

type GetUserInfoResponse struct {
	Detail *User   `json:"detail" binding:""` // 用户详情
	List   []*User `json:"list" binding:""`   // 用户列表

}

type User struct {
	Uid  int    `json:"uid" binding:""`  // 用户ID
	Name string `json:"name" binding:""` // 用户名

}
