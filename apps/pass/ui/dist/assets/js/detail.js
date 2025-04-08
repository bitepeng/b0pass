layui.use(['tree', 'table', 'carousel', 'util'], function(){
    var layer = layui.layer
    ,carousel = layui.carousel
    ,$=layui.jquery

    //Func----------------------------------------//
    openPage=function (url, param) {
          var form = '<form action="'+url+'" target="_blank"  id="windowOpen" style="display:none">';
          for(var key in param) {
          form += '<input name="'+key+'" value="'+param[key]+'"/>';
          }
          form += '</form>';
          $('body').append(form);
          $('#windowOpen').submit();
          $('#windowOpen').remove();
      }

    //Data----------------------------------------//
    // file数据
    showFile=function(f,t){
      var token = localStorage.getItem('token') || '';
      var fpath=f.substr(0,f.lastIndexOf("/")+1);
      layer.load("showFile:::"); 
      console.log("showFile:",f,t,fpath);
      if(t=="img"){
        api_ajax("/pass/file-list?f="+fpath+"&t=img","GET",{},function(res){
          var findex=0;
          var str='<div class="layui-carousel" id="imgs"><div carousel-item>';
          for(var key in res.data) {
            if(res.data[key].path==decodeURIComponent(f)){
              findex=key;
              //console.log("imglist::"+findex+"|"+res.data[key].path+"|"+f);
            }
            str+='<div><img src="/files'+res.data[key].path+'?token='+token+'" style="max-width:100%;max-height:95vh;" onclick="openPage(\'http://'+window.location.host+'/files'+res.data[key].path+'\')"></div>';
          }
          str+="</div></div>";
          $("#content").html(str);
          carousel.render({
            elem: '#imgs'
            ,width: '100%'
            ,full: true
            ,arrow: 'always'
            ,index: findex
            ,change: function(obj){
                let fsrc=$("img",obj.item).attr("src");
                let newtitle=fsrc.substr(fsrc.lastIndexOf("/")+1);
                newtitle=newtitle.replace("?token="+token,"");
                parent.layer.title(newtitle,parent.layer.getFrameIndex(window.name));
            }
          });
          layer.closeAll('loading');  
        });
          
          //$("#content").html('<a href="/files'+f+'" target="_blank"><img src="/files'+f+'" style="max-width:100%;max-height:95vh;"></a>');
          layer.closeAll('loading');
      }else{
        api_ajax("/pass/file-content?f="+f+"&token="+token,"GET",{},function(res){
          console.log(f,t,res);
          //console.log("�::",(res.data).indexOf("�"))
          $("#content").html('<textarea  style="width:100%;height:88vh">'+HtmlUtil.htmlEncode(res.data)+'</textarea>'); 
          layer.closeAll('loading');  
        });
      }
    }

    var f =window.location.href.split("=")[1];
    var t =f.split("|")[1];
    f=f.split("|")[0];
    if(f!="" && f!='undefined'){
      showFile(f,t);
    }


    $("#open_blank").on("click",function(){
      //alert("x");
      openPage("http://"+window.location.host+"/files/"+f+"&token="+token);
    });


});