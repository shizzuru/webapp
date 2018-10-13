package main

import(
	//"html/template"
	"net/http"
	"log"
)

func redirect(w http.ResponseWriter, req *http.Request){
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	http.Redirect(w,req,target,http.StatusTemporaryRedirect)
}

func main(){
	go http.ListenAndServe(":80",http.HandlerFunc(redirect))
	http.Handle("/",http.FileServer(http.Dir("/home/shizzuru/webapp")))
	log.Fatal(http.ListenAndServeTLS(":443","/etc/letsencrypt/live/catbox.space/fullchain.pem","/etc/letsencrypt/live/catbox.space/privkey.pem",nil))
}

