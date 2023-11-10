package butler

// Tools is the interface for the RPC service.
type Tools int

func (t *Tools) Greet(_ struct{}, reply *string) error {
	*reply = "Welcome, master!"
	return nil
}
