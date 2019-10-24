/**
 * 图片预览
 * @author yangjian
 */
(function($) {

	$.fn.jpreview = function() {

		$.each(this, function(i, item) {
			if($(item).attr('event-bind') == 'yes') {
				return true;
			}
			$(item).attr('event-bind', 'yes');
			$(item).on("click", function() {

				var src = $(this).attr('data-src');
				if ( !src ) {
					alert("图片地址不存在，无法预览.");
					return;
				}
				var timer = 500;
				var $lock = $('<div style="position: absolute; background: #000000; opacity: .6; z-index: 666; display: none;"></div>');
				var $imgBox = $('<div style="display: none; position: absolute; z-index: 999; "></div>');
				var $img = $('<img src="" />');
				$img.css({
					"max-width" : $(window).width() - 100 + "px",
					"max-height" : $(window).height() - 100 + "px"
				});
				var $closeBtn = $('<span style="position: absolute; font-size: 30px; border-radius: 40px; font-weight: bold;' +
					'height: 40px; width: 40px; line-height: 35px; z-index: 999; background: #ffffff; text-align: center; cursor: pointer;">x</span>');
				$closeBtn.css({
					top : -20,
					right : -20
				});
				$closeBtn.on('click', function() {
					$lock.fadeOut(timer, function() {$lock.remove()});
					$imgBox.fadeOut(timer, function() {$imgBox.remove()});
				});

				$lock.css({
					left : 0,
					top : 0,
					width : $(document).width(),
					height : $(document).height()
				});
				$lock.on('click', function() {
					$lock.fadeOut(timer, function() {$lock.remove()});
					$imgBox.fadeOut(timer, function() {$imgBox.remove()});
				});

				var _scrollTop = window.document.body.scrollTop || window.document.documentElement.scrollTop;
				$imgBox.css({
					top : $(window).height() / 2 + _scrollTop,
					left : $(window).width() / 2,
				});
				$img.attr('src', src);
				$img.on("load", function() {
					$imgBox.css({
						top : ($(window).height() - $imgBox.height()) / 2 + _scrollTop,
						left : ($(window).width() - $imgBox.width()) / 2,
					}).fadeIn(timer);
				});
				$imgBox.append($img);
				$imgBox.append($closeBtn);

				$('body').append($lock.fadeIn(timer));
				$('body').append($imgBox);

			});
		});

	}

})(jQuery);