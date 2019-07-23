package channel

import (
	"sync"
	"time"
)

type Data struct {
	sync.Mutex
}

//当定义未(d Data)test(s string){}非指针类型时候锁失效
func (d *Data) test(s string) {
	d.Lock()
	defer d.Unlock()

	for i := 0; i < 5; i++ {
		print(s, i, "\n")
		time.Sleep(time.Second)
	}
}

func Test1() {
	var wg sync.WaitGroup
	wg.Add(2)

	var d Data

	go func() {
		defer wg.Done()

		d.test("write")
	}()

	go func() {
		defer wg.Done()

		d.test("read")
	}()

	wg.Wait()
}
