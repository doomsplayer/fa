package models

import (
	. "badmintonhome/lib"
	"crypto/md5"
	"crypto/sha1"
	"os"
	"path/filepath"
	"time"
)

type File struct {
	Id   int64
	Name string
	Path string
}

func (f *File) Save() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	count, err := Engine.Id(f.Id).Count(f)
	E(err)
	if count == 0 {
		_, err = Engine.Insert(f)
		E(err)
	} else {
		_, err = Engine.Update(f)
		E(err)
	}
	return nil
}

func (f *File) SaveToFile(name string, identity string, data []byte) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	ext := filepath.Ext(name)
	name = filepath.Base(name)
	mdt := md5.Sum([]byte(time.Now().String()))
	fnameb := sha1.Sum([]byte(name + string(mdt[:])))
	fname := string(fnameb[:])
	fname += identity + ext

	folder := "static/upload/"
	if len(ext) <= 1 {
		folder += `misc/`
	} else {
		folder += ext[1:] + `/`
	}

	fb, err := os.Create(folder + fname)
	E(err)

	_, err = fb.Write(data)
	E(err)

	fb.Close()

	f.Name = name
	f.Path = folder + fname
	f.Save()
	return
}

func (f *File) GetFileById(id int64) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	err = Engine.Id(id).Find(f)
	return
}

func (f *File) Delete() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	_, err = Engine.Id(f.Id).Delete(f)
	return
}
