package service

import (
	"github.com/gin-gonic/gin"
	"superMarket-backend/common"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/response"
	"superMarket-backend/utils"
)

type IMemberService interface {
	QueryMemberPageByQo(c *gin.Context, dto dto.QueryMemberDTO)
	DelMember(c *gin.Context, id string)
	SaveMember(c *gin.Context, member model.Member)
	QueryMemberById(c *gin.Context, id string)
	UpdateMember(c *gin.Context, member model.Member)
	QueryMemberByPhone(c *gin.Context, phone string)
}

type MemberServiceImpl struct {
}

func (memberService *MemberServiceImpl) QueryMemberPageByQo(c *gin.Context, dto dto.QueryMemberDTO) {
	page := memberDao.SelectPageByQo(dto)
	response.Success(c, page, "操作成功")
}

func (memberService *MemberServiceImpl) DelMember(c *gin.Context, id string) {
	var updateModel = model.Member{
		State:    common.MEMBER_STATE_BAN,
		Integral: 0,
	}

	var queryModel = model.Member{
		ID: utils.ConvertStringToInt64(id),
	}
	memberDao.UpdateMemberByQO(updateModel, queryModel)
	response.Success(c, nil, "操作成功")
}

func (memberService *MemberServiceImpl) SaveMember(c *gin.Context, member model.Member) {
	var queryModel = model.Member{
		Phone: member.Phone,
		State: common.STATE_NORMAL,
	}
	one := memberDao.SelectOneByQo(queryModel)
	if one != (model.Member{}) {
		response.Error(c, "该用户已注册过")
	}
	member.Password = common.DEFAULT_PWD
	member.State = common.STATE_NORMAL
	member.Integral = 0
	memberDao.SaveMember(member)
	response.Success(c, nil, "操作成功")
}

func (memberService *MemberServiceImpl) QueryMemberById(c *gin.Context, id string) {
	byId := memberDao.SelectById(utils.ConvertStringToInt64(id))
	response.Success(c, byId, "操作成功")
}

func (memberService *MemberServiceImpl) UpdateMember(c *gin.Context, member model.Member) {
	if common.STATE_NORMAL == member.State {
		existMember := memberDao.SelectOneByExtraQo(member)
		if existMember != (model.Member{}) {
			response.Error(c, "已被录入")
			return
		}
	}
	var queryModel = model.Member{
		ID: member.ID,
	}
	memberDao.UpdateMemberByQO(member, queryModel)
	response.Success(c, nil, "操作成功")
}

func (memberService *MemberServiceImpl) QueryMemberByPhone(c *gin.Context, phone string) {
	var queryModel = model.Member{
		Phone: phone,
		State: common.STATE_NORMAL,
	}
	one := memberDao.SelectOneByQo(queryModel)
	if one == (model.Member{}) {
		response.Error(c, "该会员不存在")
		return
	}
	response.Success(c, one, "操作成功")
}
