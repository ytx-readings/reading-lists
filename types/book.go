package types

type BookDownloadURLs map[BookFormat]URL // key: format, value: download link
type BookFormat string
type URL string

type Book struct {
	Title        string
	Authors      []string
	Year         int
	DownloadURLs BookDownloadURLs
	GitHubRepo   string
}

const ( // book formats
	AZW  = "AZW"
	AZW3 = "AZW3"
	EPUB = "EPUB"
	FB2  = "FB2"
	MOBI = "MOBI"
	PDF  = "PDF"
)

type BooksOfYear []Book
type BooksByYearMap map[int]BooksOfYear
