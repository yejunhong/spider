package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strings"
)

type Model struct{
	Db *gorm.DB
	Db61 *gorm.DB
	DbManhua *gorm.DB
}

func InitDb(ip, user, pass, database string) *gorm.DB {
	var cdn string = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", user, pass, ip, database)
  	var err error
	db, err := gorm.Open("mysql", cdn)
	// db, err := gorm.Open("mysql", "test:test@/spider?charset=utf8")
	if err != nil {
		panic(err)
	}
	db.SingularTable(true) // 不进行转换表名
	return db
}

/**
 *
 * 批量insert
 * @param table 表名
 * @param data 需要批量入库的数据
 *
 */
func(model *Model) BatchInsert(table string, data []map[string]interface{}, duplicate []string){
	var field []string
	var values []interface{}
	var PresetValues []string
	var PresetValue []string
	for _, val := range data {
		for k, _ := range val {
			field = append(field, k)
			PresetValue = append(PresetValue, "?")
		}
		break
	}
	
	for _, val := range data {
		for _, v := range field {
			values = append(values, val[v])
		}
		PresetValues = append(PresetValues, "(" + strings.Join(PresetValue, ",") + ")")
	}
	var PresetVal string = strings.Join(PresetValues, ",")
	var fieldStr string = "`" + strings.Join(field, "`,`") + "`"
	var sql string = fmt.Sprintf(`INSERT INTO %s(%s) VALUES %s `, table, fieldStr, PresetVal)

	if len(duplicate) > 0 { // 如果设置了 存在则更新，不存在则新增
		// ON DUPLICATE KEY UPDATE info = VALUES(info)
		var duplicate_val []string
		for _, v := range duplicate {
			duplicate_val = append(duplicate_val, fmt.Sprintf("`%s` = VALUES(`%s`)", v, v))
		}
		sql = sql + " ON DUPLICATE KEY UPDATE " + strings.Join(duplicate_val, ",")
		// fmt.Println(sql)
	}

	model.Db.Exec(sql, values...)
	// fmt.Println(err)
}

/**
 *
 * 批量insert
 * @param table 表名
 * @param data 需要批量入库的数据
 *
 */
 func DbBatchInsert(Db *gorm.DB, table string, data []map[string]interface{}, duplicate []string){
	var field []string
	var values []interface{}
	var PresetValues []string
	var PresetValue []string
	for _, val := range data {
		for k, _ := range val {
			field = append(field, k)
			PresetValue = append(PresetValue, "?")
		}
		break
	}
	
	for _, val := range data {
		for _, v := range field {
			values = append(values, val[v])
		}
		PresetValues = append(PresetValues, "(" + strings.Join(PresetValue, ",") + ")")
	}
	var PresetVal string = strings.Join(PresetValues, ",")
	var fieldStr string = "`" + strings.Join(field, "`,`") + "`"
	var sql string = fmt.Sprintf(`INSERT INTO %s(%s) VALUES %s `, table, fieldStr, PresetVal)

	if len(duplicate) > 0 { // 如果设置了 存在则更新，不存在则新增
		// ON DUPLICATE KEY UPDATE info = VALUES(info)
		var duplicate_val []string
		for _, v := range duplicate {
			duplicate_val = append(duplicate_val, fmt.Sprintf("`%s` = VALUES(`%s`)", v, v))
		}
		sql = sql + " ON DUPLICATE KEY UPDATE " + strings.Join(duplicate_val, ",")
		// fmt.Println(sql)
	}

	Db.Exec(sql, values...)
	// fmt.Println(err)
}