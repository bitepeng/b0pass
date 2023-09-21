/**
 * 全局变量 默认设置
 */
var fpath = window.location.href.split("=")[1];
var servIP,servPort;
var areaBig = ['100%','100%'];
var areaSmall = ['100%','100%'];
var currPath = "/";
var showType = localStorage.getItem("showType") || 1;
var showTypes = ["图文","列表"];

/**
 * 初始化布局
 */
document.getElementById("show-type-label").innerText=showTypes[showType];
var showIco = {field:'name', title: '文件名', minWidth: 180, sort: true, templet: function(res){
  return '<i class="layui-icon layui-icon-right"></i>&nbsp; '+res.name;
}};
var showOpt = {/*fixed: 'right',*/ title:' 操作', width: 70, minWidth: 70, toolbar: '#barDT'};
if(showType==0){
  showIco={field:'name', title: '文件名', minWidth: 180, templet: function(res){
    var str="";
    if(res.type=="dir"){str+='<img src="assets/img/dir.png"  width="32" height="32">&nbsp; '+res.name
    }else if(res.type=="pdf"){str+='<img src="assets/img/pdf.png"  width="32" height="32">&nbsp; '+res.name
    }else if(res.type=="img"){str+='<img src="/files'+res.path+'"  width="200" height="150"><br>'+res.name
    }else if(res.type=="vod"){str+='<video id="video" controls="controls"  width="200" height="150"><source  src="/files'+res.path+'"></video><br>'+res.name
    }else{str+='<img src="assets/img/file.png"  width="32" height="32">&nbsp; '+res.name
    }
    return str;
  }}
}
var tableCols = [
  showIco
  ,{field:'ext', title: '类型', width: 70, sort: true,templet: function(res){return (res.ext).replace(".","");}}
  ,{field:'sizes', title: '大小', width: 70, sort: true}
  ,{field:'date', title:'修改时间', width: 100, sort: true}
  ,showOpt
];

/**
 * 根据浏览器宽度重新定义
 */
var pageWidth = document.body.clientWidth;
var uploadhtm = "upload.html";
if(pageWidth>=1280){areaBig = ['1280px','90%'];areaSmall = ['480px','500px'];
}else if(pageWidth>=900){areaBig = ['900px','90%'];areaSmall = ['480px','500px'];
}else if(pageWidth<900){
  uploadhtm = "upload_h5.html";
  areaBig = ['100%','100%'];areaSmall = ['100%','100%'];
  if(showType==0){
    showIco={field:'name', title: '文件名', minWidth: 100, height:"300",  sort: true, templet: function(res){
      var str="";
      if(res.type=="dir"){str+= '<img src="assets/img/dir.png"  width="48" height="48">&nbsp; '+res.name
      }else if(res.type=="pdf"){str+='<img src="assets/img/pdf.png"  width="32" height="32">&nbsp; '+res.name
      }else if(res.type=="img"){str+= '<img src="/files'+res.path+'"  width="200" height="150"><br>'+res.name
      }else if(res.type=="vod"){str+='<video id="video" controls="controls" width="200" height="150"><source  src="/files'+res.path+'"></video><br>'+res.name
      }else{str+= '<img src="assets/img/file.png"  width="32" height="32">&nbsp; '+res.name
      }
      return str;
    }}
  }
  tableCols = [
    showIco,{field:'date', title:'修改时间', width: 0, hide: true},showOpt
  ];
}

/**
 * LayUI APP
 */
