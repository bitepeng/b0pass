layui.use(['upload', 'element'], function(){
    var $ = layui.jquery
    ,upload = layui.upload
    ,element = layui.element
    ,layer = layui.layer;
    
    var fpath =decodeURI(window.location.href.split("=")[1]);
    if(fpath!="" && fpath!=undefined){
      fpath=fpath.replace("\/\/","\/");
    }else{
      fpath="/";
    }
    fpath=fpath.replace("//","/");
    console.log("upload fpath: ",fpath);
    $("#f").html("上传到 "+fpath +" 目录");
    
    
    //多文件列表
    var token = localStorage.getItem('token') || '';
    $(".layui-upload-list").hide();
    $("#submitAct").hide();
    var uploadListIns = upload.render({
      elem: '#selectFile'
      ,elemList: $('#fileList')
      ,url: '/pass/file-upload?f='+fpath
      ,headers:{token:token}
      ,accept: 'file'
      ,multiple: true
      ,number: 100
      ,auto: false
      ,bindAction: '#submitAct'
      ,choose: function(obj){
        var that = this;
        //将每次选择的文件追加到文件队列
        that.files = obj.pushFile(); 
        console.log(that.files);
        //显示文件列表
        $(".layui-upload-list").show();
        $("#submitAct").show();
        $("#selectFile").text("继续添加");
        
        //读取本地文件
        that.elemList.html("");
        for (var index in that.files) {
            var file=that.files[index];
            //人性化Size
            var fsize=file.size;
            var fsize_m=fsize/1024/1024;
            if(fsize_m>=1024){
              fsize=(fsize_m/1024).toFixed(1)+'G';
            }else if((fsize/1024)>=1024){
              fsize=fsize_m.toFixed(1)+'M';
            }else{
              fsize=(fsize/1024).toFixed(1)+'K';
            }
            //文件列表
            var tr = $(['<tr id="upload-'+ index +'">'
              ,'<td style="overflow: hidden;text-overflow:ellipsis;white-space:normal;">'+ file.name +'<br>',
              ,'<button class="layui-btn layui-btn-xs act-reload layui-hide">重传</button>'
              ,'<button class="layui-btn layui-btn-xs layui-btn-danger act-delete">移除</button>'
              +' Size:'+ fsize+'<br>'
              ,'</td>'
              ,'<td>'+
              '<div class="layui-progress" lay-filter="progress-file-'+ index +'">'
              ,'</td>'
            ,'</tr>'].join(''));
             //单个重传
            tr.find('.act-reload').on('click', function(e){
              var tr_ = e.currentTarget.parentElement.parentElement;
              var index_ = (tr_.id).replace("upload-","");
              obj.upload(index_, file);
            });
            //移除文件
            tr.find('.act-delete').on('click', function(e){
              var tr_ = e.currentTarget.parentElement.parentElement;
              var index_ = (tr_.id).replace("upload-","");
              delete that.files[index_];
              tr_.remove();
              uploadListIns.config.elem.next()[0].value = ''; 
            });
            //插入DOM
            that.elemList.append(tr);
            element.render('progress');
        };
      }
      ,before: function(obj){
          layer.load();
          $("#submitAct").hide();
      }
      ,done: function(res, index, upload){
          var that = this;
          var tr = that.elemList.find('tr#upload-'+ index)
          ,tds = tr.children();
          //删除文件队列已经上传成功的文件
          delete that.files[index]; 
          return;
          that.error(index, upload);
      }
      ,allDone: function(obj){ 
          layer.closeAll('loading');
          layer.msg("上传完成!");
          console.log("上传完成:",obj)
          var index=parent.layer.getFrameIndex(window.name);
          parent.location.reload(false);
          parent.layer.close(index);
      }
      ,error: function(index, upload){
          layer.close('loading');
          var that = this;
          var tr = that.elemList.find('tr#upload-'+ index)
          ,tds = tr.children();
          //显示重传
          tds.eq(0).find('.file-reload').removeClass('layui-hide'); 
      }
      ,progress: function(n, elem, e, index){
        var that = this;
        var tds = that.elemList.find('tr#upload-'+ index).children();
        tds.eq(1).html(n + '%'); 
        element.progress('progress-file-'+ index, n + '%'); 
      }
    });
    
    
  });