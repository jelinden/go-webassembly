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
	http.HandleFunc("/wasm/spinning.wasm", wasmHandle)

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
			let spin;
			let spinningReceived;
			let resolveSpinningReceived;

			async function run() {
				spinningReceived = new Promise(resolve => {
					resolveSpinningReceived = resolve;
				})

				WebAssembly.instantiateStreaming(fetch('wasm/spinning.wasm'), go.importObject).then(function(res) {
					go.run(res.instance);
				});
			}
			run();
			setInterval(function(){ spin(); }, 30);
			function runSpin(callback) {
				spin = callback;
				resolveSpinningReceived()
			}
		</script>
	</head>
	<body>
		<h2>WebAssembly content</h2>
		<div id="spinning"></div>
		<div id="counter"></div>
	</body>
</html>
`
