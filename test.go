package main
 
import (
    // "fmt"
    "spider/model"
    "spider/web/api"
)

var models *model.Model

func main(){

    models = &model.Model{Db: model.InitDb()}

    api.HttpRun(models, "4321")

    // fmt.Println(models.GetCartoonResources("酷", 1, 0))
}
