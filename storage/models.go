package storage



type Storage interface {
    Set(id string, obj interface{}) bool
    Get(id string) (interface{}, error)
    Has(id string) bool
    Del(id string) bool
}
