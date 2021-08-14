# love2d - 04 show chinese

love2d 可以通过以下语句直接打印输输出英文

```
love.graphics.print("hello world", 100, 100)
```

但是换成中文，输出的就变成几个小方格了，显示不太正常。

可以通过加载中文字体来解决这个问题。

可以选择开源的一些 ttf 字符集。

Windows 操作系统下可以在以下目录下找到系统自带字符集。

```
C:\\Windows\\Fonts
```

> 但在代码中直接加载系统绝对路径字符集会提示载入失败，拷贝一份放到项目目录下却又可以加载成功，黑人问号脸？

## 参考

- [Love2D游戏开发教程03 - 使用中文字体 (哔哩哔哩)](https://www.bilibili.com/video/av414318490)