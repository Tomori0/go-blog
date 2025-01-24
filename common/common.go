package common

import (
    "go-blog-zleng/config"
    "go-blog-zleng/model"
)

var Template model.HtmlTemplate

func LoadTemplate() {
    Template = model.InitTemplate(config.Cfg.System.CurrentDir + "/template")
}
