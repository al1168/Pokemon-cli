package pokecache

import (
	"fmt"
	// "structs"
	"time"
)
type CacheEntry struct {
	data []byte
	createdAt time.Time
}
type Cache struct{
	structure map[string]CacheEntry
	interval time.Duration
}
func NewCache(duration time.Duration) Cache{
	generatedCache := Cache{
		structure: make(map[string]CacheEntry),
	}
	generatedCache.interval = duration
	go generatedCache.reapLoop()
	return generatedCache
}
// func (c *Cache) Delete(url string) error {
// 	c.
// }
func (c *Cache) Add(url string, data []byte) error{
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
	done := make(chan bool)

	for{
		select{
			case <-done:
				return
			case <-ticker.C:
				toDelete := make([]string, len(c.structure))
				currentTime := time.Now()
				fmt.Printf("currentTime: %v", currentTime)
				keys := make([]string, 0, len(c.structure))
				for key := range c.structure {
					keys = append(keys, key)
				}
				fmt.Println(keys)
				fmt.Println()
				for key, cacheEntryStruct := range c.structure{
					if  currentTime.Sub(cacheEntryStruct.createdAt)  > c.interval{
						toDelete = append(toDelete, key)
					}
				for _, key := range toDelete{
					delete(c.structure, key)
				}
			}
		}
	}
	
}
// 	// for i := 0 ; i++ {

// 	// }
// 	for i:= 0; i<10; i++{
// 		fmt.Printf("%v", ticker.C)
// 	}
// 	defer ticker.Stop()

// }