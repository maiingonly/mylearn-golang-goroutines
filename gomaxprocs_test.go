package mylearngolanggoroutines

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGetGomaxprocs(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println("total CPU", totalCpu)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total thread", totalThread)

}
