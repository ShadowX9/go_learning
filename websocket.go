package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type job func(in, out chan interface{})

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

// example job
func logThemAll(in, out chan interface{}) {
	for v := range in {
		fmt.Printf("%v\n", v)
		out <- v
	}
}

// example job
func persist(in, out chan interface{}) {
	for range in {
		// just for example
	}
}

type start struct {
	sock *websocket.Conn
}

func (s *start) StartJob(in, out chan interface{}) {

	for {
		_, message, err := s.sock.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		// Пишем сообщения в канал
		out <- message
	}
}

func jobWorker(job job, in, out chan interface{}) {
	defer close(out)
	job(in, out)
}

func ExecutePipeline(jobs ...job) {
	in := make(chan interface{})
	for _, j := range jobs {
		out := make(chan interface{})
		go jobWorker(j, in, out)
		in = out
	}
	fmt.Println("Завершение работы конвеера")
	time.Sleep(time.Second) // ожидание завершения всех горутин
}

var startState start

func main() {

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	ListenAndProcess := func(addr string, jobs ...job) error {

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

			// io.WriteString(w, "Call HandleFunc_onListen") - для проверки
			// WriteString не нужен, иначе ни один клиент не сможет воспользоваться этим соединением

			c, err := upgrader.Upgrade(w, r, nil)
			if err != nil {
				log.Print("upgrade:", err)
				return
			}
			startState = start{sock: c}
			defer c.Close()
			// Запускаем конвеер для обработки сообщения
			ExecutePipeline(jobs...)
		})

		log.Fatal(http.ListenAndServe(addr, nil))
		var text string
		return &errorString{text} // а вообще возвращение ошибки точно такое же есть в пакете errors
		// но по сути оно тут не нужно, поскольку log.Fatal и так уже
	}

	log.Fatal(ListenAndProcess(
		":8081",
		job(startState.StartJob),
		job(logThemAll),
		job(persist),
	))
}
