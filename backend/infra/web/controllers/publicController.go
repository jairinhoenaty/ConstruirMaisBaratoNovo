package controllers

import (
	pkgbanneruc "construir_mais_barato/app/usecase/banner"
	pkgbudgetuc "construir_mais_barato/app/usecase/budget"
	pkgcityuc "construir_mais_barato/app/usecase/city"
	pkgclientuc "construir_mais_barato/app/usecase/client"
	pkgpcontactuc "construir_mais_barato/app/usecase/contact"
	pkgpproductuc "construir_mais_barato/app/usecase/product"
	pkgproductuc "construir_mais_barato/app/usecase/product"
	pkgprofessionuc "construir_mais_barato/app/usecase/profession"
	pkgprofessionaluc "construir_mais_barato/app/usecase/professional"
	pkgregionuc "construir_mais_barato/app/usecase/region"
	pkgstoreuc "construir_mais_barato/app/usecase/store"
	pkguseruc "construir_mais_barato/app/usecase/user"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"net/http"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type PublicController struct {
	FindByEmailUCParams                                 pkguseruc.FindByEmailUCParams
	SaveClientUCParams                                  pkgclientuc.SaveClientUCParams
	SaveStoreUCParams                                   pkgstoreuc.SaveStoreUCParams
	FindByPageUCParams                                  pkgbanneruc.FindByPageUCParams
	SaveContactUCParams                                 pkgpcontactuc.SaveContactUCParams
	FindByCityProductUCParams                           pkgproductuc.FindByCityUCParams
	FindDayofferProductUCParams                         pkgproductuc.FindDayofferProductUCParams
	FindAllProductUCParams                              pkgproductuc.FindAllProductUCParams
	FindByUFUCParams                                    pkgcityuc.FindByUFUCParams
	FindRegionByCityIdUCParams							pkgregionuc.FindByCityUCParams
	SaveBudgetUCParams                                  pkgbudgetuc.SaveBudgetUCParams
	UserSendEmailUCParams                               pkguseruc.UserSendEmailUCParams
	ResetPasswordUCParams                               pkguseruc.ResetPasswordUCParams
	FindProfessionUCParams                              pkgprofessionuc.FindProfessionUCParams
	FindAllProfessionUCParams                           pkgprofessionuc.FindAllProfessionUCParams
	SaveProfessionalUCParams                            pkgprofessionaluc.SaveProfessionalUCParams
	FindByCityAndProfessionUCParams                     pkgbanneruc.FindByCityAndProfessionUCParams
	FindProfessionsWithCountIdUCParams                  pkgprofessionuc.FindProfessionsWithCountIdUCParams
	FindAllWithoutPaginationProfessionParams            pkgprofessionuc.FindAllWithoutPaginationProfessionParams
	FindByProfessionalByCityAndProfessionUCParamns      pkgprofessionaluc.FindByProfessionalByCityAndProfessionUCParamns
	FindByNameProfessionalAndCityAndProfessionUCParamns pkgprofessionaluc.FindByNameProfessionalAndCityAndProfessionUCParamns
}

type PublicControllerParams struct {
	FindByEmailUCParams                                 pkguseruc.FindByEmailUCParams
	SaveClientUCParams                                  pkgclientuc.SaveClientUCParams
	SaveStoreUCParams                                   pkgstoreuc.SaveStoreUCParams
	FindByPageUCParams                                  pkgbanneruc.FindByPageUCParams
	SaveContactUCParams                                 pkgpcontactuc.SaveContactUCParams
	FindByCityProductUCParams                           pkgproductuc.FindByCityUCParams
	FindDayofferProductUCParams                         pkgproductuc.FindDayofferProductUCParams
	FindRegionByCityIdUCParams							pkgregionuc.FindByCityUCParams
	FindAllProductUCParams                              pkgproductuc.FindAllProductUCParams
	FindByUFUCParams                                    pkgcityuc.FindByUFUCParams
	SaveBudgetUCParams                                  pkgbudgetuc.SaveBudgetUCParams
	UserSendEmailUCParams                               pkguseruc.UserSendEmailUCParams
	ResetPasswordUCParams                               pkguseruc.ResetPasswordUCParams
	FindProfessionUCParams                              pkgprofessionuc.FindProfessionUCParams
	FindAllProfessionUCParams                           pkgprofessionuc.FindAllProfessionUCParams
	SaveProfessionalUCParams                            pkgprofessionaluc.SaveProfessionalUCParams
	FindByCityAndProfessionUCParams                     pkgbanneruc.FindByCityAndProfessionUCParams
	FindProfessionsWithCountIdUCParams                  pkgprofessionuc.FindProfessionsWithCountIdUCParams
	FindAllWithoutPaginationProfessionParams            pkgprofessionuc.FindAllWithoutPaginationProfessionParams
	FindByProfessionalByCityAndProfessionUCParamns      pkgprofessionaluc.FindByProfessionalByCityAndProfessionUCParamns
	FindByNameProfessionalAndCityAndProfessionUCParamns pkgprofessionaluc.FindByNameProfessionalAndCityAndProfessionUCParamns
}

func NewPublicController(params *PublicControllerParams, g *echo.Group) {
	controller := PublicController{
		FindByEmailUCParams:                                 params.FindByEmailUCParams,
		SaveClientUCParams:                                  params.SaveClientUCParams,
		SaveStoreUCParams:                                   params.SaveStoreUCParams,
		FindByPageUCParams:                                  params.FindByPageUCParams,
		SaveContactUCParams:                                 params.SaveContactUCParams,
		FindByCityProductUCParams:                           params.FindByCityProductUCParams,
		FindDayofferProductUCParams:                         params.FindDayofferProductUCParams,
		FindRegionByCityIdUCParams:							 params.FindRegionByCityIdUCParams,
		FindAllProductUCParams:                              params.FindAllProductUCParams,
		FindByUFUCParams:                                    params.FindByUFUCParams,
		SaveBudgetUCParams:                                  params.SaveBudgetUCParams,
		UserSendEmailUCParams:                               params.UserSendEmailUCParams,
		ResetPasswordUCParams:                               params.ResetPasswordUCParams,
		FindProfessionUCParams:                              params.FindProfessionUCParams,
		SaveProfessionalUCParams:                            params.SaveProfessionalUCParams,
		FindAllProfessionUCParams:                           params.FindAllProfessionUCParams,
		FindByCityAndProfessionUCParams:                     params.FindByCityAndProfessionUCParams,
		FindProfessionsWithCountIdUCParams:                  params.FindProfessionsWithCountIdUCParams,
		FindAllWithoutPaginationProfessionParams:            params.FindAllWithoutPaginationProfessionParams,
		FindByProfessionalByCityAndProfessionUCParamns:      params.FindByProfessionalByCityAndProfessionUCParamns,
		FindByNameProfessionalAndCityAndProfessionUCParamns: params.FindByNameProfessionalAndCityAndProfessionUCParamns,
	}

	g.POST("/save/budget", controller.SaveBudget)
	g.POST("/user/send-mail", controller.SendMail)
	g.POST("/user/find-by-email", controller.FindUserByEmail)
	g.POST("/reset/password", controller.ResetPassword)
	g.GET("/professions/all", controller.FindProfessionAll)
	g.POST("/save/professional", controller.SaveProfessional)
	g.POST("/cities-by-state", controller.PublicFindCitiesByState)
	g.GET("/professions/:quantityProfession", controller.FindProfessions)
	g.POST("/find/professions-with-count", controller.FindProfessionsWithCount)
	g.POST("/find-banner-city-and-profession", controller.FindByCityAndProfession)
	g.POST("/search-all-professionals-and-city-and-profession", controller.PublicFindAllProfessionalsByCityAndProfession)
	g.POST("/search-professionals-by-name-and-city-and-profession", controller.FindByNameProfessinalsAndCityAndProfession)
	g.POST("/products/dayoffer", controller.FindProductsByDayOffer)
	g.GET("/products", controller.FindAllProducts)
	g.POST("/products/findbycity", controller.FindByCity)
	g.GET("/regions/findbycity", controller.FindRegionByCityId)

	g.POST("/contact", controller.SaveContact)

	g.POST("/banners/page", controller.FindBannerbyPage)

	g.POST("/save/store", controller.SaveStore)
	g.POST("/save/client", controller.SaveClient)
	g.POST("/upload/image", controller.uploadFile)

}

func (c *PublicController) FindByCityAndProfession(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	assembler := pkgbanneruc.FindByCityIdAndProfessionIDAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	usecase := pkgbanneruc.NewFindByCityAndProfessionUC(c.FindByCityAndProfessionUCParams)
	usecase.Assembler = &assembler
	banner, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, banner)
}

