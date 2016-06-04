window.onload = function() {
	console.log("this is some js. hello");

	{{ if .Redirect}}
	window.location = "{{ .Url }}";
	{{ end }}
};
