package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type StringServer interface {
	ShowString(s string) (string,error)
	ValidString(s string) (string,error)
}

type stringServer struct{}

func(svr *stringServer) ShowString(s string)(string,error){
	if ok:= strings.Contains(s,"error");ok{
		return "",errors.New("content contains error string")
	}
	ctx := fmt.Sprintf("call ShowString method. show content: %s",s)
	fmt.Println(ctx)
	return ctx,nil
}

func(svr *stringServer) ValidString(s string)(string,error){
	if ok := strings.ContainsRune(s,rune('错')); ok{
		return "",errors.New("content contains 错 string")
	}
	ctx := fmt.Sprintf("call ValidString method. show content: %s",s)
	fmt.Println(ctx)
	return ctx,nil
}

type showStringRequest struct{
	S string `json:"s"`
}

type showStringResponst struct{
	S string `json:"s"`
	Err string `json:"err"`
}

type validStringRequest struct{
	S string `json:"s"`
}

type validStringResponst struct{
	S string `json:"s"`
	Err string `json:"err"`
}

func makeShowString(svr stringServer)endpoint.Endpoint{
	return func(_ context.Context,req interface{}) (interface{},error){
		v := req.(showStringRequest)
		s,err := svr.ShowString(v.S)
		if err != nil{
			return showStringResponst{"",err.Error()},err
		}
		return showStringResponst{s,nil},nil
	}
}

func makeValidString(svr stringServer)endpoint.Endpoint{
	return func(_ context.Context,req interface{}) (interface{},error){
		v := req.(validStringRequest)
		s,err := svr.ShowString(v.S)
		if err != nil{
			return validStringResponst{"",err.Error()},err
		}
		return validStringResponst{s,nil},nil
	}
}

func main() {
	svr := stringServer{}

	showStringHandler := httptransport.NewServer(
		makeShowString(svr),
		decodeShowString,
		encodeResponst)

	validStringHandler := httptransport.NewServer(
		makeValidString(svr),
		decodeValidString,
		encodeResponst)

	http.Handle("/uppercase", showStringHandler)
	http.Handle("/count", validStringHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func decodeShowString(_ context.Context,r *http.Request) (interface{},error){
	var request showStringResponst
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil{
		return nil,err
	}
	return request,nil
}

func decodeValidString(_ context.Context,r *http.Request) (interface{},error){
	var request validStringResponst
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil{
		return nil,err
	}
	return request,nil
}

func encodeResponst(_ context.Context,w http.ResponseWriter,response interface{}) error{
	return json.NewEncoder(w).Encode(response)
}
