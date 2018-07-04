package main

import (
	"log"
	"net/http"
	"strings"
)

const port = ":8000"

func main() {
	http.HandleFunc("/", rootHandle)
	http.HandleFunc("/wasm_exec.js", jsHandle)
	http.HandleFunc("/wasm/simple.wasm", wasmHandle)

	log.Printf("listening on %q...", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func wasmHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/wasm")
	http.ServeFile(w, r, "./simple"+strings.Replace(r.URL.Path[1:], "wasm", "", 1))
}

func jsHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	http.ServeFile(w, r, r.URL.Path[1:])
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
			WebAssembly.instantiateStreaming(fetch('wasm/simple.wasm'), go.importObject).then(function(res) {
				console.log(res.instance);
				go.run(res.instance).then(function(value) {
					document.getElementById('wasm').innerHTML = value;
				});
			});
		</script>
	</head>
	<body>
		<h2>WebAssembly content</h2>
		<div id="wasm"></div>
	</body>
</html>
`
