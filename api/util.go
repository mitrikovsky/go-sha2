package api

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

type output struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func outputJsonMessageResult(ctx *fasthttp.RequestCtx, code int, r string) {
	// Write content-type, statuscode, payload
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Response.Header.SetStatusCode(code)
	out := output{code, r}
	jsonResult, _ := json.Marshal(out)
	fmt.Fprint(ctx, string(jsonResult))
	ctx.Response.Header.Set("Connection", "close")
}

func outputJSON(ctx *fasthttp.RequestCtx, code int, result interface{}) {
	// Marshal provided interface into JSON structure
	jsonResult, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Error(err)
		outputJsonMessageResult(ctx, 500, "json.MarshalIndent: "+err.Error())
		return
	}
	// Write content-type, statuscode, payload
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Response.SetStatusCode(code)
	fmt.Fprint(ctx, string(jsonResult))
	ctx.Response.Header.Set("Connection", "close")
}
