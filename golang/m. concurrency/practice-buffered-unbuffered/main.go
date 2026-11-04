//package main
//
//import (
//	"fmt"
//	"sync"
//)
//
//// Print A B C Repeatedly.
//// Use Only three goroutines.
//// Each goroutine prints only one letter.
//// No if for counter checking
//// No if for closing channels
//// No sleeps allowed.
//// Output: ABCABCABC
//
//func printA(aChan, bChan, cChan, sem chan struct{}, wg *sync.WaitGroup) {
//	defer wg.Done()
//
//	for i := 1; i <= 3; i++ {
//		<-aChan
//
//		sem <- struct{}{}
//
//		fmt.Print("A", " ")
//
//		bChan <- struct{}{}
//
//		<-sem
//	}
//
//}
//
//func printB(aChan, bChan, cChan, sem chan struct{}, wg *sync.WaitGroup) {
//	defer wg.Done()
//
//	for i := 1; i <= 3; i++ {
//		<-bChan
//
//		sem <- struct{}{}
//
//		fmt.Print("B", " ")
//
//		cChan <- struct{}{}
//
//		<-sem
//
//	}
//}
//
//func printC(aChan, bChan, cChan, sem chan struct{}, wg *sync.WaitGroup) {
//	defer wg.Done()
//
//	for i := 1; i <= 3; i++ {
//		<-cChan
//
//		sem <- struct{}{}
//
//		fmt.Print("C", " ")
//
//		aChan <- struct{}{}
//
//		<-sem
//
//	}
//}
//
//func main() {
//	var (
//		aChan = make(chan struct{}, 1)
//		bChan = make(chan struct{})
//		cChan = make(chan struct{})
//
//		wg sync.WaitGroup
//
//		limit = 3
//		sem   = make(chan struct{}, limit)
//	)
//
//	wg.Add(3)
//
//	go printA(aChan, bChan, cChan, sem, &wg)
//	go printB(aChan, bChan, cChan, sem, &wg)
//	go printC(aChan, bChan, cChan, sem, &wg)
//
//	aChan <- struct{}{}
//
//	wg.Wait()
//
//}

package main

import (
	"fmt"
	"sync"
)

// Print A B C Repeatedly.
// Use Only three goroutines.
// Each goroutine prints only one letter.
// No if for counter checking
// No if for closing channels
// No sleeps allowed.
// Output: ABCABCABC

func printA(aChan, bChan, cChan, sem chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 3; i++ {
		<-aChan

		//	sem <- struct{}{}

		fmt.Print("A", " ")

		bChan <- struct{}{}

		//	<-sem
	}

}

func printB(aChan, bChan, cChan, sem chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 3; i++ {
		<-bChan

		//	sem <- struct{}{}

		fmt.Print("B", " ")

		cChan <- struct{}{}

		//	<-sem

	}
}

func printC(aChan, bChan, cChan, sem chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 3; i++ {
		<-cChan

		//	sem <- struct{}{}

		fmt.Print("C", " ")

		aChan <- struct{}{}

		//	<-sem

	}
}

func main() {
	var (
		aChan = make(chan struct{}, 1)
		bChan = make(chan struct{})
		cChan = make(chan struct{})

		wg sync.WaitGroup

		limit = 3
		sem   = make(chan struct{}, limit)
	)

	wg.Add(3)

	go printA(aChan, bChan, cChan, sem, &wg)
	go printB(aChan, bChan, cChan, sem, &wg)
	go printC(aChan, bChan, cChan, sem, &wg)

	aChan <- struct{}{}

	wg.Wait()

}
