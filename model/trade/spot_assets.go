package trade

import (
	"github.com/shopspring/decimal"
	"github.com/weblazy/easy/db/emysql"
	"gorm.io/gorm"
)

var SpotAssetsHandler = &SpotAssets{}

type SpotAssets struct {
	Id   int64  `json:"id" gorm:"column:id;primary_key;type:int AUTO_INCREMENT"`
	Name string `json:"name" gorm:"column:name;type:varchar(100) NOT NULL;default:'';comment:'用户名';uniqueIndex"`
}

func (t *SpotAssets) TableName() string {
	return "spot_assets"
}

func (t *SpotAssets) Insert(db *gorm.DB, data *SpotAssets) error {
	if db == nil {
		db = GetDB()
	}
	return db.Create(data).Error
}
func (t *SpotAssets) BulkInsert(db *gorm.DB, fields []string, params []map[string]interface{}) error {
	if db == nil {
		db = GetDB()
	}
	return emysql.BulkInsert(db, t.TableName(), fields, params)
}
func (t *SpotAssets) BulkSave(db *gorm.DB, fields []string, params []map[string]interface{}) error {
	if db == nil {
		db = GetDB()
	}
	return emysql.BulkSave(db, t.TableName(), fields, params)
}
func (t *SpotAssets) Delete(db *gorm.DB, where string, args ...interface{}) error {
	if db == nil {
		db = GetDB()
	}
	return db.Where(where, args...).Delete(&SpotAssets{}).Error
}
func (t *SpotAssets) Updates(db *gorm.DB, data map[string]interface{}, where string, args ...interface{}) (int64, error) {
	if db == nil {
		db = GetDB()
	}
	db = db.Model(&SpotAssets{}).Where(where, args...).Updates(data)
	return db.RowsAffected, db.Error
}
func (t *SpotAssets) GetOne(where string, args ...interface{}) (*SpotAssets, error) {
	var obj SpotAssets
	return &obj, GetDB().Where(where, args...).Take(&obj).Error
}
func (*SpotAssets) GetList(where string, args ...interface{}) ([]*SpotAssets, error) {
	var list []*SpotAssets
	return list, GetDB().Where(where, args...).Find(&list).Error
}
func (t *SpotAssets) GetListWithLimit(limit int, where string, args ...interface{}) ([]*SpotAssets, error) {
	var list []*SpotAssets
	return list, GetDB().Where(where, args...).Limit(limit).Find(&list).Error
}
func (t *SpotAssets) GetListOrderLimit(order string, limit int, where string, args ...interface{}) ([]*SpotAssets, error) {
	var list []*SpotAssets
	if limit == 0 || limit > 10000 {
		limit = 10
	}
	return list, GetDB().Where(where, args...).Order(order).Limit(limit).Find(&list).Error
}
func (t *SpotAssets) GetListPage(pageNum, limit int, where string, args ...interface{}) ([]*SpotAssets, error) {
	var list []*SpotAssets
	offset := (pageNum - 1) * limit
	return list, GetDB().Where(where, args...).Offset(offset).Limit(limit).Find(&list).Error
}
func (t *SpotAssets) GetCount(where string, args ...interface{}) (int64, error) {
	var count int64
	return count, GetDB().Model(&SpotAssets{}).Where(where, args...).Count(&count).Error
}
func (t *SpotAssets) GetSumInt64(sql string, args ...interface{}) (int64, error) {
	type sum struct {
		Num int64 `json:"num" gorm:"column:num"`
	}
	var obj sum
	return obj.Num, GetDB().Raw(sql, args...).Scan(&obj).Error
}
func (t *SpotAssets) GetSumFloat64(sql string, args ...interface{}) (float64, error) {
	type sum struct {
		Num float64 `json:"num" gorm:"column:num"`
	}
	var obj sum
	return obj.Num, GetDB().Raw(sql, args...).Scan(&obj).Error
}
func (t *SpotAssets) GetSumDecimal(sql string, args ...interface{}) (decimal.Decimal, error) {
	type sum struct {
		Num decimal.Decimal `json:"num" gorm:"column:num"`
	}
	var obj sum
	return obj.Num, GetDB().Raw(sql, args...).Scan(&obj).Error
}
