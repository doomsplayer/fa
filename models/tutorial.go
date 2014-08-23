package models

import (
	"fmt"
	"time"
)

type Tutorial struct {
	Id          int64     `form:"-"`
	Title       string    `form:"title" valid:"Required"`
	PicId       int       `form:"picid" valid:"Required"`
	Content     string    `form:"content" xorm:"text" valid:"Required"`
	Description string    `form:"description"`
	Type        string    `form:"type" valid:"Required"`
	Author      string    `form:"author"`
	Source      string    `form:"source"`
	Click       int       `form:"-"`
	Favor       int       `form:"-"`
	Time        time.Time `xorm:"updated"`
}

type TutorialSlice []Tutorial

func (t *Tutorial) Put() (err error) {
	t.Time = time.Now()
	_, err = Engine.Insert(t)
	return
}

func (t *Tutorial) Get() (err error) {
	q := Engine.Id(t.Id)
	if t.Type != `` {
		q = q.Where("type=?", t.Type)
	}
	b, err := q.Get(t)
	if !b {
		return fmt.Errorf("not exist")
	}
	return
}
func (t *Tutorial) Update() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	_, err = Engine.Id(t.Id).Cols("title", "content", "type", "description", "pic_id", "author", "source").Update(t)
	return
}

func (t *Tutorial) Delete() (err error) {
	_, err = Engine.Id(t.Id).Delete(t)
	return err
}

func (t *TutorialSlice) Hot(n, from int, tp string) (err error) {
	if n == 0 {
		n = 1
	}
	q := Engine.Desc("time").Limit(n, from)
	if tp != `` {
		q = q.Where("type=?", tp)
	}
	err = q.Find(t)
	return
}

func (t *TutorialSlice) All(n, from int, tp string) (err error) {

	q := Engine.Desc("id").Limit(n, from)
	if tp != `` {
		q = q.Where("type=?", tp)
	}
	err = q.Find(t)
	return
}

type TutorialType struct {
	Id          int64
	Name        string
	Description string
}
type TutorialTypeSlice []TutorialType

func (t *TutorialType) Put() (err error) {
	_, err = Engine.Insert(t)
	return
}

func (t *TutorialType) Delete() (err error) {
	_, err = Engine.Delete(t)
	return err
}

func (t *TutorialTypeSlice) GetAll() (err error) {
	err = Engine.Find(t)
	return
}
