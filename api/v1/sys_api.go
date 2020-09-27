package v1

import (
	"anew-server/common"
	"anew-server/dto/request"
	"anew-server/dto/response"
	"anew-server/dto/service"
	"anew-server/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 获取接口列表
func GetApis(c *gin.Context) {
	// 绑定参数
	var req request.ApiListReq
	err := c.Bind(&req)
	if err != nil {
		response.FailWithCode(response.ParmError)
		return
	}

	// 创建服务
	s := service.New(c)
	apis, err := s.GetApis(&req)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	// 转为ResponseStruct, 隐藏部分字段
	var respStruct []response.ApiListResp
	utils.Struct2StructByJson(apis, &respStruct)
	if req.Tree {
		// 转换成树结构
		tree := make([]response.ApiTreeResp, 0)
		for _,api := range respStruct{
			existIndex := -1
			children := make([]response.ApiListResp, 0)
			for index, leaf := range tree {
				if leaf.Category == api.Category {
					children = leaf.Children
					existIndex = index
					break
				}
			}
			// api结构转换
			var item response.ApiListResp
			utils.Struct2StructByJson(api, &item)
			item.Title = fmt.Sprintf("[%s] [%s] %s", item.Method, item.Desc, item.Path )
			item.Key = fmt.Sprintf("%d",item.Id)
			children = append(children, item)
			if existIndex != -1 {
				// 更新元素
				tree[existIndex].Children = children
			} else {
				// 新增元素
				tree = append(tree, response.ApiTreeResp{
					Key: api.Category,
					Title:    api.Category + " [分组]",
					Category: api.Category,
					Children: children,
				})
			}
		}

		response.SuccessWithData(tree)
		return
	}
	response.SuccessWithData(respStruct)
}


// 创建接口
func CreateApi(c *gin.Context) {
	user,_:= GetCurrentUser(c)
	// 绑定参数
	var req request.CreateApiReq
	err := c.Bind(&req)
	if err != nil {
		response.FailWithCode(response.ParmError)
		return
	}

	// 参数校验
	err = common.NewValidatorError(common.Validate.Struct(req), req.FieldTrans())
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	// 记录当前创建人信息
	req.Creator = user.Name
	// 创建服务
	s := service.New(c)
	err = s.CreateApi(&req)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	response.Success()
}

// 更新接口
func UpdateApiById(c *gin.Context) {
	// 绑定参数
	var req gin.H
	err := c.Bind(&req)
	if err != nil {
		response.FailWithCode(response.ParmError)
		return
	}

	// 获取path中的apiId
	apiId := utils.Str2Uint(c.Param("apiId"))
	if apiId == 0 {
		response.FailWithMsg("接口编号不正确")
		return
	}
	// 创建服务
	s := service.New(c)
	// 更新数据
	err = s.UpdateApiById(apiId, req)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	response.Success()
}

// 批量删除接口
func BatchDeleteApiByIds(c *gin.Context) {
	var req request.IdsReq
	err := c.Bind(&req)
	if err != nil {
		response.FailWithCode(response.ParmError)
		return
	}

	// 创建服务
	s := service.New(c)
	// 删除数据
	err = s.DeleteApiByIds(req.Ids)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	response.Success()
}
