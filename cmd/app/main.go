package main

import (
	"bufio"
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/sashaem1/WorkerPool_VK/pkg/workerPool"
)

func main() {
	data := GetData("data/test.txt")
	p := workerPool.NewWorkerPool()
	wg := &sync.WaitGroup{}

	parentCtx := context.Background()
	p.AddWorkers(6, parentCtx)

	wg.Add(1)
	go p.Run(data, wg)

	time.Sleep(time.Second * 2)
	p.DeletWorkers(3)

	time.Sleep(time.Second * 2)
	p.AddWorkers(2, parentCtx)

	wg.Wait()

}

func GetData(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}
