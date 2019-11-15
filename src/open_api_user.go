package dingtalk

import (
	"fmt"
	"net/url"
)

type UserIdResponse struct {
	OpenAPIResponse
	UserID   string `json:"userid"`
	DeviceID string `json:"deviceId"`
	IsSys    bool   `json:"is_sys"`
	SysLevel int    `json:"sys_level"`
}

type UserIdByUnionIdResponse struct {
	OpenAPIResponse
	UserID      string `json:"userid"`
	ContactType int    `json:"contactType"`
}

type UserInfoResponse struct {
	OpenAPIResponse
	UserID          string      `json:"userid"`
	UnionID         string      `json:"unionid"`
	Name            string      `json:"name"`
	Tel             string      `json:"tel"`
	WorkPlace       string      `json:"workPlace"`
	Remark          string      `json:"remark"`
	Mobile          string      `json:"mobile"`
	Email           string      `json:"email"`
	OrgEmail        string      `json:"orgEmail"`
	Active          bool        `json:"active"`
	OrderInDepts    string      `json:"orderInDepts"`
	IsAdmin         bool        `json:"isAdmin"`
	IsBoos          bool        `json:"isBoss"`
	IsLeaderInDepts string      `json:"isLeaderInDepts"`
	IsHide          bool        `json:"isHide"`
	Department      []int       `json:"department"`
	Position        string      `json:"position"`
	Avatar          string      `json:"avatar"`
	HiredDate       int64       `json:"hiredDate"`
	Jobnumber       string      `json:"jobnumber"`
	Extattr         interface{} `json:"extattr"`
	IsSenior        bool        `json:"isSenior"`
	StateCode       string      `json:"stateCode"`
	Roles           []Roles     `json:"roles"`
}

type Roles struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	GroupName string `json:"groupName"`
}

type UserIDListResponse struct {
	OpenAPIResponse
	UserIDs []string `json:"userIds"`
}

type UserSimpleListResponse struct {
	OpenAPIResponse
	HasMore  bool
	UserList []*USimpleList
}

type USimpleList struct {
	UserID string
	Name   string
}

type UserListResponse struct {
	OpenAPIResponse
	HasMore  bool
	UserList []*UDetailedList
}

type UDetailedList struct {
	UserID     string `json:"userid"`
	UnionID    string
	Order      int64
	Mobile     string
	Tel        string
	WorkPlace  string
	Remark     string
	IsAdmin    bool
	IsBoss     bool
	IsHide     bool
	IsLeader   bool
	Name       string
	Active     bool
	Department []int
	Position   string
	Email      string
	OrgEmail   string
	Avatar     string
	Jobnumber  string
	HiredDate  int64
	Extattr    interface{}
	StateCode  string
}

type UserAdminListResponse struct {
	OpenAPIResponse
	AdminList []*Admins
}

type Admins struct {
	UserID   string `json:"userid"`
	SysLevel int    `json:"sys_level"`
}

type UserCanAccessMicroappResponse struct {
	OpenAPIResponse
	CanAccess bool
}

type UserCreateResponse struct {
	OpenAPIResponse
	UserID string
}

type UserCreateRequest struct {
	UserID       string      `json:"userid,omitempty"`
	Name         string      `json:"name"`
	OrderInDepts string      `json:"orderInDepts,omitempty"`
	Department   []int       `json:"department"`
	Position     string      `json:"position,omitempty"`
	Mobile       string      `json:"mobile"`
	Tel          string      `json:"tel,omitempty"`
	WorkPlace    string      `json:"workPlace,omitempty"`
	Remark       string      `json:"remark,omitempty"`
	Email        string      `json:"email,omitempty"`
	OrgEmail     string      `json:"orgEmail,omitempty"`
	JobNumber    string      `json:"jobnumber,omitempty"`
	IsHide       bool        `json:"isHide,omitempty"`
	IsSenior     bool        `json:"isSenior,omitempty"`
	Extattr      interface{} `json:"extattr,omitempty"`
}

type UserUpdateResponse struct {
	OpenAPIResponse
}

type UserUpdateRequest struct {
	Lang         string      `json:"lang,omitempty"`
	UserID       string      `json:"userid"`
	Name         string      `json:"name"`
	OrderInDepts string      `json:"orderInDepts,omitempty"`
	Department   []int       `json:"department,omitempty"`
	Position     string      `json:"position,omitempty"`
	Mobile       string      `json:"mobile,omitempty"`
	Tel          string      `json:"tel,omitempty"`
	WorkPlace    string      `json:"workPlace,omitempty"`
	Remark       string      `json:"remark,omitempty"`
	Email        string      `json:"email,omitempty"`
	OrgEmail     string      `json:"orgEmail,omitempty"`
	JobNumber    string      `json:"jobnumber,omitempty"`
	IsHide       bool        `json:"isHide,omitempty"`
	IsSenior     bool        `json:"isSenior,omitempty"`
	Extattr      interface{} `json:"extattr,omitempty"`
}

