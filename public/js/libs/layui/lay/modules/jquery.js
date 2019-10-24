//将jQuery对象局部暴露给layui
layui.define(function(exports){
	layui.$ = jQuery;
	exports('jquery', jQuery);
});