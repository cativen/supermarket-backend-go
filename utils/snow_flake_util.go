package utils

import (
	"errors"
	"sync"
	"time"
)

const (
	epoch             int64 = 1719562568879 // 设置起始时间戳，例如2021-01-01 00:00:00 UTC
	timeBitLength     uint8 = 41            // 时间戳占用的位数
	workerIDBitLength uint8 = 5             // 工作机器ID占用的位数
	sequenceBitLength uint8 = 12            // 序列号占用的位数

	maxWorkerID   int64 = -1 ^ (-1 << workerIDBitLength) // 工作机器ID的最大值
	maxSequence   int64 = -1 ^ (-1 << sequenceBitLength) // 序列号的最大值
	timeShift     uint8 = workerIDBitLength + sequenceBitLength
	workerIDShift uint8 = sequenceBitLength
	twepoch       int64 = epoch
)

type Snowflake struct {
	mu        sync.Mutex
	timestamp int64
	workerID  int64
	sequence  int64
}

func NewSnowflake(workerID int64) (*Snowflake, error) {
	if workerID < 0 || workerID > maxWorkerID {
		return nil, errors.New("worker ID out of range")
	}
	return &Snowflake{
		timestamp: 0,
		workerID:  workerID,
		sequence:  0,
	}, nil
}

func (s *Snowflake) Generate() (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().UnixMilli() - twepoch
	if now < s.timestamp {
		return 0, errors.New("clock is moving backwards")
	}

	if now == s.timestamp {
		s.sequence = (s.sequence + 1) & maxSequence
		if s.sequence == 0 {
			for now <= s.timestamp {
				now = time.Now().UnixMilli() - twepoch
			}
		}
	} else {
		s.sequence = 0
	}

	s.timestamp = now

	id := ((now << timeShift) | (s.workerID << workerIDShift) | (s.sequence))
	return id, nil
}
