package server

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	// used to validate secret expiration and calculate totalWatchTime
	serverDate     time.Time
	timerStartTime time.Time
	totalWatchTime time.Duration
	dailyAllotment time.Duration
}

func NewServer() *Server {
	return &Server{serverDate: time.Now(), timerStartTime: time.Time{}, totalWatchTime: 0, dailyAllotment: (2 * time.Hour)}
}

func (s *Server) GetCurrentWatchTime(w http.ResponseWriter, req *http.Request) {
	msg := fmt.Sprintf("today youve watched for %v\n", s.totalWatchTime)
	w.Write([]byte(msg))
}

func (s *Server) GetRemainingWatchTime(w http.ResponseWriter, req *http.Request) {
	msg := fmt.Sprintf("today you have %v remaining\n", (s.dailyAllotment - s.totalWatchTime))
	w.Write([]byte(msg))
}

func (s *Server) StartTimer(w http.ResponseWriter, req *http.Request) {
	if s.timerStartTime.IsZero() {
		s.timerStartTime = time.Now()
	}
}

func (s *Server) EndTimer(w http.ResponseWriter, req *http.Request) {
	if !s.timerStartTime.IsZero() {
		tempEnd := time.Now()
		watchDuration := tempEnd.Sub(s.timerStartTime)
		s.totalWatchTime += watchDuration
		s.timerStartTime = time.Time{}
		http.Redirect(w, req, "/api/timer/current", http.StatusSeeOther)
	} else {
		msg := ("hmmm you shouldnt be watching rn w/o timer?\n")
		w.Write([]byte(msg))
	}
}

func (s *Server) AddBonusAllotment(w http.ResponseWriter, req *http.Request) {
	body, _ := (io.ReadAll(req.Body))
	defer req.Body.Close()

	bonus, err := strconv.Atoi(string(body))
	if err != nil {
		w.Write([]byte("bonus time was not added. only numbers accepted. please try again\n"))
	}

	s.dailyAllotment += time.Duration(bonus)
}
