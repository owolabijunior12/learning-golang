package main

import (
	"fmt"
	"sync"
	"time"
)

// COURSE 4: CONCURRENCY - GOROUTINES AND CHANNELS
// Topics covered:
// 1. Goroutines (lightweight threads)
// 2. Channels (safe communication between goroutines)
// 3. Channel operations (send, receive, close)
// 4. Channel directions (send-only, receive-only)
// 5. Select statement (multiplexing)
// 6. Buffered vs unbuffered channels
// 7. Worker pools
// 8. WaitGroup for synchronization
// 9. Timeouts and context

// ============ 1. SIMPLE GOROUTINE ============
func greet(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("Hello %s (iteration %d)\n", name, i)
		time.Sleep(100 * time.Millisecond)
	}
}

// ============ 2. CHANNEL BASICS ============
// Send numbers from 1 to n through a channel
func generateNumbers(n int, ch chan int) {
	for i := 1; i <= n; i++ {
		fmt.Printf("Generating: %d\n", i)
		ch <- i // send
		time.Sleep(100 * time.Millisecond)
	}
	close(ch) // always close channels when done
}

// Read from channel and process
func processNumbers(ch chan int) {
	for num := range ch { // receives until channel is closed
		fmt.Printf("Processing: %d, Square: %d\n", num, num*num)
	}
}

// ============ 3. BUFFERED CHANNELS ============
// Can hold multiple values without blocking
func bufferedChannelDemo() {
	ch := make(chan int, 3) // capacity of 3

	ch <- 10
	ch <- 20
	ch <- 30

	fmt.Printf("Value 1: %d\n", <-ch)
	fmt.Printf("Value 2: %d\n", <-ch)
	fmt.Printf("Value 3: %d\n", <-ch)
}

// ============ 4. SELECT STATEMENT ============
// Wait for multiple channel operations
func receiveFromMultiple(ch1, ch2 chan string) {
	for i := 0; i < 4; i++ {
		select {
		case msg := <-ch1:
			fmt.Printf("From ch1: %s\n", msg)
		case msg := <-ch2:
			fmt.Printf("From ch2: %s\n", msg)
		}
	}
}

// ============ 5. TIMEOUT WITH SELECT ============
func fetchWithTimeout(ch chan string) {
	select {
	case result := <-ch:
		fmt.Printf("Got result: %s\n", result)
	case <-time.After(2 * time.Second):
		fmt.Println("Operation timed out!")
	}
}

// ============ 6. WORKER POOL PATTERN ============
type Job struct {
	ID   int
	Data string
}

type Result struct {
	Job    Job
	Output string
}

func worker(id int, jobs <-chan Job, results chan<- Result) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job.ID)
		time.Sleep(500 * time.Millisecond)

		results <- Result{
			Job:    job,
			Output: fmt.Sprintf("Processed: %s", job.Data),
		}
	}
}

// ============ 7. SYNC.WAITGROUP ============
// Used to wait for multiple goroutines to complete
func downloadFile(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark as complete when function returns

	fmt.Printf("Downloading file %d...\n", id)
	time.Sleep(time.Duration(id) * 500 * time.Millisecond)
	fmt.Printf("File %d downloaded!\n", id)
}

