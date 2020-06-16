package api

import (
	"fmt"
	"testing"
)


func TestRequest(t *testing.T) {

	a := API{}
	fmt.Println(a.USERS())

}
