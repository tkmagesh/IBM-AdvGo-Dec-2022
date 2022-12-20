package worker

type Work interface {
	Task()
}

type Worker struct {
	/*  */
}

func (w Worker) Run() {
	/*  */
}

func (w Worker) Shutdown() {
	/*  */
}

func New(count int) *Worker {
	/* return &Worker{} */
}
