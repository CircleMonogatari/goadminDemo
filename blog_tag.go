package main

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetBlogTagTable(ctx *context.Context) table.Table {

	blogTag := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := blogTag.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).FieldFilterable()
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Created_on", "created_on", db.Int)
	info.AddField("Created_by", "created_by", db.Varchar)
	info.AddField("Modified_on", "modified_on", db.Int)
	info.AddField("Modified_by", "modified_by", db.Varchar)
	info.AddField("Deleted_on", "deleted_on", db.Int)
	info.AddField("State", "state", db.Tinyint)

	info.SetTable("blog_tag").SetTitle("BlogTag").SetDescription("BlogTag")

	formList := blogTag.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Created_on", "created_on", db.Int, form.Number)
	formList.AddField("Created_by", "created_by", db.Varchar, form.Text)
	formList.AddField("Modified_on", "modified_on", db.Int, form.Number)
	formList.AddField("Modified_by", "modified_by", db.Varchar, form.Text)
	formList.AddField("Deleted_on", "deleted_on", db.Int, form.Number)
	formList.AddField("State", "state", db.Tinyint, form.Number)

	formList.SetTable("blog_tag").SetTitle("BlogTag").SetDescription("BlogTag")

	return blogTag
}
