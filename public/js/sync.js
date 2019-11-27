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
    var data=JSON.parse(result.data)
    syncDo(data);
};


function syncSend(data){
    ws.send(data);
}