package example0

type HelloService struct {}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = request + " world!"
	return nil
}