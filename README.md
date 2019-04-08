# template
go load template

##
    1、支持模板定时更新

### go code

    temp := template.Temp("./html/index2.html")
    p := Person{Name: "safly", Age: "30", Hight: "3000"}
    tmp.Execute(w, p)

### html code

    <html>
    <head>
        <title>
        </title>
    </head>

    <body>
    <p>
        <!--p代表当前对象-->
        hello,{{.Name}} are you {{.Age}}? {{.Hight}}
    </p>
    </body>
    </html>
