package dialog

import (
	"google.golang.org/grpc"

	dto "monolith/internal/generated/rpc/clients/dialog"
)

type Client struct {
	cli dto.ServiceDialogClient
}

func New(url string) *Client {
	conn, _ := grpc.Dial(url, grpc.WithInsecure())
	return &Client{
		cli: dto.NewServiceDialogClient(conn),
	}
}
