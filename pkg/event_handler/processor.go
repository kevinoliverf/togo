package eventhandler

import (
	"context"
	"log"
	"sync"
	"time"
)

// EventProcessor is responsible for fetching and processing events
type EventProcessor struct {
	*EventRegistry
	queue          chan Event
	fetcher        EventSource
	processorCount int
}

func NewEventProcessor(fetcher EventSource, processorCount int, queueSize int) *EventProcessor {
	return &EventProcessor{
		EventRegistry:  NewEventRegistry(),
		queue:          make(chan Event, queueSize),
		fetcher:        fetcher,
		processorCount: processorCount,
	}
}

func (e *EventProcessor) Start(ctx context.Context) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		e.Fetch(context.TODO())
	}()

	for i := 0; i < e.processorCount; i++ {
		wg.Add(1)
		ctx := context.WithValue(context.Background(), "processor_id", i)
		go func() {
			e.Process(ctx)
		}()
	}

	wg.Wait()
}

func (e *EventProcessor) Fetch(ctx context.Context) {
	for {
		log.Println("Fetching events")
		event, err := e.fetcher.Fetch(ctx)
		if err != nil {
			log.Println("Error fetching event:", err)
			time.Sleep(1 * time.Second)
			continue
		}
		if event == nil {
			log.Println("No events fetched")
			time.Sleep(1 * time.Second)
			continue
		}
		for _, ev := range event {
			select {
			case e.queue <- ev:
				log.Println("Event queued")
			default:
				log.Println("Queue is full, requeueing event")
				err := e.fetcher.Push(ctx, ev)
				if err != nil {
					log.Println("Error pushing event:", err)
				}
				time.Sleep(1 * time.Second)
			}
		}
	}

}

func (e *EventProcessor) Process(ctx context.Context) {
	for {
		log.Println("Processing events")
		select {
		case event := <-e.queue:
			serializer, ok := e.handlers[event.GetType()]
			if !ok {
				log.Fatal("No handler registered for event type:", event.GetType())
				continue
			}
			command, err := serializer.Deserialize(ctx, event.GetBody())
			if err != nil {
				log.Fatal("Error deserializing event:", err)
				continue
			}
			command.Execute(ctx)
		default:
			log.Println("No events to process")
			time.Sleep(1 * time.Second)
		}
	}
}
