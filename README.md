# go-webassembly

A small project for trying out webassembly with Go.

<img src="https://github.com/jelinden/go-webassembly/raw/master/spinning/spinning.png" width="350">

10.08.2018 Edit: wasm_exec.js updated for go 1.11beta3

You can download the latest Go from [https://golang.org/dl/](https://golang.org/dl/) and
the latest wasm_exec.js file from [https://github.com/golang/go/blob/master/misc/wasm/](https://github.com/golang/go/blob/master/misc/wasm/)

## directory simple

Print hello to browser console.

```
cd simple
bash build.sh
./server
```

Go to http://localhost:8000

## directory valuetoelement

Print hello to a web page.

```
cd valuetoelement
bash build.sh
./server
```

Go to http://localhost:8000

## directory redbox

Print a red box to a web page from Go code.

```
cd redbox
bash build.sh
./server
```

Go to http://localhost:8000

## directory spinning

Spinning red box from Go code.

```
cd spinning
bash build.sh
./server
```

Go to http://localhost:8000
