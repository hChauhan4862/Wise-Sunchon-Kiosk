package function

import (
	"fmt"
	"time"

	"github.com/ReneKroon/ttlcache"
)

var Cache *ttlcache.Cache

func InitCache() {
	Cache = ttlcache.NewCache()
	Cache.SkipTtlExtensionOnHit(true)
	Cache.SetExpirationCallback(BackGroundTaskHandler)
	// defer Cache.Close()
}

func TaskAlarm(TaskTime time.Time, Format string, args ...any) {
	fmt.Println("Task Time: ", TaskTime, "DUration: ", time.Until(TaskTime), "Name: ", fmt.Sprintf(Format, args...))
	if time.Now().After(TaskTime) || TaskTime.IsZero() || time.Now().Equal(TaskTime) {
		TaskAlarm(time.Now().Add(10*time.Second), Format, args...)
		return
	}
	NAME := fmt.Sprintf(Format, args...)
	Cache.SetWithTTL("TASK#"+NAME, "TASK", time.Until(TaskTime))
}

func LockGet(key []string) bool {
	// check if the any key is already in the cache
	for _, k := range key {
		if _, ok := Cache.Get("LOCK#" + k); ok {
			return false
		}
	}
	// if not, then add the key to the cache
	for _, k := range key {
		Cache.SetWithTTL("LOCK#"+k, "LOCK", 30*time.Second) // This Lock will be released after 30 seconds if not released manually
	}
	return true
}

func LockRelease(key []string) {
	for _, k := range key {
		Cache.Remove("LOCK#" + k)
	}
}
