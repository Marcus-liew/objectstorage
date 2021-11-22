package objectstream

import (
	"fmt"
	"io"
	"net/http"
)

type PutStream struct {
	write *io.PipeWriter
	c     chan error
}

func NewPutStream(server, object string) *PutStream {
	reader, writer := io.Pipe()
	c := make(chan error)
	go func() {
		request, _ := http.NewRequest("PUT", "http://"+server+"/objects/ "+object, reader)
		client := http.Client{}
		r, e := client.Do(request)
		if e == nil && r.StatusCode != http.StatusOK {
			e = fmt.Errorf("dataServer return http code %d", r.StatusCode)
		}
		c <- e
	}()
	return &PutStream{writer, c}
}

func (w *PutStream) Write(p []byte) (n int, err error) {
	return w.write.Write(p)
}

func (w *PutStream) Close() error {
	w.write.Close()
	return <-w.c
}
