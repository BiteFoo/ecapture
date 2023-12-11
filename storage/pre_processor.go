package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type PreType int64 
//处理数据使用
const (
    ReqType PreType = iota
    RespTye 

)

type PreProcessor interface{
    Write(PreProcessorData)(error)
}
type PreProcessorData struct {
    Type PreType //请求数据类型
    Req *http.Request //解析好的请求
    Resp *http.Response //解析好的响应
    UUID string  //匹配到每个数据包的uuid 便于存放到对应的数据目录
    Length uint64 //数据的长度
}
//
type HttpsPreProcessor struct {
 
    
}
var q  chan PreProcessorData //数据队列

func init(){
    q  = make(chan PreProcessorData)
    go run()
}

func dumpToJson(obj interface{}){
    v ,err := json.MarshalIndent(obj,""," ")
    if err != nil{
        log.Println("Dump object to json error ",err)
        return 
    }
    log.Printf("%v\n",string(v))
}

//
func run(){
    //在这里进行处理封装的数据
    for data := range q {
        // log.Printf("->dataType =  %v\n",data.Type)
        if data.Type == ReqType{
            var req Request
            req.Headers = make(map[string]string)
            for k,v := range data.Req.Header{
                req.Headers[k] = strings.Join(v,"")
            }
            body ,err := io.ReadAll(data.Req.Body)
            if err != nil{
                log.Println("Reading request body error: ",err)
                continue
            }
            req.Body = string(body)
            req.Method = data.Req.Method
            req.Host = data.Req.Host
            req.Protocol = data.Req.Proto
            req.Path = data.Req.URL.Path
            req.RequestSize = data.Length
            req.UUID = data.UUID 
            // log.Printf("%v\n",req)
            dumpToJson(req)

        }else if data.Type == RespTye{
            var resp  Response
            
            body ,err := io.ReadAll(data.Resp.Body)
            if err!= nil{
                log.Println("read responsed eror ",err)
                continue
            }
            resp.Body  = string(body)
            resp.StatusCode = fmt.Sprintf("%v",data.Resp.StatusCode)
            resp.Headers =  make(map[string]string)
            for k,v := range data.Resp.Header{
                resp.Headers[k] = strings.Join(v,"")
            }
            resp.UUID = data.UUID
            resp.ResponseSize = data.Length
            // log.Printf("%v\n",resp)
            dumpToJson(resp)


            
        }
        
    }
}


func Write(pre PreProcessorData)(error){

    q <- pre

    return nil
}
