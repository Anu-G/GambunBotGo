package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"time"
)

const pool = 100

func Split() []string {
	var res []string
	splitter := strings.Split("$tl jpid text", " ")
	for i := range splitter {
		res = append(res, splitter[i])
	}

	return res
}

func Test_Split(t *testing.T) {
	exp := ""
	res := Split()

	if !reflect.DeepEqual(exp, res) {
		t.Errorf("res=%v", res[2])
	}
}

func Gacha(total_gacha, rate, bulk_draw int) int {
	rand.Seed(int64(time.Now().Nanosecond()))

	shuffle := rand.Perm(pool)
	ssr := make([]int, rate)
	for i := range ssr {
		ssr[i] = shuffle[i]
	}

	lucky_hit := 0
	for x := 0; x < total_gacha; x++ {
		pull_result := make([]int, bulk_draw)
		for i := range pull_result {
			pull_result[i] = rand.Intn(pool)
		}

		for y := range ssr {
			for z := range pull_result {
				if ssr[y] == pull_result[z] {
					lucky_hit += 1
				}
			}
		}
	}

	return int(lucky_hit)
}

func Test_Gacha(t *testing.T) {
	exp := 0
	res := Gacha(100, 3, 10)

	var arrMax []int
	var arrNorm []int
	var arrMin []int
	for i := 0; i <= 100; i++ {
		gacha := Gacha(100, 3, 10)
		if gacha >= 34 {
			arrMax = append(arrMax, gacha)
		} else if gacha < 28 {
			arrMin = append(arrMin, gacha)
		} else {
			arrNorm = append(arrNorm, gacha)
		}
	}
	fmt.Println(len(arrMax), len(arrNorm), len(arrMin))

	if !reflect.DeepEqual(exp, res) {
		t.Errorf("res=%v", res)
	}
}
