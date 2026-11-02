package web

import (
	"fmt"
	"os"
	"strings"

	"github.com/theobori/recueil/image"
	"github.com/theobori/recueil/internal"
)

func CreateRss(metadatas image.Metadatas, destinationPath string, title string, siteUrl string) string {
	b := strings.Builder{}

	b.WriteString("<rss version=\"2.0\" xmlns:atom=\"http://www.w3.org/2005/Atom\">")
	b.WriteString("<channel>")
	fmt.Fprintf(&b, "<atom:link href=\"%s\" rel=\"self\" type=\"application/rss+xml\"/>", strings.TrimSuffix(siteUrl, "/")+"/rss.xml")
	fmt.Fprintf(&b, "<title>%s</title>", title)
	fmt.Fprintf(&b, "<link>%s</link>", siteUrl)
	b.WriteString("<language>en-US</language>")
	b.WriteString("<description>NO</description>")

	for _, metadata := range metadatas {
		b.WriteString("<item>")
		fmt.Fprintf(&b, "<title>%s</title>", metadata.Title())

		metadataUrl := internal.UrlJoin(siteUrl, metadata.Destination.Path.PublicValue)
		fmt.Fprintf(&b, "<link>%s</link>", metadataUrl)
		fmt.Fprintf(
			&b,
			"<description><![CDATA[<img src=\"%s\" data-thumb=\"%s\" alt=\"%s\">]]></description>",
			metadataUrl,
			internal.UrlJoin(siteUrl, metadata.Destination.ThumbnailPath.PublicValue),
			metadata.Description(),
		)
		fmt.Fprintf(&b, "<pubDate>%s</pubDate>", metadata.Date())
		b.WriteString("</item>")
	}

	b.WriteString("</channel>")
	b.WriteString("</rss>")

	return b.String()
}

func ExecuteRss(metadatas image.Metadatas, destinationPath string, title string, siteUrl string) error {
	feedString := CreateRss(metadatas, destinationPath, title, siteUrl)

	f, err := os.Create(destinationPath)
	if err != nil {
		return err
	}

	_, err = f.WriteString(feedString)
	if err != nil {
		return err
	}

	return nil
}
