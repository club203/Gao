package cron

import (
	"time"

	"nightingale-club203/src/models"
)

const cleanerInterval = 3600 * time.Second

func CleanerLoop() {
	tc := time.Tick(cleanerInterval)

	for {
		models.AuthState{}.CleanUp()
		models.Captcha{}.CleanUp()
		<-tc
	}
}
