package main

import (
	"log"
	"net/http"
	"strings"
)

const port = ":8000"
const functionURL = "/wasm/valuetoelement.wasm"

func main() {
	http.HandleFunc("/", rootHandle)
	http.HandleFunc("/wasm_exec.js", jsHandle)
	http.HandleFunc(functionURL, wasmHandle)

	log.Printf("listening on %q...", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func wasmHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/wasm")
	http.ServeFile(w, r, strings.Replace(r.URL.Path[1:], "wasm/", "", 1))
}

func jsHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	http.ServeFile(w, r, "../js/"+r.URL.Path[1:])
}

func rootHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
	w.Write([]byte(page))
}

const page = `
<html>
	<head>
		<title>Testing WebAssembly</title>
		<script src="wasm_exec.js" type="text/javascript"></script>
		<script type="text/javascript">
			const go = new Go();
			WebAssembly.instantiateStreaming(fetch('` + functionURL + `'), go.importObject).then(function(res) {
				go.run(res.instance);
			});
		</script>
	</head>
	<body>
		<h2>WebAssembly content</h2>
		<div id="hello"></div>
	</body>
</html>
`
