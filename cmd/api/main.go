package main

import (
	"coba_dulu/config"
	"coba_dulu/internal/handler"
	"coba_dulu/internal/middleware"
	"coba_dulu/internal/repository"
	"coba_dulu/internal/usecase"
	"coba_dulu/pkg"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()

	// Repositories
	userRepository         = repository.NewUserRepository(db)
	anggotaRepository      = repository.NewAnggotaRepository(db)
	kategoriRepository     = repository.NewKategoriRepository(db)
	penerbitRepository     = repository.NewPenerbitRepository(db)
	bukuRepository         = repository.NewBukuRepository(db)
	peminjamanRepository   = repository.NewPeminjamanRepository(db)
	pengembalianRepository = repository.NewPengembalianRepository(db)

	// Usecases
	jwtService          = pkg.NewJWTService()
	userUsecase         = usecase.NewUserUsecase(userRepository)
	anggotaUsecase      = usecase.NewAnggotaUsecase(anggotaRepository)
	kategoriUsecase     = usecase.NewKategoriUsecase(kategoriRepository)
	penerbitUsecase     = usecase.NewPenerbitUsecase(penerbitRepository)
	bukuUsecase         = usecase.NewBukuUsecase(bukuRepository)
	peminjamanUsecase   = usecase.NewPeminjamanUsecase(peminjamanRepository)
	pengembalianUsecase = usecase.NewPengembalianUsecase(pengembalianRepository)

	// Handlers
	userHandler         = handler.NewUserHandler(userUsecase, jwtService)
	anggotaHandler      = handler.NewAnggotaHandler(anggotaUsecase)
	kategoriHandler     = handler.NewKategoriHandler(kategoriUsecase)
	penerbitHandler     = handler.NewPenerbitHandler(penerbitUsecase)
	bukuHandler         = handler.NewBukuHandler(bukuUsecase)
	peminjamanHandler   = handler.NewPeminjamanHandler(peminjamanUsecase)
	pengembalianHandler = handler.NewPengembalianHandler(pengembalianUsecase)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	// Public routes
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", userHandler.Login)
		authRoutes.POST("/register", userHandler.Register)
	}

	// Protected routes
	apiRoutes := r.Group("api", middleware.AuthorizeJWT(jwtService, "admin"))
	{
		// Anggota routes
		anggotaRoutes := apiRoutes.Group("/anggota")
		{
			anggotaRoutes.GET("/", anggotaHandler.GetAll)
			anggotaRoutes.GET("/:id", anggotaHandler.GetByID)
			anggotaRoutes.POST("/", anggotaHandler.Create)
			anggotaRoutes.PUT("/:id", anggotaHandler.Update)
			anggotaRoutes.DELETE("/:id", anggotaHandler.Delete)
		}
		// Kategori routes
		kategoriRoutes := apiRoutes.Group("/kategori")
		{
			kategoriRoutes.GET("/", kategoriHandler.GetAll)
			kategoriRoutes.GET("/:id", kategoriHandler.GetByID)
			kategoriRoutes.POST("/", kategoriHandler.Create)
			kategoriRoutes.PUT("/:id", kategoriHandler.Update)
			kategoriRoutes.DELETE("/:id", kategoriHandler.Delete)
		}
		// Penerbit routes
		penerbitRoutes := apiRoutes.Group("/penerbit")
		{
			penerbitRoutes.GET("/", penerbitHandler.GetAll)
			penerbitRoutes.GET("/:id", penerbitHandler.GetByID)
			penerbitRoutes.POST("/", penerbitHandler.Create)
			penerbitRoutes.PUT("/:id", penerbitHandler.Update)
			penerbitRoutes.DELETE("/:id", penerbitHandler.Delete)
		}
		// Buku routes
		bukuRoutes := apiRoutes.Group("/buku")
		{
			bukuRoutes.GET("/", bukuHandler.GetAll)
			bukuRoutes.GET("/:id", bukuHandler.GetByID)
			bukuRoutes.POST("/", bukuHandler.Create)
			bukuRoutes.PUT("/:id", bukuHandler.Update)
			bukuRoutes.DELETE("/:id", bukuHandler.Delete)
		}
		// Peminjaman routes
		peminjamanRoutes := apiRoutes.Group("/peminjaman")
		{
			peminjamanRoutes.POST("/", peminjamanHandler.Create)
		}
		// Pengembalian routes
		pengembalianRoutes := apiRoutes.Group("/pengembalian")
		{
			pengembalianRoutes.POST("/", pengembalianHandler.Create)
		}
	}

	r.Run()
}
