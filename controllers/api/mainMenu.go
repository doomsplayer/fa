package api

import (
	"badmintonhome/controllers/lib"
)

type MainMenuApi struct {
	lib.BaseController
}

// @Title mainBar
// @Description 获得首页目录
// @Success 200 {string} 目录的json
// @Failure 404 Not found
// @router /mainBar [get]
func (m *MainMenuApi) Menu() {
	// url := `#`
	menu := []*Menu{}
	me := new(Menu)
	me.Title = "促销信息"
	me.Url = "#/portfolio"
	menu = append(menu, me)

	me = new(Menu)
	me.Title = `学打羽毛球`
	me.Url = "#/learnBadminton"
	me.AddSub(`大话羽球`, `#/articleAlbum/大话羽球`)
	me.AddSub(`羽球知识`, `#/articleAlbum/羽球知识`)
	me.AddSub(`羽球技术`, `#/articleAlbum/羽球技术`)
	me.AddSub(`羽球战术`, `#/articleAlbum/羽球战术`)
	me.AddSub(`伤病防护`, `#/articleAlbum/伤病防护`)
	menu = append(menu, me)

	me = new(Menu)
	me.Title = `视频库`
	me.Url = `#/videoLib`
	me.AddSub(`国际大赛`, `#/videoAlbum/国际大赛`)
	me.AddSub(`经典对战`, `#/videoAlbum/经典对战`)
	me.AddSub(`玩转羽球`, `#/videoAlbum/玩转羽球`)
	menu = append(menu, me)

	me = new(Menu)
	me.Title = `关于我们`
	me.Url = `#/about-us`
	menu = append(menu, me)

	m.Data[`json`] = menu
	m.ServeJson()
}

type Menu struct {
	Title    string
	Url      string
	SubTitle []struct {
		Title string
		Url   string
	}
}

func (m *Menu) AddSub(title string, url string) {
	m.SubTitle = append(m.SubTitle,
		struct {
			Title string
			Url   string
		}{title, url})
}
