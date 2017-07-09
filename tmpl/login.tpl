<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="utf-8">
    <title>登入</title>
    <style>
        .fll{float: left;}
        .flr{float: right;}
        body{
            background: #fffffc;
        }
        .login{
            /*position: absolute;*/
            /*top: 40%;*/
            /*left: 50%;*/
            /*margin-top:  -150px;*/
            /*margin-left: -140px;*/
            margin: 150px auto;
            width: 280px;
            height: 300px;
            overflow: hidden;
            /*background: #333;*/
        }
 
        .loginbox{
            font-family: &#39;Segoe UI&#39;, &#39;Microsoft Yahei&#39;, Arial, Simsun, sans-serif, "宋体";
            font-size: 14px;
            padding:20px 25px 0 25px;
        }
        .loginboxtag{
            height: 30px;
            line-height: 30px;
            padding-left: 2px;
            color: #555;
            cursor: default;
            user-select: none; -moz-user-select: none;  -webkit-user-select: none;  -ms-user-select: none;
        }
        .loginboxinput{
            height: 50px;
        }
        input[type="text"], input[type="password"], textarea{ 
            padding-left: 27px;
            color: #0e8c3c;
            outline: 0px;
            height: 30px;
            width: 200px;
            border: 1px solid #ccc;
            /*border-radius: 3px;*/
            transition: all .2s;-webkit-transition: all .2s;-moz-transition: all .2s;
        }
 
        input[type="text"] {
            background: url(img/iconman.png) no-repeat #f8faf8;
        }
        input[type="password"] {
            background: url(img/iconlock.png) no-repeat #f8faf8;
        }
        input[type="text"]:focus, input[type="password"]:focus, textarea:focus { 
            border: 1px solid #c8c8c8;
            background-color: #f3f7f3;
            /*box-shadow:inset 0 0 3px rgba(40,140,210,1);*/
            /*background-color: #444;*/
        }
 
        .loginboxbtn{
            overflow: auto;zoom: 1;
            height: 40px;
            padding-top: 10px;
 
        }
        input[type="checkbox"]{
            margin: 0;
            margin-right: 10px;
        }
        .loginboxbtn .rem{
            font-size: 12px;
            padding-top: 15px;
            user-select: none; -moz-user-select: none;  -webkit-user-select: none;  -ms-user-select: none;
        }
        .loginboxbtn .rem span label{
            /*display: inline-block;
            margin-top: -5px;*/
            color: #555;
            position: relative;
            top: -2px;
            cursor: default;
        }
        .btn{
            display: inline-block;
            width: 80px;
            height: 30px;
            line-height: 30px;
            text-align: center;
            /*background: #0e9c4c;*/
            background: #46Ae00;
            color: #fff;
            /*border-radius: 5px;*/
            box-shadow: 0 0 1px rgba(0,0,0,0.3);
            cursor: pointer;
            transition: all .1s;-webkit-transition: all .1s;-moz-transition: all .1s;
            user-select: none; -moz-user-select: none;  -webkit-user-select: none;  -ms-user-select: none;
        }
        .btn:hover{
            /*background: #0e8c3c;*/
            background: #339b00;
        }
        .btn:active{
            /*background: #0e7c2c;*/
            background: #288f00;
        }
    </style>
    <script src="/static/img/html/jquery.js" type="text/javascript"></script>
</head>
<body>
    <div class="login">
        <div class="angel">
            <div class="loginbox">
                <form name="loginform" id="loginform" action="/admin/loginpost" method="POST">
                    <div class="loginboxtag">用户名：</div>
                    <div class="loginboxinput"><input type="text" name="username"></div>
                    <div class="loginboxtag">密码：</div>
                    <div class="loginboxinput"><input type="password" name="password"></div>
                    <div class="loginboxbtn">
                      <div class="fll rem"> <span><label for="rem">{{.Errmsg}}</label></span></div> 
                        <div id="submit" class="flr btn">登入</div>
                    </div>
                </form>
            </div>
        </div>
    </div>
 
 
</body>
</html>
 
 
<script type="text/javascript"> 
$(".login").fadeIn(700);
//改善表单输入用户体验
$(function () {
    $("input:text:first").focus();
 
    //回车聚焦下个输入框
    var $inp = $("input:text");
    $inp.bind("keydown", function (e) {
        var key = e.which;
        if (key == 13) {
            e.preventDefault();
            $("input:password:first").focus();
        }
    });
 
    //回车触发按钮点击
    var $pwd = $("input:password:first");
    $pwd.bind("keydown", function (e) {
        var key = e.which;
        if (key == 13) {
            e.preventDefault();
            $("#submit").trigger("click");
        }
    });
 
    //
    $("#submit").click(function() {
        
        $("#loginform").submit();
    });
});
</script>