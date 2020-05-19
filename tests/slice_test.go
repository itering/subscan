package tests

import (
	"github.com/itering/subscan/util"
	"testing"
)

var (
	str = "subscan"
	sst = []string{"subscan0", "subscan", "subscan1"}
	ssf = []string{"subscan0", "subscan1", "subscan2"}
	ssm = map[string]bool{
		"subscan0": true,
		"subscan1": true,
		"subscan2": true,
	}
	ns  = []int{0, 1, 2, 3, 4, 5, 6}
	nsd = []int{6, 5, 4, 3, 2, 1, 0}
)

func TestLookup(t *testing.T) {
	if util.StringInSlice(str, sst) == false {
		t.Errorf(
			"Lookup string in string slice failed, got %v, want %v",
			false,
			true,
		)
	}

	if util.StringInSlice(str, ssf) == true {
		t.Errorf(
			"Lookup string in string slice failed, got %v, want %v",
			true,
			false,
		)
	}
}

func TestContinuous(t *testing.T) {
	rns := util.ContinuousSlice(0, 7, "desc")
	rnsd := util.ContinuousSlice(0, 7, "sced")

	for i, _ := range rns {
		if rns[i] != ns[i] {
			t.Errorf(
				"Map string to string slice failed #%d, got %v, want %v",
				i,
				rns[i],
				ns[i],
			)
		}
	}

	for i, _ := range rnsd {
		if rnsd[i] != nsd[i] {
			t.Errorf(
				"Map string to string slice failed #%d, got %v, want %v",
				i,
				rnsd[i],
				nsd[i],
			)
		}
	}
}

func TestMap(t *testing.T) {
	res := util.MapStringToSlice(ssm)

	for i, _ := range res {
		if res[i] != ssf[i] {
			t.Errorf(
				"Map string to string slice failed #%d, got %v, want %v",
				i,
				res[i],
				ssf[i],
			)
		}
	}
}