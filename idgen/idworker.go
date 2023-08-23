// A ID generator by Snowflake

package idgen

import (
	"errors"
	"log"
	"sync"
	"time"
)

// Snowflake的优化方案，减少了原算法产生的ID数，改为每秒产生65535个ID

// 63     52              20	           15			 0
// +------+---------------+----------------+------------+
// |  0...| timestamp(s)  | worker node id | sequence	|
// +------+---------------+----------------+------------+

const (
	// 时间戳使用32位存放，为了增加上限采用存储时减去此值，此值为2020-01-01 00:00:00时间戳
	CEpoch        = 1577808000
	CWorkerIDBits = 5  // Num of WorkerID Bits
	CSequenceBits = 16 // Num of Sequence Bits
)

// IDWorker Struct
type IDWorker struct {
	workerID      int64
	lastTimeStamp int64
	sequence      int64
	sequenceMask  int64
	maxWorkerID   int64
	lock          *sync.Mutex
}

// NewIDWorker Func: Generate NewIDWorker with Given workerid
func NewIDWorker(workerID int64) (*IDWorker, error) {
	iw := new(IDWorker)

	iw.maxWorkerID = getMaxWorkerID()

	if workerID > iw.maxWorkerID || workerID < 0 {
		return nil, errors.New("worker not fit")
	}

	iw.workerID = workerID
	iw.lastTimeStamp = -1
	iw.sequence = 0
	iw.sequenceMask = getSequenceMask()
	iw.lock = new(sync.Mutex)

	return iw, nil
}

func getMaxWorkerID() int64 {
	return -1 ^ -1<<CWorkerIDBits
}

func getSequenceMask() int64 {
	return -1 ^ -1<<CSequenceBits
}

func (iw *IDWorker) timeGen() int64 {
	return time.Now().Unix()
}

func (iw *IDWorker) timeReGen(last int64) int64 {
	ts := time.Now().Unix()
	for {
		if ts <= last {
			ts = iw.timeGen()
		} else {
			break
		}
	}
	return ts
}

// NextID Generate next ID
func (iw *IDWorker) NextID() (ts int64, err error) {
	iw.lock.Lock()
	defer iw.lock.Unlock()

	ts = iw.timeGen()
	if ts < iw.lastTimeStamp {
		err = errors.New("Clock moved backwards, Refuse gen ID")
		return 0, err
	}
	if ts == iw.lastTimeStamp {
		iw.sequence = (iw.sequence + 1) & iw.sequenceMask
		if iw.sequence == 0 {
			log.Printf("maximum id reached in 1 second in epoch %d", ts)
			ts = iw.timeReGen(ts)
		}
	} else {
		iw.sequence = 0
	}

	iw.lastTimeStamp = ts
	ts = (ts-CEpoch)<<CWorkerIDBits + CSequenceBits | iw.workerID<<CSequenceBits | iw.sequence
	return ts, nil
}
