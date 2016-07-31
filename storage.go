package main


type Storage interface {
    Set(id string, obj interface{}) bool
    Get(id string) (interface{}, error)
    Has(id string) bool
    Del(id string) bool
}


type Item struct {
    value interface{}
}

func (it *Item) NewItem(v interface{})  {
    it.value = v
}

// Cache, id -> *Object
type MemStorage struct {
    objects map[string]interface{}
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
    mu.RLock()
    obj := ca.objects[id]
    mu.RUnlock()

    return obj
}

// Set object in storage
func (ca* MemStorage) Set(id string, obj interface{}) bool {
    mu.Lock()
    ca.objects[id] = obj
    mu.Unlock()
    return true
}

// Update object in storage
func (ca* MemStorage) Update(id string,obj interface{}) {
    mu.Lock()
    ca.objects[id] = obj
    mu.Unlock()
}

// Get object names with ids
func (ca* MemStorage) GetByCriteria() []interface{} {
    var names []interface{}
    mu.RLock()
    for _, obj := range ca.objects {
        names = append(names, obj)
    }
    mu.RUnlock()
    return names
}