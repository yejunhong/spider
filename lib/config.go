package lib

import (
  "fmt"
  "gopkg.in/yaml.v2"
  "io/ioutil"
)

type BookMysql struct {
  User string
  Pass string
  Host string
  Port string
	Name string
	Prefix string
}

type YamlConf struct {
  BookMysql BookMysql
}

var Conf *YamlConf

func LoadConfig() {

  yamlText, _ := ioutil.ReadFile("./config.yaml")
  // fmt.Println(string(yamlText))
  Conf = &YamlConf{}
  //把yaml形式的字符串解析成struct类型
  var err error = yaml.Unmarshal(yamlText, &Conf)
  fmt.Println(err)
  if Conf.BookMysql.Host == "" {
      fmt.Println("配置文件设置错误")
  } else {
    fmt.Println("初始数据", Conf)
  }
  
}
