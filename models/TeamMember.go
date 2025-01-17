package models

import (
	"awesomeProject3/utils/errmsg"
	"github.com/jinzhu/gorm"
	"strings"
)

type TeamMember struct{
	gorm.Model
	Team_id uint `gorm:"type:int;not null" json:"team_id"`
	Member_uid uint `gorm:"type:int;not null" json:"member_uid"`
	Member_username string `gorm:"type:varchar(255)" json:"member_username"`
	Member_group_id uint `gorm:"type:int;default:1" json:"member_group_id"`
	Name string `gorm:"type:varchar(255)" json:"name"`

}
func  SaveTeamMember (id int , uid uint, memberUsername string )(TeamMember,int){
	var teammember TeamMember
	var member_username_array []string
	var teammember2 []TeamMember
	var teamInfo Team
	var teamItem []TeamItem
	var teamitemMember TeamItemMember
	var teamitemMember2 []TeamItemMember
	err = db.Model(Team{}).Where("id =?",id).Where("uid =?",uid).Find(&teamInfo).Error
	if err != nil{
		return teammember , errmsg.ERROR
	}
	member_username_array = strings.Split(memberUsername,",")

	for _,v := range member_username_array{
		var memberInfo User
		db.Model(User{}).Where("username = ?",v).Find(&memberInfo)
		teammember.Team_id = uint(id)
		teammember.Member_uid = memberInfo.ID
		teammember.Member_username = memberInfo.Username
		teammember.Name = memberInfo.Name
		teammember2 = append(teammember2 , teammember)
	}
	for _, member := range teammember2 {
		db.Model(TeamMember{}).Create(&member)
	}

	db.Model(TeamMember{}).Where("team_id =?",id).Find(&teamItem)
	for _,v := range teamItem {
		for _, member := range teammember2 {
			teamitemMember.Team_id = v.Team_id
			teamitemMember.Member_uid = member.Member_uid
			teamitemMember.Member_username = member.Member_username
			teamitemMember.Item_id = v.Item_id
			teamitemMember.Member_group_id = 1
			teamitemMember2 = append(teamitemMember2 , teamitemMember)
		}
	}
	for _, member := range teamitemMember2 {
		db.Create(&member)
	}

// 检查该团队已经加入了那些项目
return teammember,errmsg.SUCCESE
}

func GetTeamMemberList(id int,uid uint)([]TeamMember,int){
	var teamMember []TeamMember
	if uid >0{
		err :=db.Model(TeamMember{}).Where("team_id=?",id).Joins("LEFT JOIN user on user.ID = team_member.member_uid").Order("created_at desc").Select("team_member.*,user.username as name").Find(&teamMember).Error
		if err != nil {
			return teamMember, errmsg.ERROR
		}
	}
	return teamMember,errmsg.SUCCESE
}

func DeleteTeamMember(id int,uid uint )int{
	var teamMember TeamMember
	var teamInfo Team
	var teamItemMember TeamItemMember

	err = db.Model(TeamMember{}).Where("id=?",id).Find(&teamMember).Error

	err =db.Model(Team{}).Where("id=?",teamMember.Team_id).Where("uid=?",uid).Find(&teamInfo).Error
	if err != nil {
		return errmsg.ERROR
	}
	err =db.Model(TeamItemMember{}).Where("member_uid = ?",teamMember.Member_uid).Where("team_id=?",teamMember.Team_id).Delete(&teamItemMember).Error
	err =db.Model(TeamItem{}).Where("id =?",id).Delete(&teamMember).Error

	return errmsg.SUCCESE

}