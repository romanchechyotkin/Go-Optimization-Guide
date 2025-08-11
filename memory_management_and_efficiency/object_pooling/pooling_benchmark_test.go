package perf

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"sync"
	"testing"
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

// BenchmarkWithoutPooling measures the performance of direct heap allocations.
func BenchmarkWithoutPooling(b *testing.B) {
	for b.Loop() {
		data := generateDataJson()
		obj := &Data{}
		json.Unmarshal(data, obj)
	}
}

// dataPool is a sync.Pool that reuses instances of Data to reduce memory allocations.
var dataPool = sync.Pool{
	New: func() any {
		return &Data{}
	},
}

// BenchmarkWithPooling measures the performance of using sync.Pool to reuse objects.
func BenchmarkWithPooling(b *testing.B) {
	for b.Loop() {
		data := generateDataJson()
		obj := dataPool.Get().(*Data) // Retrieve from pool
		json.Unmarshal(data, obj)
		dataPool.Put(obj) // Return object to pool for reuse
	}
}
