package mysql

import (
	"bluebell/models"
	"database/sql"

	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id,community_name from community"
	if err = db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows { //查询为空
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}

//GetCommunityDetailByID 根据ID查询社区详情
func GetCommunityDetailByID(id int64) (comDetail *models.CommunityDetail, err error) {
	comDetail = new(models.CommunityDetail)
	sqlStr := `select 
				community_id,community_name,introduction,create_time 
				from community 
				where community_id = ?`
	if err = db.Get(comDetail, sqlStr, id); err != nil {
		if err == sql.ErrNoRows { //查询为空
			err = ErrorInvalidID //id错误
		}
	}
	return
}
