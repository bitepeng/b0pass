layui.use(['layer'], function(){
    var $ = layui.jquery
    ,layer = layui.layer;
    var token = localStorage.getItem('token') || '';

    function login(){
        let code = domid("code");
        // 检查验证码输入框是否有内容
        if(!code.value.trim()){
            alert("请输入验证码");
            return;
        }else{
            $.ajax({
                url: "/pass/login?code="+code.value,
                method: "get",
                data: {},
                success: function(res) {
                    console.log(res);
                    // 检查响应数据格式
                    if(res && typeof res.code!== 'undefined' && typeof res.msg!== 'undefined'){
                        //alert(res.msg);
                        if(res.code==0){
                            let auth = (res.data).split(":")[0];
                            let token = (res.data).split(":")[1];
                            localStorage.setItem("auth", auth);
                            localStorage.setItem("token", token);
                            window.location.href = "/app/pass";
                        }else{
                            layer.msg("登录失败",{icon: 5}); 
                        }
                    } else {
                        layer.msg("服务器返回数据格式错误",{icon: 5});
                    }
                },
                // 处理请求失败的情况
                error: function(error) {
                    console.error('登录请求失败:', error);
                    layer.msg("登录请求失败，请稍后重试",{icon: 5});
                }
            });
        }
        window.location.reload();
    };

    domid("login").onclick=function(){
        login(); 
    }

    window.onload=function () {
        if(token){
            window.location.href = "/app/pass";
        }else{
            domid("login").style.display="";
            domid("code").style.display="";
        }
    };

});
