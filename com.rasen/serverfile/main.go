package main

import (
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main()  {
	svr := &http.Server{
		Addr: ":5535",
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	path,_:= os.Getwd()
	path = filepath.Join(path,"src","com.rasen")
	http.Handle("/",http.FileServer(http.Dir(path+"/static")))

	//http.HandleFunc("/ind",func(w http.ResponseWriter,r *http.Request){
		//http.FileServer(http.Dir("/static"))
		//wd,_ := os.Getwd()
		//filePath := filepath.Join(wd,"src","com.rasen","static",r.URL.Path+".html")
		//fmt.Println("url:", r.URL.Path,"|| path:",filePath)

		//fd,err := os.Open(filePath)
		//if err != nil{
		//	fmt.Println("err:",err)
		//}
		//defer fd.Close()
		//w.Header().Set("Cache-Control", "no-cache")
		//w.Header().Set("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS")
		//w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With")
		//w.Header().Set("Access-Control-Allow-Origin", "*")
		//w.Header().Set("Connection", "keep-alive")
		//http.ServeFile(w,r,filePath)
	//})

	svr.ListenAndServe()

}
