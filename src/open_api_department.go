package dingtalk

import (
	"fmt"
	"net/url"
)

type SubDepartmentListResponse struct {
	OpenAPIResponse
	SubDeptIdList []int `json:"sub_dept_id_list"`
}

type DepartmentListResponse struct {
	OpenAPIResponse
	Department []Department
}

type Department struct {
	Id              int
	Name            string
	ParentId        int
	CreateDeptGroup bool
	AutoAddUser     bool
}

type DepartmentDetailResponse struct {
	OpenAPIResponse
	Id                    int
	Name                  string
	ParentId              int
	Order                 int
	CreateDeptGroup       bool
	AutoAddUser           bool
	DeptHiding            bool
	DeptPermits           string
	UserPermits           string
	OuterDept             bool
	OuterPermitDepts      string
	OuterPermitUsers      string
	OrgDeptOwner          string
	DeptManagerUserIdList string
	SourceIdentifier      string
	GroupContainSubDept   bool
}

type DepartmentCreateResponse struct {
	OpenAPIResponse
	Id int
}

type DepartmentCreateRequest struct {
	Name             string `json:"name"`
	ParentId         string `json:"parentid"`
	Order            string `json:"order,omitempty"`
	CreateDeptGroup  bool   `json:"createDeptGroup,omitempty"`
	DeptHiding       bool   `json:"deptHiding,omitempty"`
	DeptPerimits     string `json:"deptPerimits,omitempty"`
	UserPerimits     string `json:"userPerimits,omitempty"`
	OuterDept        string `json:"outerDept,omitempty"`
	OuterPermitDepts string `json:"outerPermitDepts,omitempty"`
	OuterPermitUsers string `json:"outerPermitUsers,omitempty"`
	SourceIdentifier string `json:"sourceIdentifier,omitempty"`
}

type DepartmentUpdateResponse struct {
	OpenAPIResponse
	Id int
}

type DepartmentUpdateRequest struct {
	Lang                  string `json:"lang,omitempty"`
	Name                  string `json:"name,omitempty"`
	ParentId              string `json:"parentid,omitempty"`
	Order                 string `json:"order,omitempty"`
	Id                    string `json:"id"`
	CreateDeptGroup       bool   `json:"createDeptGroup,omitempty"`
	AutoAddUser           bool   `json:"autoAddUser,omitempty"`
	DeptManagerUseridList string `json:"deptManagerUseridList,omitempty"`
	DeptHiding            bool   `json:"deptHiding,omitempty"`
	DeptPerimits          string `json:"deptPerimits,omitempty"`
	UserPerimits          string `json:"userPerimits,omitempty"`
	OuterDept             string `json:"outerDept,omitempty"`
	OuterPermitDepts      string `json:"outerPermitDepts,omitempty"`
	OuterPermitUsers      string `json:"outerPermitUsers,omitempty"`
	OrgDeptOwner          string `json:"orgDeptOwner,omitempty"`
	SourceIdentifier      string `json:"sourceIdentifier,omitempty"`
}

type DepartmentDeleteResponse struct {
	OpenAPIResponse
}

type DepartmentListParentDeptsByDeptResponse struct {
	OpenAPIResponse
	ParentIds []int `json:"parentIds"`
}

type DepartmentListParentDeptsResponse struct {
	OpenAPIResponse
	ParentIds interface{} `json:"dep"`
}

// 获取部门id列表
func (dtc *Client) DepartmentList(id string) (*DepartmentListResponse, error) {
	data := &DepartmentListResponse{}
	params := url.Values{}
	if id != "" {
		params.Add("id", id)
	}
	err := dtc.httpRPC("department/list", params, nil, data)
	return data, err
}

// 获取子部门Id列表
func (dtc *Client) SubDepartmentList(id string) (*SubDepartmentListResponse, error) {
	data := &SubDepartmentListResponse{}
	params := url.Values{}
	if id != "" {
		params.Add("id", id)
	}
	err := dtc.httpRPC("department/list_ids", params, nil, data)
	return data, err
}

// 获取部门详情
func (dtc *Client) DepartmentDetail(id string) (*DepartmentDetailResponse, error) {
	data := &DepartmentDetailResponse{}
	params := url.Values{}
	if id != "" {
		params.Add("id", id)
	}
	err := dtc.httpRPC("department/get", params, nil, data)
	return data, err
}

// 创建部门
func (dtc *Client) DepartmentCreate(info *DepartmentCreateRequest) (DepartmentCreateResponse, error) {
	var data DepartmentCreateResponse
	err := dtc.httpRPC("department/create", nil, info, &data)
	return data, err
}

// 更新部门
func (dtc *Client) DepartmentUpdate(info *DepartmentUpdateRequest) (DepartmentUpdateResponse, error) {
	var data DepartmentUpdateResponse
	err := dtc.httpRPC("department/update", nil, info, &data)
	return data, err
}

// 删除部门
func (dtc *Client) DepartmentDelete(id int) (DepartmentDeleteResponse, error) {
	var data DepartmentDeleteResponse
	params := url.Values{}
	params.Add("id", fmt.Sprintf("%d", id))
	err := dtc.httpRPC("department/delete", params, nil, &data)
	return data, err
}

// 查询部门的所有上级父部门路径
func (dtc *Client) DepartmentListParentDeptsByDept(id int) (DepartmentListParentDeptsByDeptResponse, error) {
	var data DepartmentListParentDeptsByDeptResponse
	params := url.Values{}
	params.Add("id", fmt.Sprintf("%d", id))
	err := dtc.httpRPC("department/list_parent_depts_by_dept", params, nil, &data)
	return data, err
}

// 查询指定用户的所有上级父部门路径
func (dtc *Client) DepartmentListParentDepts(userId string) (DepartmentListParentDeptsResponse, error) {
	var data DepartmentListParentDeptsResponse
	params := url.Values{}
	params.Add("userId", userId)
	err := dtc.httpRPC("department/list_parent_depts", params, nil, &data)
	return data, err
}
