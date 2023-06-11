package common

import "log"

const (
	DBTypeRestaurant = 1
	DBTypeUser       = 2
)

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recovered from panic:", err)
	}
}
