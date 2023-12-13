layui.use(['layer'], function(){
    var $ = layui.jquery
    ,layer = layui.layer;

    function getIP(){
        var servIP;
        $.ajax({
            url: "/gateway/config",
            success: function(res) {
            console.log("::Config::",res.data);
            //linux操作系统禁用一些功能
            if(res.data.Password!="windows"){
                domid("btn_left_key").style.display="none";
            }
            ips=(res.data.ListenAddr).split(":");
            servPort=":"+(res.data.ListenAddr).split(":")[1];
            if(ips[0]!=""){
                if(ips[0]="127.0.0.1"){
                alert("如果IP被设置为127.0.0.1，意味着只能本机使用，将无法分享文件！\n若无特殊需求，请将配置文件的'ListenAddr'设置为纯端口，如':8899'。");
                }
                servIP=ips[0];
                setTextValue("http://"+servIP+servPort);
                console.log("::ServIP::",servIP);
            }
            //兼容域名情况
            if(res.data.Domain!=""){
                servIP=res.data.Domain;servPort="";
                setTextValue("http://"+servIP+servPort);
                console.log("::ServIP::",servIP);
            }else{
                $.ajax({
                url: "/pass/read-ip",
                success: function(res) {
                    servIP=res.data;
                    setTextValue("http://"+servIP+servPort);
                    console.log("::ServIP::",servIP);
                }
                });
            }
            }
        });
    }

    window.onload=function () {
        document.getElementById('text').style.display="";
        document.getElementById('selects').style.display="none";
        var ip=args('f');
        if(ip){
            document.getElementById('text').value="http://"+ip;
            makeCode();
        }else{
            getIP();
        }  
    };

    function args(name) {
        var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
        var r = window.location.search.substr(1).match(reg);
        if (r != null) return decodeURI(r[2]); return null;
    }

    //键盘
    $("#btn_left_key").on("click",function(){
        layer.open({
          title: "遥控主电脑键盘",
          area: ['100%','100%'],
          type: 1, 
          content: '<div class="padding15"><div class="layui-form-item" align="center"><p><br></p></div>'+$("#send-key").html()+'</div>',
          cancel: function () {}
        });
    })
    window.sendKey=function(k){
    $.ajax({
        url: "/pass/cmd-key?k="+k,
        method: "get",
        data: {},
        success: function(res) {
        layer.msg("按下"+k+"键");
        }
    });
    }

 });
 function setTextValue(v){
    document.getElementById('text').value=v;
    makeCode();
}
function makeCode(){
    document.getElementById("qrcode").innerHTML="";
    new QRCode(document.getElementById("qrcode"), {
        text: document.getElementById('text').value,
        width: 235,
        height: 235,
        colorDark : "#226ef1",
        colorLight : "#ffffff",
        correctLevel : QRCode.CorrectLevel.H
    });
}
