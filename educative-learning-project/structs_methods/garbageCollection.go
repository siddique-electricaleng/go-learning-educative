package main

import (
	"runtime"
	"time"
)

func garbageCollectionMain() {
	ms := runtime.MemStats{}
	runtime.ReadMemStats(&ms)

	println("Heap after GC\nUsed: ", ms.HeapInuse, "kilobytes Free:", ms.HeapIdle, "kilobytes Meta:", ms.GCSys, "kilobytes")

	time.Sleep(5 * time.Second)

}
