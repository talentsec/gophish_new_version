# gophish_new_version
增加了以下功能：
* 对在landing_pages页面中设置对钓鱼页面内增加了ajax拦截和http包转发功能。

注意：
* 本次修改相对与原gophish添加对代码主要是在./modifyHtml下。
* http转发只会转发content-type为application/json的包，如果还想转发content-type类型为其它的包请在./modifyHtml/forwardHttp.go文件内修改。
