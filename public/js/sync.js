/**
 * WebSocket
 *
 * syncDo(data)
 * syncSend(data)
 *
 */
var url = window.location.origin.replace("http://", "");
url = "ws://" + url + "/sync/web-socket";
var ws  = new WebSocket(url);
ws.onmessage = function (result) {
    var data=JSON.parse(result.data);
    //消息接收由载入页面实现
    syncDo(data);
};


function syncSend(data){
    ws.send(data);
}