func (c *PublicController) FindByNameProfessinalsAndCityAndProfession(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	assembler := pkgprofessionaluc.FindProfessionalByCityAndProfessionAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	uc := pkgprofessionaluc.NewFindByNameProfessionalAndCityAndProfessionUC(c.FindByNameProfessionalAndCityAndProfessionUCParamns)
	uc.Assembler = &assembler
	professionals, err := uc.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, professionals)
}

func (c *PublicController) ResetPassword(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	assembler := pkguseruc.ResetPasswordAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request format"})
	}

	usecase := pkguseruc.NewResetPasswordUC(c.ResetPasswordUCParams)
	usecase.Assembler = &assembler

	err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, map[string]string{"error": "Erro ao alterar a senha"})
	}
	return ctx.JSON(http.StatusOK, nil)
}

func (c *PublicController) SendMail(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	var request struct {
		Email string `json:"email"`
	}
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request format"})
	}

	usecase := pkguseruc.NewSendEmailUC(c.UserSendEmailUCParams)
	usecase.Email = request.Email

	err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, map[string]string{"error": "Erro ao enviar o email"})
	}
	return ctx.JSON(http.StatusOK, nil)
}

func (c *PublicController) SaveBudget(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgbudgetuc.NewSaveBudgetUC(c.SaveBudgetUCParams)
	budgetAssembler := pkgbudgetuc.BudgetAssembler{}
	if err := ctx.Bind(&budgetAssembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	usecase.Assembler = &budgetAssembler

	budget, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, budget)

}

