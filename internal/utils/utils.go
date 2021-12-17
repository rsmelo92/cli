package utils

import (
	"fmt"

	"github.com/enescakir/emoji"
)

func OutputSuccess(value string)  {
	fmt.Printf("%v %s\n", emoji.Parse(":white_check_mark:"), value)
}
