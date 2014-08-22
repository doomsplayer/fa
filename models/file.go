package models

import (
	. "badmintonhome/lib"
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type File struct {
	Id   int64
	Name string
	Path string
}

func (f *File) Put() (err error) {
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
	mdt := fmt.Sprintf("%x", md5.Sum([]byte(time.Now().String())))
	fnameb := fmt.Sprintf("%x", sha1.Sum([]byte(name+string(mdt[:]))))
	fname := string(fnameb[:])
	fname += identity + ext

	folder := "static/upload/"
	if len(ext) <= 1 {
		folder += `misc/`
	} else {
		folder += ext[1:] + `/`
	}
	os.MkdirAll(folder, 0755)

	fb, err := os.Create(folder + fname)
	E(err)

	_, err = fb.Write(data)
	E(err)

	fb.Close()

	f.Name = name
	f.Path = folder + fname
	f.Put()
	return
}

func (f *File) Get() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	has, err := Engine.Get(f)
	if !has {
		panic(fmt.Errorf("not exist"))
	}
	return
}

func (f *File) Delete() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	has, err := Engine.Id(f.Id).Get(f)
	E(err)
	if has != true {
		panic(fmt.Errorf("file not existed"))
	}
	err = os.Remove(f.Path)
	E(err)
	_, err = Engine.Id(f.Id).Delete(f)
	E(err)
	return
}

type FileSlice []File

func (this *FileSlice) GetAll(start, num int) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	err = Engine.Find(this)

	return
}
