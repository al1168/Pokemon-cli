package pokecache

import (
	"fmt"
	"time"
	"sync"
)
type CacheEntry struct {
	data []byte
	createdAt time.Time
}
type Cache struct{
	structure map[string]CacheEntry
	interval time.Duration
	mux *sync.Mutex
}
func NewCache(duration time.Duration) Cache{
	generatedCache := Cache{
		structure: make(map[string]CacheEntry),
		mux : &sync.Mutex{},
	}
	generatedCache.interval = duration
	go generatedCache.reapLoop()
	return generatedCache
}

func (c *Cache) Add(url string, data []byte) error{
	c.mux.Lock()
	defer c.mux.Unlock()
	if url == "" || data == nil{
		return fmt.Errorf("failed caching: Url: %v \n data: %v", url, data)
	}
	timeRightNow := time.Now()
	newCacheEntry := CacheEntry{}
	newCacheEntry.createdAt = timeRightNow
	newCacheEntry.data = data
	c.structure[url] = newCacheEntry
	return nil
}

func (c *Cache) Get(url string) ([]byte, bool, error){
	c.mux.Lock()
	defer c.mux.Unlock()
	val, ok := c.structure[url]
	if !ok {
		return nil, false, fmt.Errorf("cache element key DNE, url: %v", url)
	}
	return val.data, true, nil
}
/*
Create a cache.reapLoop() 
method that is called when the cache is 
created (by the NewCache function). 
Each time an interval (the time.Duration passed to NewCache) 
passes it should remove any entries that are older than the interval. 
This makes sure that the cache doesn't grow too large over time. 
For example, if the interval is 5 seconds, and an entry was added 7 seconds ago,
 that entry should be removed.
*/
func (c *Cache) reapLoop(){

	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.reap()
	}
}
	
func (c *Cache) reap(){
	c.mux.Lock()
	defer c.mux.Unlock()
	toDelete := make([]string, len(c.structure))
	currentTime := time.Now()
	fmt.Printf("currentTime: %v", currentTime)
	
	for key, cacheEntryStruct := range c.structure{
		if  currentTime.Sub(cacheEntryStruct.createdAt)  > c.interval{
			toDelete = append(toDelete, key)
		}
	}
	for _, key := range toDelete{
		delete(c.structure, key)
	}
	fmt.Printf("\ncache size %v\n",len(c.structure))
}