package routes

import (
	"net/http"
	"strconv"

	"example.com/dad_midia/models"
	"github.com/gin-gonic/gin"
)

func getMidias(context *gin.Context) {
	estudante, err := strconv.ParseInt(context.Query("estudante"), 10, 64)
	// estudante, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse estudante."})
		return
	}

	midias, err := models.GetAllMidias(estudante)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch midias. Try again later"})
		return
	}
	context.JSON(http.StatusOK, midias)
}

func getMidia(context *gin.Context) {
	midiaId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse midia id."})
		return
	}

	midia, err := models.GetMidiaByID(midiaId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch midia."})
		return
	}

	context.JSON(http.StatusOK, midia)
}

func createMidia(context *gin.Context) {

	var midia models.Midia
	err := context.ShouldBindJSON(&midia)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data."})
		return
	}

	userId := context.GetInt64("userId")

	midia.Estudante = userId
	err = midia.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create midia. Try again later"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Midia created!", "midia": midia})
}

func updateMidia(context *gin.Context) {
	midiaId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse midia id."})
		return
	}

	midia, err := models.GetMidiaByID(midiaId)

	userId := context.GetInt64("userId")

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the midia."})
		return
	}

	if midia.Estudante != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update midia."})
		return
	}

	var updatedMidia models.Midia
	err = context.ShouldBindJSON(&updatedMidia)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse midia id."})
		return
	}

	updatedMidia.ID = midiaId

	err = updatedMidia.Update()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not update midia."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Midia updated succesfully."})
}

func deleteMidia(context *gin.Context) {
	midiaId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse midia id."})
		return
	}

	midia, err := models.GetMidiaByID(midiaId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the midia."})
		return
	}

	userId := context.GetInt64("userId")
	if midia.Estudante != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete midia."})
		return
	}

	err = midia.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the midia."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Midia deleted succesfully."})
}
