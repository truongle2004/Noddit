package proxy

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReserveProxy(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		targetUrl := target + c.Request.RequestURI

		// create new request
		req, err := http.NewRequest(c.Request.Method, targetUrl, c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create request",
			})
			return
		}

		// copy headers
		for key, values := range c.Request.Header {
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}

		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to make request",
			})
			return
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Println("Failed to close response body:", err)
			}
		}(resp.Body)

		// Copy response header
		for k, v := range resp.Header {
			c.Writer.Header()[k] = v
		}

		c.Status(resp.StatusCode)
		io.Copy(c.Writer, resp.Body)
	}
}
