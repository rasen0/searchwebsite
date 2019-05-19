package nsq

type header struct {
	class string
	len   int32
}

type nsqData struct {
	header
	data []interface{}
}

// 网站类型 webType、 网站连接 href、 网站简述 summary
type analyseMessage struct {
	webType string
	summary string
	href    string
}
