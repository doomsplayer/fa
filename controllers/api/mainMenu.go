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
	url := `#`
	menu := []*Menu{}
	me := new(Menu)
	me.Title = "促销信息"
	me.Url = "#/portfolio"
	menu = append(menu, me)

	me = new(Menu)
	me.Title = `学打羽毛球`
	me.Url = "#/learnBadminton"
	me.AddSub(`大话羽球`, `#`)
	me.AddSub(`羽球知识`, url)
	me.AddSub(`羽球技术`, url)
	me.AddSub(`羽球战术`, url)
	me.AddSub(`伤病防护`, url)
	menu = append(menu, me)

	me = new(Menu)
	me.Title = `视频库`
	me.Url = `#/videoLib`
	me.AddSub(`国际大赛专辑`, `#/wordChampion`)
	me.AddSub(`经典大赛专辑`, url)
	me.AddSub(`玩转羽球`, url)
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
