/*
LikeWeny
Fight for love
*/
package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Student struct {
	ID   string `gorm:"primaryKey;comment:身份证号"`
	Name string `gorm:"not null;index;comment:学生姓名"`
	// Gender uint `gorm:"default:0;not null;check:gender <> '^[0,1,2]{1}$';comment:性别"`
	// Birthday uint64 `gorm:"not null;comment:生日"`
	Address string `gorm:"not null;comment:家庭地址"`
	Phone   uint   `gorm:"not null;index;check:phone <> '^[0-9]{11}$';comment:监护人电话"`
	// Native string `gorm:not null;comment:籍贯`
	Nation  string `gorm:"not null;comment:民族"`
	Outlook uint   `gorm:"not null;check:outlook <> '^[0-9]{2}$';comment:政治面貌代码"`
	GroupID string `gorm:"unique;not null;comment:外键唯一标识"`
	Model
}

type Group struct {
	Name    string  `gorm:"not null;comment:小组名称"`
	Leader  string  `gorm:"unique;not null;comment:组长姓名"`
	ClassID string  `gorm:"primaryKey;comment:外键唯一标识"`
	Student Student `gorm:"foreignkey:GroupID;references:ClassID;comment:学生表"`
	Model
}

type Class struct {
	Name           string `gorm:"not null;comment:班级名称"`
	Director       string `gorm:"not null;index;comment:主任姓名"`
	DeputyDirector string `gorm:"comment:副主任姓名"`
	Phone          uint   `gorm:"check:phone <> '^[0-9]{11}$';comment:主任联系电话"`
	Monitor        string `gorm:"unique;index;not null;comment:班长姓名"`
	DeputyMonitor  string `gorm:"unique;index;comment:副班长姓名"`
	DepartmentID   string `gorm:"primaryKey;comment:外键唯一标识"`
	Group          Group  `gorm:"foreignkey:ClassID;references:DepartmentID;comment:小组表"`
	Model
}

type Department struct {
	Name           string `gorm:"primaryKey;comment:年级部门名称"`
	Director       string `gorm:"not null;index;comment:主任姓名"`
	DeputyDirector string `gorm:"comment:副主任姓名"`
	Phone          uint   `gorm:"check:phone <> '^[0-9]{11}$';comment:部门联系电话"`
	GradeID        string `gorm:"unique;not null;comment:外键唯一标识"`
	Class          Class  `gorm:"foreignkey:DepartmentID;references:GradeID;comment:班级表"`
	Model
}

type Grade struct {
	Name       string     `gorm:"primaryKey;comment:年级名称"`
	SchoolID   string     `gorm:"unique;not null;comment:外键唯一标识"`
	Department Department `gorm:"foreignkey:GradeID;references:SchoolID;comment:年级部门表"`
	Model
}

type School struct {
	ID        string `gorm:"primaryKey;comment:唯一标识"`
	Name      string `gorm:"not null;check:name <> '^[\u4e00-\u9fa5]$';index;comment:学校名称"`
	Address   string `gorm:"not null;comment:学校地址"`
	Principal string `gorm:"not null;comment:校长姓名"`
	Phone     uint   `gorm:"check:phone <> '^[0-9]{11}$';comment:学校联系电话"`
	Grade     Grade  `gorm:"foreignkey:SchoolID;references:ID;comment:年级表"`
	Model
}

func LinkDatabase(DatabaseInfo map[string]string) (*gorm.DB, error) {
	dsn := DatabaseInfo["UserName"] + ":" + DatabaseInfo["PassWd"] + "@tcp(" + DatabaseInfo["IP"] + ":" + DatabaseInfo["Port"] + ")/" + DatabaseInfo["NameDB"] + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "sms_",
		},
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
func Database(DatabaseInfo map[string]string) (map[string]string, error) {
	db, err := LinkDatabase(DatabaseInfo)
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&School{}, &Grade{}, &Department{}, &Class{}, &Group{}, Student{})
	// db.Create(&School{ID:"9C098D19-73B1-4F0D-8DCB-356B76AE8535",Name:"山东新华电脑学院",Address:"山东省济南市历下区燕山街道文化东路55号",Principal:"我不知道",Phone:0})
	// db.Create(&Grade{Name:"信息安全",SchoolID:"9C098D19-73B1-4F0D-8DCB-356B76AE8535"})
	// db.Create(&Department{Name:"1802级",Director:"宋鹏飞",DeputyDirector:"NULL",Phone:0,GradeID:"9C098D19-73B1-4F0D-8DCB-356B76AE8535"})
	// db.Create(&Class{Name:"二班",Director:"宋江",DeputyDirector:"NULL",Phone:0,Monitor:"王二麻子",DepartmentID:"9C098D19-73B1-4F0D-8DCB-356B76AE8535"})
	// db.Create(&Group{Name:"一组",Leader:"一组",ClassID:"9C098D19-73B1-4F0D-8DCB-356B76AE8535"})
	// db.Create(&Student{ID:"370921198208026666",Name:"靳灿奇",Address:"山东省济南市历下区燕山街道文化东路55号",Phone:17615848429,Nation:"汉族",Outlook:0,GroupID:"9C098D19-73B1-4F0D-8DCB-356B76AE8535"})
	return map[string]string{}, nil
}
