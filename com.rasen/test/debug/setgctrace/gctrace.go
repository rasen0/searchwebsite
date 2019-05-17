package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)
// set GODEBUG=gctrace=1 && go run ..
func main() {
	st := struct{
		s   string
		dig int64
	}{
		"sss",
		123456789,
	}
	cmap := sync.Map{}
	countLoop := 0
	var wg sync.WaitGroup
	flag := make(chan bool)
	wg.Add(2)
	go func() {
		rangenum0 := 0
		for i := 0; i < 1000000; i++ {
			cmap.Store(i, st)
			rangenum0++
		}
		flag <- true
		fmt.Printf("save ending:%v cmap:%p\n",rangenum0,&cmap)
		wg.Done()
	}()
	var total int32
	//go func(cmap sync.Map) {
	go func() {
	JUMP:
		for {
			rangenum := 0
			cmap.Range(func(key, value interface{}) bool {
				cmap.Delete(key)
				rangenum++
				return true
			})
			countLoop++
			atomic.AddInt32(&total,int32(rangenum))
			fmt.Printf("j cmap:%p range:%v,total:%v\n",&cmap,rangenum,total)
			rangenum = 0
			var i bool
			select {
			case i=<-flag:
				fmt.Println(i)
				cmap.Range(func(key, value interface{}) bool {
					cmap.Delete(key)
					rangenum++
					return true
				})
				fmt.Println("goto jump")
				atomic.AddInt32(&total,int32(rangenum))
				break JUMP
			default:
			}
			time.Sleep(100 * time.Millisecond)
		}
		wg.Done()
	}()
	wg.Wait()
	atomic.LoadInt32(&total)
	fmt.Println("countLoop:", countLoop," total:",total)
}

//func main() {
//	st := struct{
//		s   string
//		dig int64
//	}{
//		"sss",
//		123456789,
//	}
//	cmap := sync.Map{}
//	countLoop := 0
//	var wg sync.WaitGroup
//	flag := make(chan bool)
//	wg.Add(2)
//	go func() {
//		rangenum0 := 0
//		for i := 0; i < 1000000; i++ {
//			cmap.Store(i, st)
//			rangenum0++
//		}
//		flag <- true
//		fmt.Printf("save ending:%v cmap:%p\n",rangenum0,&cmap)
//		wg.Done()
//	}()
//	var total int32
//	//go func(cmap sync.Map) {
//	go func() {
//	JUMP:
//		for {
//			rangenum := int32(0)
//			cmap.Range(func(key, value interface{}) bool {
//				cmap.Delete(key)
//				rangenum++
//				return true
//			})
//			countLoop++
//			total+=rangenum
//			fmt.Printf("j cmap:%p range:%v,total:%v\n",&cmap,rangenum,total)
//			rangenum = 0
//			var i bool
//			select {
//			case i=<-flag:
//				fmt.Println(i)
//				cmap.Range(func(key, value interface{}) bool {
//					cmap.Delete(key)
//					rangenum++
//					return true
//				})
//				fmt.Println("goto jump")
//				total+=int32(rangenum)
//				break JUMP
//			default:
//			}
//			time.Sleep(100 * time.Millisecond)
//		}
//		wg.Done()
//	}()
//	wg.Wait()
//	atomic.LoadInt32(&total)
//	fmt.Println("countLoop:", countLoop," total:",total)
//}
