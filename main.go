package main

import (
	"embed"
	_ "embed"
	"math/rand"
	"os"
	"time"
	"web/app"

	"github.com/andrewarrow/feedback/network"
	"github.com/andrewarrow/feedback/prefix"
	"github.com/andrewarrow/feedback/router"
)

//go:embed app/feedback.json
var embeddedFile []byte

//go:embed views/*.html
var embeddedTemplates embed.FS

//go:embed assets/**/*.*
var embeddedAssets embed.FS

var buildTag string

func main() {

	rand.Seed(time.Now().UnixNano())
	if len(os.Args) == 1 {
		return
	}
	arg := os.Args[1]

	prefix.FeedbackName = ""
	network.BaseUrl = os.Getenv("BACKEND")
	network.BaseHeaderKey = "AA-Service"
	network.BaseHeaderValue = "backend"
	router.BuildTag = buildTag
	router.EmbeddedTemplates = embeddedTemplates
	router.EmbeddedAssets = embeddedAssets
	r := router.NewRouter("NO_DATABASE_URL_TO_TURN_OFF_PSQL", embeddedFile)

	if arg == "init" {
		//router.InitNewApp()
	} else if arg == "run" {
		r.Paths["/"] = app.HandleWelcome
		//r.Paths["apple-app-site-association"] = app.HandleApple
		//r.Paths["admin"] = app.HandleAdmin
		//r.Paths["media"] = app.HandleMedia
		r.BearerAuthFunc = nil
		r.CookieAuthFunc = nil

		r.ListenAndServe(":" + os.Args[2])
	} else if arg == "t" {
	} else if arg == "slugs" {
		/*
			var regexSlug = regexp.MustCompile(`[^a-zA-Z0-9]+`)
			offsetInt := 0
			offset := ""
			for {
				list := r.All("tag", "order by tag", offset)
				for _, item := range list {
					tag := item["tag"].(string)
					result := regexSlug.ReplaceAllString(strings.ToLower(tag), "-")
					r.FreeFormUpdate("update tags set slug=$1 where tag=$2", result, tag)
					fmt.Println(result)
				}
				if len(list) == 0 {
					break
				}
				offsetInt += 30
				offset = fmt.Sprintf("%d", offsetInt)
			}*/
	} else if arg == "tr" {
	}
}
