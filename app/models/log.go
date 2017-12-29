package models

import (
	"bzppx-codepub/app/utils"
	"github.com/snail007/go-activerecord/mysql"
)

const Table_Log_Name = "log"

type Log struct {
}

var LogModel = Log{}

// 根据 log_id 获取日志
func (l *Log) GetLogByLogId(logId string) (log map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Log_Name).Where(map[string]interface{}{
		"log_id": logId,
	}))
	if err != nil {
		return
	}
	log = rs.Row()
	return
}

// 插入
func (l *Log) Insert(log map[string]interface{}) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Exec(db.AR().Insert(Table_Log_Name, log))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// 根据关键字分页获取日志
func (l *Log) GetLogsByKeywordAndLimit(keyword string, limit int, number int) (logs []map[string]string, err error) {
	
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Log_Name).Where(map[string]interface{}{
		"message LIKE": "%" + keyword + "%",
	}).Limit(limit, number).OrderBy("log_id", "DESC"))
	if err != nil {
		return
	}
	logs = rs.Rows()
	
	return
}

// 分页获取日志
func (l *Log) GetLogsByLimit(limit int, number int) (logs []map[string]string, err error) {
	
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			From(Table_Log_Name).
			Limit(limit, number).
			OrderBy("log_id", "DESC"))
	if err != nil {
		return
	}
	logs = rs.Rows()
	
	return
}

// 获取日志总数
func (l *Log) CountLogs() (count int64, err error) {
	
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			Select("count(*) as total").
			From(Table_Log_Name))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

// 根据关键字获取日志总数
func (l *Log) CountLogsByKeyword(keyword string) (count int64, err error) {
	
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().
		Select("count(*) as total").
		From(Table_Log_Name).
		Where(map[string]interface{}{
		"message LIKE": "%" + keyword + "%",
	}))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}
