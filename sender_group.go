package pusher

import "sync"

type SenderGroup struct {
	data sync.Map
}

func (s *SenderGroup) Add(sender MessageSender) {
	s.data.Store(sender.GetName(), sender)
}
func (s *SenderGroup) Update(sender MessageSender) {
	s.Delete(sender.GetName())
	s.Add(sender)
}
func (s *SenderGroup) Get(name string) (MessageSender, bool) {
	value, found := s.data.Load(name)
	if found {
		return value.(MessageSender), found
	}
	return nil, found
}
func (s *SenderGroup) Delete(name string) {
	sender, found := s.Get(name)
	if found {
		sender.Close()
	}
	s.data.Delete(name)
}
func (s *SenderGroup) GetSpeed() uint64 {
	var speedCount uint64 = 0
	s.data.Range(func(key, value any) bool {
		speed := value.(MessageSender).GetSpeed()
		speedCount += speed
		return true
	})
	return speedCount
}
func (s *SenderGroup) DeleteAll() {
	s.data.Range(func(key, value any) bool {
		sender := value.(MessageSender)
		s.Delete(sender.GetName())
		return true
	})
}
