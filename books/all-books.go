package books

import "github.com/ytx-readings/reading-lists/types"

var (
	Book_RZJX = &types.Book{
		Title:   "认知觉醒：开启自我改变的原动力",
		Authors: []string{"周岭"},
		Year:    2020,
		DownloadURLs: types.BookDownloadURLs{
			types.EPUB: "https://bit.ly/rzjx",
		},
		GitHubRepo: "https://github.com/ytx-readings/zl-books",
	}

	Book_RZJX_Q = &types.Book{
		Title:   "认知觉醒：伴随一生的学习方法论（青少年学习版）",
		Authors: []string{"周岭"},
		Year:    2022,
		DownloadURLs: types.BookDownloadURLs{
			types.EPUB: "https://bit.ly/rzjx-q",
		},
		GitHubRepo: "https://github.com/ytx-readings/zl-books",
	}

	Book_RZQD = &types.Book{
		Title:   "认知驱动：做成一件对他人很有用的事",
		Authors: []string{"周岭"},
		Year:    2021,
		DownloadURLs: types.BookDownloadURLs{
			types.EPUB: "https://bit.ly/rzqd-epub",
		},
		GitHubRepo: "https://github.com/ytx-readings/zl-books",
	}
)
