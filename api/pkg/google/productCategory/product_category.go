package productCategory

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

type Category struct {
	Id   int64
	Name string
	Path []string
}

func GetProductCategories(r io.Reader) []Category {
	lineScanner := bufio.NewScanner(r)
	out := []Category{}
	for lineScanner.Scan() {
		line := lineScanner.Text()
		if line[0] == '#' {
			continue
		}
		sections := strings.Split(line, " - ")
		idRaw := sections[0]
		categoryNames := sections[1]

		id, err := strconv.ParseInt(idRaw, 10, 64)
		if err != nil {
			log.Default().Printf("Error: failed to parse categoryId: %s", err)
		}

		categories := strings.Split(categoryNames, " > ")

		category := Category{
			Id:   id,
			Name: categories[len(categories)-1],
			Path: categories[:len(categories)-1],
		}
		out = append(out, category)
	}
	return out
}
