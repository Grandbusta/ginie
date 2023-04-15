package lib

import (
	"context"
	"log"
	"strings"

	"github.com/chromedp/chromedp"
)

func ScrapeFootball() {
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
	defer cancel()

	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.skysports.com/football"),
		chromedp.WaitVisible(".block-header__title"),
		chromedp.Text(".news-list block", &res, chromedp.NodeVisible),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(strings.TrimSpace(res))
}
