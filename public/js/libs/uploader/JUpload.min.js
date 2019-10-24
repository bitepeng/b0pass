/**
 * 基于iframe实现的异步上传插件
 * https://git.oschina.net/blackfox/ajaxUpload
 * @author yangjian<yangjian102621@gmail.com>
 * @version 1.0.0
 * @since 2016.06.02
 */
(function($) {
	var html5Support = false;
	if ( typeof(Worker) !== "undefined" ) {
		html5Support = true;
	}
	if ( Array.prototype.remove == undefined ) {
		Array.prototype.remove = function(item) {
			for ( var i = 0; i < this.length; i++ ) {
				if ( this[i] == item ) {
					this.splice(i, 1);
					break;
				}
			}
		}
	}
	//图片裁剪
	if ( !$.fn.imageCrop ) {
		$.fn.imageCrop = function(__width, __height) {
			$(this).on("load", function () {

				var width, height, left, top;
				var orgRate = this.width/this.height;
				var cropRate = __width/__height;
				if ( orgRate >= cropRate ) {
					height = __height;
					width = __width * orgRate;
					top = 0;
					left = (width - __width)/2;
				} else {
					width = __width;
					height = __height / orgRate;
					left = 0;
					//top = (height - __height)/2;
					top = 0;
				}
				$(this).css({
					"position" : "absolute",
					top : -top + "px",
					left : -left + "px",
					width : width + "px",
					height : height + "px"
				});
			});
		}
	}

	// 加载 css 文件
	var js = document.scripts, script = js[js.length - 1], jsPath = script.src;
	var cssPath = jsPath.substring(0, jsPath.lastIndexOf("/") + 1)+"jupload.min.css"
	$("head:eq(0)").append('<link href="'+cssPath+'" rel="stylesheet" type="text/css" />');

	//单个上传文件
	$.fn.JUpload = function(__options) {
		var options = $.extend({
			src : "src",
			url : null,
			maxFileSize : 2, //最大文件：2M
			onStart : function () {  //开始上传

			},
			onComplete : function () { //完成上传

			},
			onSuccess : function(data) { //上传一张图片成功回调
				//console.log(data);
			},
			onError : function () { //出错时回调

			},
			onRemove : function(data) { //删除一张图片回调
				//console.log(data);
			}, //删除一张图片回调
			messageHandler : function (message) {
				alert(message);
			},
			image_container : null,
			maxFileNum : 5, //最多上传文件的个数
			extAllow : "jpg|png|gif|jpeg",
			extRefuse : "exe|txt",
			datas : [], //初始化已上传图片
			twidth : 113,
			theight : 113
		}, __options);
		var images = []; //已经上传的图片列表
		var hasUoloaded = 0; //已上传图片数
		if ( options.datas.length > 0 ) {
			//添加图片
			for ( var i = 0; i < options.datas.length; i++ ) {
				addImage(options.datas[i]);
			}
			images = options.datas;
			hasUoloaded = images.length;
		}
		var frameName = "iframe_"+Math.random();
		var $form = $('<form action="'+options.url+'" target="'+frameName+'" enctype="multipart/form-data" method="post"></form>');
		var $input = $('<input type="file" name="'+options.src+'" class="upload-input" />');
		var $iframe = $('<iframe name="'+frameName+'" class="upload-iframe"></iframe>');
		//给按钮绑定点击事件
		$(this).on("click", function() {
			$input.trigger("click");
		});
		if (html5Support) { //html5 上传

			$input.on("change",  function () {
				if ( options.maxFileNum > 0 && hasUoloaded >= options.maxFileNum ) {
					__error__("您最多允许上传"+options.maxFileNum+"张图片。");
					return false;
				}
				options.onStart();
				setTimeout(function () {
					uploadFile($input[0].files[0]);
				}, 200);
			});

		} else { //通过iframe上传

			$input.on("change", function() {
				if ($input.val() == "") {
					__error__("请选择文件.");
					return false;
				}
				if ( options.maxFileNum > 0 && hasUoloaded >= options.maxFileNum ) {
					__error__("您最多允许上传"+options.maxFileNum+"张图片。");
					return false;
				}
				options.onStart();
				setTimeout(function() {
					$form[0].submit();
				}, 200);
			});
			$iframe.on("load", function() {
				try {
					var html = this.contentWindow.document.getElementsByTagName("pre")[0].innerHTML;
					if ( !html ) return false;
					var res = $.parseJSON(html);
					if ( res.code == "000" ) {
						if ( options.image_container != null ) {
							addImage(res.data.url);
						}
						hasUoloaded++;
						options.onSuccess(res.data);
						//清空input已选文件
						$input.val("");

					} else {
						__error__(res.message);
					}
					options.onComplete();

				} catch (e) {
					//console.log(e);
				}
			});
		}

		$form.append($input);
		$('body').append($form);
		$('body').append($iframe);
		if ( options.image_container ) {
			$("#"+options.image_container).addClass("clearfix");
		}

		//添加图片
		function addImage(src) {
			var builder = new StringBuilder();
			builder.append('<div class="img-wrapper"><div class="img-container" style="width: '+options.twidth+'px; height: '+options.theight+'px">');
			builder.append('<img src="'+src+'" data-src="'+src+'">');
			builder.append('<div class="file-opt-box clearfix"><span class="remove">删除</span></div></div></div>');
			var $image = $(builder.toString());
			$("#"+options.image_container).append($image);
			$image.find("img").imageCrop(options.twidth, options.theight);
			//删除图片
			$image.find(".remove").on("click", function() {
				try {
					var src = $(this).parent().prev().attr("src");
					images.remove(src);
					$image.remove();
					options.onRemove(images);
					hasUoloaded--;
					//清空input已选文件
					$input.val("");
				} catch (e) {console.log(e);}
			});

			images.push(src);
		}

		//upload file function(文件上传主函数), 使用 Html5 上传
		function uploadFile(file) {

			if ( !fileCheckHandler(file) ) {
				return;
			}
			// prepare XMLHttpRequest
			var xhr = new XMLHttpRequest();
			xhr.open('POST', options.url);
			//upload successfully
			xhr.addEventListener('load',function(e) {

				var res = $.parseJSON(e.target.responseText);
				if ( res.code == "000" ) {
					if ( options.image_container != null ) {
						addImage(res.data.url);
					}
					hasUoloaded++;
					options.onSuccess(res.data);
				} else {
					__error__("上传失败");
					options.onError();
				}

			}, false);

			// file upload complete
			xhr.addEventListener('loadend', function () {
				options.onComplete();
			}, false);

			xhr.addEventListener('error', function(e) {
				//log errors here
			}, false);

			xhr.upload.addEventListener('progress', function(e) {
				//set progress
			}, false);

			xhr.upload.addEventListener('loadstart', function(e) {

			}, false);

			// prepare FormData
			var formData = new FormData();
			formData.append(options.src, file);
			xhr.send(formData);

		}

		//file check handler(文件检测处理函数)
		function fileCheckHandler(file) {

			//检查文件大小
			var maxsize = options.maxFileSize * 1024 * 1024;
			if ( maxsize > 0 && file.size > maxsize ) {
				__error__("文件大小不能超过 "+options.maxFileSize + "MB");
				options.onError();
				return false;
			}

			//检查文件后缀名
			var ext = getFileExt(file.name);
			if ( ext && options.extAllow.indexOf(ext) != -1
				&& options.extRefuse.indexOf(ext) == -1 ) {
				return true;
			} else {
				__error__("非法的文件后缀 "+ext);
				options.onError();
				return false;
			}

		}

		//获取文件后缀名
		function getFileExt(filename) {

			var position = filename.lastIndexOf('.')
			if ( position != -1 ) {
				return filename.substr(position+1).toLowerCase();
			}
			return false;
		}

		//显示错误信息
		function __error__(message) {
			options.messageHandler(message);
		}
	}


	//string builder
	var StringBuilder = function() {
		var buffer = new Array();
		StringBuilder.prototype.append = function(str) {
			buffer.push(str);
		}
		StringBuilder.prototype.toString = function () {
			return buffer.join("");
		}

	}
})(jQuery);
