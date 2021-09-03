package api

import (
	"birthdayGiftVoucher/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service
type Server struct {
	router  *gin.Engine
	svc service.Service
}

func NewServer() *Server {
	router := gin.Default()


	server := &Server{
		router: router,
		svc: service.NewBirthdayService(),
	}

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "home page")
	})

	router.GET("/view", server.GetGiftView)

	router.GET("/gift", server.GetGift)
	return server
}

// Start runs the HTTP server on a specific address.
func (svr *Server) Start(address string) error {
	return svr.router.Run(address)
}

func (svr *Server) GetGiftView(ctx *gin.Context) {

	pageContent := svr.svc.GiftView()

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(pageContent))
}

func (svr *Server) GetGift(ctx *gin.Context) {

	displayMsg := ""
	m := ctx.Request.URL.Query()
	inpName := m.Get("firstname")
	if inpName != "" {
		voucher, err := svr.svc.GetVoucher(inpName)
		if err != nil {
			displayMsg = "sorry! this gift is not for you"
		} else {
			displayMsg = voucher
		}
	} else {
		displayMsg = "do not play here"
	}

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(displayMsg))
}