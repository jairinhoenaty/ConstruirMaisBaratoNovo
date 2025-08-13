package controllers

import (
	pkgbanneruc "construir_mais_barato/app/usecase/banner"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
)

type BannerController struct {
	SaveBannerUCParams pkgbanneruc.SaveBannerUCParams
	DeleteBannerUCParams  pkgbanneruc.DeleteBannerUCParams
	FindByPageUCParams pkgbanneruc.FindByPageUCParams

}

type BannerControllerParams struct {
	SaveBannerUCParams pkgbanneruc.SaveBannerUCParams
	DeleteBannerUCParams pkgbanneruc.DeleteBannerUCParams
	FindByPageUCParams pkgbanneruc.FindByPageUCParams
}

func NewBannerController(params *BannerControllerParams, g *echo.Group) {
	controller := BannerController{
		SaveBannerUCParams: params.SaveBannerUCParams,
		DeleteBannerUCParams:  params.DeleteBannerUCParams,
		FindByPageUCParams: params.FindByPageUCParams,
	}

	g.POST("/banner", controller.Save)
	g.POST("/banners/page", controller.FindbyPage)
	g.DELETE("/banner/:id", controller.Delete)


}



func (c *BannerController) FindbyPage(ctx echo.Context) error {

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

func (c *BannerController) Save(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	// Parse form values
	assembler := pkgbanneruc.BannerAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}
	/*bannerId, err := strconv.ParseUint(bannerIdStr, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid bannerId"})
	}
	accessLink := ctx.FormValue("accessLink")
*/
	professionsStr := ctx.Request().PostForm["professions"]

	// Convert professions to a slice of ints
	var professions []uint
	for _, professionStr := range professionsStr {
		professionID, err := strconv.ParseUint(professionStr, 10, 32)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid profession ID"})
		}
		professions = append(professions, uint(professionID))
	}
/*
	// Get the file(s)
	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid file upload"})
	}

	// Open the file
	src, err := file.Open()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	defer src.Close()

	// Read the file into a byte slice
	fileBytes, err := io.ReadAll(src)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
*/
	/*assembler := pkgbanneruc.BannerAssembler{
		ID:      uint(bannerId),
		AccessLink:  accessLink,
		Professions: professions,
		//Image:       ,
	}*/

	// assembler := pkgbanneruc.BannerAssembler{}
	// if err := ctx.Bind(&assembler); err != nil {
	// 	return ctx.JSON(http.StatusPreconditionFailed, err)
	// }

	usecase := pkgbanneruc.NewSaveBannerUC(c.SaveBannerUCParams)
	usecase.Assembler = &assembler
	banner, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed,err)
	}
	return ctx.JSON(http.StatusOK, banner)

}


func (c *BannerController) Delete(ctx echo.Context) error {

	defer ctx.Request().Body.Close()
	usecase := pkgbanneruc.NewDeleteBannerUC(c.DeleteBannerUCParams)
	idAssembler := ctx.Param("id")
	id, err := strconv.ParseUint(idAssembler, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	uintID := uint(id)

	usecase.ID = &uintID
	err = usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, nil)

}

