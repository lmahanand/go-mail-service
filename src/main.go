package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	ctrl "./controller"
	m "./model"
	s "./service"
	"github.com/gorilla/mux"
	"github.com/robfig/cron"
)

var mu sync.Mutex

//MsgScheduler ...
type MsgScheduler struct {
	At     time.Time
	Every  time.Duration
	Emails map[string][]m.Email
	Sender string
}

//Next ...
func (s *MsgScheduler) Next(t time.Time) time.Time {
	if t.After(s.At) {
		return t.Add(s.Every)
	}

	return s.At
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/email/list", ctrl.GetEmails).Methods("GET")
	router.HandleFunc("/email", ctrl.SendEmail).Methods("POST")

	c := cron.New()

	//Execute every 10 seconds after a certain time (2 second from now)
	now := time.Now()
	at := now.Add(2 * time.Second)
	s := &MsgScheduler{at, 10 * time.Second, s.Emails, s.Sender}
	c.Schedule(s, cron.FuncJob(
		func() {
			messageScheduler()
		}))

	c.Start()
	defer c.Stop()

	log.Println("Server started at port 8081")
	log.Fatal(http.ListenAndServe(":8081", router))

}

func messageScheduler() {
	t := time.Now().UTC()
	fmt.Printf("UTC time is %v\n", t)
	es := s.EmailService{}
	es.SendScheduledEmails()
}
