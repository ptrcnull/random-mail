package main

import (
	"encoding/csv"
	"github.com/gofiber/fiber"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const minCount = 1000
const domain = "example.tld"
const maxRequestCount = 1000

func readCsv(filename string) ([]string, error) {
	file, err := os.Open("data/" + filename)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(file)
	_, _ = r.Read()

	var records []string

	for {
		record, err := r.Read()
		if err == io.EOF {
			return records, nil
		}
		if err != nil {
			return nil, err
		}
		count, _ := strconv.ParseInt(record[1], 10, 16)
		if count < minCount {
			continue
		}
		records = append(records, depolishify(strings.ToLower(record[0])))
	}
}

func depolishify(str string) string {
	str = strings.ReplaceAll(str, "ą", "a")
	str = strings.ReplaceAll(str, "ć", "c")
	str = strings.ReplaceAll(str, "ę", "e")
	str = strings.ReplaceAll(str, "ś", "s")
	str = strings.ReplaceAll(str, "ń", "n")
	str = strings.ReplaceAll(str, "ó", "o")
	str = strings.ReplaceAll(str, "ł", "l")
	str = strings.ReplaceAll(str, "ź", "z")
	str = strings.ReplaceAll(str, "ż", "z")
	return str
}

func randBool() bool {
	return rand.Uint64()&(1<<63) == 0
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	app := fiber.New()

	surnamesMale, err := readCsv("surnames_male.csv")
	if err != nil {
		panic(err)
	}
	namesMale, err := readCsv("names_male.csv")
	if err != nil {
		panic(err)
	}

	surnamesFemale, err := readCsv("surnames_female.csv")
	if err != nil {
		panic(err)
	}
	namesFemale, err := readCsv("names_female.csv")
	if err != nil {
		panic(err)
	}

	getOne := func(gender string) string {
		gender = string(strings.ToLower(gender)[0])
		var names []string
		var surnames []string
		if gender == "m" || (gender == "r" && randBool()) {
			names = namesMale
			surnames = surnamesMale
		} else {
			names = namesFemale
			surnames = surnamesFemale
		}

		name := names[rand.Intn(len(names))]
		surname := surnames[rand.Intn(len(surnames))]

		sep := []string{"", ".", "-", "_"}[rand.Intn(4)]
		year := strconv.FormatInt(int64(rand.Intn(40)+1960), 10)
		if randBool() {
			year = year[2:]
		}

		return name + sep + surname + year + "@" + domain
	}

	app.Get("/", func(ctx *fiber.Ctx) {
		gender := ctx.Query("gender", "random")
		count64, err := strconv.ParseInt(ctx.Query("count", "1"), 10, 32)
		count := int(count64)

		if err != nil {
			ctx.SendString("error: " + err.Error())
			return
		}

		if count > maxRequestCount {
			ctx.SendString("error: too many names requested")
			return
		}

		sep := ctx.Query("sep", "\n")
		res := ""

		for i := 0; i < count; i++ {
			res += getOne(gender) + sep
		}

		ctx.SendString(res)
	})

	err = app.Listen(":8081")
	if err != nil {
		panic(err)
	}
}
