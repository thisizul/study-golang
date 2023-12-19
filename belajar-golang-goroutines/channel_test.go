package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	chanel := make(chan string)
	defer close(chanel)
	go func() {
		time.Sleep(2 * time.Second)
		chanel <- "Jhon Smith Lance"
		fmt.Println("Selesai Mengirim Data ke Chanel")
	}()
	data := <-chanel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func GiveMeRespond(chanel chan string) {
	time.Sleep(2 * time.Second)
	chanel <- "Give me Respond Jhon Smith Lance"
}

func TestChannelAsParameter(t *testing.T) {
	chanel := make(chan string)
	defer close(chanel)
	go GiveMeRespond(chanel)
	data := <-chanel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}