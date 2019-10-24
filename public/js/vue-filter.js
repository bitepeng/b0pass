/**
 * @author yangjian
 * @since 18-8-23 上午9:55.
 */

/**
 * Vue 时间过滤器
 */
// 微信时间格式
Vue.filter('wxDateFormat', function (dateValue) {
	if (!dateValue) {
		return "";
	}
	var now = new Date();
	var d = new Date();
	d.setTime(dateValue);
	if ( dateValue > now.getTime() ) {
		return d.getFullYear()+'年'+formatNumber(d.getMonth()+1, 2)+'月'+formatNumber(d.getDate(), 2)+'日';
	}

	var t = Math.ceil((now.getTime() - dateValue)/1000);
	if ( t < 10 )        return '刚刚';                            //just now
	if ( t < 60 )        return t+'秒前';                        //under one minuts
	if ( t < 3600 )        return Math.floor(t/60)+'分钟前';        //under one hour
	if ( t < 86400 && d.getDate() == new Date().getDate() ) return Math.floor(t/3600)+'小时前';        //under one day
	if ( t < 86400 && d.getDate() < new  Date().getDate() ) return "昨天";    //yesterday
	if ( t < 864000 )    return Math.floor(t/86400)+'天前';        //under 10 days

	if ( t < 31104000 && d.getFullYear() == new Date().getFullYear() ) { //under one year
		return (d.getMonth()+1)+'月'+d.getDate()+'日';
	}

	return d.getFullYear()+'年'+(d.getMonth()+1)+'月'+d.getDate()+'日';

})
// 普通时间格式
Vue.filter('dateFormat', function (dateValue, format) {
	if (!dateValue) {
		return "";
	}
	var date = new Date();
	date.setTime(dateValue);
	format = format || 'yyyy-MM-dd hh:mm:ss';

	var map = {
		"y": date.getFullYear(), //年份
		"M": formatNumber(date.getMonth() + 1, 2), //月份
		"d": formatNumber(date.getDate(), 2), //日
		"h": formatNumber(date.getHours(), 2), //小时
		"m": formatNumber(date.getMinutes(), 2), //分
		"s": formatNumber(date.getSeconds(), 2), //秒
		"q": Math.floor((date.getMonth() + 3) / 3), //季度
		"S": date.getMilliseconds() //毫秒
	};
	return format.replace(/([yMdhmsqS])+/g, function(all, t){
		return[map[t]];
	});

})

/**
 * 格式化数字为指定长度，不够前面补 0
 * @param num
 * @param length
 */
function formatNumber(num, length) {
	var str = [num];
	for (var i = 1; i < length; i++) {
		if (num < Math.pow(10, i)) {
			str.unshift(0);
		}
	}
	return str.join('');
}