package main

import "github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"

// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
//
// http://{{config.Domain}}:{{Port}}/{{config.Prefix}}/info/{{key}}
//
// example:
//
// "blog_article" => http://localhost:9033/admin/info/blog_article
// "blog_auth" => http://localhost:9033/admin/info/blog_auth
// "blog_tag" => http://localhost:9033/admin/info/blog_tag
//
// example end
//
var Generators = map[string]table.Generator{
	"blog_article": GetBlogArticleTable,
	"blog_auth":    GetBlogAuthTable,
	"blog_tag":     GetBlogTagTable,

	"license/management": nil,
	"license/application":nil,
	// generators end
}
