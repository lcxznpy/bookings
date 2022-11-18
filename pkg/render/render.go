package render

import (
	"bytes"
	"fmt"
	"github.com/lcxznpy/bookings/pkg/config"
	"github.com/lcxznpy/bookings/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

// 添加数据
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	//如果已经使用了cache,则直接读取缓冲区内容
	if app.UseCache {
		tc = app.TemplateCache
	} else { //没有则创建template cache
		tc, _ = CreateTemplateCache()
	}
	//获取模板map
	//从app config 获得template cache,从这里获得太费时间
	//tc, err := CreateTemplateCache()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//如果map-tc中有tmpl对应的value值,ok为true,反之为false
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("cannot get template from templatecache")
	}
	//创建缓冲区
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	//将渲染模板的数据读入缓冲区
	_ = t.Execute(buf, td)
	//WriteTo从缓冲中读取数据直到缓冲内没有数据或遇到错误，并将这些数据写入w。
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error wrinting template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	//用来看是否访问了某个page,tmpl文件
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		//返回page路径中的最后一个元素，对应key值
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		//找到与字符串匹配的文件,对应value值
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}

		}
		myCache[name] = ts

	}
	return myCache, nil

}
