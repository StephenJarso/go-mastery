package concurrency

import (
	"fmt"
	"time"
)

// Channels are the pipes that connect concurrent goroutines.
// You can send values into channels from one goroutine and receive those values
// in another goroutine, providing both communication and synchronization.

// UnbufferedChannels demonstrates channels without capacity.
// Sends and receives block until both sender and receiver are ready.
func UnbufferedChannels() {
	fmt.Println("=== Unbuffered Channels ===")

	// Create an unbuffered channel of strings using make(chan Type)
	ch := make(chan string)

	// Launch a goroutine that sends a message.
	go func() {
		fmt.Println("Sender: Preparing to send message...")
		// This send blocks until the receiver is ready to read from ch.
		ch <- "Hello from Sender!"
		fmt.Println("Sender: Message sent successfully!")
	}()

	// Wait a moment in main before receiving to demonstrate blocking behavior.
	time.Sleep(50 * time.Millisecond)

	fmt.Println("Receiver: Ready to receive...")
	// This receive blocks until a sender sends a message to ch.
	msg := <-ch
	fmt.Printf("Receiver: Received message: %q\n", msg)
}

// BufferedChannels demonstrates channels with capacity.
// Sends are non-blocking until the buffer is full.
// Receives are non-blocking until the buffer is empty.
func BufferedChannels() {
	fmt.Println("\n=== Buffered Channels ===")

	// Create a buffered channel of integers with a capacity of 3.
	ch := make(chan int, 3)

	// We can send up to 3 values without blocking, even without a receiver.
	ch <- 10
	ch <- 20
	ch <- 30

	fmt.Println("Sent 3 values to buffered channel (capacity 3) without blocking.")

	// Sending a 4th value would block here because the buffer is full.
	// We'll spawn a receiver to empty the channel in the background.
	go func() {
		time.Sleep(50 * time.Millisecond)
		fmt.Printf("Async Receiver: Received %d\n", <-ch)
	}()

	fmt.Println("Sending 4th value (will block briefly)...")
	ch <- 40
	fmt.Println("Sent 4th value successfully!")

	// Empty the remaining items.
	fmt.Printf("Receiver: Received %d\n", <-ch)
	fmt.Printf("Receiver: Received %d\n", <-ch)
	fmt.Printf("Receiver: Received %d\n", <-ch)
}

// ChannelClosingAndRange demonstrates closing a channel and reading until it's closed.
func ChannelClosingAndRange() {
	fmt.Println("\n=== Channel Closing and Range ===")

	jobs := make(chan int, 5)

	// Producer goroutine
	go func() {
		for i := 1; i <= 5; i++ {
			jobs <- i
			fmt.Printf("Producer: Sent job %d\n", i)
		}
		// Closing a channel notifies receivers that no more values will be sent.
		// Only the sender should close a channel, never the receiver.
		// Sending on a closed channel causes a panic.
		close(jobs)
		fmt.Println("Producer: Closed the channel.")
	}()

	// Consumer using 'range' loop
	// 'range' automatically terminates when the channel is closed and empty.
	fmt.Println("Consumer: Starting to process jobs...")
	for job := range jobs {
		fmt.Printf("Consumer: Received job %d\n", job)
		time.Sleep(5 * time.Millisecond)
	}
	fmt.Println("Consumer: Finished processing all jobs.")

	// Comma-ok syntax to check if a channel is open
	// Reading from a closed channel returns the zero value and false.
	val, ok := <-jobs
	fmt.Printf("Reading from closed channel: val = %d, ok = %t\n", val, ok)
}

func RunChannelsDemo() {
	UnbufferedChannels()
	BufferedChannels()
	ChannelClosingAndRange()
}
