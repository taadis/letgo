# linux namespace

## linux namespace的产生背景

Linux namespace（命名空间）的产生背景和原因主要可以归结为对系统资源隔离的需求，以及随着虚拟化技术的发展而兴起的一种轻量级虚拟化技术。

- 资源隔离的需求：
在多用户、多任务的环境下，系统需要为不同的用户或任务提供相互隔离的资源环境，以防止它们之间的相互干扰。Linux namespace 技术应运而生，为进程和应用程序提供了一个独立的资源视图，从而实现了资源的有效隔离。
- 虚拟化技术的发展：
随着虚拟化技术的兴起，人们希望能够在单一物理机上运行多个相互隔离的虚拟环境。传统的虚拟化技术（如完全虚拟化、半虚拟化）虽然能够提供较好的隔离性，但往往伴随着较高的资源开销。Linux namespace 作为一种轻量级虚拟化技术，能够在提供隔离性的同时，减少资源消耗，提高系统效率。
- 容器技术的兴起：
容器技术（如Docker）的快速发展，对资源隔离提出了更高的要求。Linux namespace 技术为容器技术提供了底层支持，使得容器能够拥有独立的网络、进程、文件系统等资源视图，从而实现了更加高效的资源隔离和管理。

## 什么是linux namespace?

Linux namespace是Linux内核提供的一种特性，它允许在单个宿主机系统中创建多个隔离的环境。每个namespace可以拥有一套独立的系统资源，例如进程ID、网络接口、挂载点等，从而实现资源的隔离和管理。这种隔离机制使得每个namespace看起来像一个独立的系统，尽管它们实际上都运行在同一个物理系统上。

## linux有哪些namespace类型?

- PID Namespace (进程标识符命名空间)：用于隔离进程的标识符和进程层次结构，进程标识符是唯一的，PID Namespace 允许不同 Namespace 中的进程拥有相同的 PID，使得每个namespace中的进程都认为自己是PID为1的进程，独立的pid空间。
- UTS Namespace (UNIX时间系统命名空间)：用于隔离内核中维护的 UNIX 系统主机名和域名，使得每个namespace可以拥有独立的hostname和domainname，独立的UTS namespace中的进程认为自己是在不同的主机上。
- Mount Namespace (文件系统挂载命名空间)：用于隔离进程对文件系统上不同目录的访问，使得每个namespace中的进程认为自己只有自己的根目录，独立的mount namespace中的进程可以。允许每个进程在自己的namespace中拥有独立的文件系统挂载点，从而实现文件系统的隔离和管理。
- IPC Namespace (进程间通信命名空间)：用于隔离内核中进程间通信资源，使得每个namespace中的进程认为自己只有自己的System V IPC，AIX的semget()与Unix9。用于隔离进程间通信的机制，包括共享内存、信号量、消息队列等，确保不同namespace中的进程无法直接访问其他namespace中的IPC资源。
- Network Namespace (网络命名空间)：用于将一个网络栈隔离到一组进程，使得每个namespace中的进程认为自己只有自己的网卡、路由表、防火墙规则等，独立于其他namespace中的。用于隔离网络栈和网络资源，包括网络设备、IP地址、路由表、网络协议栈等，使得每个namespace可以有自己独立的网络配置和环境。
- User Namespace (用户命名空间)：用于将系统用户和用户组映射到指定的用户和用户组，使得每个namespace中的进程认为自己有独立的root用户，独立于其他namespace中的。用于隔离用户权限，允许每个namespace有自己的用户和用户组ID空间，从而实现用户级别的隔离和安全性。
- Cgroup Namespace (cgroup命名空间)：用于将进程组隔离到指定的cgroup，使得每个namespace中的进程认为自己有独立的cgroup，独立于其他namespace中的。用于隔离cgroup资源，包括CPU、内存、磁盘I/O等，确保不同namespace中的进程无法直接访问其他namespace中的cgroup资源。用于隔离和管理控制组（cgroup）层次结构，使得在不同的namespace中的进程可以有各自独立的cgroup视图，从而实现资源的隔离和管理。

这些namespace类型在Linux内核中提供了强大的隔离能力，是容器技术如Docker得以实现的基础。通过组合使用这些namespace，可以为每个容器创建一个独立、隔离的运行环境。

## go中如何使用namespace?

见代码.
