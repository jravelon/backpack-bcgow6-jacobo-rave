package main

import (
	"fmt"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

func main() {
	file := file.File{"tickets.csv"}
	tickets, _ := file.Read()
	booking := service.NewBookings(tickets)
	t, _ := booking.Read(2)
	booking.Delete(2)
	booking.Delete(3)
	_, err := booking.Read(2)
	fmt.Println(t)
	fmt.Println(err)
}
