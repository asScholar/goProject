/*
第一节，Go语言简介
Go语言是Google在2007年开发的一种开源编程语言，于2009年11月10日向全球公开。Go是非常年轻的一门语言，它的主要目标是“兼具Python等动态语言的开发速度和C/C++等编译型语言的性能于安全性”。
Go语言不但能让程序员访问底层操作系统，它还提供了强大的网络编程和并发编程支持。Go语言的用途众多，可以进行网络编程，系统编程，开发编程，分布式编程。

第二节，Go语言为并发而生
如今，众核CPU已经成为服务器的标配，挖掘众核运算能力是由程序员人工设计算法和框架来完成的。虽然一些编程语言不断地提高多核资源使用效率，例如Java的Netty等，但仍然需要开发人员花费大量的时间和精力才能搞懂这些框架。而Go语言在众核并发上拥有原生的设计优势。
Go语言的并发是基于goroutine的，goroutine类似于线程，但并非线程。可以将goroutine理解为一种虚拟线程。Go语言运行时会参与调度goroutine，并将goroutine合理地分配到每个CPU，最大限度地使用CPU性能。
多个goroutine中，Go语言使用通道（channel）进行通信，程序可以将需要并发的环节设计为生产者和消费者的模式，将数据放入通道。通道的另外一端的代码将这些数据进行并发计算并返回结果。
下面代码展示了，生产者每秒生成一个字符串，然后通过通道传给消费者。生产者使用两个goroutine并发运行，消费者在main()函数的gorountine中进行处理。
整段代码没有线程创建，没有线程池也没有加锁，仅仅通过关键字go实现goroutine，和通道实现数据交换。
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*数据生产者*/
func producer(header string, channel chan<- string) {
	/*无限循环，不停地生产数据*/
	for {
		/*将随机数和字符串格式化为字符串发送给通道*/
		channel <- fmt.Sprintf("%s:%v", header, rand.Int31())

		/*等待1秒*/
		time.Sleep(time.Second)
	}
}

/*数据消费者*/
func customer(channel <-chan string) {
	/*不停地获取数据*/
	for {
		/*从通道中读取数据，此处会阻塞指导通道中返回数据*/
		message := <-channel

		/*打印数据*/
		fmt.Println(message)
	}
}

func main() {
	/*创建一个字符串类型的通道*/
	channel := make(chan string)

	/*创建producer()函数的并发goroutine*/
	go producer("cat", channel)
	go producer("dog", channel)

	/*数据消费者函数*/
	customer(channel)
}

/*
第三节，哪些项目使用Go语言开发
Docker，
Kubernetes，
etcd等

第四节，Go语言性能如何
Go语言在性能上更接近于Java语言。Go语言在未来的版本中会通过不断地版本优化提高单核运行性能。

第五节，Go语言标准库强大
Go语言标准库覆盖网络，系统，加密，编码，图形等各个方面，可以直接使用标准的http包进行HTTP协议的收发处理;网络库基于高性能的操作系统通信模型（Linux的epoll，Windows的IOCP ）;所有的加密，编码都内建支持，不需要再从第三方开发者处获取。

第六节，Go语言上手简单
Go语言到底有多么简单？下面从实现一个HTTP服务器开始了解。
HTTP文件服务器是常见的Web服务器之一，开发阶段为了测试，需要自行安装Apache或Nginx服务器，下载安装配置需要大量的时间，使用Go语言实现一个简单的HTTP服务器只需要几行代码，如下所示。
*/

/*
package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", nil)
}
*/
/*
执行上述代码，然后在浏览器中输入http://127.0.0.1:808即可浏览文件，这些文件正是当前目录在HTTP服务器上的映射目录。
Go语言工程结构简单。
Go语言编译速度快。

第七节，Go语言代码风格清晰，简单
1, for循环可以不加括号
2, if表达式可以不加括号
等
*/
