FROM alpine:3.16.2
RUN apk add --no-cache mariadb-client
CMD ["sleep", "infinity"]
#ENTRYPOINT ["mysql"]

# 参考:https://hub.docker.com/_/alpine
# 使用alpine构建基础工具,查看镜像大小?

# 构建命令
# docker build . --tag taadis/tool:alpine3.16.2-mysql-client --no-cache
# docker images | grep mysql-client
# taadis/tool   alpine3.16.2-mysql-client   859932e0ad3b   37 seconds ago   40.3MB

# 还是比想象中大很多,能不能再小点?
# 构建过程中
#(1/7) Installing mariadb-common (10.6.9-r0)
#(2/7) Installing libgcc (11.2.1_git20220219-r2)
#(3/7) Installing ncurses-terminfo-base (6.3_p20220521-r0)
#(4/7) Installing ncurses-libs (6.3_p20220521-r0)
#(5/7) Installing libstdc++ (11.2.1_git20220219-r2)
#(6/7) Installing mariadb-client (10.6.9-r0)
#(7/7) Installing mysql-client (10.6.9-r0)
# 可以看到同时拉取了 mysql-client和mariadb-client还有几个其他依赖包?可以再简化下嘛?
