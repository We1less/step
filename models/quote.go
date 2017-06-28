// 引用
package models

import (
	"time"
	
	"github.com/astaxie/beego/orm"
)

type Quote struct {
	Id      int64
	Author  string
	Content string
	Extra   string
	Status  int
	Ctime   int64
	Utime   int64
}

// QuoteStatus 引用状态
func QuoteStatus() map[int]string {
	return map[int]string{
		1: "可用",
		2: "不可用",
	}
}

// QuoteStatusDesc 引用状态描述
func QuoteStatusDesc(id int) string {
	if desc, ok := QuoteStatus()[id]; ok {
		return desc
	}
	return "未知"
}

// QuoteSave 添加引用
func QuoteSave(q *Quote) (int64, error) {
	o := orm.NewOrm()
	q.Status = 1
	q.Ctime = time.Now().Unix()
	q.Utime = time.Now().Unix()
	return o.Insert(q)
}

// QuoteUpdate 更新引用
func QuoteUpdate(q * Quote) (int64, error) {
	o := orm.NewOrm()
	q.Utime = time.Now().Unix()
	return o.Update(q)
}

// QuoteList 引用列表
func QuoteList() []*Quote {
	var quote Quote
	var quotes []*Quote
	o := orm.NewOrm()
	o.QueryTable(quote).RelatedSel().Filter("Status", 1).All(&quotes)
	return quotes
}

// QuoteInfo 引用详情
func QuoteInfo(id int64) (Quote, error) {
	var q Quote
	o := orm.NewOrm()
	err := o.QueryTable(q).RelatedSel().Filter("Id", id).One(&q)
	return q, err
}

/*
CREATE TABLE `quote` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `author` varchar(64) NOT NULL COMMENT '作者',
  `content` varchar(255) NOT NULL COMMENT '引用内容',
  `extra` varchar(255) NOT NULL COMMENT '引用描述, 解说',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态1可用2不可用',
  `ctime` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `utime` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
*/