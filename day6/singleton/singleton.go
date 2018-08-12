package singleton

type singleton struct {
	count int
}
type Singleton interface {
	Increment() int
}

var instance *singleton

func (s *singleton) Increment() int {
	s.count++
	return s.count
}

func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}

	// fmt.Println(instance)
	return instance
}
