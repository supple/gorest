package storage

import "sync"

type Item struct {
    value interface{}
}

func (it *Item) NewItem(v interface{})  {
    it.value = v
}

// Cache, id -> *Object
type MemStorage struct {
    objects map[string]interface{}
    mu *sync.RWMutex
}

func NewMemStorage() *MemStorage {
    ca := MemStorage{}
    //ca.objects = make(map[string]*interface{}, len(objs))
    //for _, obj := range objs {
    //    id := randSeq(5)
    //    ca.objects[id] = obj
    //}

    return &ca
}

// Get returns object, or nil if there's no one.
func (ca* MemStorage) Get(id string) interface{} {
    ca.mu.RLock()
    obj := ca.objects[id]
    ca.mu.RUnlock()

    return obj
}

// Set object in storage
func (ca* MemStorage) Set(id string, obj interface{}) bool {
    ca.mu.Lock()
    ca.objects[id] = obj
    ca.mu.Unlock()
    return true
}

// Update object in storage
func (ca* MemStorage) Update(id string,obj interface{}) {
    ca.mu.Lock()
    ca.objects[id] = obj
    ca.mu.Unlock()
}

// Get object names with ids
func (ca* MemStorage) GetByCriteria() []interface{} {
    var names []interface{}
    ca.mu.RLock()
    for _, obj := range ca.objects {
        names = append(names, obj)
    }
    ca.mu.RUnlock()
    return names
}