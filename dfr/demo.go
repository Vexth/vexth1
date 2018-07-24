package dfr

import (
	"fmt"

	db "github.com/vexth1/db/mysql"
)

// SPU 结构
type SPU struct {
	ID       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	SPUCode  string `json:"SPUcode" form:"SPUcode"`
	Abstract string `json:"abstract" form:"abstract"`
	Details  string `json:"details" form:"details"`
	ClassID  int    `json:"classid" form:"classid"`
	BrandID  int    `json:"brandid" form:"brandid"`
}

// AddSPU 新增
func (s *SPU) AddSPU() bool {
	rs, err := db.SQLDB.Exec("INSERT INTO spu(id, name, SPUcode, abstract, details, classid, brandid) VALUES (?, ?, ?, ?, ?, ?, ?)", s.ID, s.Name, s.SPUCode, s.Abstract, s.Details, s.ClassID, s.BrandID)
	if err != nil {
		return false
	}
	id, err := rs.LastInsertId()
	fmt.Println(id)
	if err != nil {
		return false
	}

	return true
}

// EditSPU 修改记录
func (s *SPU) EditSPU() bool {
	rs, err := db.SQLDB.Exec("UPDATE sup set name=?,SPUcode=?,abstract=?,details=? where id=?", s.Name, s.SPUCode, s.Abstract, s.Details, s.ID)
	if err != nil {
		return false
	}
	id, err := rs.RowsAffected()
	fmt.Println(id)
	if err != nil {
		return false
	}
	return true
}

// DeleteSPU 删除记录
func DeleteSPU(ID int) bool {
	rs, err := db.SQLDB.Exec("Delete From sup where id=?", ID)
	if err != nil {
		return false
	}
	id, err := rs.RowsAffected()
	fmt.Println(id)
	if err != nil {
		return false
	}
	return true
}

// GetSUPList 得到记录列表
func GetSUPList(pageno, pagesize int, search string) (persons []SPU) {

	fmt.Println("搜索参数:" + search)
	persons = make([]SPU, 0)
	//SQL查询分页语句
	if search != "" {
		// rows, err := db.SQLDB.Query("SELECT * FROM person")
		rows, err := db.SQLDB.Query("SELECT id name SPUcode abstract details classid brandid FROM spu where 1=1 and classid like '%"+search+"%' or brandid like '%"+search+"%' limit ?,?", (pageno-1)*pagesize, pagesize)
		if err != nil {
			return nil
		}
		defer rows.Close()

		//数据添加到数据集中
		for rows.Next() {
			var s SPU
			rows.Scan(&s.ID, &s.ClassID, &s.BrandID)
			persons = append(persons, s)
		}
		if err = rows.Err(); err != nil {
			return nil
		}

	} else {
		rows, err := db.SQLDB.Query("SELECT id name SPUcode abstract details classid brandid FROM spu where 1=1  limit ?,?", (pageno-1)*pagesize, pagesize)
		if err != nil {
			return nil
		}
		defer rows.Close()

		// 数据添加到数据集中
		for rows.Next() {
			fmt.Println(rows.Next())
			var s SPU
			rows.Scan(&s.ID, &s.ClassID, &s.BrandID)
			persons = append(persons, s)
		}
		if err = rows.Err(); err != nil {
			return nil
		}
	}
	return persons
}

// GetRecordNum 得到记录数
func GetRecordNum(search string) int {
	num := 0

	//SQL查询分页语句
	if search != "" {
		rows, err := db.SQLDB.Query("SELECT id name SPUcode abstract details classid brandid FROM spu where 1=1 and classid like '%?%' or brandid '%?%'", search, search)
		if err != nil {
			return 0
		}
		defer rows.Close()

		//数据添加到数据集中
		for rows.Next() {
			num++
		}

	} else {
		rows, err := db.SQLDB.Query("SELECT id name SPUcode abstract details classid brandid FROM spu where 1=1")
		if err != nil {
			return 0
		}
		defer rows.Close()

		//数据添加到数据集中
		for rows.Next() {
			num++
		}

	}
	return num
}
