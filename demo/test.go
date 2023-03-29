package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jacyluo/fytPlus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()
	r.GET("/", GetDirectUrl)
	r.Run(":8000")
}

type FangYiTong struct {
	ApiUrl string `yaml:"ApiUrl"`
	Appid  string `yaml:"Appid"`
	Token  string `yaml:"Token"`
	Key    string `yaml:"Key"`
}
type Database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
}

type Config struct {
	FangYiTong FangYiTong `yaml:"FangYiTong"`
	Database   Database   `yaml:"database"`
}

func GetConfig(conf *Config) error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	path += "/demo/config.yml"
	configFile, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(configFile, conf)
	if err != nil {
		return err
	}
	return nil
}

func GetDirectUrl(c *gin.Context) {
	var conf Config
	err := GetConfig(&conf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	fyt := fytPlus.FangYiTong{
		ApiUrl: conf.FangYiTong.ApiUrl,
		Appid:  conf.FangYiTong.Appid,
		Key:    conf.FangYiTong.Key,
		Token:  conf.FangYiTong.Token,
	}
	var res fytPlus.FytRes

	req := fytPlus.RedirectUrlReq{
		State:  "xinzhi",
		Attach: "xinzhiAgent",
		Scope:  "snsapi_base",
	}

	err = fyt.GetRedirectUrl(&req, &res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": res,
		//"config": conf,
	})
}
