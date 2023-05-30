package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

// 获取tags
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	//maps用于指定过滤条件，可以传入一个结构体或映射，表示要应用的过滤条件。
	//使用 Where 方法将传入的 maps 参数应用为过滤条件，以筛选出符合条件的标签。
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

// 获取tag总数
func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

// 查找标签
func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)

	return tag.ID > 0
}

// 增加标签
func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})
	return true
}
