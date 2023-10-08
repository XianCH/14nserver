---
title: "Go基础标准库"
date: 2023-10-05T15:37:25+08:00
lastmod: 2023-10-05T15:37:25+08:00
author: ["x14n"]
keywords: 
- 
categories: # 没有分类界面可以不填写
- 
tags: # 标签
- 
description: ""
weight:
slug: ""
draft: false # 是否为草稿
comments: true # 本页面是否显示评论
reward: false # 打赏
mermaid: true #是否开启mermaid
showToc: true # 显示目录
TocOpen: true # 自动展开目录
hidemeta: false # 是否隐藏文章的元信息，如发布日期、作者等
disableShare: true # 底部不显示分享栏
showbreadcrumbs: true #顶部显示路径
cover:
    image: "" #图片路径例如：posts/tech/123/123.png
    zoom: # 图片大小，例如填写 50% 表示原图像的一半大小
    caption: "" #图片底部描述
    alt: ""
    relative: false
---



# os

`os` 包是 Go 语言标准库的一部分，主要用于与操作系统进行交互，提供了访问操作系统底层功能的功能。下面是关于 `os` 包的一些重要功能和用法：

1. **文件和目录操作**：
   - `os.Create`：创建文件。
   - `os.Open`：打开文件。
   - `os.OpenFile`：以指定模式打开文件。
   - `os.Remove`：删除文件或目录。
   - `os.Rename`：重命名文件或目录。
   - `os.Mkdir`：创建目录。
   - `os.MkdirAll`：递归创建目录。
   - `os.Chdir`：改变工作目录。
   - `os.Getwd`：获取当前工作目录。
   - `os.Stat`：获取文件或目录的信息。
2. **文件描述符**：
   - `os.Stdin`、`os.Stdout`、`os.Stderr`：标准输入、标准输出和标准错误的文件描述符。
   - `os.File`：表示文件的结构，包括文件描述符、文件名等信息。
3. **环境变量**：
   - `os.Getenv`：获取环境变量的值。
   - `os.Setenv`：设置环境变量的值。
   - `os.Unsetenv`：删除环境变量。
4. **进程控制**：
   - `os.Args`：命令行参数。
   - `os.Getpid`：获取当前进程的 PID。
   - `os.Getppid`：获取父进程的 PID。
   - `os.Exit`：退出当前进程。
5. **信号处理**：
   - `os.Signal`：表示操作系统信号的类型。
   - `os.Notify`：用于接收指定的操作系统信号。
6. **用户和组信息**：
   - `os.User`：表示用户的结构。
   - `os.LookupEnv`：根据环境变量名查找环境变量值。
   - `os.Hostname`：获取主机名。
7. **权限和文件信息**：
   - `os.Chmod`：改变文件或目录的权限。
   - `os.Chtimes`：改变文件或目录的访问和修改时间。
   - `os.Symlink`：创建符号链接。
   - `os.Readlink`：读取符号链接目标路径。
   - `os.Lstat`：获取符号链接信息，而不是目标文件信息。
8. **文件路径操作**：
   - `os.PathSeparator`：表示文件路径中的路径分隔符。
   - `os.Join`：连接路径元素以创建有效的文件路径。
   - `os.Split`：分割文件路径为目录和文件名。
   - `os.IsExist`：判断错误是否表示文件已存在。
   - `os.IsNotExist`：判断错误是否表示文件不存在。
9. **执行外部命令**：
   - `os.Exec`：执行外部命令并获取其输出。
   - `os.StartProcess`：以指定的属性启动新进程。
10. **文件锁**：
    - `os.FileLock`：文件锁结构，用于文件的读写锁定。
11. **平台相关功能**：
    - `os.Environ`：获取当前进程的所有环境变量。
    - `os.Hostname`：获取主机名。



# io

`io` 包是 Go 语言标准库中的一个核心包，提供了对输入和输出操作的抽象和通用接口。这个包包含了许多接口和类型，用于处理文件、网络连接、内存缓冲、压缩、解压缩等 I/O 相关的操作。以下是一些 `io` 包的重要接口和类型：

1. **`Reader` 接口**：

   `Reader` 接口定义了读取数据的方法，允许实现该接口的类型从不同的数据源读取数据。

   - `Read(p []byte) (n int, err error)`：从数据源读取数据并将其存储到 `p` 中，返回读取的字节数和可能的错误。

   标准库中的一些类型实现了 `Reader` 接口，例如 `os.File`、`strings.Reader` 等。

2. **`Writer` 接口**：

   `Writer` 接口定义了写入数据的方法，允许实现该接口的类型将数据写入不同的目标。

   - `Write(p []byte) (n int, err error)`：将 `p` 中的数据写入目标，返回写入的字节数和可能的错误。

   标准库中的一些类型实现了 `Writer` 接口，例如 `os.File`、`bytes.Buffer` 等。

3. **`Closer` 接口**：

   `Closer` 接口定义了关闭资源的方法，通常用于释放资源，如文件句柄、网络连接等。

   - `Close() error`：关闭资源并返回可能的错误。

   例如，`os.File` 实现了 `Closer` 接口，可以用来关闭文件。

4. **`Seeker` 接口**：

   `Seeker` 接口定义了在数据源中定位的方法，通常用于随机访问文件或其他数据源。

   - `Seek(offset int64, whence int) (int64, error)`：根据 `whence` 指定的方式定位到 `offset` 处，返回新的偏移量和可能的错误。

   例如，`os.File` 实现了 `Seeker` 接口，可以用来在文件中定位。

5. **`ReadWriter` 接口**：

   `ReadWriter` 接口组合了 `Reader` 和 `Writer` 接口，表示同时支持读和写操作的对象。

6. **`ReadCloser` 和 `WriteCloser` 接口**：

   这些接口组合了 `Reader` 或 `Writer` 接口和 `Closer` 接口，表示同时支持读或写和关闭的对象。

7. **`MultiReader` 和 `MultiWriter`**：

   这些类型实现了多个 `Reader` 或 `Writer` 的复合操作，允许多个数据源进行串联。

8. **`io.Copy` 函数**：

   `io.Copy` 函数用于将数据从一个 `Reader` 复制到一个 `Writer` 中，通常用于文件复制、网络传输等操作。

9. **`io/ioutil` 包**：

   `io/ioutil` 包提供了一些便捷的函数，用于简化文件读写、临时文件创建等操作。例如，`ioutil.ReadFile` 用于读取整个文件到内存中，`ioutil.WriteFile` 用于将数据写入文件。

10. **`io.Pipe` 类型**：

    `io.Pipe` 类型用于在不同的 goroutine 之间创建一个管道，允许一个 goroutine 写入数据，另一个 goroutine 读取数据。

11. **`io.Reader` 和 `io.Writer` 的适配器**：

    `io` 包还提供了许多适配器类型，用于将其他类型（例如字符串、字节数组）转换为 `Reader` 或 `Writer` 接口，以便进行 I/O 操作
