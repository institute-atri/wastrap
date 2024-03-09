package cmd

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

type URLHandler interface {
	HandleWordPressError()
	HandleUrlError(err error)
	HandleOpenUrlError(err error)
}

type defaultUrlHandler struct{}

var urlCmd = &cobra.Command{
	Use:   "url",
	Short: "Open the specified WordPress URL in the default web browser",
	Run: func(cmd *cobra.Command, args []string) {
		handleUrl(cmd, args, &defaultUrlHandler{})
	},
}

func (d *defaultUrlHandler) HandleWordPressError() {
	fmt.Println()
	fmt.Printf("-[Error]: %s\n", formatMessage("This is not a WordPress site. Aborting.", 1))
}
func (d *defaultUrlHandler) HandleUrlError(err error) {
	fmt.Println()
	fmt.Printf("-[Error]: %s %v\n", formatMessage("Checking if the site is based on WordPress: ", 1), err)
}
func (d *defaultUrlHandler) HandleOpenUrlError(err error) {
	fmt.Println()
	fmt.Printf("-[Error]: %s %v\n", formatMessage("Error opening URL:", 1), err)
}

func handleUrl(cmd *cobra.Command, args []string, handler URLHandler) {
	urlStr, _ := cmd.Flags().GetString("url")
	if urlStr == "" && len(args) > 0 {
		urlStr = args[0]
	}
	if urlStr == "" {
		_ = cmd.Help()
		return
	}
	fmt.Printf("Opening URL: %s\n", urlStr)
	loading()

	isWordPress, err := isWordPress(urlStr)
	if err != nil {
		handler.HandleUrlError(err)
		return
	}
	if !isWordPress {
		handler.HandleWordPressError()
		return
	}
	err = openURL(urlStr)
	if err != nil {
		handler.HandleOpenUrlError(err)
		return
	}
}

func isWordPress(urlStr string) (bool, error) {
	pageContent, err := fetchPageContent(urlStr)
	if err != nil {
		return false, err
	}

	wpIdentifiers := []string{"wp-admin", "wp-login", "wp-content"}
	for _, wpID := range wpIdentifiers {
		if strings.Contains(pageContent, wpID) {
			return true, nil
		}
	}
	return false, nil
}

func fetchPageContent(urlStr string) (string, error) {
	resp, err := http.Get(urlStr)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func openURL(urlStr string) error {
	_, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", urlStr)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", urlStr)
	case "linux":
		cmd = exec.Command("xdg-open", urlStr)
	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	return cmd.Run()
}

func init() {
	urlCmd.Flags().StringP("url", "u", "", "WordPress URL")
	rootCmd.AddCommand(urlCmd)
}
