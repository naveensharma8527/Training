package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

const (
	SERVER_HOST = "192.168.1.19"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

var users []string

type Chats struct {
	Chat []string
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	CheckError(err)

	userInfo := strings.TrimSpace(string(buffer[:n]))

	if strings.HasPrefix(userInfo, "/previouschat:") {
		prevMessages(conn)
	} else if strings.HasPrefix(userInfo, "/sendchat:") {
		splitResult := strings.SplitN(userInfo, ":", 2)
		if len(splitResult) == 2 {
			result := strings.TrimSpace(splitResult[1])
			manageChat(result)

		} else {
			fmt.Println("Invalid input format")
		}
	} else {
		var temp []string
		_ = json.Unmarshal([]byte(userInfo), &temp)

		fmt.Println(temp)

		file, err := os.Open("./userData.txt")
		CheckError(err)
		defer file.Close()

		scan := bufio.NewScanner(file)
		scan.Scan()
		line := scan.Text()

		var data = make(map[string]string)

		_ = json.Unmarshal([]byte(line), &data)

		_, ok := data[temp[0]]

		var message string

		if temp[2] == "Signup" {
			if ok {
				message = "Username already exist, try to create with another username."
				conn.Write([]byte(message))
			} else {
				data[temp[0]] = temp[1]

				t, err := json.Marshal(data)
				CheckError(err)

				finalOut := string(t)

				myfile, err := os.OpenFile("./userData.txt", os.O_CREATE|os.O_WRONLY, 0644)
				CheckError(err)

				_, err = io.WriteString(myfile, finalOut)
				CheckError(err)

				message = "Account created succesfully!!!!"
				conn.Write([]byte(message))

				myfile.Close()
				users = append(users, temp[0])
			}
		} else {
			if ok && temp[1] == data[temp[0]] {
				message = "Login Successful"
				conn.Write([]byte(message))
				users = append(users, temp[0])
			} else {
				message = "Invalid Credential"
				conn.Write([]byte(message))
			}
		}
	}

}

func manageChat(message string) {
	file, err := os.Open("./chatdata.txt")
	CheckError(err)
	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Scan()
	line := scan.Text()

	var data = make(map[string]Chats)

	_ = json.Unmarshal([]byte(line), &data)

	opt1 := users[0] + ":" + users[1]
	opt2 := users[1] + ":" + users[0]

	if len(data) == 0 {
		data[opt1] = Chats{Chat: []string{message}}
	} else {
		_, ok1 := data[opt1]
		_, ok2 := data[opt2]

		if ok1 {
			mydata := data[opt1]
			mydata.Chat = append(mydata.Chat, message)
			data[opt1] = mydata
		} else if ok2 {
			mydata := data[opt2]
			mydata.Chat = append(mydata.Chat, message)
			data[opt2] = mydata
		} else {
			mydata := data[opt1]
			mydata.Chat = append(mydata.Chat, message)
			data[opt1] = mydata
		}
	}
	t, err := json.Marshal(data)
	CheckError(err)

	finalOut := string(t)

	myfile, err := os.OpenFile("./chatdata.txt", os.O_CREATE|os.O_WRONLY, 0644)
	CheckError(err)

	_, err = io.WriteString(myfile, finalOut)
	CheckError(err)

	myfile.Close()

}

func main() {
	listener, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	CheckError(err)
	defer listener.Close()

	fmt.Println("Server started. Waiting for clients...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		fmt.Println()
		go handleConnection(conn)
	}

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func prevMessages(conn net.Conn) {

	file, err := os.Open("./chatdata.txt")

	CheckError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	var data = make(map[string]Chats)
	_ = json.Unmarshal([]byte(line), &data)

	opt1 := users[0] + ":" + users[1]
	opt2 := users[1] + ":" + users[0]

	_, ok1 := data[opt1]
	_, ok2 := data[opt2]

	var arr []string

	if ok1 {
		for _, val := range data[opt1].Chat {
			arr = append(arr, val)
		}
	} else if ok2 {
		for _, val := range data[opt2].Chat {
			arr = append(arr, val)
		}
	}

	t, err := json.Marshal(arr)
	finalOutput := string(t)

	finalOutput = "/allchat:" + finalOutput

	fmt.Println("Final ", finalOutput)

	_, err = conn.Write([]byte(finalOutput))
	CheckError(err)

}
