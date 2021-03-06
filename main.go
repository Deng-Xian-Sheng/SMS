package main

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"sms/service"
)

type DefaultYaml struct {
	Database struct {
		UserName string
		PassWd   string
		IP       string
		Port     string
		NameDB   string
	}
	Service struct {
		Port string
	}
}

func GrabYaml() (map[string]string, error) {
	FilePath := "./config.yml"
	OsFile, err := os.OpenFile(FilePath, os.O_RDONLY, os.ModeAppend|os.ModePerm)
	if err != nil {
		var StartData = `database:
    username: #数据库用户名
    passwd: #数据库用户密码
    ip: #数据库远程连接IP
    port: #数据库远程端口
    namedb: #数据库名
service:
    port: #HTTP服务端口
`
		CreateFile, err := os.OpenFile(FilePath, os.O_RDWR|os.O_CREATE, os.ModeAppend|os.ModePerm)
		if err != nil {
			return map[string]string{"code": "1", "data": "创建配置文件失败，请检查当前目录权限"}, err
		}
		defer CreateFile.Close()
		//写入文件时，使用带缓存的 *Writer
		WriteFile := bufio.NewWriter(CreateFile)
		WriteFile.WriteString(StartData + "\n")
		//Flush将缓存的文件真正写入到文件中
		WriteFile.Flush()
		return map[string]string{"code": "2", "data": "已在当前目录下创建配置文件config.yml，请填写配置后重新运行程序"}, nil
	}
	defer OsFile.Close()
	//获取文件内容
	FileScanner := bufio.NewScanner(OsFile)
	var FileData string
	// 设置分词方式(按行读取)
	FileScanner.Split(bufio.ScanLines)
	for FileScanner.Scan(){
		if FileData == "" {
			FileData = FileData + FileScanner.Text()
		} else {
			FileData = FileData + "\n" + FileScanner.Text()
		}
	}

	var defaultYaml DefaultYaml
	err = yaml.Unmarshal([]byte(FileData), &defaultYaml)
	if err != nil {
		return map[string]string{"code": "3", "data": "配置文件内容有误，无法解析配置文件"}, err
	}
	if defaultYaml.Database.UserName == "" || defaultYaml.Database.PassWd == "" || defaultYaml.Database.IP == "" || defaultYaml.Database.Port == "" || defaultYaml.Database.NameDB == "" || defaultYaml.Service.Port == "" {
		return map[string]string{"code": "4", "data": "配置文件必填项不能为空"}, nil
	}
	return map[string]string{"code": "0", "DatabaseUserName": defaultYaml.Database.UserName, "DatabasePassWd": defaultYaml.Database.PassWd, "DatabaseIP": defaultYaml.Database.IP, "DatabasePort": defaultYaml.Database.Port, "DatabaseNameDB": defaultYaml.Database.NameDB, "ServicePort": defaultYaml.Service.Port}, nil
}

func main() {
	YamlOut, err := GrabYaml()
	var DatabaseInfo map[string]string
	var ServiceInfo map[string]string
	if err != nil {
		panic(YamlOut["data"] + "\n" + fmt.Sprint(err))
	}
	if YamlOut["code"] == "2" {
		fmt.Println(YamlOut["data"])
		os.Exit(0)
	} else if YamlOut["code"] == "4" {
		fmt.Println(YamlOut["data"])
		os.Exit(1)
	} else if YamlOut["code"] == "0" {
		DatabaseInfo = map[string]string{"UserName": YamlOut["DatabaseUserName"], "PassWd": YamlOut["DatabasePassWd"], "IP": YamlOut["DatabaseIP"], "Port": YamlOut["DatabasePort"], "NameDB": YamlOut["DatabaseNameDB"]}
		ServiceInfo = map[string]string{"Port": YamlOut["ServicePort"]}
	}
	service.Service(ServiceInfo, DatabaseInfo)
}
