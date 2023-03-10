package main

import (
	"fmt"
	"net/http"
)

func AdvanceLearn() {
	fmt.Println("*********************Advance Learn Golang********************")

	//region web
	// Start()
	//endregion

	fmt.Println("*************************Advance END*************************")
}

// region web 服务
func Start() {
	http.HandleFunc("/", Index)
	http.ListenAndServe("0.0.0.0:8081", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Wold!")
	fmt.Println("控制台：hello world！")
}

//endregion

//region

//endregion

//region

//endregion

//region

//endregion
