package main

import (
	"fmt"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Print("Ocurrio un error: ", err)
		}
	}()

	var ticket []service.Ticket

	file := file.File{"tickets.csv"}
	data, err := file.Read()
	if err != nil {
		panic(err)
	}

	ticket = append(ticket, data...)

	fmt.Print("Cargando tickets")
	fmt.Print("......\n")

	booking := service.NewBookings(ticket)
	fmt.Printf("hay %d tickets", len(ticket))
	fmt.Print("......\n")

	newTicket := service.Ticket{
		Id:          1222,
		Names:       "Joao da silva",
		Email:       "Joao@email.com",
		Destination: "Seul",
		Date:        "12:3",
		Price:       1000,
	}
	a, err := booking.Create(newTicket)
	if err != nil {
		panic(err)
	}
	fmt.Print("Nuevo ticket guardado ")
	fmt.Print(a, "......\n")

	t, err := booking.Read(2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ticket leido %v", t)
	fmt.Print("......\n")

	err = file.Write(&newTicket)
	if err != nil {
		panic(err)
	}
	fmt.Print("cambios guardados.........\n")

	b, err := booking.Read(1222)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ticket leido %v", b)
	fmt.Print("......\n")

}
