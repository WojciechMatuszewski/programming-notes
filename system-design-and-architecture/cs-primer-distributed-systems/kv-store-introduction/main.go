package main

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
)

const PORT = 8000
const IP = "127.0.0.1"

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	listener := NewUDPListener(IP, PORT)
	err := listener.Listen(ctx)
	if err != nil {
		panic(err)
	}

	store := NewStoreCommandHandler()

	buf := make([]byte, 1024)
	done := make(chan struct{})

	go func() {
		defer close(done)

		for {
			data, addr, err := listener.Read(buf)
			if err != nil {
				if ctx.Err() != nil {
					return
				}

				panic(err)
			}

			command, err := ParseCommand(data)
			if err != nil {
				panic(err)
			}

			resp, err := command.Execute(store)
			if err != nil {
				panic(err)
			}

			err = listener.Send([]byte(resp.payload), addr)
			if err != nil {
				panic(err)
			}
		}

	}()

	<-ctx.Done()
	listener.Close()
	<-done
}

type UDPListener struct {
	address *net.UDPAddr
	conn    net.PacketConn
}

func NewUDPListener(ip string, port int) UDPListener {
	address := &net.UDPAddr{IP: net.ParseIP(ip), Port: port}

	return UDPListener{
		address: address,
		conn:    nil,
	}
}

func (l *UDPListener) Listen(ctx context.Context) error {
	cfg := net.ListenConfig{}

	conn, err := cfg.ListenPacket(ctx, "udp", l.address.String())
	if err != nil {
		return err
	}

	l.conn = conn
	return nil
}

func (l *UDPListener) Close() error {
	if l.conn != nil {
		return l.conn.Close()
	}

	return nil
}

func (l *UDPListener) Read(buf []byte) ([]byte, net.Addr, error) {
	n, addr, err := l.conn.ReadFrom(buf)
	return buf[:n], addr, err
}

func (l *UDPListener) Send(buf []byte, addr net.Addr) error {
	_, err := l.conn.WriteTo(buf, addr)
	return err
}

type Response struct {
	payload string
}

type CommandVerb string

const (
	GetCommandVerb CommandVerb = "get"
	SetCommandVerb CommandVerb = "set"
)

type CommandHandler interface {
	HandleGet(command GetCommand) (Response, error)
	HandleSet(command SetCommand) (Response, error)
}

type StoreCommandHandler struct {
	store map[string]string
}

func NewStoreCommandHandler() *StoreCommandHandler {
	return &StoreCommandHandler{store: map[string]string{}}
}

func (h *StoreCommandHandler) HandleGet(command GetCommand) (Response, error) {
	payload := h.store[command.key]
	return Response{payload: payload}, nil
}

func (h *StoreCommandHandler) HandleSet(command SetCommand) (Response, error) {
	h.store[command.key] = command.payload
	return Response{payload: "OK"}, nil
}

type Command interface {
	Execute(handler CommandHandler) (Response, error)
}

type GetCommand struct {
	key string
}

func (gc GetCommand) Execute(handler CommandHandler) (Response, error) {
	// Double dispatch pattern.
	// This allows us to get rid of the type-switch, because each command knows how to dispatch itself.

	// It's called "double-dispatch" because we rely on two dispatches happening.
	// One at the caller level: `command.Execute`
	// One here, where we choose which method on the handler to call.
	return handler.HandleGet(gc)
}

func ParseGetCommand(args []byte) (GetCommand, error) {
	return GetCommand{key: string(args)}, nil
}

type SetCommand struct {
	key     string
	payload string
}

func (sc SetCommand) Execute(handler CommandHandler) (Response, error) {
	return handler.HandleSet(sc)
}

func ParseSetCommand(args []byte) (SetCommand, error) {
	key, payload, found := bytes.Cut(args, []byte{' '})
	if !found {
		return SetCommand{}, fmt.Errorf("malformed SET command")
	}

	return SetCommand{key: string(key), payload: string(payload)}, nil
}

func ParseCommand(b []byte) (Command, error) {
	if len(b) < 3 {
		return nil, fmt.Errorf("malformed command")
	}

	verb := CommandVerb(string(b[:3]))
	args := b[3:]

	switch verb {
	case GetCommandVerb:
		return ParseGetCommand(args)
	case SetCommandVerb:
		return ParseSetCommand(args)
	default:
		return nil, fmt.Errorf("unknown command verb: %s", verb)
	}
}
