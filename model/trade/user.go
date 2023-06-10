package trade

import (
	"github.com/shopspring/decimal"
	"github.com/weblazy/easy/db/emysql"
	"gorm.io/gorm"
)

var UserHandler = &User{}

type User struct {
	Id   int64  `json:"id" gorm:"column:id;primary_key;type:int AUTO_INCREMENT"`
	Name string `json:"name" gorm:"column:name;type:varchar(100) NOT NULL;default:'';comment:'用户名';uniqueIndex"`
}

func (t *User) TableName() string {
	return "user"
}

func (t *User) Insert(db *gorm.DB, data *User) error {
	if db == nil {
		db = GetDB()
	}
	return db.Create(data).Error
}
func (t *User) BulkInsert(db *gorm.DB, fields []string, params []map[string]interface{}) error {
	if db == nil {
		db = GetDB()
	}
	return emysql.BulkInsert(db, t.TableName(), fields, params)
}
func (t *User) BulkSave(db *gorm.DB, fields []string, params []map[string]interface{}) error {
	if db == nil {
		db = GetDB()
	}
	return emysql.BulkSave(db, t.TableName(), fields, params)
}
func (t *User) Delete(db *gorm.DB, where string, args ...interface{}) error {
	if db == nil {
		db = GetDB()
	}
	return db.Where(where, args...).Delete(&User{}).Error
}
func (t *User) Updates(db *gorm.DB, data map[string]interface{}, where string, args ...interface{}) (int64, error) {
	if db == nil {
		db = GetDB()
	}
	db = db.Model(&User{}).Where(where, args...).Updates(data)
	return db.RowsAffected, db.Error
}
func (t *User) GetOne(where string, args ...interface{}) (*User, error) {
	var obj User
	return &obj, GetDB().Where(where, args...).Take(&obj).Error
}
func (*User) GetList(where string, args ...interface{}) ([]*User, error) {
	var list []*User
	return list, GetDB().Where(where, args...).Find(&list).Error
}
func (t *User) GetListWithLimit(limit int, where string, args ...interface{}) ([]*User, error) {
	var list []*User
	return list, GetDB().Where(where, args...).Limit(limit).Find(&list).Error
}
func (t *User) GetListOrderLimit(order string, limit int, where string, args ...interface{}) ([]*User, error) {
	var list []*User
	if limit == 0 || limit > 10000 {
		limit = 10
	}
	return list, GetDB().Where(where, args...).Order(order).Limit(limit).Find(&list).Error
}
func (t *User) GetListPage(pageNum, limit int, where string, args ...interface{}) ([]*User, error) {
	var list []*User
	offset := (pageNum - 1) * limit
	return list, GetDB().Where(where, args...).Offset(offset).Limit(limit).Find(&list).Error
}
func (t *User) GetCount(where string, args ...interface{}) (int64, error) {
	var count int64
	return count, GetDB().Model(&User{}).Where(where, args...).Count(&count).Error
}
func (t *User) GetSumInt64(sql string, args ...interface{}) (int64, error) {
	type sum struct {
		Num int64 `json:"num" gorm:"column:num"`
	}
	var obj sum
	return obj.Num, GetDB().Raw(sql, args...).Scan(&obj).Error
}
func (t *User) GetSumFloat64(sql string, args ...interface{}) (float64, error) {
	type sum struct {
		Num float64 `json:"num" gorm:"column:num"`
	}
	var obj sum
	return obj.Num, GetDB().Raw(sql, args...).Scan(&obj).Error
}
func (t *User) GetSumDecimal(sql string, args ...interface{}) (decimal.Decimal, error) {
	type sum struct {
		Num decimal.Decimal `json:"num" gorm:"column:num"`
	}
	var obj sum
	return obj.Num, GetDB().Raw(sql, args...).Scan(&obj).Error
}
