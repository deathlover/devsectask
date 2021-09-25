package main

import (
	"fmt"

	"github.com/deteact_devsecops/client"
)

func main() {
	job := prepareJob()
	job.GetList()
	job.GetAssetsID()
	job.GetTranshLinks()

	fmt.Printf("0x001 AssetsID runner starting...\n\n")
	job.ExecAssetsURL()

	fmt.Printf("0x002 TransactionID runner starting...\n\n")
	job.ExecTranshURL()

	fmt.Printf("0x003 Stupid Attack Start...\n\n")
	job.StupidAttacker()

	fmt.Printf("0x004 Mind Attack Start...\n\n")
	job.MindAttacker()
}

func prepareJob() *client.Job {
	runner := client.NewJob()
	return runner
}
