package main

const MAIN_HTML string = `
<html>
<head>
	<script type="text/javascript" src="http://libs.baidu.com/jquery/2.0.0/jquery.min.js"></script>
</head>
<body>
	<img id="qrcode" src="https://login.weixin.qq.com/qrcode/{{.}}">
	<div id="contacts" style="display:none">
		<div>Friends</div>
		<div id="friends"></div>
		<div>Groups</div>
		<div id="groups"></div>
		<div>Publics</div>
		<div id="publics"></div>
	</div>
	<script type="text/javascript" src="https://rawgit.com/tidezyc/wx_robot/master/main.js"></script>
</body>
</html>`
