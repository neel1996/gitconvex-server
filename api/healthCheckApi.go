package api

import (
	"fmt"
	"runtime"
)

func HealthCheckApi(){
	platform := runtime.GOOS
	fmt.Printf("The OS is : %v", platform)
}
