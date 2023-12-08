package storage

import (
	"log"
	"net/http"
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
 
    Q chan PreProcessorData //数据队列
}

func init(){
    log.Println("call preprocessor")
}
//
func NewPreProcessor()PreProcessor{
    proceser := &HttpsPreProcessor{

        Q: make(chan PreProcessorData),
    }
    return proceser

}
//
func (hr *HttpsPreProcessor)Run(){
    //在这里进行处理封装的数据
    for data := range hr.Q{
        log.Printf("->dataType =  %v\n",data.Type)
        
    }
}


func (hr *HttpsPreProcessor)Write(pre PreProcessorData)(error){

    hr.Q <- pre

    return nil
}