func (c *PublicController) FindProfessionsWithCount(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgprofessionuc.NewFindProfessionsWithCountIdUC(c.FindProfessionsWithCountIdUCParams)
	professions, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, professions)
}

func (c *PublicController) SaveProfessional(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgprofessionaluc.NewSaveProfessionalUC(c.SaveProfessionalUCParams)
	assembler := pkgprofessionaluc.ProfessionalAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, map[string]string{
			"error":   "Failed to bind request data",
			"details": err.Error(),
		})
	}
	usecase.Assembler = &assembler
	professional, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, professional)

}

func (c *PublicController) SaveStore(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgstoreuc.NewSaveStoreUC(c.SaveStoreUCParams)
	assembler := pkgstoreuc.StoreAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, map[string]string{
			"error":   "Failed to bind request data",
			"details": err.Error(),
		})
	}
	usecase.Assembler = &assembler
	store, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, store)

}

func (c *PublicController) SaveClient(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgclientuc.NewSaveClientUC(c.SaveClientUCParams)
	assembler := pkgclientuc.ClientAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, map[string]string{
			"error":   "Failed to bind request data",
			"details": err.Error(),
		})
	}
	usecase.Assembler = &assembler
	store, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, store)

}

func (c *PublicController) FindProfessions(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	quantityProfessionAssembler := ctx.Param("quantityProfession")

	usecase := pkgprofessionuc.NewFindProfessionUC(c.FindProfessionUCParams)
	quantProf, err := strconv.ParseUint(quantityProfessionAssembler, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	uintQuantProf := uint(quantProf)
	usecase.QuantityProfessions = uintQuantProf

	professions, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, professions)
}

