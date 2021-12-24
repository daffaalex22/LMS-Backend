package konversi

import (
	"backend/helper/err"
	"strconv"
)

func StringToUint(str string) (uint, error) {
	convInt, errConvert := strconv.Atoi(str)
	if errConvert != nil {
		return 0, err.ErrConvertId
	}
	return uint(convInt), nil
}
