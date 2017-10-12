package admin

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rakyll/statik/fs"
	"github.com/urfave/cli"

	"github.com/petersondmg/vigilant-pi/lib/config"
	libhttp "github.com/petersondmg/vigilant-pi/lib/http"
	// import embedded static files
	_ "github.com/petersondmg/vigilant-pi/statik"
)

// Command ...
func Command() cli.Command {
	return cli.Command{
		Name:        "admin",
		Description: "starts a web admin interface",
		Action: func(c *cli.Context) error {
			server()
			return nil
		},
	}
}

func server() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/api/cameras", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("{}"))
	})
	mux.Handle("/", http.FileServer(statikFS))

	admin := config.Current.Admin
	addr := fmt.Sprintf(":%d", admin.Port)
	log.Printf("Starting server on %s", addr)

	s := &http.Server{
		Addr:           addr,
		Handler:        libhttp.NewAuthMiddleware(mux, admin.User, admin.Pass),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
