package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type searchContent struct{
	Content string `json:"content",omit`
}

func SearchHandler(ctx *gin.Context){
	var ctt searchContent
	err := ctx.Bind(&ctt)
	if err != nil{
		fmt.Println("err:",err)
	}
	fmt.Println("ctt",ctt)
	ctx.JSON(200,gin.H{
		"resp ctt":"recv "+ctt.Content,
})
}
