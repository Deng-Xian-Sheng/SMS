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
	ID           string `gorm:"type:varchar(191);primaryKey;comment:唯一标识"`
	IdentityCard string `gorm:"not null;unique;index;comment:身份证号"`
	Name         string `gorm:"not null;index;check:name <> '^[\u4e00-\u9fa5]{4}$';comment:学生姓名"`
	// Gender uint `gorm:"default:0;not null;check:gender <> '^[0,1,2]{1}$';comment:性别"`
	// Birthday uint64 `gorm:"not null;check: birthday <> '^[0-9]$';comment:生日"`
	Address string `gorm:"not null;comment:家庭地址"`
	Phone   uint   `gorm:"not null;index;check:phone <> '^[0-9]{11}$';comment:监护人电话"`
	// Native string `gorm:not null;check: native <> '^[\u4e00-\u9fa5]$';comment:籍贯`
	Nation  string `gorm:"not null;check:nation <> '^[\u4e00-\u9fa5]$';comment:民族"`
	Outlook uint   `gorm:"not null;check:outlook <> '^[0-9]{2}$';comment:政治面貌代码"`
	GroupID string `gorm:"not null;comment:外键"`
	Role    []Role `gorm:"foreignKey:StudentID;references:ID;comment:角色表"`
	Model
}

type Group struct {
	ID      string    `gorm:"type:varchar(191);primaryKey;comment:唯一标识"`
	Name    string    `gorm:"not null;check:name <> '^{16}$';comment:小组名称"`
	Leader  string    `gorm:"unique;not null;check:leader <> '^[\u4e00-\u9fa5]{4}$';comment:组长姓名"`
	ClassID string    `gorm:"not null;comment:外键"`
	Student []Student `gorm:"foreignKey:GroupID;references:ID;comment:学生表"`
	Role    []Role    `gorm:"foreignKey:GroupID;references:ID;comment:角色表"`
	Model
}

type Class struct {
	ID             string  `gorm:"type:varchar(191);primaryKey;comment:唯一标识"`
	Name           string  `gorm:"not null;check:name <> '^{16}$';comment:班级名称"`
	Director       string  `gorm:"not null;index;check:director <> '^[\u4e00-\u9fa5]{4}$';comment:主任姓名"`
	DeputyDirector string  `gorm:"check:deputydirector <> '^[\u4e00-\u9fa5]{4}$';comment:副主任姓名"`
	Phone          uint    `gorm:"check:phone <> '^[0-9]{11}$';comment:主任联系电话"`
	Monitor        string  `gorm:"unique;index;not null;check:monitor <> '^[\u4e00-\u9fa5]{4}$';comment:班长姓名"`
	DeputyMonitor  string  `gorm:"unique;index;check:deputymonitor <> '^[\u4e00-\u9fa5]{4}$';comment:副班长姓名"`
	DepartmentID   string  `gorm:"not null;comment:外键"`
	Group          []Group `gorm:"foreignKey:ClassID;references:ID;comment:小组表"`
	Role           []Role  `gorm:"foreignKey:ClassID;references:ID;comment:角色表"`
	Model
}

type Department struct {
	ID             string  `gorm:"type:varchar(191);primaryKey;comment:唯一标识"`
	Name           string  `gorm:"not null;check:name <> '^{16}$';comment:年级部门名称"`
	Director       string  `gorm:"not null;index;check:director <> '^[\u4e00-\u9fa5]{4}$';comment:主任姓名"`
	DeputyDirector string  `gorm:"check:deputydirector <> '^[\u4e00-\u9fa5]{4}$';comment:副主任姓名"`
	Phone          uint    `gorm:"check:phone <> '^[0-9]{11}$';comment:部门联系电话"`
	GradeID        string  `gorm:"not null;comment:外键"`
	Class          []Class `gorm:"foreignKey:DepartmentID;references:ID;comment:班级表"`
	Role           []Role  `gorm:"foreignKey:DepartmentID;references:ID;comment:角色表"`
	Model
}

type Grade struct {
	ID         string       `gorm:"type:varchar(191);primaryKey;comment:唯一标识"`
	Name       string       `gorm:"not null;unique;check:name <> '^{16}$';comment:年级名称"`
	SchoolID   string       `gorm:"not null;comment:外键"`
	Department []Department `gorm:"foreignKey:GradeID;references:ID;comment:年级部门表"`
	Role       []Role       `gorm:"foreignKey:GradeID;references:ID;comment:角色表"`
	Model
}

type School struct {
	ID              string  `gorm:"type:varchar(191);primaryKey;comment:唯一标识"`
	Name            string  `gorm:"not null;check:name <> '^[\u4e00-\u9fa5]{16}$';index;comment:学校名称"`
	Address         string  `gorm:"not null;comment:学校地址"`
	Principal       string  `gorm:"not null;check:principal <> '^[\u4e00-\u9fa5]{4}$';comment:校长姓名"`
	DeputyPrincipal string  `gorm:"check:deputyprincipal <> '^[\u4e00-\u9fa5]{4}$';comment:副校长姓名"`
	Phone           uint    `gorm:"check:phone <> '^[0-9]{11}$';comment:学校联系电话"`
	Grade           []Grade `gorm:"foreignKey:SchoolID;references:ID;comment:年级表"`
	Role            []Role  `gorm:"SchoolID:UserNumber;references:ID;comment:角色表"`
	Model
}

type User struct {
	ID       string `gorm:"type:varchar(191);primaryKey;comment:唯一标识"`
	Name     string `gorm:"not null;unique;check:name <> '^[a-zA-Z0-9_-]{2,16}$';comment:用户名"`
	Passwd   string `gorm:"comment:用户密码"`
	NickName string `gorm:"check nickname <> '^{16}$';comment:昵称"`
	RoleID   string `gorm:"not null;comment:外键"`
	Model
}

type Role struct {
	ID           string `gorm:"type:varchar(191);primaryKey;comment:唯一标识"`
	Name         uint   `gorm:"not null;comment:角色代码"` //以0开始，0代表学校；逐渐递增
	SchoolID     string `gorm:"comment:外键"`
	GradeID      string `gorm:"comment:外键"`
	DepartmentID string `gorm:"comment:外键"`
	ClassID      string `gorm:"comment:外键"`
	GroupID      string `gorm:"comment:外键"`
	StudentID    string `gorm:"comment:外键"`
	User         []User `gorm:"foreignKey:RoleID;references:ID;comment:用户表"`
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
	db.AutoMigrate(&School{}, &Grade{}, &Department{}, &Class{}, &Group{}, &Student{}, &Role{}, &User{})
	return db, nil
}

// func Database(DatabaseInfo map[string]string) (map[string]string, error) {
// 	db, err := LinkDatabase(DatabaseInfo)
// 	if err != nil {
// 		return nil, err
// 	}
	
// 	return map[string]string{}, nil
// }
