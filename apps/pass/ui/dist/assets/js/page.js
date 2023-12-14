/**
 * 全局函数 $+layer
 */
var $,layer;
layui.use(['layer'], function(){
    $ = layui.jquery,layer = layui.layer;
    $(".ver").html('v2.0.5');
});


/**
 * 页面header(因为闪动，暂未启动)
 */
var page_header = function(active_title){
    let menu = [
        ['link.html','layui-icon-cols','连接','扫码连接'],
        ['index.html','layui-icon-release','文件','支持超大文件秒传'],
        ['text.html','layui-icon-list','文本','多端文本传输'],
        ['help.html','layui-icon-rate','帮助','获取帮助'],
    ];
    let ver = "2.0.5";li_big="",li_small="";
    for(let i=menu.length;i>0;i--){
        let key=i-1;
        let val=menu[key];
        let active=(active_title==menu[key])?"layui-this":"";
        li_big+=`
        <li class="layui-nav-item pull-right ${active}"><a href="${val[0]}" title="${val[3]}">
        <i class="layui-icon ${val[1]}"></i> ${val[2]}</a></li>
        `;
    }
    for(let i=0;i<menu.length;i++){
        let val=menu[i];
        let active=(active_title==menu[i][2])?"layui-this":"";
        li_small+=`
        <li class="layui-nav-item pull-right ${active}"><a href="${val[0]}" title="${val[3]}">
        <i class="layui-icon ${val[1]}"></i> ${val[2]}</a></li>
        `;
    }
    return `<div class="layui-header layui-hide-xs">
        <ul class="layui-nav layui-bg-blue main-div" lay-bar="disabled">
            <li class="layui-nav-item logo">
                <a href="http://4bit.cn/p/b0pass" target="_blank">百灵快传 
                <span class="layui-badge layui-bg-gray">v${ver}</span></a> 
            </li>
            ${li_big}
        </ul>
    </div>
    <div class="layui-header layui-hide-md">
            <ul class="layui-nav layui-bg-blue main-div text-center">
            ${li_small}
            </ul>
    </div>`;
}