// A ID generator by Snowflake

package idgen

import (
	"testing"
)

func BenchmarkIDWorker(b *testing.B) {
	iw, err := NewIDWorker(1)
	if err != nil {
		b.Error(err)
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		iw.NextID()
	}
}

func BenchmarkIDWorkerHit(b *testing.B) {
	iw, err := NewIDWorker(1)
	if err != nil {
		b.Error(err)
		return
	}
	hit := make(map[int64]struct{}, b.N)
	b.Log(len(hit))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		id, err := iw.NextID()
		if err != nil {
			continue
		}
		hit[id] = struct{}{}
	}
	delete(hit, 0)
	b.Log(len(hit))
}
