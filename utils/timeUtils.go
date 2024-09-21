package utils

import (
	"fmt"
	"time"
)

func GetCurrentKoreaTime() time.Time {
	//아래 코드는 앱에서 콘솔 찍어보면 현재시간 잘 찍고있음.
	//그러나
	//Go에서 한국 시간(KST)을 올바르게 설정하고 있어도, MongoDB는 데이터를 UTC 시간으로 변환하여 저장합니다.
	//즉, +09:00 시간대의 한국 시간이 맞지만, MongoDB에 저장될 때는
	//이를 UTC로 변환하여 -09:00 시간만큼 차감된 시간이 저장됩니다.
	// => 이말은 현재시간에 + 9시간으로 찍으면 되는거아닌가

	// location, _ := time.LoadLocation("Asia/Seoul")
	// return time.Now().In(location)

	now := time.Now()
	fmt.Println("now:", now)

	koreaTime := now.Add(9 * time.Hour)
	fmt.Println("koreaTime:", koreaTime)
	return koreaTime
}
