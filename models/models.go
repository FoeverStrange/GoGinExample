package models

//数据库连接与初始化
import (
	"GoGinExample/pkg/setting"
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// 为所有模型，如文章标签提供基础结构，便于其他模型共享
type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}

// updateTimeStampForCreateCallback在创建时设定CreatedOn, ModifiedOn
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `ModifyTime` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	//如果没有手动指定要更新的列，那么会执行SetColumn
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("modifiedOn", time.Now().Unix())
	}
}
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}
		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
		//判断执行硬删除还是软删除，如果Search.Unscoped为默认(false自动过滤软删除
		//获取我们约定的删除字段，若存在则 UPDATE 软删除，若不存在则 DELETE 硬删除
		if !scope.Search.Unscoped && hasDeletedOnField {
			//scope.Raw生成原始SQL字段
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
}

func CloseDB() {
	defer db.Close()
}
