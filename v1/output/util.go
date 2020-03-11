package output

import "net/http"

// 设置内容类型
func writeContentType(rw http.ResponseWriter, value []string) {
	// 获取response header
	header := rw.Header()

	//fmt.Printf("%v %v", header, value)
	// 如果没有找到 就去设置
	if val := header["Content-Type"]; len(val)==0 {
		header["Content-Type"] = value
	}
}
