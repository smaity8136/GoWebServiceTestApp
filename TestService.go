package main

import (
 "net/http")


type testHandler struct {

	Message string
}

func (t* testHandler)  ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Write([]byte(t.Message))
}

func testBHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("testb handler"))

}
func main(){
	
	http.Handle("/test", &testHandler{Message: "testMsg"})
	http.HandleFunc("/testB", testBHandler)
    http.ListenAndServe(":5000",nil)

}


