setTimeout(function () {
	Push();
}, 2000);
var id = setInterval(function () {
	Push();

}, 1000);

function Push() {
	$.ajax({
		type: "GET",
		url: "/contacts",
		success: function (data) {
			var obj = eval("(" + data + ")");
			if (obj.ret == 0) {
				clearInterval(id);
				$('#qrcode').remove()
				for (i in obj.friends) {
					var f = obj.friends[i]
					$('#friends').append('<a target="_black" href="/msg?username=' + f.UserName + '">' + f.NickName + '</a><br/>')
				}
				for (i in obj.groups) {
					var g = obj.groups[i]
					$('#groups').append('<a target="_black" href="/msg?username=' + g.UserName + '">' + g.NickName + '</a><br/>')
				}
				for (i in obj.publics) {
					var p = obj.publics[i]
					$('#publics').append('<a target="_black" href="/msg?username=' + g.UserName + '">' + p.NickName + '</a><br/>')
				}
				$('#contacts').show()
			}
		}
	});
}
