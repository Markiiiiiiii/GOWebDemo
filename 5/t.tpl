<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <h1>嵌套测试</h1>
    {{/*嵌套另外一个define的模板*/}}
{{template "ul.tpl"}}
<br>
{{/*嵌套外部的模板文件*/}}
{{template "li.tpl"}}
</body>
</html>
{{define "ul.tpl"}}
<ul>
    <li>----->>></li>
    <li>----->>></li>
    <li>----->>></li>
    <li>----->>></li>
    <li>----->>></li>
</ul>
{{end}}