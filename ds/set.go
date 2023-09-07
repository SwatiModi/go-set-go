package ds

type Set interface {
	Add(item interface{}) bool
	Remove(item interface{}) bool
	Contains(item interface{}) bool
}