type UserDeleteResponse struct {
	OpenAPIResponse
}

type UserBatchDeleteResponse struct {
	OpenAPIResponse
}

type UserGetOrgUserCountResponse struct {
	OpenAPIResponse
	Count int
}

// 通过Code换取userid
func (dtc *Client) UserIdByCode(code string) (*UserIdResponse, error) {
	data := &UserIdResponse{}
	params := url.Values{}
	params.Add("code", code)
	err := dtc.httpRPC("user/getuserinfo", params, nil, data)
	return data, err
}

// 通过UnionId获取UserId
func (dtc *Client) UserIdByUnionId(unionID string) (*UserIdByUnionIdResponse, error) {
	data := &UserIdByUnionIdResponse{}
	params := url.Values{}
	params.Add("unionid", unionID)
	err := dtc.httpRPC("user/getUseridByUnionid", params, nil, data)
	return data, err
}

// 通过userid 换取 用户详细信息
func (dtc *Client) UserInfoByUserId(userID string) (*UserInfoResponse, error) {
	data := &UserInfoResponse{}
	params := url.Values{}
	params.Add("userid", userID)
	err := dtc.httpRPC("user/get", params, nil, data)
	return data, err
}

// UserIDList 获取部门用户userid列表
func (dtc *Client) UserIDList(departmentID string) (*UserIDListResponse, error) {
	data := &UserIDListResponse{}
	params := url.Values{}
	params.Add("deptId", departmentID)
	err := dtc.httpRPC("user/getDeptMember", params, nil, data)
	return data, err

}

// 获取部门成员（简化版）
func (dtc *Client) UserSimpleList(departmentID int) (*UserSimpleListResponse, error) {
	data := &UserSimpleListResponse{}
	params := url.Values{}
	params.Add("department_id", fmt.Sprintf("%d", departmentID))
	err := dtc.httpRPC("user/simplelist", params, nil, data)
	return data, err
}

// 获取部门成员（详情版）
func (dtc *Client) UserList(departmentID, size, offset int) (*UserListResponse, error) {
	data := &UserListResponse{}
	params := url.Values{}
	params.Add("department_id", fmt.Sprintf("%d", departmentID))
	if size > 0 {
		if size > 100 {
			size = 100
		}
		params.Add("size", fmt.Sprintf("%d", size))
		params.Add("offset", fmt.Sprintf("%d", offset))
	}
	err := dtc.httpRPC("user/list", params, nil, data)
	return data, err
}

// 获取管理员列表
func (dtc *Client) UserAdminList() (*UserAdminListResponse, error) {
	data := &UserAdminListResponse{}
	err := dtc.httpRPC("user/get_admin", nil, nil, data)
	return data, err
}

// 获取管理员的微应用管理权限
func (dtc *Client) UserCanAccessMicroapp(appID string, userID string) (*UserCanAccessMicroappResponse, error) {
	var data UserCanAccessMicroappResponse
	params := url.Values{}
	params.Add("appId", appID)
	params.Add("userId", userID)
	err := dtc.httpRPC("user/can_access_microap", params, nil, &data)
	return &data, err
}

// 创建成员
func (dtc *Client) UserCreate(info *UserCreateRequest) (*UserCreateResponse, error) {
	var data UserCreateResponse
	err := dtc.httpRPC("user/create", nil, info, &data)
	return &data, err
}

// 更新成员
func (dtc *Client) UserUpdate(info *UserUpdateRequest) (*UserUpdateResponse, error) {
	var data UserUpdateResponse
	err := dtc.httpRPC("user/update", nil, info, &data)
	return &data, err
}

// 删除成员
func (dtc *Client) UserDelete(userID string) (*UserDeleteResponse, error) {
	var data UserDeleteResponse
	params := url.Values{}
	params.Add("userid", userID)
	err := dtc.httpRPC("user/delete", params, nil, &data)
	return &data, err
}

// 批量删除成员
func (dtc *Client) UserBatchDelete(userIdList []string) (*UserBatchDeleteResponse, error) {
	var data UserBatchDeleteResponse
	body := map[string][]string{
		"useridlist": userIdList,
	}
	err := dtc.httpRPC("user/batchdelete", nil, body, &data)
	return &data, err
}

// 获取企业员工人数
func (dtc *Client) UserGetOrgUserCount(onlyActive int) (*UserGetOrgUserCountResponse, error) {
	var data UserGetOrgUserCountResponse
	params := url.Values{}
	params.Add("onlyActive", fmt.Sprintf("%d", onlyActive))
	err := dtc.httpRPC("user/get_org_user_count", params, nil, &data)
	return &data, err
}
