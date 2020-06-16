package datamodel
//
//import (
//	"github.com/GoAdminGroup/go-admin/modules/config"
//	template2 "github.com/GoAdminGroup/go-admin/template"
//	"github.com/GoAdminGroup/go-admin/template/types"
//	"html/template"
//)
//
//func GetContent() (types.Panel, error) {
//
//	components := template2.Get(config.Get().THEME)
//	colComp := components.Col()
//
//	infobox := components.InfoBox().
//		SetText("CPU TRAFFIC").
//		SetColor("blue").
//		SetNumber("41,410").
//		SetIcon("ion-ios-gear-outline").
//		GetContent()
//
//	var size = map[string]string{"md": "3", "sm": "6", "xs": "12"}
//	infoboxCol1 := colComp.SetSize(size).SetContent(infobox).GetContent()
//	row1 := components.Row().SetContent(infoboxCol1).GetContent()
//
//	return types.Panel{
//		Content:     row1,
//		Title:       "Dashboard",
//		Description: "this is a example",
//	}, nil
//}