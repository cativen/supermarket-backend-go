package dao

import (
	"superMarket-backend/common"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/utils"
	"superMarket-backend/vo"
)

type MemberDao interface {
	SelectById(id int64) model.Member
	SelectMemberListByIds(ids []int64) []model.Member
	SelectPageByQo(dto dto.QueryMemberDTO) vo.Page[model.Member]
	UpdateMemberByQO(updateModel model.Member, queryModel model.Member)
	SelectOneByQo(query model.Member) model.Member
	SaveMember(member model.Member)
	SelectOneByExtraQo(member model.Member) model.Member
	SelectListByStateGroupById() []model.Member
	SelectListByGreaterEqual(integral uint) []model.Member
	SelectListByGreater(integral uint) []model.Member
}

type MemberDaoImpl struct {
}

func (memberDao *MemberDaoImpl) SelectById(id int64) model.Member {
	var member model.Member
	db := common.GetDB()
	db.First(&member, id)
	return member
}

func (memberDao *MemberDaoImpl) SelectOneByExtraQo(member model.Member) model.Member {
	db := common.GetDB()
	var res model.Member
	db.Where("phone = ? and state = ? and id != ?", member.Phone, common.STATE_NORMAL, member.ID).First(&res)
	return res
}

func (memberDao *MemberDaoImpl) SelectMemberListByIds(ids []int64) []model.Member {
	var members []model.Member
	db := common.GetDB()
	db.Where("id in (?)", ids).Find(&members)
	return members
}

func (memberDao *MemberDaoImpl) SelectPageByQo(dto dto.QueryMemberDTO) vo.Page[model.Member] {
	var page vo.Page[model.Member]
	db := common.GetDB()
	//分页查询出所有的销售记录
	db = db.Table("t_member").Select("*")
	phone := dto.Phone
	if phone != "" && len(phone) > 0 {
		db = db.Where("phone like ?", "%"+phone+"%")
	}

	name := dto.Name
	if name != "" && len(name) > 0 {
		db = db.Where("name like ?", "%"+name+"%")
	}

	state := dto.State
	if state != "" && len(state) > 0 {
		db = db.Where("state = ?", state)
	}

	// 查询会员列表
	var records []model.Member
	count, records := utils.Paginate(db, dto.CurrentPage, dto.PageSize, records)

	page.Records = records
	page.Total = count
	page.Current = dto.CurrentPage
	page.Size = dto.PageSize
	return page
}

func (memberDao *MemberDaoImpl) UpdateMemberByQO(updateModel model.Member, queryModel model.Member) {
	db := common.GetDB()
	db.Where(&queryModel).Updates(&updateModel)
}

func (memberDao *MemberDaoImpl) SelectOneByQo(query model.Member) model.Member {
	var members model.Member
	db := common.GetDB()
	db.Where(query).First(&members)
	return members
}

func (memberDao *MemberDaoImpl) SaveMember(query model.Member) {
	db := common.GetDB()
	db.Create(&query)
}

func (memberDao *MemberDaoImpl) SelectListByStateGroupById() []model.Member {
	db := common.GetDB()
	var members []model.Member
	db.Where("state = 0").Group("id").Find(&members)
	return members
}

func (memberDao *MemberDaoImpl) SelectListByGreaterEqual(integral uint) []model.Member {
	db := common.GetDB()
	var members []model.Member
	db.Where("integral > ? and state=0", integral).Find(&members)
	return members
}

func (memberDao *MemberDaoImpl) SelectListByGreater(integral uint) []model.Member {
	db := common.GetDB()
	var members []model.Member
	db.Where("integral >= ? and state=0", integral).Find(&members)
	return members
}
