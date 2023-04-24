package framework

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"text/template"
)

type IResponse interface {
	// json输出
	Json(obj any) IResponse

	// jsonp输出
	Jsonp(obj any) IResponse

	// xml输出
	Xml(obj any) IResponse

	// html输出
	Html(file string, obj any) IResponse

	// string
	Text(format string, values ...any) IResponse

	// 重定向
	Redirect(path string) IResponse

	// header
	SetHeader(key string, val string) IResponse

	// Cookie
	SetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse

	// 设置状态码
	SetStatus(code int) IResponse

	// 设置200状态
	SetOkStatus() IResponse
}

func (ctx *Context) Json(obj any) IResponse {
	byt, err := json.Marshal(obj)
	if err != nil {
		return ctx.SetStatus(http.StatusInternalServerError)
	}

	ctx.SetHeader("Content-Type", "application/json")
	ctx.responseWriter.Write(byt)

	return ctx
}

func (ctx *Context) Jsonp(obj any) IResponse {
	// 获取请求参数callback
	callbackFunc, _ := ctx.QueryString("callback", "callback_function")
	ctx.SetHeader("Content-Type", "application/javascript")

	// 输出到前端页面的时候需要注意下进行字符过滤，否则有可能造成xss攻击
	callback := template.JSEscapeString(callbackFunc)

	// 输出函数名
	_, err := ctx.responseWriter.Write([]byte(callback))
	if err != nil {
		return ctx
	}

	// 输出左括号
	_, err = ctx.responseWriter.Write([]byte("("))

	if err != nil {
		return ctx
	}

	// 数据函数参数
	ret, err := json.Marshal(obj)
	if err != nil {
		return ctx
	}

	_, err = ctx.responseWriter.Write(ret)

	if err != nil {
		return ctx
	}

	// 输出右括号
	_, err = ctx.responseWriter.Write([]byte(")"))
	if err != nil {
		return ctx
	}

	return ctx
}

func (ctx *Context) Xml(obj any) IResponse {
	byt, err := xml.Marshal(obj)
	if err != nil {
		return ctx.SetStatus(http.StatusInternalServerError)
	}
	ctx.SetHeader("Content-Type", "application/xml")
	ctx.responseWriter.Write(byt)

	return ctx
}

// html输出
func (ctx *Context) Html(file string, obj any) IResponse {
	t, err := template.New("output").ParseFiles(file)
	if err != nil {
		return ctx
	}

	if err := t.Execute(ctx.responseWriter, obj); err != nil {
		return ctx
	}

	ctx.SetHeader("Content-Type", "application/html")

	return ctx
}

// string
func (ctx *Context) Text(format string, values ...any) IResponse {
	out := fmt.Sprintf(format, values...)
	ctx.SetHeader("Content-Type", "application/text")
	ctx.responseWriter.Write([]byte(out))

	return ctx
}

// 重定向
func (ctx *Context) Redirect(path string) IResponse {
	http.Redirect(ctx.responseWriter, ctx.request, path, http.StatusMovedPermanently)
	return ctx
}

func (ctx *Context) SetHeader(key string, val string) IResponse {
	ctx.responseWriter.Header().Add(key, val)
	return ctx
}

func (ctx *Context) SetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse {
	if path == "" {
		path = "/"
	}
	http.SetCookie(ctx.responseWriter, &http.Cookie{
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

func (ctx *Context) SetStatus(code int) IResponse {
	ctx.responseWriter.WriteHeader(code)
	return ctx
}

// 设置200状态
func (ctx *Context) SetOkStatus() IResponse {
	ctx.responseWriter.WriteHeader(http.StatusOK)
	return ctx
}