// ============ 8. PRODUCER-CONSUMER PATTERN ============
func producer(ch chan<- int, count int) {
	for i := 1; i <= count; i++ {
		fmt.Printf("Producing: %d\n", i)
		ch <- i
		time.Sleep(200 * time.Millisecond)
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for value := range ch {
		fmt.Printf("Consuming: %d\n", value)
	}
}

// ============ 9. FAN-OUT FAN-IN PATTERN ============
func fanOut(input <-chan int, numWorkers int) []<-chan int {
	channels := make([]<-chan int, numWorkers)
	for i := 0; i < numWorkers; i++ {
		ch := make(chan int)
		go func(id int, ch chan<- int) {
			for val := range input {
				fmt.Printf("Worker %d received: %d\n", id, val)
				ch <- val * val
			}
			close(ch)
		}(i, ch)
		channels[i] = ch
	}
	return channels
}

func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for val := range c {
				out <- val
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// ============ COURSE FOUR MAIN FUNCTION ============
func courseFour() {
	fmt.Println("=== CONCURRENCY: GOROUTINES AND CHANNELS ===\n")

	// ============ 1. BASIC GOROUTINES ============
	fmt.Println("1. BASIC GOROUTINES")
	fmt.Println("---")

	// Without goroutines - sequential execution
	fmt.Println("Sequential (takes 3 seconds):")
	greet("Alice")

	// With goroutines - concurrent execution
	fmt.Println("\nConcurrent (takes ~1 second):")
	go greet("Bob")
	go greet("Charlie")
	time.Sleep(1 * time.Second) // Give goroutines time to complete
	fmt.Println()

	// ============ 2. UNBUFFERED CHANNELS ============
	fmt.Println("2. UNBUFFERED CHANNELS (Synchronous)")
	fmt.Println("---")

	ch := make(chan int) // unbuffered

	go generateNumbers(3, ch)
	processNumbers(ch)
	fmt.Println()

	// ============ 3. BUFFERED CHANNELS ============
	fmt.Println("3. BUFFERED CHANNELS (Asynchronous)")
	fmt.Println("---")
	bufferedChannelDemo()
	fmt.Println()

	// ============ 4. SELECT STATEMENT ============
	fmt.Println("4. SELECT STATEMENT (Multiplexing)")
	fmt.Println("---")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "Message from ch1"
		ch1 <- "Another from ch1"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "Message from ch2"
		ch2 <- "Another from ch2"
	}()

	receiveFromMultiple(ch1, ch2)
	fmt.Println()

	// ============ 5. TIMEOUT ============
	fmt.Println("5. TIMEOUT PATTERN")
	fmt.Println("---")

	slowChannel := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		slowChannel <- "This will timeout"
	}()

	fetchWithTimeout(slowChannel)
	fmt.Println()

	// ============ 6. WORKER POOL ============
	fmt.Println("6. WORKER POOL PATTERN")
	fmt.Println("---")

	jobs := make(chan Job, 5)
	results := make(chan Result, 5)

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Submit jobs
	for j := 1; j <= 5; j++ {
		jobs <- Job{ID: j, Data: fmt.Sprintf("Job %d data", j)}
	}
	close(jobs)

	// Collect results
	fmt.Println("Results:")
	for i := 0; i < 5; i++ {
		result := <-results
		fmt.Printf("  Job %d: %s\n", result.Job.ID, result.Output)
	}
	fmt.Println()

	// ============ 7. SYNC.WAITGROUP ============
	fmt.Println("7. SYNC.WAITGROUP")
	fmt.Println("---")

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go downloadFile(i, &wg)
	}

	wg.Wait()
	fmt.Println("All downloads complete!\n")

	// ============ 8. PRODUCER-CONSUMER ============
	fmt.Println("8. PRODUCER-CONSUMER PATTERN")
	fmt.Println("---")

	producerCh := make(chan int)
	go producer(producerCh, 5)
	consumer(producerCh)
	fmt.Println()

	// ============ 9. FAN-OUT / FAN-IN ============
	fmt.Println("9. FAN-OUT / FAN-IN PATTERN")
	fmt.Println("---")

	// Simplified fan-out/fan-in
	input := make(chan int, 4)
	for i := 1; i <= 4; i++ {
		input <- i
	}
	close(input)

	// Create 2 workers
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for val := range input {
			ch1 <- val * val
		}
		close(ch1)
	}()

	go func() {
		for val := range input {
			ch2 <- val * val
		}
		close(ch2)
	}()

	// Merge results
	fmt.Println("Squared results from workers:")
	for i := 0; i < 4; i++ {
		select {
		case val := <-ch1:
			fmt.Printf("  Worker 1: %d\n", val)
		case val := <-ch2:
			fmt.Printf("  Worker 2: %d\n", val)
		}
	}
	fmt.Println()

	fmt.Println("=== END OF COURSE 4: CONCURRENCY ===")
}

// Helper types and functions for concurrency patterns
type workerResult struct {
	data chan int
}
}

// KEY TAKEAWAYS:
// 1. Goroutines are lightweight - you can have thousands
// 2. Channels are the way to communicate between goroutines
// 3. Send to channel: ch <- value
// 4. Receive from channel: value := <-ch
// 5. Always close channels when done sending (only sender should close)
// 6. Range on channels waits until closed
// 7. Unbuffered channels block until both sides are ready
// 8. Buffered channels allow sending without immediate receiver
// 9. Select statement lets you wait on multiple channel operations
// 10. WaitGroup synchronizes goroutines - crucial for cleanup
// 11. Use context for cancellation and timeouts (advanced)
// 12. Avoid goroutine leaks - always ensure they terminate
// 13. Don't share memory; communicate through channels
// 14. Channel direction: <-chan (receive-only), chan<- (send-only)
// 15. nil channel blocks forever - be careful with channel initialization
// 16. Close a closed channel = panic
// 17. Send on closed channel = panic
// 18. Receive on closed channel = zero value + false
