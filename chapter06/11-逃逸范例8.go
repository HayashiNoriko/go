package main

func main() {
	ch := make(chan []string)

	s := []string{"hello"}

	go func() {
		ch <- s
	}()
}
