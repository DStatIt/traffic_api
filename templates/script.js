var xhttp = new XMLHttpRequest();


window.onload = function() {
	xhttp.open("GET", "http://{{ .Host }}/update/?userID={{ .UserID }}&url=" + window.location.href, true);
	xhttp.send();
	{{ if .Redirect}}
	conole.log("{{ .Url }}");
	{{ end }}
};
