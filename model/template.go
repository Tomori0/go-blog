package model

import (
	"html/template"
	"io"
	"log"
	"time"
)

type TemplateBlog struct {
	*template.Template
}

type HtmlTemplate struct {
	Index      TemplateBlog
	Category   TemplateBlog
	Custom     TemplateBlog
	Detail     TemplateBlog
	Login      TemplateBlog
	Pigeonhole TemplateBlog
	Writing    TemplateBlog
}

func InitTemplate(templateDir string) HtmlTemplate {
	tp := readTemplate(
		[]string{
			"index",
			"category",
			"custom",
			"detail",
			"login",
			"pigeonhole",
			"writing",
		},
		templateDir,
	)
	htmlTemplate := HtmlTemplate{
		Index:      tp[0],
		Category:   tp[1],
		Custom:     tp[2],
		Detail:     tp[3],
		Login:      tp[4],
		Pigeonhole: tp[5],
		Writing:    tp[6],
	}
	return htmlTemplate
}

func (t *TemplateBlog) WriteData(w io.Writer, data interface{}) {
	err := t.Execute(w, data)
	if err != nil {
		log.Fatal(err)
		w.Write([]byte(err.Error()))
	}
}

func readTemplate(templates []string, templateDir string) []TemplateBlog {
	var tbs []TemplateBlog
	for _, view := range templates {
		viewHtml := view + ".html"
		t := template.New(viewHtml)
		t.Funcs(template.FuncMap{"isEven": IsEven, "getNextName": GetNextName, "date": Date, "dateDay": DateDay})
		home := templateDir + "/home.html"
		header := templateDir + "/layout/header.html"
		footer := templateDir + "/layout/footer.html"
		pagination := templateDir + "/layout/pagination.html"
		personal := templateDir + "/layout/personal.html"
		post := templateDir + "/layout/post-list.html"
		t, err := t.ParseFiles(templateDir+"/"+viewHtml, home, header, footer, personal, post, pagination)
		if err != nil {
			log.Println(err)
		}
		tbs = append(tbs, TemplateBlog{t})
	}
	return tbs
}

func IsEven(index int) bool {
	return index%2 == 0
}

func GetNextName(arr []string, index int) string {
	return arr[index+1]
}

func Date(format string) string {
	return time.Now().Format(format)
}

func DateDay(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}
