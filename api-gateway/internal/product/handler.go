package product

type Handler struct {
	client *Client
}

func NewHandler(client *Client) *Handler {
	return &Handler{
		client: client,
	}
}
