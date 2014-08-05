package models

import (
	"fmt"
	"time"
)

type Promotion struct {
	Id           int64
	Title        string
	Title2       string
	Description  string
	Type         string
	Link         string
	Description2 string
	Detail       string `xorm:"text"`
	PicId        int64
	UpdateTime   time.Time
	StartTime    time.Time
	EndTime      time.Time
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
	f.Id = p.PicId
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

	_, err = Engine.Insert(p)
	return
}

type PromotionSlice []Promotion

func (p *PromotionSlice) Recent(n int) (err error) {
	err = Engine.Desc("start_time").Limit(n).Find(p)
	return
}
func (p *PromotionSlice) Fresh(n int) (err error) {
	err = Engine.Desc("end_time").Limit(n).Find(p)
	return
}
func (p *PromotionSlice) Ending(n int) (err error) {
	err = Engine.Where("end_time > ?", time.Now()).Incr("end_time").Limit(n).Find(p)
	return
}
func (p *PromotionSlice) Hot(n int) (err error) {
	err = Engine.Desc("start_time").Limit(n).Find(p)
	return
}
