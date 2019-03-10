package main

type Response struct {
	Protocol       string                 `json:"protocol"`
	ResHttpCode    int                    `json:"res_http_code"`
	ResContentType string                 `json:"res_content_type"`
	ResBody        map[string]interface{} `json:"res_body"`
}

func NewDefaultRes()*Response {
	return &Response{
		Protocol:"http",
		ResHttpCode:200,
		ResContentType:ContentTypeJson,
		ResBody:map[string]interface{}{
			"err_no":0,
			"err_msg":"success",
		},
	}
}

func NewNotFoundUrlRes()*Response {
	return &Response{
		Protocol:"http",
		ResHttpCode:200,
		ResContentType:ContentTypeJson,
		ResBody:map[string]interface{}{
			"err_no":"1000",
			"err_msg":"url hit nothing",
		},
	}
}

type ResConf map[string]*Response


const (
	ContentTypeJson = "application/json"
	ProtocolHTTP = "http"
)

