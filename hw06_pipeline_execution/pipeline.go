package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func worker(in Bi, done In, out Out) {
	for {
		select {
		case <-done:
			close(in)
			return
		case v, ok := <-out:
			if !ok {
				close(in)
				return
			} else {
				in <- v
			}
		}
	}
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	for _, stage := range stages {
		chanNext := make(Bi)
		go worker(chanNext, done, in)
		in = stage(chanNext)
	}
	return in
}
