package file

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

type File struct {
	Path string
}

func (f *File) Read() ([]service.Ticket, error) {
	data, err := os.ReadFile(f.Path)
	if err != nil {
		return nil, err
	}
	dataSplited := strings.Split(string(data), "\n")
	//csvFile := csv.NewReader(bytes.NewBuffer(data))
	ListTickets := []service.Ticket{}
	for _, line := range dataSplited {
		//line, err := csvFile.Read()
		if err != nil {
			break
		}
		EachLine := strings.Split(line, ",")

		id, err := strconv.Atoi(EachLine[0])
		if err != nil {
			return nil, err
		}
		price, err := strconv.Atoi(EachLine[5])
		if err != nil {
			return nil, err
		}
		ticket := service.Ticket{
			Id:          id,
			Names:       EachLine[1],
			Email:       EachLine[2],
			Destination: EachLine[3],
			Date:        EachLine[4],
			Price:       price,
		}
		ListTickets = append(ListTickets, ticket)
	}
	return ListTickets, nil
}

func (f *File) Write(s *service.Ticket) error {
	newTicket := fmt.Sprintf("\n%d,%s,%s,%s,%s,%d", s.Id, s.Names, s.Email, s.Destination, s.Date, s.Price)
	write, err := os.OpenFile(f.Path, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	_, err2 := write.Write([]byte(newTicket))
	if err2 != nil {
		return err2
	}
	write.Close()
	return nil
}