func (c *PublicController) PublicFindCitiesByState(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	assembler := pkgcityuc.UFCityAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	usecase := pkgcityuc.NewFindByUFUC(c.FindByUFUCParams)
	usecase.Assembler = &assembler
	cities, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, cities)
}

func (c *PublicController) PublicFindAllProfessionalsByCityAndProfession(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	assembler := pkgprofessionaluc.FindProfessionalByCityAndProfessionAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	usecase := pkgprofessionaluc.NewFindByProfessionalByCityAndProfessionUC(c.FindByProfessionalByCityAndProfessionUCParamns)
	usecase.Assembler = &assembler
	professionals, total, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	// Estrutura de resposta
	response := struct {
		Professionals []*pkgprofessionaluc.ProfessionalPresenter `json:"profissionais"`
		Total         int64                                      `json:"total"`
	}{
		Professionals: professionals,
		Total:         total,
	}

	return ctx.JSON(http.StatusOK, response)
}

func (c *PublicController) FindProfessionAll(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgprofessionuc.NewFindAllWithoutPaginationProfessionUC(c.FindAllWithoutPaginationProfessionParams)
	professions, err := usecase.Execute()

	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, professions)
}

func (c *PublicController) FindProductsByDayOffer(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgproductuc.NewFindDayofferProductUC(c.FindDayofferProductUCParams)

	product, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, product)
}

func (c *PublicController) FindAllProducts(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	assembler := pkgpproductuc.FindWithPaginationProductAssembler{}

	limit := ctx.QueryParam("limit")
	offset := ctx.QueryParam("offset")
	professionalID := ctx.QueryParam("professional_id")
	professionalIDInt, err := strconv.Atoi(professionalID)
	storeID := ctx.QueryParam("store_id")
	storeIDInt, err := strconv.Atoi(storeID)

	// Converter os parâmetros para inteiros, com valores padrão se não forem fornecidos
	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt <= 0 {
		limitInt = 20 // valor padrão
	}

	offsetInt, err := strconv.Atoi(offset)
	if err != nil || offsetInt < 0 {
		offsetInt = 0 // valor padrão
	}

	// Definir os parâmetros de paginação no assembler
	assembler.Limit = limitInt
	assembler.Offset = offsetInt
	assembler.ProfessionalID = professionalIDInt
	assembler.StoreID = storeIDInt
	usecase := pkgproductuc.NewFindAllProductUC(c.FindAllProductUCParams)
	usecase.Assembler = assembler

	product, total, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	response := struct {
		Products *[]pkgproductuc.ProductPresenter `json:"products"`
		Total    int64                            `json:"total"`
	}{
		Products: product,
		Total:    total,
	}
	return ctx.JSON(http.StatusOK, response)

}

func (c *PublicController) FindByCity(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	assembler := pkgproductuc.FindByCityAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}
	usecase := pkgproductuc.NewFindByCityUC(c.FindByCityProductUCParams)
	usecase.Assembler = &assembler

	product, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, product)
}

func (c *PublicController) FindRegionByCityId(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgregionuc.NewFindByCityUC(c.FindRegionByCityIdUCParams)
	idAssembler := ctx.QueryParam("cityId")
	id, err := strconv.ParseUint(idAssembler, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	uintID := uint(id)

	usecase.ID = &uintID
	region, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, region)
}


func (c *PublicController) SaveContact(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgpcontactuc.NewSaveContactUC(c.SaveContactUCParams)
	contactAssembler := pkgpcontactuc.ContactAssembler{}
	if err := ctx.Bind(&contactAssembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	usecase.Assembler = &contactAssembler

	contact, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, contact)

}

