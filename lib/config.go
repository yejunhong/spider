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
  Db_caiji BookMysql
  Db_xiaoshuo BookMysql
  Db_manhua BookMysql
}

func LoadConfig() *YamlConf {

  yamlText, _ := ioutil.ReadFile("./config.yaml")
  // fmt.Println(string(yamlText))
  var Conf *YamlConf = &YamlConf{}
  //把yaml形式的字符串解析成struct类型
  var err error = yaml.Unmarshal(yamlText, &Conf)
  fmt.Println(err)
  if Conf.Db_caiji.Host == "" {
      fmt.Println("配置文件设置错误")
  } else {
    fmt.Println("初始数据", Conf)
  }
  return Conf
}
