package models

import (
	"fmt"
	"time"
)

type Video struct {
	Id          int64     `form:"-"`
	Title       string    `form:"title" valid:"Required"`
	PicId       int       `form:"picid" valid:"Required"`
	Content     string    `form:"content" xorm:"text" valid:"Required"`
	Type        string    `form:"type" valid:"Required"`
	Author      string    `form:"author"`
	Source      string    `form:"source"`
	Click       int       `form:"-"`
	Favor       int       `form:"-"`
	Description string    `form:"description"`
	Time        time.Time `xorm:"updated"`
}

type VideoSlice []Video

func (t *Video) Put() (err error) {
	t.Time = time.Now()
	_, err = Engine.Insert(t)
	return
}

func (t *Video) Get() (err error) {
	b, err := Engine.Get(t)
	if !b {
		return fmt.Errorf("not exist")
	}
	return
}

func (t *Video) Delete() (err error) {
	_, err = Engine.Id(t.Id).Delete(t)
	return err
}

func (t *VideoSlice) Hot(n, from int, tp string) (err error) {
	if n == 0 {
		n = 1
	}
	q := Engine.Desc("time").Limit(n, from)
	if tp != `` {
		q = q.Where("type=?", tp)
	}
	err = q.Find(t)
	fmt.Println(t)
	return
}

func (t *VideoSlice) All(n, from int, tp string) (err error) {
	q := Engine.Desc("id").Limit(n, from)
	if tp != `` {
		q = q.Where("type=?", tp)
	}
	err = q.Find(t)
	return
}

type VideoType struct {
	Id          int64
	Name        string
	Description string
}
type VideoTypeSlice []VideoType

func (t *VideoType) Put() (err error) {
	_, err = Engine.Insert(t)
	return
}

func (t *VideoType) Delete() (err error) {
	_, err = Engine.Delete(t)
	return err
}

func (t *VideoTypeSlice) GetAll() (err error) {
	err = Engine.Find(t)
	return
}
