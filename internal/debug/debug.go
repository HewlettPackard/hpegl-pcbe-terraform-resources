package debug

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httputil"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-http-go"
)

type InspectionOptions struct {
	Enabled         bool
	RequestHeaders  *abstractions.RequestHeaders
	ResponseHeaders *abstractions.ResponseHeaders
}
type InspectionHandler struct {
	options InspectionOptions
}

func (middleware InspectionHandler) Intercept(
	pipeline nethttplibrary.Pipeline,
	middlewareIndex int,
	req *http.Request,
) (*http.Response, error) {
	var bodyBytes []byte
	var err error

	if req.Body != nil {
		bodyBytes, err = io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}
	reqBytes, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		return nil, err
	}

	tflog.Info(req.Context(),
		"\n\n->TX\n\n"+
			string(reqBytes)+
			"--\n")

	resp, err := pipeline.Next(req, middlewareIndex)

	respBytes, _ := httputil.DumpResponse(resp, true)
	tflog.Info(req.Context(),
		"\n\n<-RX\n\n"+
			string(respBytes)+
			"\n--\n")

	return resp, err
}

func NewInspectionOptions() *InspectionOptions {
	return &InspectionOptions{
		Enabled:         false,
		RequestHeaders:  abstractions.NewRequestHeaders(),
		ResponseHeaders: abstractions.NewResponseHeaders(),
	}
}

func NewInspectionHandlerWithOptions(
	options InspectionOptions,
) *InspectionHandler {
	return &InspectionHandler{options: options}
}
