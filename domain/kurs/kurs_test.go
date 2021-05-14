package kurs_test

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/wendylau87/xfers2021/domain/kurs"
	"github.com/wendylau87/xfers2021/entities"
	"github.com/wendylau87/xfers2021/infrastructure"
	"github.com/wendylau87/xfers2021/infrastructure/logger"
	"github.com/wendylau87/xfers2021/infrastructure/sqlhandler"
	"os"
	"path"
	"runtime"
	"testing"
)

var(
	dom kurs.DomainItf
)


func TestMain(t *testing.M){
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..", "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}

	log := logger.NewLogger()
	infrastructure.Load(*log)
	sql, err := sqlhandler.NewSQLHandler(*log)
	if err != nil{
		log.LogError("Failed initiated database : %s", err)
	}else{
		dom = kurs.InitKursDomain(*log, sql)

		exitVal := t.Run()
		os.Exit(exitVal)
	}

}

func TestDomain_GetKurs(t *testing.T) {
	Convey("Get Kurs", t, func() {
		testCases := []struct {
			testID   int
			testType string
			testDesc string
			startDate string
			endDate string
		}{
			{
				testID:   1,
				testDesc: "success get kurs",
				testType: "P",
				startDate: "2021-05-13",
				endDate: "2021-05-14",
			},
		}
		for _, tc := range testCases {
			Convey(fmt.Sprintf("%d - [%s] : %s", tc.testID, tc.testType, tc.testDesc), func() {
				_, err := dom.GetKursByDate(tc.startDate, tc.endDate)
				if tc.testType == "P" {
					So(err, ShouldBeNil)


				} else {
					So(err, ShouldNotBeNil)
				}
			})
		}
	})
}

func TestDomain_GetKursWithName(t *testing.T) {
	Convey("Get Kurs with date", t, func() {
		testCases := []struct {
			testID   int
			testType string
			testDesc string
			startDate string
			endDate string
			name string
		}{
			{
				testID:   1,
				testDesc: "success get kurs with name",
				testType: "P",
				startDate: "2021-05-13",
				endDate: "2021-05-14",
				name:"USD",
			},
		}
		for _, tc := range testCases {
			Convey(fmt.Sprintf("%d - [%s] : %s", tc.testID, tc.testType, tc.testDesc), func() {
				_, err, _ := dom.GetKursByName(tc.name, tc.startDate, tc.endDate)
				if tc.testType == "P" {
					So(err, ShouldBeNil)
				} else {
					So(err, ShouldNotBeNil)
				}
			})
		}
	})
}

func TestDomain_CreateKurs(t *testing.T) {
	Convey("create Kurs", t, func() {
		testCases := []struct {
			testID   int
			testType string
			testDesc string
			payload  entities.CreateKurs
		}{
			{
				testID:   1,
				testDesc: "success create",
				testType: "P",
				payload: entities.CreateKurs{
					Name:  "SGD",
					ValidDate: "2021-05-14",
					ERate: entities.CreateERate{
						Buy:  10180,
						Sell: 10280,
					},
					TTCounter: entities.CreateTTCounter{
						Buy:  10181,
						Sell: 10281,
					},
					BankNote: entities.CreateBankNote{
						Buy:  10182,
						Sell: 10282,
					},
				},
			},
		}
		for _, tc := range testCases {
			Convey(fmt.Sprintf("%d - [%s] : %s", tc.testID, tc.testType, tc.testDesc), func() {
				_, err := dom.CreateKurs(tc.payload)
				if tc.testType == "P" {
					So(err, ShouldBeNil)

				} else {
					So(err, ShouldNotBeNil)
				}
			})
		}
	})
}

func TestDomain_UpdateKurs(t *testing.T) {
	Convey("update Kurs", t, func() {
		testCases := []struct {
			testID   int
			testType string
			testDesc string
			payload  entities.PutKurs
		}{
			{
				testID:   1,
				testDesc: "success update",
				testType: "P",
				payload: entities.PutKurs{
					Name:  "SGD",
					ValidDate: "2021-05-14",
					ERate: entities.CreateERate{
						Buy:  11180,
						Sell: 11280,
					},
					TTCounter: entities.CreateTTCounter{
						Buy:  11181,
						Sell: 11281,
					},
					BankNote: entities.CreateBankNote{
						Buy:  11182,
						Sell: 11282,
					},
				},
			},
		}
		for _, tc := range testCases {
			Convey(fmt.Sprintf("%d - [%s] : %s", tc.testID, tc.testType, tc.testDesc), func() {
				_, err, _ := dom.UpdateKurs(tc.payload)
				if tc.testType == "P" {
					So(err, ShouldBeNil)

				} else {
					So(err, ShouldNotBeNil)
				}
			})
		}
	})
}

func TestDomain_DeleteKurs(t *testing.T) {
	Convey("Get Kurs", t, func() {
		testCases := []struct {
			testID   int
			testType string
			testDesc string
			validDate string
		}{
			{
				testID:   1,
				testDesc: "success delete kurs",
				testType: "P",
				validDate: "2021-05-13",
			},
		}
		for _, tc := range testCases {
			Convey(fmt.Sprintf("%d - [%s] : %s", tc.testID, tc.testType, tc.testDesc), func() {
				err := dom.DeleteKurs(tc.validDate)
				if tc.testType == "P" {
					So(err, ShouldBeNil)
				} else {
					So(err, ShouldNotBeNil)
				}
			})
		}
	})
}