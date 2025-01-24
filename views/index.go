package views

import (
	"go-blog-zleng/common"
	"go-blog-zleng/config"
	"go-blog-zleng/model"
	"net/http"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	//model.InitTemplate(config.Cfg.System.CurrentDir)

	categories := []model.Category{
		{
			Cid:  1,
			Name: "Go",
		},
	}

	postMores := []model.PostMore{
		{
			Pid:          1,
			Title:        "Go Blog",
			Content:      "Content",
			UserName:     "Zeff",
			ViewCount:    123,
			CreateAt:     "2025-01-22",
			CategoryId:   1,
			CategoryName: "Go",
			Type:         0,
		},
	}

	hr := &model.HomeResp{
		Viewer:    config.Cfg.Viewer,
		Categorys: categories,
		Posts:     postMores,
		Total:     1,
		Page:      1,
		Pages:     []int{1},
		PageEnd:   true,
	}
	index.WriteData(w, hr)
	//t.Execute(w, hr)
}
