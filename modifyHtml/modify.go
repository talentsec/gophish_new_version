package modifyHtml

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)
var (
	ajaxSrc = "https://unpkg.com/ajax-hook@2.0.3/dist/ajaxhook.min.js"
	ajaxScript = "\n<script src=\"https://unpkg.com/ajax-hook@2.0.3/dist/ajaxhook.min.js\"></script>\n"
	callAjax = "\n<script>\nah.proxy({\nonRequest:(config,handler)=>{\nconfig.body=config.body+\"&__original_url__=\"+config.url;\nconfig.url=window.location.href;\nconsole.log(config);\nhandler.next(config);\n},\nonError:(err,handler)=>{\nconsole.log(err);\nhandler.next(err);\n},\nonResponse:(response,handler)=>{\nconsole.log(response);\nhandler.next(response);\n}\n})\n</script>\n"
	callAjaxContent = "\nah.proxy({\nonRequest:(config,handler)=>{\nconfig.body=config.body+\"&__original_url__=\"+config.url;\nconfig.url=window.location.href;\nconsole.log(config);\nhandler.next(config);\n},\nonError:(err,handler)=>{\nconsole.log(err);\nhandler.next(err);\n},\nonResponse:(response,handler)=>{\nconsole.log(response);\nhandler.next(response);\n}\n})\n"
	jquerySrc = "http://ajax.aspnetcdn.com/ajax/jQuery/jquery-1.8.0.js"
	jqueryScript = "\n<script src=\"http://ajax.aspnetcdn.com/ajax/jQuery/jquery-1.8.0.js\"></script>\n"
	callJqueryContent = "\n$(window).ready(function(){\nvar int=self.setInterval(function(){\n$(\"form\").attr(\"action\", \"\");\n},1000);\n});\n"
	callJquery = "\n<script>\n$(window).ready(function(){\nvar int=self.setInterval(function(){\n$(\"form\").attr(\"action\", \"\");\n},1000);\n});\n</script>\n"
	)


func ModifyHtml(myHtml string) string {
	isHaveAjaxSrc := false
	isHaveAjaxCode := false
	isHaveJquerySrc := false
	isHaveJqueryCode := false
	d, _ := goquery.NewDocumentFromReader(strings.NewReader(myHtml))
	myScript := d.Find("script")
	myScript.Each(func(i int, f *goquery.Selection) {
		if t, _ := f.Attr("src"); strings.EqualFold(t, ajaxSrc) {
			isHaveAjaxSrc = true
		}
		if u, _ := f.Attr("src"); strings.EqualFold(u, jquerySrc) {
			isHaveJquerySrc = true
		}
		if s := f.Text(); strings.EqualFold(deleteSpace(s), callAjaxContent) {
			isHaveAjaxCode = true
		}
		if v := f.Text(); strings.EqualFold(v, callJqueryContent) {
			isHaveJqueryCode = true
		}
	})

	newHtml := myHtml
	if !isHaveAjaxSrc {
		bodyIndex := strings.Index(newHtml, "</body>")
		newHtml = newHtml[:bodyIndex] + ajaxScript + newHtml[bodyIndex:]
	}

	if !isHaveAjaxCode {
		bodyIndex := strings.Index(newHtml, "</body>")
		newHtml = newHtml[:bodyIndex] + callAjax + newHtml[bodyIndex:]
	}

	if !isHaveJquerySrc {
		bodyIndex := strings.Index(newHtml, "</body>")
		newHtml = newHtml[:bodyIndex] + jqueryScript + newHtml[bodyIndex:]
	}

	if !isHaveJqueryCode {
		bodyIndex := strings.Index(newHtml, "</body>")
		newHtml = newHtml[:bodyIndex] + callJquery + newHtml[bodyIndex:]
	}

	return newHtml
}

func deleteSpace(s string) string {
	s = strings.Replace(s, " ", "", -1)
	return s
}
