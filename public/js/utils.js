/**
 * 发送异步 http GET 请求
 * @param url
 * @param data
 * @param success
 * @param fail
 */
function httpGet(url, data, success, fail) {
	if (isObject(data)) { //传入了 data 参数
		url = buildQueryUrl(url, data);
	} else { // 没有传入 data 参数
		fail = success;
		success = data;
	}
	request(url, "GET", true, null, success, fail);
}

/**
 * 发送同步 http GET 请求
 * @param url
 * @param data
 * @param success
 * @param fail
 */
function asyncHttpGet(url, data, success, fail) {

	if (isObject(data)) { //传入了 data 参数
		url = buildQueryUrl(url, data);
	} else { // 没有传入 data 参数
		fail = success;
		success = data;
	}
	request(url, "GET", false, null, success, fail);
}

/**
 * 发送异步 http POST 请求
 * @param url
 * @param data
 * @param success
 * @param fail
 */
function httpPost(url, data, success, fail) {

	if( !data ){ // 没有传入 data 参数
		fail = success;
		success = data;
	}
	if (typeof success != "function") {
		success = function() {}
	}
	if (typeof fail != "function") {
		fail = function() {}
	}
	request(url, "POST", true, data, success, fail);
}

/**
 * 发送 http 请求
 * @param url 地址
 * @param method 请求方式, POST, GET
 * @param async 是否同步
 * @param data 数据
 * @param success 成功时候回调
 * @param fail 失败时候回调
 * @returns {*}
 */
function request(url, method, async, data, success, fail) {
	console.log("request url"+url+">>>>");

	if (typeof success != "function") {
		success = function() {}
	}
	if (typeof fail != "function") {
		fail = function() {}
	}
	var options = {
		type: method,
		url: url,
		async: async,
		data: data,
		dataType: "json",
		success: function(result) {
			console.log("request url"+url+"::: "+JSON.stringify(result));
			if (result.err == 0) {
				success(result);
			} else {
				fail(result.message);
			}
		},
		error: function(error) {
			fail(error);
			console.log("request fail url"+url+"::: "+JSON.stringify(error));
		}
	};
	$.ajax(options);
}

/* 合并对象，使用 dist 覆盖 src */
function mergeObject(src, dist) {

	if (!isObject(src) || !isObject(dist)) {
		return;
	}
	for (var key in dist) {
		src[key] = dist[key];
	}
}

/* 判断是否是 javascript 对象 */
function isObject(obj) {
	return Object.prototype.toString.call(obj) == "[object Object]";
}

/* 判断是否是 javascript 对象 */
function isArray(arr) {
	return Object.prototype.toString.call(arr) == "[object Array]";
}

/* build query url */
function buildQueryUrl(url, params) {

	var p = [];
	for (var key in params) {
		p.push(key + "=" + params[key]);
	}
	if (url.indexOf("?") == -1) {
		url += "?" + p.join("&");
	} else {
		url += "&" + p.json("&");
	}
	return url;
}

// 成功提示
function messageOk(message, callback) {
	var layer = parent.layer === undefined ? layui.layer : top.layer;
	layer.msg(message, {icon:1}, callback);
}
// 失败提示
function messageError(message, callback) {
	var layer = parent.layer === undefined ? layui.layer : top.layer;
	layer.msg(message, {icon:2, anim:6}, callback);
}
function messageInfo(message, callback) {
	var layer = parent.layer === undefined ? layui.layer : top.layer;
	layer.msg(message, callback);
}
// 加载提示
function messageLoading(message) {
	message = message ? message : "正在处理中，请稍后..."
	var layer = parent.layer === undefined ? layui.layer : top.layer;
	return layer.msg(message,{icon: 16,time:false,shade:0.6});
}
// 关闭提示框
function loadingClose(index) {
	var layer = parent.layer === undefined ? layui.layer : top.layer;
	return layer.close(index);
}

//获取url中的参数
function getUrlParam(name)
{
	var query = window.location.search.substring(1);
	var vars = query.split("&");
	for (var i=0;i<vars.length;i++) {
		var pair = vars[i].split("=");
		if(pair[0] == name){return pair[1];}
	}
	return(false);
}
/* 删除数组中的某个元素 */
function removeDataByKey(arr, key, value) {

	for (var i = 0; i < arr.length; i++) {
		if (arr[i][key] == value) {
			arr.splice(i, 1);
		}
	}
}

function getDataByKey(arr, key, value) {

	for (var i = 0; i < arr.length; i++) {
		if (arr[i][key] == value) {
			return arr[i];
		}
	}
}

/**
 * // Usage
	 getUserIP(function(ip){
		alert("Got IP! :" + ip);
	});
 * @param onNewIP
 */
function getUserIP(onNewIP) { //  onNewIp - your listener function for new IPs
	//compatibility for firefox and chrome
	var myPeerConnection = window.RTCPeerConnection || window.mozRTCPeerConnection || window.webkitRTCPeerConnection;
	var pc = new myPeerConnection({
			iceServers: []
		}),
		noop = function() {},
		localIPs = {},
		ipRegex = /([0-9]{1,3}(\.[0-9]{1,3}){3}|[a-f0-9]{1,4}(:[a-f0-9]{1,4}){7})/g,
		key;
	function iterateIP(ip) {
		if (!localIPs[ip]) onNewIP(ip);
		localIPs[ip] = true;
	}
	//create a bogus data channel
	pc.createDataChannel("");
	// create offer and set local description
	pc.createOffer().then(function(sdp) {
		sdp.sdp.split('\n').forEach(function(line) {
			if (line.indexOf('candidate') < 0) return;
			line.match(ipRegex).forEach(iterateIP);
		});
		pc.setLocalDescription(sdp, noop, noop);
	}).catch(function(reason) {
		// An error occurred, so handle the failure to connect
	});
	//sten for candidate events
	pc.onicecandidate = function(ice) {
		if (!ice || !ice.candidate || !ice.candidate.candidate || !ice.candidate.candidate.match(ipRegex)) return;
		ice.candidate.candidate.match(ipRegex).forEach(iterateIP);
	};
}
