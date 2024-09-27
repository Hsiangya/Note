package main

import (
	"chatRoom/internal/models"
	"fmt"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("listenError", err)
		return
	}

	go broadcast() // 全局广播

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("acceptError", err)
			continue
		}
		go handler(conn)
	}
}

func broadcast() {
	defer fmt.Println("broadcast end")
	for {
		info := <-models.Message
		for _, user := range models.AllUser {
			user.Msg <- info
		}
	}
}

func handler(conn net.Conn) {
	clientAddr := conn.RemoteAddr().String()
	msgChan := make(chan string, 10)
	newUser := models.NewUser(clientAddr, clientAddr, msgChan)
	models.AllUser[newUser.Id] = newUser // TODO map不支持并发读写 这里会有问题 需要修改

	var isQuit = make(chan bool)
	var restTimer = make(chan bool)
	go watch(newUser, conn, isQuit, restTimer)
	go writeBackToClient(newUser, conn)
	loginInfo := fmt.Sprintf("[%s]:[%s] login\n", newUser.Id, newUser.Name)
	fmt.Println(loginInfo)
	models.Message <- loginInfo

	for {
		buf := make([]byte, 1024)
		cnt, err := conn.Read(buf)
		if cnt == 0 || err != nil {
			fmt.Println("用户主动退出")
			isQuit <- true
			return
		}
		userInput := string(buf[:cnt])
		models.Message <- userInput
		restTimer <- true
	}
}

func writeBackToClient(user *models.User, conn net.Conn) {
	for data := range user.Msg {
		_, _ = conn.Write([]byte(data))
	}
}

func watch(user *models.User, conn net.Conn, isQuit, restTimer <-chan bool) {
	for {
		select {
		case <-isQuit:
			delete(models.AllUser, user.Id)
			logout := fmt.Sprintf("%s logout\n", user.Name)
			models.Message <- logout
			fmt.Println(logout)
			_ = conn.Close()
			return
		case <-time.After(3 * time.Second):
			delete(models.AllUser, user.Id)
			logout := fmt.Sprintf("%s timeout logout \n", user.Name)
			models.Message <- logout
			fmt.Println(logout)
			_ = conn.Close()
			return
		case <-restTimer:
			fmt.Println(user.Name, "写入了数据")
		}
	}
}
