package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

func TestGenerate(t *testing.T) {
	buf, err := os.ReadFile("shifts.csv")
	if err != nil {
		panic(err.Error())
	}
	rows := strings.Split(string(buf), "\n")
	startDate, err := time.Parse(time.DateOnly, "2024-02-05")
	if err != nil {
		panic(err.Error())
	}
	finishDate, err := time.Parse(time.DateOnly, "2024-02-11")
	if err != nil {
		panic(err.Error())
	}
	deltaDays := int(finishDate.Sub(startDate).Hours() / 24)
	_ = deltaDays
	for _, row := range rows {
		//for i := 0; i < deltaDays; i++ {
		for iDay := startDate; finishDate.Sub(iDay).Hours() > 0; iDay = iDay.Add(24 * time.Hour) {
			items := strings.Split(row, ",")
			times := strings.Split(items[4], "-")
			tFrom, err := time.Parse(time.TimeOnly, times[0]+":00")
			if err != nil {
				panic(err.Error())
			}
			tTo, err := time.Parse(time.TimeOnly, times[1]+":00")
			if err != nil {
				panic(err.Error())
			}

			from := iDay.Add(time.Duration(tFrom.Hour()-3) * time.Hour)
			to := iDay.Add(time.Duration(tTo.Hour()-3) * time.Hour)
			res := iDay.Format(time.DateOnly) + "," + from.Format(time.RFC3339) + "," + to.Format(time.RFC3339) + "," + items[1] + "," + items[2] + "," + items[3] + "," + items[0] + "," + "Flex"
			//fmt.Println(iDay.Format(time.DateOnly), ",", from.Format(time.RFC3339), ",", to.Format(time.RFC3339), items[1], ",", items[1], ",", items[2], ",", items[3], ",", items[0])
			fmt.Println(res)
		}
	}
}

// date, t_from, t_to, size, lat, lon, address, polygon

/*2024-01-16,2024-01-16 07:00:00.000000,2024-01-16 15:00:00.000000,1,55.792809,37.586104,Савеловская,Flex
2024-01-17,2024-01-17 07:00:00.000000,2024-01-17 15:00:00.000000,1,55.792809,37.586104,Савеловская,Flex
2024-01-18,2024-01-18 07:00:00.000000,2024-01-18 15:00:00.000000,1,55.792809,37.586104,Савеловская,Flex*/
