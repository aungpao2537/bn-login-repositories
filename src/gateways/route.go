package gateways

import (
	"github.com/labstack/echo/v4"
)

// func GatewaySpeaker(gateway HTTPGateway, c *echo.Echo) {
// 	s := c.Group("/db/marketplace")

// 	s.GET("/get_all_speakers", gateway.GAllSpeakers)
// 	s.GET("/set_all_speakers", gateway.SAllSpeakers)
// 	s.GET("/get_private_speakers", gateway.GPrivateSpeakers)
// 	s.GET("/set_private_speakers", gateway.SPrivateSpeakers)
// 	s.GET("/get_public_speakers", gateway.GPublicSpeakers)
// 	s.GET("/set_public_speakers", gateway.SPublicSpeakers)
// }

// func GetewayFirebase(gateway HTTPGateway, c *echo.Echo){
// 	s := c.Group("/firebase_auth")
// 	s.GET("/firebase_auth2",gateway.FireBaseService)

// }

func GetewayLogin(gateway HTTPGateway, c *echo.Echo) {
	s := c.Group("/api/dashboard")
	// s.GET("/gettest", gateway.GAllSpeakers)

}
