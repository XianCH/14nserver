package log

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Log struct {
	EntriesNum int
	Writer     io.Writer
	Interval   time.Duration

	i       int
	ch      chan entry
	entries []entry
	t       *time.Ticker
}

type entry struct {
	r    http.Request
	code string
	msg  string
	time time.Time
}

func (l *Log) InitLog() {
	l.t = time.NewTicker(l.Interval)
	l.ch = make(chan entry, 2*l.EntriesNum)
	l.entries = make([]entry, l.EntriesNum)
}

func (l *Log) Close() {
	l.t.Stop()
	close(l.ch)
}

func (l *Log) Log(req http.Request, code string, msg string) {
	l.ch <- entry{req, code, msg, time.Now()}
}

func (l *Log) Loop() {
	for {
		select {
		case e := <-l.ch:
			l.entries[l.i] = e
			l.i++
			if l.i == l.EntriesNum {
				l.flush()
			}
		case _ = <-l.t.C:
			l.flush()
		}
	}
}

func (l *Log) flush() {
	if l.i == 0 {
		return
	}

	var w bytes.Buffer

	for i := 0; i < l.i; i++ {
		e := l.entries[i]
		fmt.Fprintf(&w, "%s %s %s %s %s %s  \"%s\"\n",
			e.time.Format(time.RFC3339),
			or(e.r.Method, "-"),
			or(e.r.URL.Path, "-"),
			or(e.code, "-"),
			or(e.msg, "-"),
			or(e.r.Header.Get("Referer"), "-"),
			or(e.r.Header.Get("User-Agent"), "-"),
		)
	}

	l.Writer.Write(w.Bytes())
	l.i = 0
}

func or(in, or string) string {
	if in == "" {
		return or
	} else {
		return in
	}
}
