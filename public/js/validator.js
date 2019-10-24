/**
 * 表单验证工具
 * @author yangjian
 */
if (!String.prototype.trim) {
	String.prototype.trim = function () {
		return this.replace(/^\s+|\s+$/, '');
	}
}
var Validator = function () {

	// 验证规则
	var pattern = {
		//用户名
		uname: function (value) {
			return /^[0-9|a-z|_]{4,20}$/i.test(value);
		},
		required: function(value) {
			return value != null && value.toString().trim().length > 0
		},
		//邮箱
		email: function (value) {
			return /^[a-z0-9]\w{1,18}@[a-z0-9]{1,20}(\.[a-z]{1,6}){1,3}$/i.test(value);
		},
		//网址url
		url: function (value) {
			return /^https?:\/\/(www\.)?\w+(\.[a-z|0-9]+){1,2}/i.test(value);
		},
		//域名
		domain: function (value) {
			return /^\w+(\.[a-z|0-9]+){1,2}$/i.test(value);
		},
		//手机号码
		mobile: function (value) {
			return /^1[3|4|5|7|8][0-9]{9}$/.test(value);
		},
		//电话号码
		phone: function (value) {
			return /^[0-9]{2,5}[-][0-9]{7,8}$/.test(value);
		},
		//整数
		number: function (value) {
			return /^[0-9]+$/.test(value);
		},
		//浮点数
		float: function (value) {
			return /^[0-9]+\.[0-9]+/.test(value);
		},
		//中文
		cn : function (value) {
			return /^[\u4E00-\u9FA5]+$/.test(value);
		},
		//英文
		en : function (value) {
			return /^\w+$/.test(value);
		},
		//IP
		ip: function (value) {
			return /^([0-9]{1,3})\.([0-9]{1,3})\.([0-9]{1,3})\.([0-9]{1,3})$/.test(value);
		},
		// 身份证
		id : function (value) {
			if (value.length != 18) return false;
			//加权因子
			var wi = [7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2, 1];
			//校验码对应值
			var vi = [1, 0, 'X', 9, 8, 7, 6, 5, 4, 3, 2];
			var ai = new Array(17);
			var sum = 0;
			var remaining, verifyNum;

			for (var i = 0; i < 17; i++) ai[i] = parseInt(value.substring(i, i + 1));

			for (var m = 0; m < ai.length; m++) sum = sum + wi[m] * ai[m];

			remaining = sum % 11;
			if (remaining == 2) verifyNum = "X";
			else verifyNum = vi[remaining];

			return (verifyNum == value.substring(17, 18));
		}
	}

	// 是否验证成功
	this.success = true;

	/**
	 * 校验表单元素
	 * @param value 要验证的值
	 * @param ruleStr 规则字符串，多个规则之间用 | 隔开
	 */
	Validator.prototype.validate = function (value, ruleStr) {
		var rules = ruleStr.split("|");
		var ret = true;
		for(var i = 0; i < rules.length; i++) {
			ret = ret && pattern[rules[i]](value);
		}
		this.success = this.success && ret;
		return ret;
	}

	/**
	 * 获取密码强度
	 * 1.如果密码少于6位，那么就认为这是一个太弱密码。返回0
	 * 2.如果密码只由数字、小写字母、大写字母或其它特殊符号当中的一种组成，则认为这是一个弱密码。返回1
	 * 3.如果密码由数字、小写字母、大写字母或其它特殊符号当中的两种组成，则认为这是一个中度安全的密码。返回2
	 * 4.如果密码由数字、小写字母、大写字母或其它特殊符号当中的三种以上组成，则认为这是一个比较安全的密码。返回3
	 * @param value
	 * @returns {number}
	 */
	Validator.prototype.passwordRank = function (value) {
		if ( value.length <= 5 ) return 0;
		var mode = 0;
		//获取该字符串的所有组成模式
		for ( var i = 0; i < str.length; i++ ) {
			mode |= charMode(str.charCodeAt(i));
		}
		return getModeNum(mode);
	}

	Validator.prototype.isSuccess = function() {
		return this.success;
	}

	/**
	 * 计算一个字符所属的类型
	 * 数字|小写字母|大写字母|特殊字符
	 */
	function charMode( code ) {
		if ( code >= 48 && code <= 57 ) //数字 00000000 00000000 00000000 00000001
			return 1;
		if ( code >= 65 && code <= 90 ) //大写字母 00000000 00000000 00000000 00000010
			return 2;
		if ( code >= 97 && code <= 122 ) //小写 00000000 0000000 00000000 00000100
			return 4;
		else
			return 8; //特殊字符    0000000 0000000 00000000 00001000
	};

	/**
	 * 获取一共有多少种组合模式，并转换为十进制的表示模式
	 * @param number 模式总数
	 * 00000010
	 * 00000001
	 */
	function getModeNum( number ) {
		var modes = 0;
		for ( var i = 0; i < 4; i++ ) {
			if ( number & 1 ) modes++;
			number>>>=1;    //向右移动一位
		}
		return modes;
	};
}
