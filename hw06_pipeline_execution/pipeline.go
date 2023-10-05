package main

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := in

	if len(stages) == 0 {
		return out
	}

	for _, stage := range stages {
		stream := make(Bi)
		go func(inStage Bi, outStage Out) {
			defer close(stream)
			for {
				select {
				case <-done:
					return
				case stage, ok := <-outStage:
					if !ok {
						return
					}
					inStage <- stage
				}
			}
		}(stream, out)
		out = stage(stream)
	}

	return out
}
