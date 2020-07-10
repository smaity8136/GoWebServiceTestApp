package main

import (
 "net/http")


type testHandler struct {

	Message string
}

func (t* testHandler)  ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Write([]byte(t.Message))
}
func main(){
	//fmt.Println("hello")
	http.Handle("/test", &testHandler{Message: "testMsg"})
    http.ListenAndServe(":5000",nil)

}

