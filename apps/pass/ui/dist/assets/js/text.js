window.onload = function () {
    var conn;
    var msg = document.getElementById("msg");
    var log = document.getElementById("log");
    var data = JSON.parse(localStorage.getItem("txtdata")) || [];

    /** Editor **/
    var editor_set = ice.editor('editor',function(){
        this.maxWindow = false;
        this.height = '200px';
        this.menu=[
            'fontSize','foreColor','bold','line',
            'insertOrderedList', 'insertUnorderedList','line',//'removeFormat',
            'paste','code'
            ];
        if(!data.length){
            this.setValue('<p> 欢迎使用<em>百灵快传</em>(<b>B0Pass</b>)文本传输功能。<br/>注意：文本存在浏览器缓存，请自行保存。</p>');
        }else{
            this.setValue('');
        }
        this.create();
    });

     /** 载入缓存 **/
     if(data.length>0){
        //alert(JSON.stringify(data));
        for (let i = 0; i < data.length; i++) {
            echomsg(data[i]['key'],data[i]['val']);
        }
    }
    
    /** SendMsg **/
    document.getElementById("send").onclick = function () {
        let send_msg=editor_set.getHTML();
        if (!conn) {return false;}
        if (!send_msg) {return false;}
        conn.send(send_msg);
        editor_set.setValue('');
        return false;
    };
    document.getElementById("clear").onclick = function () {
        let msg = "将删除所有本地缓存文本，请确认！"; 
        if (confirm(msg)==true){ 
            localStorage.setItem("txtdata","[]");
            log.innerHTML="";
            //editor_set.setValue('');
        }
    };

    /** WebSocket **/
    if (window["WebSocket"]) {
        conn = new WebSocket("ws://" + document.location.host + "/ws");
        conn.onclose = function (evt) {
            var item = document.createElement("div");
            item.innerHTML = "<b>Connection closed.</b>";
            appendLog(item);
        };
        conn.onmessage = function (evt) {
            var messages = evt.data.split("\n");
            for (var i = 0; i < messages.length; i++) {
                var ntime = new Date( +new Date() + 8 * 3600 * 1000 ).toJSON().substr(0,19).replace("T"," ");
                var nblockid = Date.now();
                echomsg(nblockid,messages[i]);
                data.push({"key":nblockid,"val":messages[i]});
                localStorage.setItem("txtdata",JSON.stringify(data));
            }
        };
    } else {
        var item = document.createElement("div");
        item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
        appendLog(item);
    }

    /** 文本显示 **/
    function echomsg(id,msg){
        var item = document.createElement("div");
        var nblockid = "block_"+id;
        var barhtml =  "<div class='bar' ><i class='layui-icon layui-icon-triangle-r'></i>MSG_"+id+" "
        barhtml = barhtml + "<span onclick=\"delb('"+nblockid+"')\"><i class='layui-icon layui-icon-close'></i> &nbsp; </span>";
        barhtml = barhtml + "<span onclick=\"copyb('"+nblockid+"',0)\">【复制文字】</span>";
        barhtml = barhtml + "<span onclick=\"copyb('"+nblockid+"',1)\">【复制HTML】</span></div>";
        item.innerHTML = barhtml + "<div id='"+nblockid+"' class='btext'>"+msg+"</div><hr>"
        appendLog(item);
    }
    function appendLog(item) {
        var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
        log.appendChild(item);
        if (doScroll) {
            log.scrollTop = log.scrollHeight - log.clientHeight;
        }
    }
};

/** 文本删除 **/
function delb(e){
    let _data = JSON.parse(localStorage.getItem("txtdata")) || [];
    for (let i = 0; i < _data.length; i++) {
        let key =e.replace("block_","");
           if(_data[i]['key']==key){
                _data = JSON.parse(localStorage.getItem("txtdata")) || [];
                _data.splice(i, 1);
                console.log(e,i,_data);
                localStorage.setItem("txtdata",JSON.stringify(_data));
                document.querySelector("#"+e).parentNode.style.display="none";
           }
    }
    layer.msg("删除成功");
}

/** 文本复制 **/
function copyb(e,t){
    var copyText;
    if(t==1){
        copyText = document.querySelector("#"+e).innerHTML;
    }else{
        copyText = document.querySelector("#"+e).textContent;
    }
    console.log(copyText);
    copyToClipBoard(copyText);
    layer.msg("复制成功");
}
function copyToClipBoard(str){
    const el = document.createElement('textarea');
    el.value = str;
    document.body.appendChild(el);
    el.select();
    document.execCommand('copy');
    document.body.removeChild(el);
};