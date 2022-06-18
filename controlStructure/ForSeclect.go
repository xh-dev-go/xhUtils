package controlStructure

func ForSelect[V any](msg chan V, done chan struct{}, operation func(input V)) {
	var isDone = false
	for {
		if isDone {
			break
		}
		select {
		case m := <-msg:
			operation(m)
		case <-done:
			isDone = true
		}
	}
}
