package trade

import (
	"github.com/shopspring/decimal"
	"github.com/weblazy/easy/db/emysql"
	"gorm.io/gorm"
)

var FutureOrderHandler = &FutureOrder{}

type FutureOrder struct {
	Id   int64  `json:"id" gorm:"column:id;primary_key;type:int AUTO_INCREMENT"`
	Name string `json:"name" gorm:"column:name;type:varchar(100) NOT NULL;default:'';comment:'用户名';uniqueIndex"`
}

func (t *FutureOrder) TableName() string {
	return "future_order"
}

func (t *FutureOrder) Insert(db *gorm.DB, data *FutureOrder) error {
	if db == nil {
		db = GetDB()
	}
	return db.Create(data).Error
}
func (t *FutureOrder) BulkInsert(db *gorm.DB, fields []string, params []map[string]interface{}) error {
	if db == nil {
		db = GetDB()
	}
	return emysql.BulkInsert(db, t.TableName(), fields, params)
}
func (t *FutureOrder) BulkSave(db *gorm.DB, fields []string, params []map[string]interface{}) error {
	if db == nil {
		db = GetDB()
	}
	return emysql.BulkSave(db, t.TableName(), fields, params)
}
func (t *FutureOrder) Delete(db *gorm.DB, where string, args ...interface{}) error {
	if db == nil {
		db = GetDB()
	}
	return db.Where(where, args...).Delete(&FutureOrder{}).Error
}
func (t *FutureOrder) Updates(db *gorm.DB, data map[string]interface{}, where string, args ...interface{}) (int64, error) {
	if db == nil {
		db = GetDB()
	}
	db = db.Model(&FutureOrder{}).Where(where, args...).Updates(data)
	return db.RowsAffected, db.Error
}
func (t *FutureOrder) GetOne(where string, args ...interface{}) (*FutureOrder, error) {
	var obj FutureOrder
	return &obj, GetDB().Where(where, args...).Take(&obj).Error
}
func (*FutureOrder) GetList(where string, args ...interface{}) ([]*FutureOrder, error) {
	var list []*FutureOrder
	return list, GetDB().Where(where, args...).Find(&list).Error
}
func (t *FutureOrder) GetListWithLimit(limit int, where string, args ...interface{}) ([]*FutureOrder, error) {
	var list []*FutureOrder
	return list, GetDB().Where(where, args...).Limit(limit).Find(&list).Error
}
func (t *FutureOrder) GetListOrderLimit(order string, limit int, where string, args ...interface{}) ([]*FutureOrder, error) {
	var list []*FutureOrder
	if limit == 0 || limit > 10000 {
		limit = 10
	}
	return list, GetDB().Where(where, args...).Order(order).Limit(limit).Find(&list).Error
}
func (t *FutureOrder) GetListPage(pageNum, limit int, where string, args ...interface{}) ([]*FutureOrder, error) {
	var list []*FutureOrder
	offset := (pageNum - 1) * limit
	return list, GetDB().Where(where, args...).Offset(offset).Limit(limit).Find(&list).Error
}
func (t *FutureOrder) GetCount(where string, args ...interface{}) (int64, error) {
	var count int64
	return count, GetDB().Model(&FutureOrder{}).Where(where, args...).Count(&count).Error
}
func (t *FutureOrder) GetSumInt64(sql string, args ...interface{}) (int64, error) {
	type sum struct {
		Num int64 `json:"num" gorm:"column:num"`
	}
	var obj sum
	return obj.Num, GetDB().Raw(sql, args...).Scan(&obj).Error
}
func (t *FutureOrder) GetSumFloat64(sql string, args ...interface{}) (float64, error) {
	type sum struct {
		Num float64 `json:"num" gorm:"column:num"`
	}
	var obj sum
	return obj.Num, GetDB().Raw(sql, args...).Scan(&obj).Error
}
func (t *FutureOrder) GetSumDecimal(sql string, args ...interface{}) (decimal.Decimal, error) {
	type sum struct {
		Num decimal.Decimal `json:"num" gorm:"column:num"`
	}
	var obj sum
	return obj.Num, GetDB().Raw(sql, args...).Scan(&obj).Error
}
