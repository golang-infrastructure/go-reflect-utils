package reflect_utils

import "testing"

func Test(t *testing.T) {

}

func TestIsZero(t *testing.T) {

	var a interface{}
	t.Log(IsZero(a))
	t.Log(IsZero(0))
	t.Log(IsZero(1))

}
