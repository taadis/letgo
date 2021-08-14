# 02_conf

首先,新建一个文件夹,并在该文件夹下新建文件 main.lua 和 conf.lua

conf.lua 文件源码内容如下:

```
function love.conf(t)
    -- 设置标题和窗口大小
    t.title = "_02_conf"
    t.screen.width = 240
    t.screen.height = 320
end

```

main.lua 文件源码内容如下:

```
-- 加载资源函数,仅初始化时调用一次
function love.load()

end

-- 绘图函数,每周期都会调用
function love.draw()

end

-- 更新函数,每周期都会调用
function love.update(dt)

end

-- 键盘检测函数,当键盘事件触发时调用
function love.keypressed(key)

end

```

在当前目录下执行以下命令启动游戏

```
love .
```

不出意外的话，会看到一个黑乎乎的游戏窗口。

> 为什么我看到的是一个会报错的蓝屏?

注意看窗口的标题和大小是否同配置里是一致的，可以调整标题文件和大小再试试看看效果。

## 解惑

首先 conf.lua 会被先加载，你可以在 conf.lua 中覆盖 love 的默认配置或加入你自己的配置。

love 的所有默认配置如下。

> 禁用一些不用的模块，可以提高载入速度。但是最好不要禁用这俩模块 love,love.filesystem

```
function love.conf(t)
    t.title = "Untitled" -- The title of the window the game is in (string)
    t.author = "Unnamed" -- The author of the game (string)
    t.url = nil -- The website of the game (string)
    t.identity = nil -- The name of the save directory (string)
    t.version = "0.8.0" -- The LÖVE version this game was made for (string)
    t.console = false -- Attach a console (boolean, Windows only)
    t.release = false -- Enable release mode (boolean)
    t.screen.width = 800 -- The window width (number)
    t.screen.height = 600 -- The window height (number)
    t.screen.fullscreen = false -- Enable fullscreen (boolean)
    t.screen.vsync = true -- Enable vertical sync (boolean)
    t.screen.fsaa = 0 -- The number of FSAA-buffers (number)
    t.modules.joystick = true -- Enable the joystick module (boolean)
    t.modules.audio = true -- Enable the audio module (boolean)
    t.modules.keyboard = true -- Enable the keyboard module (boolean)
    t.modules.event = true -- Enable the event module (boolean)
    t.modules.image = true -- Enable the image module (boolean)
    t.modules.graphics = true -- Enable the graphics module (boolean)
    t.modules.timer = true -- Enable the timer module (boolean)
    t.modules.mouse = true -- Enable the mouse module (boolean)
    t.modules.sound = true -- Enable the sound module (boolean)
    t.modules.physics = true -- Enable the physics module (boolean)
end

```

main.lua 里使我们要处理的游戏逻辑，主要是一些回调函数，会被 love 自动调用。

主要的回调函数如下：

```
love.draw -- Callback function used to draw on the screen every frame.
love.focus -- Callback function triggered when window receives or loses focus.
love.joystickpressed -- Called when a joystick button is pressed.
love.joystickreleased -- Called when a joystick button is released.
love.keypressed -- Callback function triggered when a key is pressed.
love.keyreleased -- Callback function triggered when a key is released.
love.load -- This function is called exactly once at the beginning of the game.
love.mousepressed -- Callback function triggered when a mouse button is pressed.
love.mousereleased -- Callback function triggered when a mouse button is released.
love.quit -- Callback function triggered when the game is closed.
love.run -- The main function, containing the main loop. A sensible default is used when left out.
love.update -- Callback function used to update the state of the game every frame.

```
