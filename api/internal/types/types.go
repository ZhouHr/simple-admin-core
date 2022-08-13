// Code generated by goctl. DO NOT EDIT.
package types

type RoleInfo struct {
	Id            uint64 `json:"id"`
	Name          string `json:"name"`
	Value         string `json:"value"`
	DefaultRouter string `json:"defaultRouter"`
	Status        uint32 `json:"status"`
	Remark        string `json:"remark"`
	OrderNo       uint32 `json:"orderNo"`
	CreateAt      int64  `json:"createAt"`
}

type RoleListResp struct {
	PageList
	Data []RoleInfo `json:"data"`
}

type SetStatusReq struct {
	Id     uint64 `json:"id"`
	Status uint32 `json:"status"`
}

type BaseMsg struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

type BaseResp struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

type SimpleMsg struct {
	Msg string `json:"msg"`
}

type PageInfo struct {
	Page     uint64 `json:"page"`
	PageSize uint64 `json:"pageSize"`
}

type PageList struct {
	Total uint64   `json:"total"`
	Data  []string `json:"data"`
}

type BaseInfo struct {
	ID        uint  `json:"id"`
	CreatedAt int64 `json:"createdAt"`
	UpdatedAt int64 `json:"updatedAt"`
	DeletedAt int64 `json:"deletedAt"`
}

type IdReq struct {
	ID uint `json:"id"`
}

type UUIDReq struct {
	UUID string `json:"uuid"`
}

type LoginReq struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	CaptchaId string `json:"captchaId"`
	Captcha   string `json:"captcha"`
}

type LoginResp struct {
	UserId string         `json:"userId"`
	Role   RoleInfoSimple `json:"role"`
	Token  string         `json:"token"`
	Expire uint64         `json:"expire"`
}

type RoleInfoSimple struct {
	RoleName string `json:"roleName"`
	Value    string `json:"value"`
}

type RegisterReq struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	CaptchaId string `json:"captchaId"`
	Captcha   string `json:"captcha"`
	Email     string `json:"email"`
}

type ChangePasswordReq struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

type ChangePasswordResp struct {
	BaseMsg
}

type ModifyInfoReq struct {
	UserId   int64  `json:"userId"`
	Nickname string `json:"nickname"`
	Mobile   string `json:"mobile"`
	RoleId   uint32 `json:"roleId"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	SideMode string `json:"sideMode"`
}

type UserInfoResp struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Mobile   string `json:"mobile"`
	RoleId   uint32 `json:"roleId"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	SideMode string `json:"sideMode"`
	Status   int32  `json:"status"`
	CreateAt int64  `json:"createAt"`
	UpdateAt int64  `json:"updateAt"`
}

type GetUserInfoResp struct {
	UserId   string          `json:"userId"`
	Username string          `json:"username"`
	Nickname string          `json:"nickname"`
	Avatar   string          `json:"avatar"`
	Roles    GetUserRoleInfo `json:"roles"`
}

type GetUserRoleInfo struct {
	RoleName string `json:"roleName"`
	Value    string `json:"value"`
}

type UserListResp struct {
	PageList
	Data []UserInfoResp `json:"data"`
}

type PermCodeResp struct {
	Data []string `json:"data"`
}

type CreateOrUpdateUserReq struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
	RoleId   uint32 `json:"roleId"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Status   int32  `json:"status"`
}

type GetUserListReq struct {
	Page     uint64 `json:"page"`
	PageSize uint64 `json:"pageSize"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	RoleId   uint64 `json:"roleId"`
}

type Menu struct {
	BaseInfo
	MenuType  uint32 `json:"type"`
	ParentId  uint   `json:"parentId"`
	MenuLevel uint32 `json:"level"`
	Path      string `json:"path,omitempty"` // index path
	Name      string `json:"name"`           // index name
	Redirect  string `json:"redirect,omitempty"`
	Component string `json:"component"` // the path of vue file
	OrderNo   uint32 `json:"orderNo"`
	Disabled  bool   `json:"disabled"`
	Meta
	Children []*Menu `json:"children"`
}

type Meta struct {
	KeepAlive         bool   `json:"keepAlive"` // save in cache
	HideMenu          bool   `json:"hideMenu"`
	HideBreadcrumb    bool   `json:"hideBreadcrumb"`
	CurrentActiveMenu string `json:"currentActiveMenu"`
	Title             string `json:"title"`    // menu name
	Icon              string `json:"icon"`     // menu icon
	CloseTab          bool   `json:"closeTab"` // auto close tab
}

type MenuParam struct {
	BaseInfo
	MenuId uint   `json:"menuId"`
	Type   string `json:"type"`  // pass parameters via params or query
	Key    string `json:"key"`   // the key of parameters
	Value  string `json:"value"` // the value of parameters
}

type MenuListResp struct {
	PageList
	Data []*Menu `json:"data"`
}

type GetMenuListBase struct {
	MenuType  uint32             `json:"type"`
	MenuLevel uint32             `json:"level"`
	Path      string             `json:"path"` // index path
	Name      string             `json:"name"` // index name
	Redirect  string             `json:"redirect,omitempty"`
	Component string             `json:"component"` // the path of vue file
	OrderNo   uint32             `json:"orderNo"`
	Disabled  bool               `json:"disabled"`
	Meta      Meta               `json:"meta"` // extra parameters
	Children  []*GetMenuListBase `json:"children"`
}

type CreateOrUpdateMenuReq struct {
	Id        uint32 `json:"id"`
	MenuType  uint32 `json:"type"`
	ParentId  uint32 `json:"parentId"`
	Path      string `json:"path,omitempty"` // index path
	Name      string `json:"name"`           // index name
	Redirect  string `json:"redirect"`
	Component string `json:"component"` // the path of vue file
	OrderNo   uint32 `json:"orderNo"`
	Disabled  bool   `json:"disabled"`
	Meta
}

type CaptchaInfoResp struct {
	BaseMsg
	Data CaptchaInfo `json:"data"`
}

type CaptchaInfo struct {
	CaptchaId string `json:"captchaId"`
	ImgPath   string `json:"imgPath"`
}

type ApiInfo struct {
	Id          uint64 `json:"id"`
	CreateAt    int64  `json:"createAt"`
	Path        string `json:"path"`
	Description string `json:"description"`
	Group       string `json:"group"`
	Method      string `json:"method"`
}

type ApiListResp struct {
	PageList
	Data []ApiInfo `json:"data"`
}

type ApiListReq struct {
	PageInfo
	Path        string `json:"path"`
	Description string `json:"description"`
	Group       string `json:"group"`
	Method      string `json:"method"`
}

type ApiAuthorityInfo struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

type CreateOrUpdateApiAuthorityReq struct {
	RoleId uint64             `json:"roleId"`
	Data   []ApiAuthorityInfo `json:"data"`
}

type ApiAuthorityListResp struct {
	PageList
	Data []ApiAuthorityInfo `json:"data"`
}

type MenuAuthorityInfo struct {
	RoleId  uint64   `json:"roleId"`
	MenuIds []uint64 `json:"menuIds"`
}
