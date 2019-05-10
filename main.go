package main

import (
    "fmt"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
    "context"
    "time"
)

var client *mongo.Client

func main() {
    var err error
    
    client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil { 
        panic(err)  
    }
    ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
    defer cancel()
    err = client.Connect(ctx)
    if err != nil { 
        panic(err)  
    }

    UpdateOrInsertFindByKey("test", "name", map[string]interface{}{"name": "pi1", "age1": 20})
    
}

/*
ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
err = client.Ping(ctx, readpref.Primary())
*/

func Find(){
    var result struct {
        Value float64
    }
    filter := bson.M{"name": "pi"}
    ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
    err = collection.FindOne(ctx, filter).Decode(&result)
    if err != nil {
        log.Fatal(err)
    }
}

func Insert(table string) string{
    collection := client.Database("test").Collection(table)
    ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
    res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
    return res.InsertedID
}

/**
 *
 * @param table 集合 | 表名
 * @param key 查找字段
 * @param val 需要修改的值
 * @param insert 是否新增
 *
 */
func Update(table string, key map[string]interface{}, val map[string]interface{}, insert bool){
    // 选择数据库 -> 表
    collection := client.Database("test").Collection(table)
    ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
    res, err := collection.UpdateOne(ctx, key, bson.D{{"$set", val}}, &options.UpdateOptions{Upsert: &insert})
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(res)
    }

}
