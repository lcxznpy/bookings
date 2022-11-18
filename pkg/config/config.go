package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template //每次访问网页,不需要访问磁盘来获取template,而是直接通过config中的Template获取整个template
	InfoLog       *log.Logger                   // 记录日志
	InProduction  bool                          //记录是否处理生产状态,是为true
	Session       *scs.SessionManager
}
