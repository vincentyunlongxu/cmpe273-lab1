package sleep

import "time"

func sleep(time int){
	<-time.After(time.Second * time.Duration(time))
}
