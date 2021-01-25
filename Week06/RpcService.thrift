namespace go demo.rpc
struct EchoReq {
    1: string msg;
}
// 测试服务
service RpcService {
    // 发起远程调用
    list<string> funCall(1: EchoReq req),
}
