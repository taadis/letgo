# love2d - 03 background image

准备好背景图片 background.jpg 放到 images 目录下。

新建文件 main.lua，源码内容如下：

```
background = love.graphics.newImage("images/background.jpg")

function love.draw()
    love.graphics.draw(background, 0, 0)
end

```

在 main.lua 所在目录下执行以下命令启动游戏

```
love .
```

不出意外的话，你会看到一个游戏窗口，上面是一片绿油油的草地...