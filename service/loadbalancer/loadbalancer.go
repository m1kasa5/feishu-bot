package loadbalancer

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type API struct {
	Key       string
	Times     uint32
	Available bool
}

type LoadBalancer struct {
	apis []*API
	mu   sync.RWMutex
}

func (lb *LoadBalancer) SetAvailability(key string, available bool) {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	for _, api := range lb.apis {
		if api.Key == key {
			api.Available = available
			return
		}
	}
}

func (lb *LoadBalancer) SetAvailabilityForAll(available bool) {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	for _, api := range lb.apis {
		api.Available = available
	}
}

func (lb *LoadBalancer) RegisterAPI(key string) {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	if lb.apis == nil {
		lb.apis = make([]*API, 0)
	}
	lb.apis = append(lb.apis, &API{Key: key})
}

func (lb *LoadBalancer) GetAPIs() []*API {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	apis := make([]*API, len(lb.apis))
	copy(apis, lb.apis)
	return apis
}

func (lb *LoadBalancer) GetAPI() *API {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	var availableAPIs []*API

	for _, api := range lb.apis {
		if api.Available {
			availableAPIs = append(availableAPIs, api)
		}
	}

	if len(availableAPIs) == 0 {
		// 随机复活一个
		fmt.Printf("No available API, revive one randomly\n")
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(len(lb.apis))
		lb.apis[index].Available = true
		return lb.apis[index]
	}

	selectAPI := availableAPIs[0]
	minTimes := selectAPI.Times

	// 选用时最短的api
	for _, api := range lb.apis {
		if api.Times < minTimes {
			selectAPI = api
			minTimes = api.Times
		}
	}

	selectAPI.Times++
	return selectAPI
}

func NewLoadBalancer(keys []string) *LoadBalancer {
	lb := &LoadBalancer{}

	for _, key := range keys {
		lb.apis = append(lb.apis, &API{Key: key})
	}

	lb.SetAvailabilityForAll(true)
	return lb
}
