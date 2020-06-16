package main

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetBlogArticleTable(ctx *context.Context) table.Table {

	blogArticle := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := blogArticle.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).FieldFilterable()
	info.AddField("Tag_id", "tag_id", db.Int)
	info.AddField("Title", "title", db.Varchar)
	info.AddField("Desc", "desc", db.Varchar)
	info.AddField("Content", "content", db.Text)
	info.AddField("Cover_image_url", "cover_image_url", db.Varchar)
	info.AddField("Created_on", "created_on", db.Int)
	info.AddField("Created_by", "created_by", db.Varchar)
	info.AddField("Modified_on", "modified_on", db.Int)
	info.AddField("Modified_by", "modified_by", db.Varchar)
	info.AddField("Deleted_on", "deleted_on", db.Int)
	info.AddField("State", "state", db.Tinyint)

	info.SetTable("blog_article").SetTitle("BlogArticle").SetDescription("BlogArticle")

	formList := blogArticle.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField("Tag_id", "tag_id", db.Int, form.Number)
	formList.AddField("Title", "title", db.Varchar, form.Text)
	formList.AddField("Desc", "desc", db.Varchar, form.Text)
	formList.AddField("Content", "content", db.Text, form.RichText)
	formList.AddField("Cover_image_url", "cover_image_url", db.Varchar, form.Text)
	formList.AddField("Created_on", "created_on", db.Int, form.Number)
	formList.AddField("Created_by", "created_by", db.Varchar, form.Text)
	formList.AddField("Modified_on", "modified_on", db.Int, form.Number)
	formList.AddField("Modified_by", "modified_by", db.Varchar, form.Text)
	formList.AddField("Deleted_on", "deleted_on", db.Int, form.Number)
	formList.AddField("State", "state", db.Tinyint, form.Number)

	formList.SetTable("blog_article").SetTitle("BlogArticle").SetDescription("BlogArticle")

	return blogArticle
}
