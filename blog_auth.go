package main

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetBlogAuthTable(ctx *context.Context) table.Table {

	blogAuth := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := blogAuth.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).FieldFilterable()
	info.AddField("Username", "username", db.Varchar)
	info.AddField("Password", "password", db.Varchar)

	info.SetTable("blog_auth").SetTitle("博客用户").SetDescription("博客用户")

	formList := blogAuth.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField("Username", "username", db.Varchar, form.Text)
	formList.AddField("Password", "password", db.Varchar, form.Password)

	formList.SetTable("blog_auth").SetTitle("博客用户").SetDescription("博客用户")

	return blogAuth
}
