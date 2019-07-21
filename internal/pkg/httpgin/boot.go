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

	"github.com/gin-gonic/gin/binding"
	"github.com/mmpereira2github/financial-go/internal/pkg/config"
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

// handleServiceInvocation identifies and invokes the service specified in URL
func handleServiceInvocation(c *gin.Context) {
	debugEnabled := c.GetHeader("debug") != ""
	servicename := c.Param("servicename")
	debugMap := make(map[string]string)
	if debugEnabled {
		debugMap["debug.servicename"] = servicename
	}

	if entry := services.Manager.GetServiceEntryById(servicename); entry != nil {
		var err error
		var input = entry.InputFactory()

		if debugEnabled {
			err = c.ShouldBindBodyWith(input, binding.JSON)

			if buf, ok := c.Get(gin.BodyBytesKey); ok {
				jsonAsString := string(buf.([]byte)[0:])
				debugMap["debug.input"] = jsonAsString
			}
		} else {
			err = c.ShouldBindJSON(input)
		}

		if err == nil {
			status, output := entry.Invoke(input)
			if status.Code == 0 {
				c.JSON(http.StatusOK, output)
			} else {
				debugMap["status"] = fmt.Sprintln(status)
				c.JSON(http.StatusInternalServerError, debugMap)
			}
		} else {
			debugMap["error"] = fmt.Sprintln(err)
			c.JSON(http.StatusBadRequest, debugMap)
		}
	} else {
		status := services.Status{
			Code:  services.ServiceNotFound,
			Error: fmt.Errorf("Service %s not found", servicename),
		}
		log.Println(status)
		c.JSON(http.StatusNotFound, status)
	}
}

// Boot boots the GIN framework to handle HTTP requests
func Boot() *gin.Engine {
	engine := gin.Default()

	engine.POST(buildAPIPath(), func(c *gin.Context) { handleServiceInvocation(c) })

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

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	go func() {
		// Wait for interrupt signal to gracefully shutdown the server with
		// a timeout of 5 seconds.
		quit := make(chan os.Signal)
		// kill (no param) default send syscall.SIGTERM
		// kill -2 is syscall.SIGINT
		// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		Shutdown(false)
	}()

	return engine
}
