package storage

//数据类型

//请求
type Request struct {
    Host string 
    Protocol string 
    Method string 
    Path string 
    Headers map[string]string
    ReuestSize uint64 
    Body string  //请求body base64 
    ReqId string  //请求id
    UUID string //匹配到ecapture的uuid功能

}

//响应
type Response struct {
    Headers map[string]string
    StatusCode string 
    Body string // base64 
    UUID  string //匹配到ecapture的响应数据uuid

}
//http server数据

