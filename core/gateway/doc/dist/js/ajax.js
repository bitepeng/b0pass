function ajax(url, data, method, success) {
    // 异步对象
    var ajax = new XMLHttpRequest();

    // get 跟post  需要分别写不同的代码
    if (method == 'get') {
        // get请求
        if (data) {
            // 如果有值
            url += '?';
            url += data;
        } else {

        }
        // 设置 方法 以及 url
        ajax.open(method, url);

        // 需要设置请求报文
        //ajax.setRequestHeader("Content-type","application/x-www-form-urlencoded");
        ajax.setRequestHeader("Content-type", "application/json; charset=utf-8");

        // send即可
        ajax.send();
    } else {
        // post请求
        // post请求 url 是不需要改变
        ajax.open(method, url);

        // 需要设置请求报文
        ajax.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
        //ajax.setRequestHeader("Content-type","application/json; charset=utf-8");

        // 判断data send发送数据
        if (data) {
            // 如果有值 从send发送
            ajax.send(data);
        } else {
            // 没有值 直接发送即可
            ajax.send();
        }
    }

    // 注册事件
    ajax.onreadystatechange = function () {
        // 在事件中 获取数据 并修改界面显示
        if (ajax.readyState == 4 && ajax.status == 200) {
            e = JSON.parse(ajax.responseText);
            console.log(e);
            success(e);
        }
    }

}


function formData(form) {
    var arr = {};
    for (var i = 0; i < form.elements.length; i++) {
        var feled = form.elements[i];
        switch (feled.type) {
            case undefined:
            case 'button':
            case 'file':
            case 'reset':
            case 'submit':
                break;
            case 'checkbox':
            case 'radio':
                if (!feled.checked) { break; }
            default:
                if (arr[feled.name]) {
                    arr[feled.name] = arr[feled.name] + ',' + feled.value;
                } else {
                    arr[feled.name] = feled.value;
                }
        }
    }
    arr = JSON.stringify(arr);
    return arr
}