func (c *PublicController) FindBannerbyPage(ctx echo.Context) error {

	defer ctx.Request().Body.Close()
	
	assembler := pkgbanneruc.FindByPageAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}
	
	uc := pkgbanneruc.NewFindByPageUC(c.FindByPageUCParams)
	uc.Assembler = &assembler
	banners, err := uc.Execute()
	
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, banners)
}

func (c *PublicController) FindUserByEmail(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	var request struct {
		Email string `json:"email"`
	}
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request format"})
	}

	usecase := pkguseruc.NewFindByEmailUC(c.FindByEmailUCParams)
	usecase.Email = &request.Email

	user, err := usecase.Execute()
	if err != nil {
		//return ctx.JSON(http.StatusPreconditionFailed, map[string]string{"error": "E-mail não encontrado"})
		return ctx.JSON(http.StatusOK, false)
	}
	if user != nil {
		return ctx.JSON(http.StatusOK, true)
	}

	return ctx.JSON(http.StatusOK, false)

}

/*
func (c *PublicController) uploadImage(ctx echo.Context) error {
 file, err := ctx.FormFile("image")

  if err != nil {
   log.Println("Error in uploading Image : ", err)
   return ctx.JSON(http.StatusInternalServerError,"Server error")

  }

  uniqueId := uuid.New()

  filename := strings.Replace(uniqueId.String(), "-", "", -1)

  fileExt := strings.Split(file.Filename, ".")[1]

  image := fmt.Sprintf("%s.%s", filename, fileExt)

  // Generate a unique filename
  newFile, err := os.Create("./images/" + image)

  //err = ctx.SaveFile(file, fmt.Sprintf("./images/%s", image))

  if err != nil {
   log.Println("Error in saving Image :", err)
//   return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	return ctx.JSON(http.StatusInternalServerError,"Server error")

  }
  defer newFile.Close()

    _, err = io.Copy(newFile, file)
    if err != nil {
        http.Error(w, "Error copying the file", http.StatusInternalServerError)
        return
    }

  imageUrl := fmt.Sprintf("http://localhost:3000/images/%s", image)

  data := map[string]interface{}{

   "imageName": image,
   "imageUrl":  imageUrl,
   "header":    file.Header,
   "size":      file.Size,
  }

//  return c.JSON(fiber.Map{"status": 201, "message": "Image uploaded successfully", "data":
// data})
	return ctx.JSON(http.StatusCreated,data)

}
*/

func (c *PublicController) uploadFile(ctx echo.Context) error {
	fmt.Println("File Upload Endpoint Hit")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	dir_frontend := os.Getenv("DIR_FRONTEND")
	if dir_frontend == "" {
		dir_frontend = "/frontend/images/upload/"	
	}

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	//r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, err := ctx.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	fmt.Printf("Uploaded File: %+v\n", file.Filename)
	fmt.Printf("File Size: %+v\n", file.Size)
	fmt.Printf("MIME Header: %+v\n", file.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := os.CreateTemp(dir_frontend, "upload-*.png")
	//tempFile, err := os.CreateTemp("public/images/upload/", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()
	fileData, err := file.Open()
	if err != nil {
		//file.Close()
	}

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := io.ReadAll(fileData)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)

	// Define as novas permissões
	novoPermissoes := os.FileMode(0755)
	// Altera as permissões
	err = os.Chmod(tempFile.Name(), novoPermissoes)
	if err != nil {
		fmt.Println("Erro ao alterar permissões:", err)
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	uri := "/images/upload/" + strings.Replace(tempFile.Name(), dir_frontend, "", -1)
	// return that we have successfully uploaded our file
	//fmt.Fprintf(w, "Successfully Uploaded File\n")
	return ctx.JSON(http.StatusCreated, map[string]string{
		"message": "Successfully Uploaded File",
		"uri":     uri})

}
