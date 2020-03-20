package command

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

func init() {
	rootCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Get access to REST API",
	Run:   googleToken,
}

var config = &oauth2.Config{
	RedirectURL:  "http://localhost:8080/callback",
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/gmail.readonly"},
	Endpoint:     google.Endpoint,
}

func googleToken(cmd *cobra.Command, args []string) {
	m := http.NewServeMux()
	s := http.Server{Addr: ":8080", Handler: m}
	m.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		dir, _ := os.Getwd()
		openbrowser(fmt.Sprintf("file://%s/success.html", dir))
		token, _ := config.Exchange(oauth2.NoContext, r.FormValue("code"))
		viper.Set("token", token.AccessToken)
		viper.WriteConfig()
		log.Print("Token updated")
		s.Shutdown(context.Background())

	})
	url := config.AuthCodeURL("pseudo-random")

	openbrowser(url)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