layui.use(['tree', 'table','form','dropdown','util'], function(){
      var $=layui.jquery
      ,layer = layui.layer
      ,tree = layui.tree
      ,table = layui.table
      ,form = layui.form
      ,dropdown = layui.dropdown
      ,util = layui.util

      //Func----------------------------------------//
      // 新窗口打开
      var openPage=function (url, param) {
          var form = '<form action="'+url+'" target="_blank"  id="windowOpen" style="display:none">';
          for(var key in param) {
          form += '<input name="'+key+'" value="'+param[key]+'"/>';
          }
          form += '</form>';
          $('body').append(form);
          $('#windowOpen').submit();
          $('#windowOpen').remove();
      }
      // 删除操作
      var NodeDelete=function(obj){
        $.ajax({
          url: "/pass/file-list?f="+obj.data.path,
          method: "get",
          data: {},
          success: function(res) {
            if(res.data){
              layer.msg("请先删除文件夹下所有文件");
            }else{
              layer.confirm('真的要删除吗？', function(index){
                  obj.del();
                  $.ajax({
                    url: "/pass/node-delete?f="+obj.data.path,
                    method: "get",
                    data: {},
                    success: function(res) {
                      layer.msg(obj.data.path+" 删除完成");
                    }
                  });
                  layer.close(index);
              });
            }
          },
        });
        return ;
        
      }
      // 浏览器打开
      var BrowserFile=function(obj){
        layer.msg("打开文件到浏览器");
        openPage("http://"+servIP+servPort+"/files"+obj.data.path,{});
      }
      // 主电脑打开
      var OpenFile=function(obj){
        $.ajax({
          url: "/pass/cmd-open?f="+obj.data.path,
          method: "get",
          data: {},
          success: function(res) {
            if(res){
              layer.msg(res.msg);
            }else{
              layer.msg("已在主电脑打开目录");
            }
          }
        });
      }
      // 文件扫码
      var ScanFile=function(obj){
        layer.open({
            title: "文件扫码",
            type: 2,
            area: areaSmall,
            content: 'qrcode.html?f='+servIP+servPort+"/files/"+encodeURIComponent(obj.data.path)
        });
      }
      // 下载操作
      var DownLoad=function(obj){
        //openPage("/pass/file-download",{"f":obj.data.path});
        var url = '/pass/file-download?f='+obj.data.path, fileName = '未知文件';
	      const a = document.createElement('a');
	      a.style.display = 'none';
	      a.setAttribute('target', '_blank');
	      fileName && a.setAttribute('download', fileName);
	      a.href = url;
	      document.body.appendChild(a);
	      a.click();
	      document.body.removeChild(a);
	    }
      // 重命名操作
      var ReName=function(obj){
        let fname=(obj.data.path).replace(currPath+"/","");
        let str=($("#node-rename").html()).replace("$$",fname);
        str=str.replace("$$",fname);
        layer.open({
          title: "重命名："+obj.data.path,
          area: areaSmall,
          type: 1, 
          content: '<div class="padding15"><div class="layui-form-item" align="center"><p><br></p></div>'+str+'</div>',
          cancel: function () {}
        });
      }
      //重命名Do
      form.on('submit(node-rename-form)', function(data){
          doc_src=data.field['doc-name-src'];
          doc_name=data.field['doc-name-input'];
          var loading = layer.msg('正在重命名', {icon: 16, shade: 0.3, time:0});
            $.ajax({
                url:'/pass/node-rename?f='+currPath+'/'+doc_src+'&n='+currPath+'/'+doc_name,
                type: 'get',
                dataType:"json",
                headers : {'Content-Type' : 'application/json;charset=utf-8'}, 
                success:function(rs){
                  layer.close(loading);
                  if(rs.code!=0){layer.alert(rs.msg);}else{tableRender(currPath);}
                },
                error: function(){layer.close(loading);}
            })
          layer.closeAll();
          return false;
      });
      //新建目录
      var NodeAdd=function(){
        layer.open({
          title: "新建目录："+currPath,
          area: areaSmall,
          type: 1, 
          content: '<div class="padding15"><div class="layui-form-item" align="center"><p><br></p></div>'+$("#node-add").html()+'</div>',
          cancel: function () {}
        });
      }
       //新建目录Do
       form.on('submit(node-add-form)', function(data){
        doc_add=data.field['doc-add-input'];
        var loading = layer.msg('正在创建', {icon: 16, shade: 0.3, time:0});
            $.ajax({
                url:'/pass/node-add?f='+currPath+"/"+doc_add+"/",
                type: 'get',
                dataType:"json",
                headers : {'Content-Type' : 'application/json;charset=utf-8'}, 
                success:function(rs){
                  layer.close(loading);
                  if(rs.code!=0){layer.alert(rs.msg);}else{tableRender(currPath);}
                },
                error: function(){layer.close(loading);}
            })
          layer.closeAll();
          return false;
      });

      //Data----------------------------------------//
      //显示详情
      var showDetail=function(obj){
          if(obj.data.type=="dir"){
            tableRender(obj.data.path);
          }else{
            var type = obj.data.type;
            if(type=="img" || type=="html" || type=="code"){
                layer.open({
                  title: obj.data.name,
                  type: 2,
                  area: areaBig,
                  maxWidth:1280,
                  content: 'detail.html?f='+obj.data.path+"|"+obj.data.type
              });
            }else{
              openPage("/files"+obj.data.path,{});
            }
          } 
      };

      //menu-----------------------------------------//
      dropdown.render({
        elem: '#show-type'
        ,data: [{title: '图文',id: 0},{title: '列表',id: 1}]
        ,id: 'show-type'
        //菜单被点击的事件
        ,click: function(obj){
          layer.msg("切换到"+showTypes[obj.id]+"模式");
          localStorage.setItem("showType",obj.id)
          $("#show-type-label").text(showTypes[obj.id]);
          setTimeout(function(){
            window.location.reload(true);
          },1000);
        }
      });


      //tree----------------------------------------//
      //tree数据
      var treeRender=function(f){
          $.ajax({
            url: "/pass/node-tree?f="+f,
            method: "get",
            data: {},
            success: function(res) {
                console.log("::treeRender::","/pass/node-tree?f="+f);//,res.data
                tree.render({
                    elem: '#tree-left'
                    ,data: [res.data]
                    ,accordion: true 
                    ,isJump: true
                    ,click: function(obj){
                        tableRender(obj.data.path);
                    }
                });
            },
          });
      }

      //table----------------------------------------//
      // 创建渲染表格实例
      var tableRender=function(f){
          //path
          currPath = f;
          currPath=currPath.replace("//","/");
          currPath=currPath.replace("#","");
          updateUrl("f",currPath) 
          treeRender(currPath);
          $("#crrPath").html(f);
          //data
          table.render({
              elem: '#dataTable'
              ,url:'/pass/file-list?f='+f
              ,height: 'full'
              ,cellMinWidth: 80
              ,initSort: {
                field: 'date'
                ,type: 'desc'
              }
              ,page: false
              ,cols: [tableCols]
              ,done: function(){
                console.log("::tableRender::",'/pass/file-list?f='+f);//,this
              }
              ,error: function(res, msg){
                console.log(res, msg)
              }
          });
          // 工具栏事件
          table.on('toolbar(dataTable)', function(obj){
              var id = obj.config.id;
              var checkStatus = table.checkStatus(id);
              var othis = lay(this);
              switch(obj.event){
                case 'getCheckData':
                  var data = checkStatus.data;
                  layer.alert(layui.util.escape(JSON.stringify(data)));
                break;
            };
          });
        
          //触发单元格工具事件
          table.on('tool(dataTable)', function(obj){
            var that = this;
            if(obj.event === 'more'){
              //更多下拉菜单
              dropdown.render({
                elem: that
                ,show: true //外部事件触发即显示
                ,data: 
              [/*{title: '详情',id: 'datail'}, */
                {title: (obj.data.ext).replace(".",""),id: '-'},
                {title: obj.data.sizes,id: '-'},
                {title: obj.data.date,id: '-'},{type:'-'},
                {title: "浏览器打开",id: 'browser'},
                {title: "主电脑打开",id: 'open'},
                {title: "扫码",id: 'scan'},
                {title: "下载",id: 'down'},{type:'-'},
                {title: "重命名",id: 'rename'},
                {title: '删除',id: 'del'}
              ]
                ,click: function(data, othis){
                  //根据 id 做出不同操作
                  if(data.id === 'datail'){     //详情
                    showDetail(obj);
                  }else if(data.id === 'browser'){ //浏览器打开
                    BrowserFile(obj)
                  }else if(data.id === 'open'){ //主电脑打开
                    OpenFile(obj)
                  }else if(data.id === 'scan'){ //扫码
                    ScanFile(obj)
                  }else if(data.id === 'down'){ //下载
                    DownLoad(obj)
                  }else if(data.id === 'rename'){//重命名
                    ReName(obj)
                  }else if(data.id === 'del'){   //删除
                    NodeDelete(obj);
                  }
                }
                ,align: 'right'
                ,style: 'box-shadow: 1px 1px 10px rgb(0 0 0 / 12%);'
              }); 
            }

          });
        
          //触发表格复选框选择
          table.on('checkbox(dataTable)', function(obj){
            console.log(obj)
          });
        
          // 行单击事件
          table.on('row(dataTable)', function(obj){
              showDetail(obj);
          });
          // 行双击事件
          table.on('rowDouble(dataTable)', function(obj){
            console.log(obj);
          });
        
      }


      //tools----------------------------------------//

      $(".reload").on("click",function(){
        layer.msg("刷新中");
        window.location.reload(true);
      })

      //网址二维码
      $("#link_phone").on("click",function(){
        layer.open({
            title: "主电脑参数",
            type: 2,
            area: areaSmall,
            content: 'qrcode.html?f='+servIP+servPort
        });
      })

      //浏览按钮
      $("#btn_left_open").on("click",function(){
        layer.msg("打开文件目录到浏览器");
        openPage("/files"+currPath,{});
      })

      //打开按钮
      $("#btn_left_dir").on("click",function(){
        $.ajax({
          url: "/pass/cmd-open?f="+currPath,
          method: "get",
          data: {},
          success: function(res) {
            layer.msg("已在主电脑打开目录");
          }
        });
      })

      //上传按钮
      $("#btn_main_upload").on("click",function(){
          layer.open({
              title: "上传文件",
              type: 2,
              area: areaBig,
              content: uploadhtm+'?f='+currPath+"/",
              cancel: function () {
                tableRender(currPath);
              }
          });
      })

      //添加按钮
      $(".btn_main_new").click(function(){
        //if(currPath!="/"){currPath=currPath+"/";}
        NodeAdd();
      });

     
      //----------------------------------------------//

      //刷新重载页面数据
      console.log("::currPath::",fpath);
      if(fpath!="" && fpath!=undefined){
        tableRender(decodeURI(fpath));
      }else{
        tableRender(currPath);
      }
      
      $.ajax({
        url: "/gateway/config",
        success: function(res) {
          console.log("::Config::",res.data);
          ips=(res.data.ListenAddr).split(":")
          if(ips[0]!=""){
            if(ips[0]="127.0.0.1"){
              alert("如果IP被设置为127.0.0.1，意味着只能本机使用，将无法分享文件！\n若无特殊需求，请将配置文件的'ListenAddr'设置为纯端口，如':8899'。");
            }
            servIP=ips[0];
            console.log("::ServIP::",servIP);
          }
          //兼容域名情况
          if(res.data.Domain!=""){
            servIP=res.data.Domain;servPort="";
            console.log("::ServIP::",servIP);
          }else{
            $.ajax({
              url: "/pass/read-ip",
              success: function(res) {
                servIP=res.data;
                console.log("::ServIP::",servIP);
              }
            });
            servPort=":"+(res.data.ListenAddr).split(":")[1];
            console.log("::ServPort::",servPort);
          }
        }
      });

      //----------------------------------------------//

});


/**
 * 不刷新页面更新URL参数
 */
function updateUrl( key, value){
    var newurl = updateQueryStringParameter(key, value)
    //向当前url添加参数，没有历史记录
    window.history.replaceState({
        path: newurl
    }, '', newurl);
  }

function updateQueryStringParameter(key, value) {
    var uri = window.location.href
    if(!value) {return uri;}
    var re = new RegExp("([?&])" + key + "=.*?(&|$)", "i");
    var separator = uri.indexOf('?') !== -1 ? "&" : "?";
    if (uri.match(re)) {
      return uri.replace(re, '$1' + key + "=" + value + '$2');
    }else {
      return uri + separator + key + "=" + value;
    }
 }