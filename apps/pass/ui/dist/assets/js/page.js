/**
 * 全局变量
 */
var token = localStorage.getItem('token') || '';
var auth = localStorage.getItem('auth') || '';

/**
 * 全局函数
 */
var $,layer;
layui.use(['layer'], function(){
    $ = layui.jquery,layer = layui.layer;
    $(".ver").html('v2.0.6');
});
var domid = function(id){ 
    return document.getElementById(id); 
}

/**
 * 统一API请求函数
 * @param {string} url - 请求地址
 * @param {string} method - GET/POST
 * @param {object} data - 请求数据
 * @param {function} success - 成功回调
 * @param {function} error - 错误回调(可选)
 */
var api_ajax = function (url, method, data, success, error) {
    // 从localStorage获取token
    var token = localStorage.getItem('token') || '';
    
    $.ajax({
        url: url,
        type: method,
        data: method === 'GET' ? data : JSON.stringify(data),
        contentType: 'application/json',
        headers: {
            'token': token,
            'Content-Type' : 'application/json;charset=utf-8'
        },
        success: function(res) {
            console.log(url,res,token);
            if(res.code === 400 || res.code === 403){
                layer.msg(res.msg,{icon: 5});
                return;
            }else if (res.code === 401) {
                // 401未授权跳转到登录页
                localStorage.removeItem('token'); // 清除token
                localStorage.removeItem('auth'); // 清除auth
                layer.msg("请先登录",{icon: 5});
                window.location.href = '/app/pass/login.html';
                return;  
            }else{
                success && success(res);
            }
        },
        error: function(xhr) {
            error && error(xhr);
        }
    });
}

