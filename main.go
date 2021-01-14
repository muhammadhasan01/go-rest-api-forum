// @Version 1.0.0
// @Title Backend API Discussion Forum
// @Description The Backend API which handles RESTful API that provides the backend for a discussion forum with Users, Threads and Posts
// @ContactName Muhammad Hasan
// @ContactEmail muhammad.hasan@pinhome.id
// @LicenseName MIT
// @LicenseURL https://en.wikipedia.org/wiki/MIT_License
// @Security AuthorizationHeader read write
// @SecurityScheme AuthorizationHeader JWT Token Authorization
package main

import (
	"backend-forum/router"
	"backend-forum/utils"
)

func main() {
	utils.PrepareLog()
	router.StartAPI()
}
