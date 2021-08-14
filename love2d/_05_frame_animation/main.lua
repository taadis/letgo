local anim8 = require("anim8")
local backgroundImage
local doorImage
local doorSpriteBatch
local doorAnimation

function love.load()
    backgroundImage = love.graphics.newImage("images/background.png")
    doorImage = love.graphics.newImage("images/door.png")
    --doorSpriteBatch = love.graphics.newSpriteBatch(doorImage, 10, "dynamic")
    local doorImageWidth = doorImage:getWidth()
    local doorImageHeight = doorImage:getHeight()
    local doorGrid = anim8.newGrid(496, 361, doorImageWidth, doorImageHeight)
    local frames = doorGrid("3-2", 1)
    local ss = doorGrid:getFrames("1-4", 1, "1-4", 2, "1-2", 3)
    doorAnimation = anim8.newAnimation(ss, 0.15)

end

function love.update(dt)
    doorAnimation:update(dt)
end

function love.draw()
    love.graphics.draw(backgroundImage)
    -- love.graphics.draw(doorImage, 154, 96)
    doorAnimation:draw(doorImage, 154, 96)
end

