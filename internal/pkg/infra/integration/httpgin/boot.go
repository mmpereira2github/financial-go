package httpgin

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/mmpereira2github/financial-go/internal/pkg/config"
	"github.com/mmpereira2github/financial-go/internal/pkg/infra/integration"
	"github.com/mmpereira2github/financial-go/internal/pkg/services"

	"github.com/gin-gonic/gin"
)

// Shutdown GIN http server
var Shutdown = func(bool) {
	log.Printf("No GIN HTTP Server started to shutdown")
}

var engine *gin.Engine

func getHostAndPort() string {
	port := config.Config.HTTPServer.Port
	if port <= 0 {
		port = 8080
	}
	return fmt.Sprintf(":%d", port)
}

func trim(s string) string { return strings.Trim(s, " \t") }

func buildAPIPath() string {
	apiPath := trim(config.Config.HTTPServer.APIPath)
	if apiPath == "" {
		apiPath = "/financial/api/:servicename"
	} else {
		apiPath = apiPath + "/:servicename"
	}
	return apiPath
}

func isContentTypeJSON(c *gin.Context) bool {
	return c.GetHeader("content-type") == "application/json"
}

// handleServiceInvocation identifies and invokes the service specified in URL
func handleAPIPath(c *gin.Context) {
	if isContentTypeJSON(c) {
		jsonAPIPathHandler(c)
	} else {
		status := &services.Status{
			Code:  services.InvalidInput,
			Error: fmt.Errorf("content-type=%s not supported", c.GetHeader("content-type")),
		}
		c.JSON(http.StatusBadRequest, integration.ErrorDetail{Status: status, ServiceContext: nil})
	}
}

func startListeningSignals() {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	Shutdown(false)
}

func startHTTPServer(e *gin.Engine) {
	srv := &http.Server{
		Addr:    getHostAndPort(),
		Handler: engine,
	}

	Shutdown = func(immediate bool) {
		var ctx context.Context
		var cancel context.CancelFunc
		if immediate {
			log.Println("Shutting down Server immediately...")
			ctx, cancel = context.WithCancel(context.Background())
			cancel()
		} else {
			log.Println("Shutdown Server ...")
			ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
		}
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}
		// catching ctx.Done(). timeout of 5 seconds.
		select {
		case <-ctx.Done():
			if !immediate {
				log.Println("timeout of 5 seconds.")
			}
		}
		log.Println("Server exiting")
	}

	// service connections
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

// Boot boots the GIN framework to handle HTTP requests
func Boot() *gin.Engine {
	engine := gin.Default()

	engine.POST(buildAPIPath(), func(c *gin.Context) { handleAPIPath(c) })

	go startHTTPServer(engine)

	go startListeningSignals()

	return engine
}
