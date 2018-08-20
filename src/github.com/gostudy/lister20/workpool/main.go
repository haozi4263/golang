package main

import (
	"fmt"
	"math/rand"
)

type Job struct {
	Number int
	Id int
}
type Result struct {
	job *Job
	sum int
}

func calc(job *Job, result chan *Result)  {
	var sum int
	number := job.Number
	for number !=0{
		tmp := number % 10  //得到个位数
		sum += tmp
		number /= 10  //获取十位数
	}

	r := &Result {
		job: job,
		sum: sum,
	}

	result <- r
}

func Worker(jobChan chan *Job, resultChan chan *Result)  {
	for job := range jobChan {
		calc(job,resultChan)
	}
}

func startWorkPool(num int, jobChan chan *Job, resultChan chan *Result)  {
	for i := 0; i < num; i++ {
		go Worker(jobChan,resultChan)
	}
}

func main()  {
	jobChan := make(chan *Job, 1000)
	resultChan :=make(chan *Result, 1000)

	startWorkPool(128, jobChan,resultChan)
	for i :=0; i < 128; i++ {
		go calc()
	}
	for {
		number := rand.int32()
		job := Job{

		}

	}
}