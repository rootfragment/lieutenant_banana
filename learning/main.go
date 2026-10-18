package main

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -type event -target bpfel probe probe.c -- -I.

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"os/user"
	"strconv"
	"syscall"

	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/ringbuf"
	"github.com/cilium/ebpf/rlimit"
)

func main() {
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatalf("remove memlock: %v", err)
	}

	var objs probeObjects
	if err := loadProbeObjects(&objs, nil); err != nil {
		log.Fatalf("load objects: %v", err)
	}
	defer objs.Close()

	tp, err := link.Tracepoint("sched", "sched_process_exec", objs.HandleExec, nil)
	if err != nil {
		log.Fatalf("attach tracepoint: %v", err)
	}
	defer tp.Close()

	rd, err := ringbuf.NewReader(objs.Events)
	if err != nil {
		log.Fatalf("open ringbuf: %v", err)
	}
	defer rd.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-stop
		rd.Close() // makes rd.Read return ringbuf.ErrClosed
	}()

	fmt.Println("watching for ls ... Ctrl+C to stop")

	for {
		rec, err := rd.Read()
		if err != nil {
			if errors.Is(err, ringbuf.ErrClosed) {
				return
			}
			log.Printf("read: %v", err)
			continue
		}

		var e probeEvent
		if err := binary.Read(bytes.NewReader(rec.RawSample), binary.NativeEndian, &e); err != nil {
			log.Printf("decode: %v", err)
			continue
		}

		name := strconv.Itoa(int(e.Uid))
		if u, err := user.LookupId(name); err == nil {
			name = u.Username
		}

		comm := e.Comm[:]
		if i := bytes.IndexByte(comm, 0); i >= 0 {
			comm = comm[:i]
		}
		fmt.Printf("%s (pid %d) run by %s (uid %d)\n", comm, e.Pid, name, e.Uid)
	}
}
