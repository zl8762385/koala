package koala

import (
	output "koala/v1/output"
	"net/http"
)

// 输出接口
type IOutput interface {
	// 文本
	Content(rw http.ResponseWriter, value interface{}) error
}

// 初始化接口组件
var (
	Text IOutput = output.Text{}
	Json IOutput = output.Json{}
)

