package models

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"strings"
	"time"
)

type Promotion struct {
	Id           int64     `form:"-"`
	Title        string    `form:"title" valid:"Required;MaxSize(256)"`
	Title2       string    `form:"title2" valid:"Required;MaxSize(256)"`
	Description  string    `form:"description" valid:"Required;"`
	Type         string    `form:"type" valid:"Required;"`
	Link         string    `form:"link" valid:""`
	Description2 string    `form:"description2" valid:"Required;"`
	Detail       string    `xorm:"text" form:"detail" valid:""`
	PicId        int       `form:"picid" valid:"Required;"`
	UpdateTime   time.Time `xorm:"updated" form:"-"`
	StartTime    time.Time `form:"-"`
	EndTime      time.Time `form:"-"`
}

func (p *Promotion) Get() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	has, err := Engine.Get(p)
	if !has {
		panic(fmt.Errorf("not exist"))
	}
	return
}

func (p *Promotion) Delete() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	f := new(File)
	f.Id = int64(p.PicId)
	f.Delete()

	_, err = Engine.Delete(p)
	return
}

func (p *Promotion) Put() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	p.UpdateTime = time.Now()
	_, err = Engine.Insert(p)
	return
}

type PromotionSlice []Promotion

func (p *PromotionSlice) Recent(n int, from int, tp string) (err error) {
	tp = strings.TrimSpace(tp)

	var query *xorm.Session
	if tp != `` {
		query = Engine.Where("type=?", tp).Limit(n, from)
	} else {
		query = Engine.Limit(n, from)
	}
	err = query.Desc("start_time").Find(p)
	return
}
func (p *PromotionSlice) Fresh(n int, from int, tp string) (err error) {
	tp = strings.TrimSpace(tp)

	var query *xorm.Session
	if tp != `` {
		query = Engine.Where("type=?", tp).Limit(n, from)
	} else {
		query = Engine.Limit(n, from)
	}

	err = query.Desc("end_time").Find(p)
	return
}
func (p *PromotionSlice) Ending(n int, from int, tp string) (err error) {
	tp = strings.TrimSpace(tp)

	var query *xorm.Session
	if tp != `` {
		query = Engine.Where("type=?", tp).Limit(n, from)
	} else {
		query = Engine.Limit(n, from)
	}

	err = query.Where("end_time > ?", time.Now()).Incr("end_time").Find(p)
	return
}
func (p *PromotionSlice) Hot(n int, from int, tp string) (err error) {
	tp = strings.TrimSpace(tp)

	var query *xorm.Session
	if tp != `` {
		query = Engine.Where("type=?", tp).Limit(n, from)
	} else {
		query = Engine.Limit(n, from)
	}
	err = query.Desc("start_time").Find(p)
	return
}

func (p *PromotionSlice) GetAll(from, num int) (err error) {
	err = Engine.Limit(num, from).Find(p)
	return
}

type PromotionType struct {
	Id   int64
	Name string
}

func (p *PromotionType) Put() (err error) {
	_, err = Engine.Insert(p)
	return
}
func (p *PromotionType) Delete() (err error) {
	_, err = Engine.Id(p.Id).Delete(p)
	return
}

type PromotionTypeSlice []PromotionType

func (p *PromotionTypeSlice) GetAll() (err error) {
	err = Engine.Find(p)
	return
}
