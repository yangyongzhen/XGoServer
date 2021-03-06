package main

/**

作者(Author): 林冠宏 / 指尖下的幽灵

Created on : 2018/2/10

*/

import "testing"
import (
	"github.com/XGoServer/core"
	"github.com/XGoServer/model"
	"github.com/XGoServer/util"
)

func TestInit(t *testing.T) {
	if !core.SimpleInit() {
		return
	}

	if !core.CreateTables(model.Book{}, model.Comment{}){
		return
	}
	affect,_ := core.Engine.Insert()
	if affect > 0 {
		util.LogInfo("insert success !")
	}else{
		util.LogInfo("insert failed")
	}
}