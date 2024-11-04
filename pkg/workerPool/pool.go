package workerPool

import (
	"context"
	"fmt"
	"sync"
)

type WorkerPool struct {
	wg              *sync.WaitGroup
	workers         []Worker
	isProcessInWork bool
	workedChannal   chan string
}

func NewWorkerPool() *WorkerPool {
	return &WorkerPool{
		wg:              &sync.WaitGroup{},
		workers:         []Worker{},
		isProcessInWork: false,
		workedChannal:   make(chan string),
	}
}

func (p *WorkerPool) AddWorkers(count int, parentCtx context.Context) {
	for i := 0; i < count; i++ {
		ctx, cancel := context.WithCancel(parentCtx)
		p.workers = append(p.workers, *NewWorker(len(p.workers)+1, nil, ctx, cancel))
		fmt.Printf("Информация: работник номер %d был добавлен\n", p.workers[len(p.workers)-1].ID)
		if p.isProcessInWork {
			StartNewProcces(p.workedChannal, p.wg, p.workers[len(p.workers)-1])
		}
	}

}

func (p *WorkerPool) DeletWorkers(count int) error {
	if len(p.workers) >= count {
		for i := len(p.workers) - 1; i >= len(p.workers)-count; i-- {
			p.workers[i].canselFunc()
			fmt.Printf("Информация: работник номер %d был удалён\n", p.workers[i].ID)
		}
		p.workers = p.workers[:len(p.workers)-count]

	} else {
		return fmt.Errorf("Попытка удалить Больше рабоотчников, чем существует на данный момент")
	}

	return nil
}

// запуск обработки файла
func (p *WorkerPool) Run(strs []string, parentWg *sync.WaitGroup) error {
	defer parentWg.Done()
	p.isProcessInWork = true
	defer func() {
		p.isProcessInWork = false
	}()

	go func() {
		for _, str := range strs {
			p.workedChannal <- str
		}
		close(p.workedChannal)
	}()

	//обход всех имеющихся работников
	//и запуск их работы
	for _, val := range p.workers {
		StartNewProcces(p.workedChannal, p.wg, val)

	}
	p.wg.Wait()

	return nil
}

func StartNewProcces(inputChan chan string, wg *sync.WaitGroup, worker Worker) {
	wg.Add(1)
	go worker.Process(inputChan, wg)
}
