package pool

import "io"

type Pool struct {
	/*  */
}

func (p *Pool) Acquire() {

}

func (p *Pool) Release(resource io.Closer) {

}

func (p *Pool) Close() {

}

func New(poolSize int, factory func() (io.Closer, error)) (*Pool, error) {
	/*  */
	return &Pool{
		/*  */
	}, nil
}
