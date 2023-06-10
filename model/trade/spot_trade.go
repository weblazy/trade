package trade

import (
	"github.com/shopspring/decimal"
	"github.com/weblazy/easy/db/emysql"
	"gorm.io/gorm"
)

var SpotTradeHandler = &SpotTrade{}

type SpotTrade struct {
	Id   int64  `json:"id" gorm:"column:id;primary_key;type:int AUTO_INCREMENT"`
	Name string `json:"name" gorm:"column:name;type:varchar(100) NOT NULL;default:'';comment:'用户名';uniqueIndex"`
}

func (t *SpotTrade) TableName() string {
	return "spot_trade"
}

func (t *SpotTrade) Insert(db *gorm.DB, data *SpotTrade) error {
	if db == nil {
		db = GetDB()
	}
	return db.Create(data).Error
}
func (t *SpotTrade) BulkInsert(db *gorm.DB, fields []string, params []map[string]interface{}) error {
	if db == nil {
		db = GetDB()
	}
	return emysql.BulkInsert(db, t.TableName(), fields, params)
}
func (t *SpotTrade) BulkSave(db *gorm.DB, fields []string, params []map[string]interface{}) error {
	if db == nil {
		db = GetDB()
	}
	return emysql.BulkSave(db, t.TableName(), fields, params)
}
func (t *SpotTrade) Delete(db *gorm.DB, where string, args ...interface{}) error {
	if db == nil {
		db = GetDB()
	}
	return db.Where(where, args...).Delete(&SpotTrade{}).Error
}
func (t *SpotTrade) Updates(db *gorm.DB, data map[string]interface{}, where string, args ...interface{}) (int64, error) {
	if db == nil {
		db = GetDB()
	}
	db = db.Model(&SpotTrade{}).Where(where, args...).Updates(data)
	return db.RowsAffected, db.Error
}
func (t *SpotTrade) GetOne(where string, args ...interface{}) (*SpotTrade, error) {
	var obj SpotTrade
	return &obj, GetDB().Where(where, args...).Take(&obj).Error
}
func (*SpotTrade) GetList(where string, args ...interface{}) ([]*SpotTrade, error) {
	var list []*SpotTrade
	return list, GetDB().Where(where, args...).Find(&list).Error
}
func (t *SpotTrade) GetListWithLimit(limit int, where string, args ...interface{}) ([]*SpotTrade, error) {
	var list []*SpotTrade
	return list, GetDB().Where(where, args...).Limit(limit).Find(&list).Error
}
func (t *SpotTrade) GetListOrderLimit(order string, limit int, where string, args ...interface{}) ([]*SpotTrade, error) {
	var list []*SpotTrade
	if limit == 0 || limit > 10000 {
		limit = 10
	}
	return list, GetDB().Where(where, args...).Order(order).Limit(limit).Find(&list).Error
}
func (t *SpotTrade) GetListPage(pageNum, limit int, where string, args ...interface{}) ([]*SpotTrade, error) {
	var list []*SpotTrade
	offset := (pageNum - 1) * limit
	return list, GetDB().Where(where, args...).Offset(offset).Limit(limit).Find(&list).Error
}
func (t *SpotTrade) GetCount(where string, args ...interface{}) (int64, error) {
	var count int64
	return count, GetDB().Model(&SpotTrade{}).Where(where, args...).Count(&count).Error
}
func (t *SpotTrade) GetSumInt64(sql string, args ...interface{}) (int64, error) {
	type sum struct {
		Num int64 `json:"num" gorm:"column:num"`
	}
	var obj sum
	return obj.Num, GetDB().Raw(sql, args...).Scan(&obj).Error
}
func (t *SpotTrade) GetSumFloat64(sql string, args ...interface{}) (float64, error) {
	type sum struct {
		Num float64 `json:"num" gorm:"column:num"`
	}
	var obj sum
	return obj.Num, GetDB().Raw(sql, args...).Scan(&obj).Error
}
func (t *SpotTrade) GetSumDecimal(sql string, args ...interface{}) (decimal.Decimal, error) {
	type sum struct {
		Num decimal.Decimal `json:"num" gorm:"column:num"`
	}
	var obj sum
	return obj.Num, GetDB().Raw(sql, args...).Scan(&obj).Error
}
