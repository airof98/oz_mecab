package ozmecab

import (
	"../oz_mecab"
	"fmt"
	"testing"
)

func TestMecabHelloworld(t *testing.T) {
	e := ozmecab.Init()
	if e != nil {
		t.Error(e)
	}
	v, e := ozmecab.Parse("정부가 늘어나는 신규 전력수요를 충당하기 위해 2029년까지 원자력발전소")
	if e != nil {
		t.Error(e)
	}

	defer ozmecab.Fin()
	fmt.Printf("%q\n", v)
}
