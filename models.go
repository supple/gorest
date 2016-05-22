package main

import (
    "sync"
    "math/rand"
)

var (
    mu     sync.RWMutex
)

// Helper random string generator
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
func RandSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

// # Model: AppService
type AppService struct {
    ca *CacheArr
}

// # Model: Campaign object
type Campaign struct {
    Id string `json:"id"`
    Name string `json:"name"`
}

func NewCampaign(id string, name string) *Campaign {
    return &Campaign{Id: id, Name: name}
}

// # Model: Cache, campaign id -> *Campaign
type CacheArr struct {
    campaigns map[string]*Campaign
}

func NewCacheArr(names []string) *CacheArr {
    ca := CacheArr{}
    ca.campaigns = make(map[string]*Campaign, len(names))
    for _, name := range names {
        id := RandSeq(5)
        ca.campaigns[id] = NewCampaign(id, name)
    }

    return &ca
}

// Get returns campaign, or nil if there's no one.
func (ca* CacheArr) Get(id string) *Campaign {
    mu.RLock()
    cmp := ca.campaigns[id]
    mu.RUnlock()

    return cmp
}

// Set campaign in cache
func (ca* CacheArr) Set(pCmp *Campaign) {
    mu.Lock()
    ca.campaigns[pCmp.Id] = pCmp
    mu.Unlock()
}

// Get campaign names with ids
func (ca* CacheArr) GetNames() map[string]string {
    names := make(map[string]string)
    mu.RLock()
    for _, cmp := range ca.campaigns {
        names[cmp.Id] = cmp.Name
    }
    mu.RUnlock()
    return names
}

// # Model: User

type User struct {
    id string
    appId string
    os string
    appVersion string
    token string
    email string
}