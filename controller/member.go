package controller

import (
	"github.com/gin-gonic/gin"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/response"
	"superMarket-backend/service"
)

var memberService service.IMemberService = &service.MemberServiceImpl{}

func QueryMemberPageByQo(c *gin.Context) {
	var dto dto.QueryMemberDTO
	err := c.ShouldBind(&dto)
	if err != nil {
		response.Error(c, "参数解析失败")
		return
	}
	memberService.QueryMemberPageByQo(c, dto)
}

func DelMember(c *gin.Context) {
	id, _ := c.GetPostForm("id")
	memberService.DelMember(c, id)
}

func SaveMember(c *gin.Context) {
	var model model.Member
	err := c.ShouldBind(&model)
	if err != nil {
		response.Error(c, "参数解析失败")
		return
	}
	memberService.SaveMember(c, model)
}

func QueryMemberById(c *gin.Context) {
	id := c.Query("id")
	memberService.QueryMemberById(c, id)
}

func UpdateMember(c *gin.Context) {
	var model model.Member
	err := c.ShouldBind(&model)
	if err != nil {
		response.Error(c, "参数解析失败")
		return
	}
	memberService.UpdateMember(c, model)
}

func QueryMemberByPhone(c *gin.Context) {
	phone := c.Query("phone")
	memberService.QueryMemberByPhone(c, phone)
}
