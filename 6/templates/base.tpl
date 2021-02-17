<!DOCTYPE html>
<html lang="en">
<head>
    <title>模板继承</title>
    <style>
        *{
            margin: 0;
        }
        .nav{
            height: 50px;
            width: 100%;
            position: fixed;
            top: 0;
            background-color: burlywood;
        }
        .main{
            margin-top: 50px;
        }
        .menu{
            width: 20%;
            height: 100%;
            position: fixed;
            left: 0;
            background-color: cornflowerblue;
        }
        .center{
            text-align: center;
        }
    </style>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <div class="nav">
        <div class="main">
        <div class="menu"></div>
        <div class="content center">
                {{block "content" .}}{{end}}
        </div>
    </div>
    </div>
</body>
</html>