package mediator

import "sync"

// 站长
type StationManager struct {
	isPlatformFree bool
	lock           *sync.Mutex
	trainQueue     []train
}

func NewStationManger() *StationManager {
	return &StationManager{
		isPlatformFree: true,
		lock:           &sync.Mutex{},
	}
}

func (s *StationManager) CanLand(t train) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *StationManager) IsLanded(t train) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	for _, item := range s.trainQueue {
		if t == item {
			return true
		}
	}
	return false
}

func (s *StationManager) NotifyFree() {
	s.lock.Lock()
	defer s.lock.Unlock()
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		firstTrainInQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainInQueue.PermitArrival()
	}
}
