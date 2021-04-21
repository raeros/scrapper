package main

// see https://siongui.github.io/2016/05/10/go-get-html-title-via-net-html/

import (
	"eros/scrapper"
	"net/http"
	"log"
)

func main(){
	log.SetPrefix("[TECHCRUNCH] ")
	log.SetFlags(0)

	// convert code bellow to valid Golang and "net/html".
	//
	document, err := http.Get("https://techcrunch.com/")
	
	if err != nil {
		panic(err)
	}

	defer document.Body.Close()

	//content := document.getElementsByClassName("content-wrap");
	//content_headline := _content[0].innerHTML;
	//content_headline := _content[0].innerHTML;

	if title, ok := scrapper.GetHtmlTitle(document.Body); ok {
		println(title)
	} else {
		println("Fail to get HTML title")
	}
}
