package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const addr = ":8081"

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	tcpServer := NewServer(addr)

	// tcp server
	g.Go(func() error {
		return tcpServer.Run(ctx, g)
	})

	// signal
	g.Go(func() error {
		exitSignals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT} // SIGTERM is POSIX specific
		sig := make(chan os.Signal, len(exitSignals))
		signal.Notify(sig, exitSignals...)
		for {
			fmt.Println("signal")
			select {
			case <-ctx.Done():
				fmt.Println("signal ctx done")
				return ctx.Err()
			case <-sig:
				// do something
				fmt.Println("get stop sianel")
				return errors.New("receive signal stop")
			}
		}
	})

	g.Go(func() error {
		<-ctx.Done()
		fmt.Println("get ctx done")
		return nil
	})

	//print server connect number
	g.Go(func() error {
		t := time.NewTicker(time.Second * 10)
		for {
			select {
			case <-ctx.Done():
				return nil
			case <-t.C:
				log.Printf("current connect number:%d", connectNumber)
			}
		}
	})
	log.Println("server start success,listen at:", addr)
	err := g.Wait() // first error return
	fmt.Println(err)
}

var connectNumber int64 //链接数

//tcp server
type Server struct {
	addr   string
	listen net.Listener
}

func NewServer(addr string) *Server {
	return &Server{addr: addr}
}

func (s *Server) Run(ctx context.Context, g *errgroup.Group) (err error) {

	s.listen, err = net.Listen("tcp", s.addr)
	if err != nil {
		return
	}

	s.Listen(ctx, g)

	return
}

func (s *Server) Listen(ctx context.Context, g *errgroup.Group) {

	//listen stop
	g.Go(func() error {
		<-ctx.Done()
		if err := s.listen.Close(); err != nil {
			return err
		}
		log.Println("end 01")
		return nil
	})

	for {
		conn, err := s.listen.Accept()
		if err != nil {
			log.Printf("connect fail:%v", err)
			//当前tcplisten 已关闭//退出监听
			if strings.Contains(err.Error(), "use of closed network connection") {
				log.Println("end 02")
				log.Println("listen return")
				return
			}
			continue
		}
		c := NewChannel(conn)
		atomic.AddInt64(&connectNumber, 1)
		//接收
		g.Go(func() error {
			log.Printf("receive init,id:%d", c.id)
			c.AccepMessage()
			log.Println("end 03")
			return nil
		})

		//消费
		g.Go(func() error {
			log.Printf("comsume init,id:%d", c.id)
			c.ComsumeMessage()
			log.Println("end 04")
			return nil
		})

		//stop信号处理
		g.Go(func() error {
			log.Printf("stop init, id:%d", c.id)
			<-ctx.Done()
			c.Stop()
			log.Println("end 05")
			return nil
		})

		select {
		case <-ctx.Done():
			fmt.Println("tcp listen routine stoped")
			return
		default:
			continue
		}
	}
}

//当前管道状态
type State int

const (
	NORMAL   State = 0 //0正常
	NEEDSTOP State = 1 // 接受关闭信号，拒绝接收新消息，消费完已存在的消息后关闭
	STOP     State = 2 // 已关闭

	MsgNumber = 10 //管道大小
)

//message channel
type Channel struct {
	id    int // channel id 随机数
	conn  net.Conn
	msgs  chan *Message
	mu    sync.Mutex
	state State //当前状态
}

func NewChannel(conn net.Conn) *Channel {
	id := rand.Intn(1000)
	return &Channel{id: id, conn: conn, msgs: make(chan *Message, MsgNumber)}
}

//accept message
func (c *Channel) AccepMessage() {
	var buf [4096]byte
	for {
		if c.state != NORMAL {
			log.Printf("id:%d,accpeMessage stop", c.id)
			return
		}
		n, err := c.conn.Read(buf[:])
		if err != nil {
			log.Printf("id:%d, read from %s msg faild err:[%v]\n", c.id, c.conn.RemoteAddr().String(), err)
			c.Stop()
			break
		}
		log.Printf("id:%d, rev data from %s msg:%s\n", c.id, c.conn.RemoteAddr().String(), string(buf[:n]))
		m := NewMessage(string(buf[:n]))
		c.msgs <- m
	}
}

//comsume message
func (c *Channel) ComsumeMessage() {
	for {
		msg, ok := <-c.msgs
		if !ok {
			log.Printf("id:%d, comsume message stop", c.id)
			c.ComusumeFish()
			return
		}
		//TODO
		//comsume message
		log.Printf("id:%d, comsume message conntent:%v", c.id, msg.content)
	}
}

//all message has comusume and channel has close
func (c *Channel) ComusumeFish() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.state = STOP
}

//Need Stop
func (c *Channel) Stop() (err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.state != NORMAL {
		return nil
	}
	c.state = NEEDSTOP
	err = c.conn.Close()
	close(c.msgs)
	atomic.AddInt64(&connectNumber, -1)
	return
}

//message
type Message struct {
	content string
}

func NewMessage(content string) *Message {
	return &Message{content: content}
}
