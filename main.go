package main

import (
	"io"
	"net/http"
)


var htmlString = `<html>
  <head>
    <meta name="google-site-verification" content="IIG5Ye_PV0xICl2aXCNG599pYppwfp661cdXTx52Qcw" />
    <title>My Title</title>
  </head>
  <body>
    page contents
  </body>
</html>`

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, htmlString)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8081", nil)
}
