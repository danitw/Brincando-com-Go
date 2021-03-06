package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/exec"
	"time"
)

type Response struct {
	TradeID int
	Price   string
	Size    string
	Bid     string
	Ask     string
	Volume  string
	Time    string
}

var mux map[string]func(http.ResponseWriter, *http.Request)

type myHandler struct{}

func main() {
	d()
}

// Sem criatividade pra criar um nome pra essa função
func semNome() {
	text := inputKey()
	a := validInput(text)

	if a == true {
		panic("Empty String")
	}

	arr := getArray()

	travelArray(arr)
}

func inputKey() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	return text
}

func validInput(input string) bool {
	if input == " " {
		fmt.Println("empty string")
		return true
	}

	return false
}

func scopeLocal() string {

	local := "Sou uma variavel local então só existo nesse bloco de codigo\n"

	return local
}

func getArray() map[int]string {
	arr := make(map[int]string)

	arr[0] = scopeLocal()
	arr[1] = inputKey()

	return arr
}

func travelArray(arr map[int]string) {

	count := 1

	for i := 0; i <= count; i++ {
		fmt.Println(arr[i])
	}
}

func oneUpToTen() {
	i := 1
	max := 10

	for i <= max {
		fmt.Println(i)
		i++
	}
}

func fileExists() {
	if _, err := os.Stat("senha.txt"); err == nil {
		fmt.Printf("File exists\n")
	} else {
		fmt.Printf("File does not exist\n")
	}
}

func readFile(path string) string {
	b, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println(err)
	}

	str := string(b)

	return str
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeFile(path string, words string) error {
	file, err := os.Create(path)

	if err != nil {
		return err
	}

	defer file.Close()

	_, _ = file.WriteString(words)

	return nil
}

func renameFile(path string, newPath string) error {
	err := os.Rename(path, newPath)

	if err != nil {
		return err
	}

	return nil
}

func numbersInFull() {
	numbers := make(map[int]string)

	numbers[1] = "one"
	numbers[2] = "two"
	numbers[3] = "three"
	numbers[4] = "four"
	numbers[5] = "five"

	fmt.Println(numbers)
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func pointers() {
	x := 5

	var pointer *int

	pointer = &x

	fmt.Printf("Memory address of variable x: %x\n", &x)
	fmt.Printf("Memory address stored in ipointer variable: %x\n", pointer)
	fmt.Printf("Contents of *ipointer variable: %d\n", *pointer)
}

func splitHelloWorld() map[string]string {
	split := make(map[string]string)

	text := "hello world"

	split["h"] = text[0:5]
	split["w"] = text[6:11]

	return split
}

func hi(name string) {
	defer fmt.Print(name)
	fmt.Print("Hi ")
}

func sum(numbers ...int) {
	fmt.Println(numbers)
}

func getContent() Response {
	// json data
	url := "https://api.gdax.com/products/BTC-EUR/ticker"

	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	var data Response

	err2 := json.Unmarshal(body, &data)

	if err2 != nil {
		panic(err2.Error())
	}

	return data
}

func speak(text string) {
	cmd := exec.Command("espeak", text)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func socketServer() {
	fmt.Println("Start server...")

	// listen on port 8000
	ln, _ := net.Listen("tcp", ":8000")

	// accept connection
	conn, _ := ln.Accept()

	// run loop forever (or until ctrl-c)
	for {
		// get message, output
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received:", string(message))
	}
}

func socketClient() {
	// connect to server
	conn, _ := net.Dial("tcp", "127.0.0.1:8000")
	for {
		// what to send?
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// send to server
		_, _ = fmt.Fprintf(conn, text+"\n")
		// wait for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Implement route forwarding
	if h, ok := mux[r.URL.String()]; ok {
		//Implement route forwarding with this handler, the corresponding route calls the corresponding func.
		h(w, r)
		return
	}
	_, _ = io.WriteString(w, "URL: "+r.URL.String())
}

func Tmp(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "version 3")
}

func d() {
	server := http.Server{
		Addr:        ":8080",
		Handler:     &myHandler{},
		ReadTimeout: 5 * time.Second,
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/tmp"] = Tmp
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
