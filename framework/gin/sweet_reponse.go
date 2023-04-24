// Copyright 2023 ranen.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package gin

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"text/template"
)

type IResponse interface {
	// json输出
	IJson(obj any) IResponse

	// jsonp输出
	IJsonp(obj any) IResponse

	// xml输出
	IXml(obj any) IResponse

	// html输出
	IHtml(file string, obj any) IResponse

	// string
	IText(format string, values ...any) IResponse

	// 重定向
	IRedirect(path string) IResponse

	// header
	ISetHeader(key string, val string) IResponse

	// Cookie
	ISetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse

	// 设置状态码
	ISetStatus(code int) IResponse

	// 设置200状态
	ISetOkStatus() IResponse
}

func (ctx *Context) IJson(obj any) IResponse {
	byt, err := json.Marshal(obj)
	if err != nil {
		return ctx.ISetStatus(http.StatusInternalServerError)
	}

	ctx.ISetHeader("Content-Type", "application/json")
	ctx.Writer.Write(byt)

	return ctx
}

func (ctx *Context) IJsonp(obj any) IResponse {
	// 获取请求参数callback
	callbackFunc := ctx.Query("callback")
	ctx.ISetHeader("Content-Type", "application/javascript")

	// 输出到前端页面的时候需要注意下进行字符过滤，否则有可能造成xss攻击
	callback := template.JSEscapeString(callbackFunc)

	// 输出函数名
	_, err := ctx.Writer.Write([]byte(callback))
	if err != nil {
		return ctx
	}

	// 输出左括号
	_, err = ctx.Writer.Write([]byte("("))

	if err != nil {
		return ctx
	}

	// 数据函数参数
	ret, err := json.Marshal(obj)
	if err != nil {
		return ctx
	}

	_, err = ctx.Writer.Write(ret)

	if err != nil {
		return ctx
	}

	// 输出右括号
	_, err = ctx.Writer.Write([]byte(")"))
	if err != nil {
		return ctx
	}

	return ctx
}

func (ctx *Context) IXml(obj any) IResponse {
	byt, err := xml.Marshal(obj)
	if err != nil {
		return ctx.ISetStatus(http.StatusInternalServerError)
	}
	ctx.ISetHeader("Content-Type", "application/xml")
	ctx.Writer.Write(byt)

	return ctx
}

// html输出
func (ctx *Context) IHtml(file string, obj any) IResponse {
	t, err := template.New("output").ParseFiles(file)
	if err != nil {
		return ctx
	}

	if err := t.Execute(ctx.Writer, obj); err != nil {
		return ctx
	}

	ctx.ISetHeader("Content-Type", "application/html")

	return ctx
}

// string
func (ctx *Context) IText(format string, values ...any) IResponse {
	out := fmt.Sprintf(format, values...)
	ctx.ISetHeader("Content-Type", "application/text")
	ctx.Writer.Write([]byte(out))

	return ctx
}

// 重定向
func (ctx *Context) IRedirect(path string) IResponse {
	http.Redirect(ctx.Writer, ctx.Request, path, http.StatusMovedPermanently)
	return ctx
}

func (ctx *Context) ISetHeader(key string, val string) IResponse {
	ctx.Writer.Header().Add(key, val)
	return ctx
}

func (ctx *Context) ISetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse {
	if path == "" {
		path = "/"
	}
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     key,
		Value:    val,
		MaxAge:   maxAge,
		Path:     path,
		Domain:   domain,
		Secure:   secure,
		HttpOnly: httpOnly,
	})

	return ctx
}

func (ctx *Context) ISetStatus(code int) IResponse {
	ctx.Writer.WriteHeader(code)
	return ctx
}

// 设置200状态
func (ctx *Context) ISetOkStatus() IResponse {
	ctx.Writer.WriteHeader(http.StatusOK)
	return ctx
}
