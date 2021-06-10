package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func generateFile(name string, size int) error {
	f, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for i := 0; i < size; i++ {
		word := RandStringRunes(8)

		_, err = f.WriteString(word + " ")
		if err != nil {
			fmt.Println(err)
			f.Close()
			return err
		}
        //fmt.Println(l, "bytes written successfully")
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func randomFile(count, size int) {
	wg := sync.WaitGroup{}
	for i := 0 ; i < count; i++ {
		fname := fmt.Sprintf("random_%d.txt", i)
		wg.Add(1)
		go func(){
			defer wg.Done()
			generateFile(fname, size)
		}()
	}
	wg.Wait()
}

func main() {
	count := flag.Int("n", 1,"random files count")
	size  := flag.Int("s", 1000, "random files size")
	flag.Parse()
	randomFile(*count, *size)
}
