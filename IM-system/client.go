package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

//定义一个客户端的type
type Client struct {
	ClientIP   string
	ClientPort int
	ClientName string
	conn       net.Conn
	flag       int
}

//创建一个定义客户端的func
func NewClient(clientip string, clientport int) *Client {
	newclient := &Client{
		ClientIP:   clientip,
		ClientPort: clientport,
		flag:       999,
	}
	//将客户端进行连接
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", clientip, clientport))
	if err != nil {
		fmt.Println("net.dial error:", err)
		return nil
	}
	newclient.conn = conn
	//返回对象
	return newclient
}

//建立一个func来接收server回应的消息，直接显示标准输出 不是很明白
func (newclient *Client) DealResponse() {
	//一旦clien.conn有数据，就直接copy到stdout标准输出上，永久阻塞监听
	io.Copy(os.Stdout, newclient.conn)
}

//将收到的内容封装到一个updatename的函数中
func (newclient *Client) UpdateName() bool {

	// 从用户对话框里面接收用户名信息fmt.scanln直接存入newclient的Name里面
	fmt.Println(">>>>请输入用户名:")
	fmt.Scanln(&newclient.ClientName)

	//将接收到的信息封装成发给server的信息，并且发给server
	sendMsg := "rename|" + newclient.ClientName + "\n"
	_, err := newclient.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return false
	}
	return true

}

//公聊模式
func (newclient *Client) PublicChat() {
	//获取用户输入的内容
	var chatMsg string

	fmt.Println(">>>>Please input your content, or use \"exit\" to quit.")
	fmt.Scanln(&chatMsg)
	//循环发给客户端
	for chatMsg != "exit" {
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := newclient.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn Write err:", err)
				break
			}
		}
		chatMsg = ""
		fmt.Println(">>>>Please input your content, or use \"exit\" to quit.")
		fmt.Scanln(&chatMsg)
	}

}

//先确认有谁在线 server就是需要发送who就可以
func (newclient *Client) SelectUsers() {
	sendMsg := "who"
	_, err := newclient.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn Write err:", err)
		return
	}

}

//私聊模式
func (newclient *Client) PrivateChat() {
	var chatName string
	var chatMsg string

	newclient.SelectUsers()
	//从交互界面获取想要聊天的用户名，存到chatName里面去
	fmt.Println(">>>>input username to chat, or exit")
	fmt.Scanln(&chatName)

	if chatName != "exit" {
		//从交互界面读取信息
		fmt.Println(">>>>input chat content, or exit")
		fmt.Scanln(&chatMsg)
		for chatMsg != "exit" {
			//将消息发送给server
			if len(chatMsg) != 0 {
				sendMsg := "to|" + chatName + "|" + chatMsg + "\n"
				_, err := newclient.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("conn Write err:", err)
					break
				}
			}
			chatMsg = ""
			fmt.Println(">>>>input chat content, or exit")
			fmt.Scanln(&chatMsg)
		}
		newclient.SelectUsers()
		fmt.Println(">>>>请输入聊天对象[用户名], exit退出:")
		fmt.Scanln(&chatName)

	}

}

//建立用户的操作界面
func (newclient *Client) Manu() bool {
	var flag int
	fmt.Println("1.Public Talking")
	fmt.Println("2.Private Talking")
	fmt.Println("3.Rename")
	fmt.Println("0.Exit")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		newclient.flag = flag
		return true
	} else {
		fmt.Println(">>>>Please input number from 0 to 3<<<<")
		return false
	}
}

//建立一个run函数来判断输入的数字在不在合理范围内
func (newclient *Client) Run() {
	//只要他不退出就一直循环
	for newclient.flag != 0 {
		//对于用户输入的数字不在0-3之间
		for newclient.Manu() != true {

		}
		//对于用户输入的数字在0-3之间，根据用户输入的数据选择业务，具体的业务又封装到其他的函数里面
		switch newclient.flag {
		case 1:
			// 1.Public Talking
			newclient.PublicChat()
			break
		case 2:
			// 2.Private Talking
			newclient.PrivateChat()
			break
		case 3:
			// 3.Rename
			newclient.UpdateName()
			break
		}

	}
}

//建立两个全局函数
var serverIP string
var serverPort int

//在init函数里面先进行命令行解析
//因为是一个进程的初始化工作所以就放在init函数中
func init() {
	flag.StringVar(&serverIP, "ip", "127.0.0.1", "Setting server ip address(default:127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "Setting server port(default:8888)")
}

//main func里面进行操作
func main() {
	//命令行解析
	flag.Parse()
	client := NewClient(serverIP, serverPort)
	if client == nil {
		fmt.Println(">>>>>>Fail connection")
		return
	}

	//只要连接成功了就开始接收server发的内容
	//单独用一个goroutine去处理server的回执消息
	go client.DealResponse()

	fmt.Println(">>>>>>Successful Connection")

	//启动客户端的业务
	client.Run()

	//阻塞管道
	select {}
}
