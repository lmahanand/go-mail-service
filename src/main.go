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

	//Execute every 2 seconds after a certain time (5 second from now)
	now := time.Now()
	at := now.Add(5 * time.Second) //In your case, this should be: Sep 1st, 2017
	s := &MsgScheduler{at, 2 * time.Second, s.Emails, s.Sender}
	c.Schedule(s, cron.FuncJob(
		func() {
			cur := time.Now()
			fmt.Printf("  [%v] CRON job executed after %v\n", cur, cur.Sub(now))
			fmt.Printf("Every minute number of emails %v\n", len(s.Emails[s.Sender]))
		}))

	fmt.Printf("Now: %v\n", now)
	c.Start()
	defer c.Stop()

	log.Println("Server started at port 8081")
	log.Fatal(http.ListenAndServe(":8081", router))

}
