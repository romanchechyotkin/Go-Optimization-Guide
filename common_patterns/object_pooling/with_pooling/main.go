package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"sync"
	"time"
)

type Data struct {
	UserID    int64  `json:"userId"`
	MessageID int64  `json:"messageId"`
	Ts        int64  `json:"ts"`
	Channel   string `json:"channel"`
	Type      string `json:"type"`
}

var channelMap = map[int]string{
	0: "MAIL",
	1: "PUSH",
	2: "SMS",
	3: "FC",
}

var typesMap = map[int]string{
	0: "SEND",
	1: "DELIVERY",
	2: "OPEN",
	3: "CLICK",
}

func generateDataJson() []byte {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	channel := channelMap[r.Intn(4)]
	t := typesMap[r.Intn(4)]

	userID := r.Intn(math.MaxInt64)
	messageID := r.Intn(math.MaxInt64)

	ts := time.Now().Unix()

	return []byte(fmt.Sprintf(`{
    "userId": %d,
    "messageId": %d,
    "ts": %d,
    "channel": "%s",
    "type": "%s"
}`, userID, messageID, ts, channel, t))
}

var dataPool = sync.Pool{
	New: func() any {
		return &Data{}
	},
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			data := generateDataJson()

			obj := dataPool.Get().(*Data) // Retrieve from pool
			*obj = Data{}

			log.Printf("iteration #%d; obj before unmarshal %+v\n", i, obj)
			json.Unmarshal(data, obj)
			log.Printf("iteration #%d; obj after unmarshal %+v\n", i, obj)

			dataPool.Put(obj) // Return object to pool for reuse
		}()
	}

	wg.Wait()

	fmt.Println("Done")
}
