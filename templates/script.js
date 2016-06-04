window.onload = function() {
	console.log("this is some js. hello");

	{{ if .Redirect}}
	conole.log("{{ .Url }}");
	{{ end }}
};
