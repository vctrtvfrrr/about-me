package info

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

// Data ...
type Data struct {
	Info   Info    `json:"info"`
	Badges []Badge `type:"badges"`
}

// Info ...
type Info struct {
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	Location string `json:"location"`
}

// Badge ...
type Badge struct {
	Label     string `type:"label"`
	Value     string `type:"value"`
	Color     string `type:"color"`
	Link      string `type:"link"`
	Linebreak bool   `type:"linebreak"`
}

func getAge(birthday string) string {
	now := time.Now()

	birthDate, err := time.Parse("2006-01", birthday)
	if err != nil {
		fmt.Println(err)
	}

	// Get the year number change since the player's birth.
	years := now.Year() - birthDate.Year()

	// If the date is before the date of birth, then not that many years have elapsed.
	birthDay := getAdjustedBirthDay(birthDate, now)
	if now.YearDay() < birthDay {
		years--
	}

	return strconv.Itoa(years) + " anos"
}

// Gets the adjusted date of birth to work around leap year differences.
func getAdjustedBirthDay(birthDate time.Time, now time.Time) int {
	birthDay := birthDate.YearDay()
	currentDay := now.YearDay()
	if isLeap(birthDate) && !isLeap(now) && birthDay >= 60 {
		return birthDay - 1
	}
	if isLeap(now) && !isLeap(birthDate) && currentDay >= 60 {
		return birthDay + 1
	}
	return birthDay
}

// Works out if a time.Time is in a leap year.
func isLeap(date time.Time) bool {
	year := date.Year()
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	}
	return false
}

func generateInfoList(data Info) string {
	infolist := []string{
		fmt.Sprintf("- **Nome:** %s", data.Name),
		fmt.Sprintf("- **Idade:** %s", getAge(data.Birthday)),
		fmt.Sprintf("- **Localização:** %s", data.Location),
	}

	return strings.Join(infolist[:], "\n")
}

func generateBadgesList(badges []Badge) string {
	var fbadges []string

	for _, b := range badges {
		badge := fmt.Sprintf(
			"[![](https://img.shields.io/badge/%s-%s-%s)](%s)",
			b.Label,
			b.Value,
			b.Color,
			b.Link,
		)

		if b.Linebreak {
			badge += "  \n"
		} else {
			badge += " "
		}

		fbadges = append(fbadges, badge)
	}

	return strings.Join(fbadges[:], "")
}

// GetInfo returns the formated profile informations.
func GetInfo() string {
	jsonFile, err := os.Open("data/info.json")
	defer jsonFile.Close()
	if err != nil {
		return ""
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return ""
	}

	var data Data
	json.Unmarshal(byteValue, &data)

	return fmt.Sprintf(
		"## Informações\n\n%s\n\n%s",
		generateInfoList(data.Info),
		generateBadgesList(data.Badges),
	)
}
