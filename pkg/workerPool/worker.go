package workerPool

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	ID         int
	workChan   chan string
	ctx        context.Context
	canselFunc context.CancelFunc
}

func NewWorker(ID int, workChan chan string, ctx context.Context, canselFunc context.CancelFunc) *Worker {
	return &Worker{
		ID:         ID,
		workChan:   workChan,
		ctx:        ctx,
		canselFunc: canselFunc,
	}
}

// процесс обработки информации работником
func (w *Worker) Process(inputChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for str := range inputChan {
		if err := w.ctx.Err(); err != nil {
			fmt.Printf("Информация: работник номер %d завершил работу досрочно\n", w.ID)
			return
		}
		fmt.Printf("Данные: %s, номер работника: %d\n", str, w.ID)
		time.Sleep(time.Second / 2)
	}
}
