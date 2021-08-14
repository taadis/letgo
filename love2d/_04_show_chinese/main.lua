function love.load()
    -- 使用系统绝对路径加载字符集为啥会失败哈?
    --Font = love.graphics.newFont("C:\\Windows\\Fonts\\msyh.ttc", 20)
    Font = love.graphics.newFont("fonts/msyh.ttc", 20)
end

function love.draw()
    love.graphics.setFont(Font)
    love.graphics.print("你好哇", 300, 120)
end
