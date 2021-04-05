package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/enescakir/emoji"
	"github.com/mgutz/ansi"
	"github.com/olekukonko/tablewriter"
)

var (
	NoQueryFoundError = errors.New(fmt.Sprint(ansi.Red, emoji.SmilingFaceWithTear, " No query was found!", ansi.Reset))
)

func PrintQueryAsTable(logger *log.Logger, data url.Values) {
	table := tablewriter.NewWriter(logger.Writer())
	headers := []string{"field", "value"}

	table.SetHeader(headers)
	for k, v := range data {
		table.Append([]string{k, strings.Join(v, ",")})
	}

	table.Render()
}

func CreateHandler(ops Options) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if (ops.ShowQuery) || !(ops.ShowBody || ops.ShowHeader) {
			query := r.URL.Query()

			if len(query) <= 0 {
				ops.Logger.Println(NoQueryFoundError.Error())
			} else {
				ops.Logger.Println(emoji.ThumbsUp, ansi.Green+"Query received"+ansi.Reset)
				PrintQueryAsTable(ops.Logger, query)
			}
		}
	})
}
