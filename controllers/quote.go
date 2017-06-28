// 引用
package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"

	"step/models"
)

type QuoteController struct {
	beego.Controller
}

// Index 引用列表
func (q *QuoteController) Index() {
	q.Data["quotes"] = models.QuoteList()
	q.Layout = "base.html"
	q.TplName = "quote/index.html"
}

// Create 添加引用
func (q *QuoteController) Create() {
	if q.Ctx.Request.Method == "POST" {
		author := q.Input().Get("author")
		content := q.Input().Get("content")
		extra := q.Input().Get("extra")
		if strings.TrimSpace(author) == "" {
			q.Redirect("/quote/create", 302)
		}
		if strings.TrimSpace(content) == "" {
			q.Redirect("/quote/create", 302)
		}
		if strings.TrimSpace(extra) == "" {
			q.Redirect("/quote/create", 302)
		}
		var id int64
		var err error
		if id, err = models.QuoteSave(&models.Quote{
			Author:  author,
			Content: content,
			Extra:   extra,
		}); err != nil {
			q.Redirect("/quote/create", 302)
		}
		q.Redirect("/quote/"+strconv.FormatInt(id, 10), 302)
	}

	q.Data["isNewRecord"] = true
	q.Data["quote"] = models.Quote{}
	q.Layout = "base.html"
	q.TplName = "quote/create.html"
}

// View 引用详情
func (q *QuoteController) View() {
	id, err := strconv.Atoi(q.Ctx.Input.Param(":id"))
	if err != nil || id <= 0 {
		q.Redirect("/quote", 302)
	}
	quote, err := models.QuoteInfo(int64(id))
	if err != nil {
		q.Redirect("/user", 302)
	}
	q.Data["quote"] = quote
	q.Data["status"] = models.QuoteStatusDesc(quote.Status)
	q.Data["ctime"] = time.Unix(quote.Ctime, 0).Format("2006-01-02 15:04:05")
	q.Data["utime"] = time.Unix(quote.Utime, 0).Format("2006-01-02 15:04:05")
	q.Layout = "base.html"
	q.TplName = "quote/view.html"
}

// Update 编辑引用
func (q *QuoteController) Update() {
	q.Layout = "base.html"
	q.TplName = "quote/update.html"
}

// Delete 删除引用
func (q *QuoteController) Delete() {
	id, err := strconv.Atoi(q.Ctx.Input.Param(":id"))
	if err != nil || id <= 0 {
		q.Redirect("/quote", 302)
	}
	quote, err := models.QuoteInfo(int64(id))
	if err != nil {
		q.Redirect("/quote", 302)
	}
	quote.Status = 2
	if _, err := models.QuoteUpdate(&quote); err != nil {
		fmt.Println(err)
	}
	q.Redirect("/quote", 302)
}