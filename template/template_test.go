package template

import "testing"

func TestHTTPDownloade(t *testing.T) {
	var downloader Downloader = NewHTTPDownloader()
	downloader.Download("http://example.com/abc.zip")

	t.Log("test http download")
}

func TestFTPDownloader(t *testing.T) {
	var downloader Downloader = NewFTPDownloader()
	downloader.Download("ftp://example.com/abc.zip")
	t.Log("test FTP download")
}
