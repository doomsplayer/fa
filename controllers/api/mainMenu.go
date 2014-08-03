package api

import (
	"badmintonhome/controllers/lib"
)

type MainMenu struct {
	lib.BaseController
}

// @Title mainBar
// @Description 获得首页目录
// @Success 200 {string} 目录json
// @Failure 404 Not found
// @router /mainBar [get]
func (m *MainMenu) Menu() {
	type MainMenu struct {
		Title    string
		Url      string
		SubTitle []struct {
			Title string
			Url   string
		}
	}
	menu := []MainMenu{}

	menu = append(menu,
		MainMenu{
			"促销信息",
			"",
			[]struct {
				Title string
				Url   string
			}{},
		})

	m.Data[`json`] = menu
	m.ServeJson()
}
