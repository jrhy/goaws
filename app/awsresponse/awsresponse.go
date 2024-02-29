package awsresponse

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type Encoder interface {
	Encode(v any) error
}

func NewEncoder(req *http.Request, w http.ResponseWriter) Encoder {
	contentType := req.Header.Get("Content-Type")
	switch contentType {
	case "application/x-amz-json-1.0":
		enc := json.NewEncoder(w)
		enc.SetIndent("  ", "    ")
		w.Header().Set("Content-Type", "application/json")
		return enc
	default: // presumably "application/x-www-form-urlencoded":
		enc := xml.NewEncoder(w)
		enc.Indent("  ", "    ")
		w.Header().Set("Content-Type", "application/xml")
		return enc
	}
}
