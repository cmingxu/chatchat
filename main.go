package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-analyze/charts"

	tele "gopkg.in/telebot.v4"
)

const JettonRateUrl = "https://tonapi.io/v2/rates/chart?token=TOKEN&currency=usd&points_count=100&start_date=START_DATE&end_date=END_DATE"

func main() {
	pref := tele.Settings{
		Token:   os.Getenv("TOKEN"),
		Poller:  &tele.LongPoller{Timeout: 10 * time.Second},
		Verbose: true,
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(c tele.Context) error {
		return c.Send(c.Message().Payload)
	})

	b.Handle("/get-image", func(c tele.Context) error {
		return c.Send("Acc")
	})

	b.Handle("/price", func(c tele.Context) error {
		jetton := c.Message().Payload
		if jetton == "" {
			return c.Send("Please provide a jetton address.")
		}

		url := strings.Replace(JettonRateUrl, "TOKEN", jetton, -1)
		url = strings.Replace(url, "START_DATE", fmt.Sprintf("%d", time.Now().Add(-10*time.Minute).Unix()), -1)
		url = strings.Replace(url, "END_DATE", fmt.Sprintf("%d", time.Now().Unix()), -1)

		resp, err := http.Get(url)
		if err != nil {
			c.Send("Unable to fetch price data.")
		}

		fmt.Println("Fetching price data from:", url)
		fmt.Println(resp)

		values := [][]float64{
			{120, 132, 101, charts.GetNullValue(), 90, 230, 210},
			{220, 182, 191, 234, 290, 330, 310},
			{150, 232, 201, 154, 190, 330, 410},
			{320, 332, 301, 334, 390, 330, 320},
			{820, 932, 901, 934, 1290, 1330, 1320},
		}

		opt := charts.NewLineChartOptionWithData(values)
		opt.Title.Text = "Line"
		opt.Title.FontStyle.FontSize = 16
		opt.XAxis.Labels = []string{
			"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun",
		}

		opt.Legend.SeriesNames = []string{
			"Email", "Union Ads", "Video Ads", "Direct", "Search Engine",
		}
		opt.Legend.Padding = charts.Box{
			Left: 100,
		}
		opt.Symbol = charts.SymbolCircle
		opt.LineStrokeWidth = 1.2

		p := charts.NewPainter(charts.PainterOptions{
			OutputFormat: charts.ChartOutputPNG,
			Width:        600,
			Height:       400,
		})

		if err := p.LineChart(opt); err != nil {
			fmt.Println("Error generating chart:", err)
			return c.Send("Error generating chart: ")
		}

		buf, err := p.Bytes()
		if err != nil {
			fmt.Println("Error generating bytes from chart:", err)
			return c.Send("Error " + err.Error())
		}

		tmpPath := "./tmp"
		if err := os.MkdirAll(tmpPath, 0700); err != nil {
			fmt.Println("Error creating tmp directory:", err)
			return c.Send("Error " + err.Error())
		}

		file := filepath.Join(tmpPath, "line-chart-1-basic.png")
		if err := os.WriteFile(file, buf, 0600); err != nil {
			fmt.Println("Error writing file:", err)
			return c.Send("Error " + err.Error())
		}

		photo := tele.Photo{
			File: tele.FromDisk("./tmp/line-chart-1-basic.png"),
		}

		return c.SendAlbum(tele.Album{&photo})
	})

	b.Start()
}
