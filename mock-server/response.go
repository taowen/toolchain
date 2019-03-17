package main

import (
	"encoding/json"
	"net/http"
)

func genHttpResponse(res *Response,w http.ResponseWriter) {
	var bRes []byte
	switch res.ResContentType {
	case ContentTypeJson:
		bRes, _ = json.Marshal(res.ResBody)
	}
	w.Header().Set("content-type",res.ResContentType)
	w.WriteHeader(res.ResHttpCode)
	w.Write(bRes)
}

func genResponse(res *Response,w http.ResponseWriter) {
	switch res.Protocol {
	case ProtocolHTTP:
		genHttpResponse(res,w)
	}
}