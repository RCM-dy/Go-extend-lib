package Go_extend_lib

import (
	"bufio"
	"os"
)

type TooLongInput struct{}

func (t *TooLongInput) Error() string {
	return "can't read it"
}
func Input(s string) (string, error) {
	ir := bufio.NewReader(os.Stdin)
	print(s)
	i, ok, err := ir.ReadLine()
	if ok {
		return "", &TooLongInput{}
	}
	if err != nil {
		return "", err
	}
	return string(i), nil
}
