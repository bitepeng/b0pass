//initialize code here
layui.config({debug: false}).use(['form', 'element'], function() {
	var form = layui.form,
		$ = layui.$,
		element = layui.element;

	//initailize the tab component
	var tab = {
		activeWindowId: 0,
		tabAdd: function(title, url, id) {
			//create a new tab window
			element.tabAdd('xbs_tab', {
				title: title,
				content: '<iframe tab-id="' + id + '" id="iframe_' + id + '" frameborder="0" src="' + url + '" scrolling="yes" class="x-iframe"></iframe>',
				id: id
			});
			this.activeWindowId = id;
		},
		tabDelete: function(othis) {
			// remove the specified window
			element.tabDelete('xbs_tab', '44');
			othis.addClass('layui-btn-disabled');
		},
		tabChange: function(id) {
			// do tab change
			element.tabChange('xbs_tab', id);
		}
	};

	/* bind button event for closing the tab window */
	$('.layui-tab-close').click(function(event) {
		$('.layui-tab-title li').eq(0).find('i').remove();
	});

	//listening for tab change event
	element.on('tab(xbs_tab)', function() {
		tab.activeWindowId = this.getAttribute('lay-id');
	});

	/* bind event for checked all the checkbox in data list table */
	form.on('checkbox(select-all)', function(data) {

		var checked = data.elem.checked;
		if (checked) {
			$('input[name="id"]').prop("checked", true);
			//re-render checkbox element
			form.render('checkbox');
		} else {
			$('input[name="id"]').prop("checked", false);
			form.render('checkbox');
		}
	});

	/* bind event to reload the iframe window */
	$('#frame-reload').on("click", function() {
		document.getElementById("iframe_" + tab.activeWindowId).contentWindow.location.reload(true);

	});


	$('.container .left_open i').click(function(event) {
		if ($('.left-nav').css('left') == '0px') {
			$('.left-nav').animate({
				left: '-221px'
			}, 100);
			$('.page-content').animate({
				left: '0px'
			}, 100);
			$('.page-content-bg').hide();
		} else {
			$('.left-nav').animate({
				left: '0px'
			}, 100);
			$('.page-content').animate({
				left: '221px'
			}, 100);
			if ($(window).width() < 768) {
				$('.page-content-bg').show();
			}
		}

	});

	$('.page-content-bg').click(function(event) {
		$('.left-nav').animate({
			left: '-221px'
		}, 100);
		$('.page-content').animate({
			left: '0px'
		}, 100);
		$(this).hide();
	});

	$("tbody.x-cate tr[fid!='0']").hide();
	// 栏目多级显示效果
	$('.x-show').click(function() {
		if ($(this).attr('status') == 'true') {
			$(this).html('&#xe625;');
			$(this).attr('status', 'false');
			cateId = $(this).parents('tr').attr('cate-id');
			$("tbody tr[fid=" + cateId + "]").show();
		} else {
			cateIds = [];
			$(this).html('&#xe623;');
			$(this).attr('status', 'true');
			cateId = $(this).parents('tr').attr('cate-id');
			getCateId(cateId);
			for (var i in cateIds) {
				$("tbody tr[cate-id=" + cateIds[i] + "]").hide().find('.x-show').html('&#xe623;').attr('status', 'true');
			}
		}
	})

	//左侧菜单效果
	// $('#content').bind("click",function(event){
	$('.left-nav #nav li').click(function(event) {

		if ($(this).children('.sub-menu').length) {
			if ($(this).hasClass('open')) {
				$(this).removeClass('open');
				$(this).find('.nav_right').html('&#xe697;');
				$(this).children('.sub-menu').stop().slideUp();
				$(this).siblings().children('.sub-menu').slideUp();
			} else {
				$(this).addClass('open');
				$(this).children('a').find('.nav_right').html('&#xe6a6;');
				$(this).children('.sub-menu').stop().slideDown();
				$(this).siblings().children('.sub-menu').stop().slideUp();
				$(this).siblings().find('.nav_right').html('&#xe697;');
				$(this).siblings().removeClass('open');
			}
		} else {

			var url = $(this).children('a').attr('_href');
			var title = $(this).find('cite').html();
			var index = $('.left-nav #nav li').index($(this));

			for (var i = 0; i < $('.x-iframe').length; i++) {
				if ($('.x-iframe').eq(i).attr('tab-id') == index + 1) {
					tab.tabChange(index + 1);
					event.stopPropagation();
					return;
				}
			};

			tab.tabAdd(title, url, index + 1);
			tab.tabChange(index + 1);
		}

		event.stopPropagation();

	})

});

/*弹出层*/
/*
    参数解释：
    title   标题
    url     请求的url
    id      需要操作的数据id
    w       弹出层宽度（缺省调默认值）
    h       弹出层高度（缺省调默认值）
*/
function x_admin_open(title, url, w, h) {
    if (title == null || title == '') {
        title = false;
    };
    if (url == null || url == '') {
        url = "404.html";
    };
    if (w == null || w == '') {
        w = ($(window).width() * 0.9);
    };
    if (h == null || h == '') {
        h = ($(window).height() - 50);
    };
	var layer = parent.layer === undefined ? layui.layer : top.layer;
    layer.open({
        type: 2,
        area: [w + 'px', h + 'px'],
        fix: true, //不固定
        //maxmin: true,
        shadeClose: false,
        shade: 0.4,
        title: title,
        content: url
    });
}

// 打开一最大化的窗口
function x_open_full(title, url) {
	var index = layer.open({
		type: 2,
		area: ['100%','100%'],
		title: title,
		content: url,
		maxmin: false,
		success : function(layero, index){
			/*setTimeout(function(){
				layui.layer.tips('返回', '.layui-layer-setwin .layui-layer-close', {
					tips: 3
				});
			},500)*/
		}
	});
}

/*关闭弹出框口*/
function x_admin_close(callback) {
    var index = parent.layer.getFrameIndex(window.name);
    parent.layer.close(index);
    switch (callback) {

        case "reload": //刷新页面
            parent.location.reload();
            break;
        case "render": //重新渲染列表组件
            parent.render();
            break;
    }
}

/**
 * IsPC
 * true为PC端，false为手机端
 * @return {boolean}
 */
function IsPC() {
	var userAgentInfo = navigator.userAgent;
	var Agents = ["Android", "iPhone",
		"SymbianOS", "Windows Phone",
		"iPad", "iPod"];
	var flag = true;
	for (var v = 0; v < Agents.length; v++) {
		if (userAgentInfo.indexOf(Agents[v]) > 0) {
			flag = false;
			break;
		}
	}
	return flag;
}

/**
 * cookie
 * @param name
 * @param value
 */
function setCookie(name,value) {
	var Days = 90;
	var exp = new Date();
	exp.setTime(exp.getTime() + Days*24*60*60*1000);
	document.cookie = name + "="+ escape (value.trim()) + ";expires=" + exp.toGMTString();
}

/**
 * getCookie
 * @param name
 * @returns {string|null}
 */
function getCookie(name) {
	var arr,reg=new RegExp("(^| )"+name+"=([^;]*)(;|$)");
	if(arr=document.cookie.match(reg))
		return unescape(arr[2]);
	else
		return null;
}

/**
 * delCookie(name)
 * @param name
 */
function delCookie(name) {
	var exp = new Date();
	exp.setTime(exp.getTime() - 1);
	var cval=getCookie(name);
	if(cval!=null)
		document.cookie= name + "="+cval+";expires="+exp.toGMTString();
}