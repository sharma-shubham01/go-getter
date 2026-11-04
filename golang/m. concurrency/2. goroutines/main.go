package main

// Threads
// 1. Thread is a concept in computer science to achieve concurrency
// 2. A thread is a single sequence of instructions that can be executed concurrently with other threads
// 3. Example - a single cpu core is a worker and threads are like individual tasks assigned to that worker
// 4. CPU switch between these tasks rapidly creating an illusion of running them simultaneously
// 5. Threads are unit of execution
// 6. Threads belong to a process
// 7. Threads are managed by OS - it creates, destroys and schedules threads

// Green Threads
// 1. Green threads are type of threads that are managed by the program runtime environment like VM instead of OS
// 2. Green threads are managed by user-level programs by the runtime
// 3. Green threads are lightweight compared to threads
// 4. This allows us to create many more without any resource overhead
// 5. Green threads are cooperative - they rely on the program to decide when to switch between them

// Go Routines
// 1. Go routines are cooperative
// 2. Go routines are lightweight
// 3. Go routines help with the structure of the program
// 4. There is atleast one go routine for every program
// 5. The keyword go can be used to start another go routine, it accepts a function to run concurrently
// 6. Go routines are managed by go runtime

// Go Scheduler
// 1. responsible for managing go routines
