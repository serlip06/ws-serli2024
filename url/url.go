package url

import (
	"github.com/serlip06/ws-serli2024/controller"
	//"net/http"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger" //swager handler
)


func Web(page *fiber.App) {
	// page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	// page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth

	page.Get("/", controller.Sink)
	page.Post("/", controller.Sink)
	page.Put("/", controller.Sink)
	page.Patch("/", controller.Sink)
	page.Delete("/", controller.Sink)
	page.Options("/", controller.Sink)

	page.Get("/checkip", controller.Homepage) //ujicoba panggil package musik
	page.Get("/presensi", controller.GetPresensi)
	//di buat untuk memanggil url buat get by idnya 
	page.Get("/presensi/:id", controller.GetPresensiID) //menampilkan data presensi berdasarkan id
	page.Post("/insert", controller.InsertDataPresensi)//menginsert data 
	page.Put("/update/:id", controller.UpdateData)//update data 
	page.Delete("/delete/:id", controller.DeletePresensiByID)//delete data 

	// login 
	page.Post("/Admin/login", controller.LoginAdmin)//login admin 
	page.Post("/signup", controller.SignupHandler)
	page.Post("/signin", controller.SigninHandler)
	// page.Get("/pengguna/:username", controller.GetPenggunaByUsername)
	//page.Post("/confirm-registration/:id", controller.ConfirmRegistrationHandler)


	//link untuk swager 
	page.Get("/docs/*", swagger.HandlerDefault)
}
