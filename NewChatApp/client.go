package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

const (
	SERVER_HOST = "192.168.1.19"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

var username string
var password string
var receiverIp string

func main() {
	fmt.Println()
	fmt.Println("1. Login")
	fmt.Println("2. Signup")
	fmt.Println()
	fmt.Print("Enter the option number : ")

	var choice int
	_, err := fmt.Scanln(&choice)
	CheckError(err)

	switch choice {
	case 1:
		LoginUser()
	case 2:
		SignupUser()
	default:
		fmt.Println("Invalid Choice! Enter 1 or 2")
	}
	fmt.Println()
}

func LoginUser() {
	fmt.Println()
	fmt.Println("Login Details")
	for {
		fmt.Print("Enter your username : ")
		fmt.Scanln(&username)

		fmt.Print("Enter your password : ")
		fmt.Scanln(&password)

		if len(username) > 0 && len(password) > 0 {

			var data = []string{username, password, "Login"}

			conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
			CheckError(err)

			defer conn.Close()

			t, err := json.Marshal(data)

			finalOutput := string(t)

			_, err = conn.Write([]byte(finalOutput))

			buffer := make([]byte, 1024)
			n, err := conn.Read(buffer)
			CheckError(err)

			msg := string(buffer[:n])

			if msg == "Login Successful" {
				fmt.Println()
				fmt.Println(msg)
				clientAddr := conn.RemoteAddr().String()

				go func() { serverStart(clientAddr, conn) }()

				fmt.Print("\nConnect to : ")
				var host string
				fmt.Scanln(&host)

				conn1, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
				CheckError(err)

				msg1 := "/previouschat:" + "1234"
				conn1.Write([]byte(msg1))

				buffer1 := make([]byte, 1024)
				n1, err := conn1.Read(buffer1)

				CheckError(err)

				val := strings.TrimSpace(string(buffer1[:n1]))

				if strings.HasPrefix(val, "/allchat:") {
					splitResult := strings.SplitN(val, ":", 2)
					if len(splitResult) == 2 {
						result := strings.TrimSpace(splitResult[1])
						var data []string

						_ = json.Unmarshal([]byte(result), &data)

						for _, res := range data {
							fmt.Println(res)
						}

					} else {
						fmt.Println("Invalid input format")
					}
				}

				for {
					scanner := bufio.NewScanner(os.Stdin)
					scanner.Scan()
					mymsg := scanner.Text()

					fmt.Println(mymsg)

					currentTime := time.Now()

					timeString := currentTime.Format("2006-01-02 15:04:05")

					text := username + " : " + timeString + " " + mymsg

					SendMessage(text, host)

				}

				break

			} else {
				fmt.Println()
				fmt.Println(msg)
				continue
			}
		} else {
			fmt.Println("Fill all the details")
		}
	}
}

func SignupUser() {
	fmt.Println()
	fmt.Println("Signup Details")

	for {
		fmt.Print("Enter unique username : ")

		fmt.Scanln(&username)

		fmt.Print("Enter password : ")

		fmt.Scanln(&password)

		if len(username) > 0 && len(password) > 0 {

			var data = []string{username, password, "Signup"}

			conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
			CheckError(err)

			defer conn.Close()

			t, err := json.Marshal(data)

			finalOutput := string(t)

			_, err = conn.Write([]byte(finalOutput))
			CheckError(err)

			buffer := make([]byte, 1024)
			n, err := conn.Read(buffer)
			CheckError(err)

			msg := string(buffer[:n])

			if msg == "Account created succesfully!!!!" {
				fmt.Println()
				fmt.Println(msg)
				clientAddr := conn.RemoteAddr().String()

				go func() { serverStart(clientAddr, conn) }()

				fmt.Print("\nConnect to : ")
				var host string
				fmt.Scanln(&host)

				for {

					scanner := bufio.NewScanner(os.Stdin)
					scanner.Scan()
					mymsg := scanner.Text()

					currentTime := time.Now()

					timeString := currentTime.Format("2006-01-02 15:04:05")

					text := username + " : " + timeString + " " + mymsg

					SendMessage(text, host)

				}
			} else {
				fmt.Println()
				fmt.Println(msg)
				continue
			}
		} else {
			fmt.Println("Fill all the details")
		}
	}

}

func SendMessage(message, host string) {

	CONNECT := host
	c, err := net.Dial("tcp", CONNECT)
	CheckError(err)

	conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)

	msg := "/sendchat:" + message
	conn.Write([]byte(msg))

	c.Write([]byte(message + "\n"))

	CheckError(err)

}

func HandleNewMsg(c, connection net.Conn) {

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')

		if err != nil {
			fmt.Println(err)
			return
		}

		msg := strings.TrimSpace(string(netData)) //naveen : hey

		fmt.Println(msg)
	}
}

func serverStart(clientAddr string, connection net.Conn) {

	PORT := ":" + os.Args[1]

	l, err := net.Listen("tcp", PORT)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go HandleNewMsg(c, connection)

	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
