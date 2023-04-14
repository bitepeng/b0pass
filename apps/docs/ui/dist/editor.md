# Markdown

## editor.md
http://editor.md.ipandao.com/

```
init之后配置更新

.config() // 单个更改可以写成 .config('xxx' , xxx) , 多个可以传一个对象

取值

getValue() // 取值
getMarkdown() // 获取 Markdown 源码
getHTML(); // 获取 Textarea 保存的 HTML 源码
getPreviewedHTML(); // 获取预览窗口里的 HTML，在开启 watch 且没有开启 saveHTMLToTextarea 时使用
preview.html() // 获取预览区的html

设置值

setValue()
appendMarkdown() // 插入Markdown
setMarkdown() // 设置markdown内容

预览

watch() // 开启预览
unwatch() // 关闭预览

显示隐藏

show()/hide() // 因为基于jQuery 可以直接使用

工具条

showToolbar()/hideToolbar() // 显示隐藏工具条
setToolbarAutoFixed() // true/false 设置工具条固定

跳转到指定行数

gotoLine(num)
gotoLine('first') // 回到第一行
gotoLine('last') // 调到最后一行

设置主题

setTheme() // 工具条主题
setCodeMirrorTheme() // markdown编辑区主题
setPreviewTheme() // 预览区主题
编辑器主题可以看官网 demo

光标位置

setCursor() // {line:1, ch:2} 设置光标位置
getCursor() // 获取当前光标位置
setSelection() // {line:1, ch:0}, {line:5, ch:100} 设置选中文本
getSelection() // 获取选中文本内容
replaceSelection('xxx') // 替换选中文本为xxx
insertValue('xxx') // 在光标出插入文本xxx

全屏

fullscreen() // 全屏预览

内置事件

onload // 图片上传完成
onwatch/onunwatch // 打开预览/关闭预览
onchange // 内容变化
onscroll // 滚动
onpreviewscroll // 预览时滚动
onfullscreen/onfullscreenExit // 全屏/退出全屏
onresize // 尺寸变化
onpreviewing/onpreviewed // 预览/退出预览
```