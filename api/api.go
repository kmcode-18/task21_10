package api

import (
	"encoding/json"
	"github.com/kmcode-18/task21_10/s3"
	"github.com/kmcode-18/task21_10/store"
	"github.com/kmcode-18/task21_10/utils"
	"log"
	"net/http"
)

var (
	imageFormatTable = []string{
		"image/tif",
		"image/tiff",
		"image/jpg",
		"image/jpeg",
		"image/gif",
		"image/png",
		"image/bmp",
		"image/eps",
	}
	sortOrderList = []string{
		"asc",
		"desc",
	}
	sortByList = []string{
		"created_at",
		"image_name",
	}
)

func AddImage(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("file")
	if err != nil {
		err := map[string]interface{}{"message": "enter add file in form data"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	defer file.Close()
	if !utils.StrInListStatus(handler.Header.Get("Content-Type"), imageFormatTable) {
		err := map[string]interface{}{"message": "please upload a image"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	err = s3.UploadFileToS3(handler.Filename, file)
	if err != nil {
		err := map[string]interface{}{"message": err.Error()}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	success := map[string]interface{}{"message": "image uploaded successfully"}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success)
	return
}
func GetImages(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	var qParms store.QueryDetails
	var ok bool
	imageId := queryParams.Get("id")
	if qParms.ImageId, ok = utils.CheckIntValue(imageId); !ok && imageId != "" {
		err := map[string]interface{}{"message": "id should be a integer value"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	offSet := queryParams.Get("offset")
	if qParms.OffSet, ok = utils.CheckIntValue(offSet); !ok && offSet != "" {
		err := map[string]interface{}{"message": "offset should be a integer value"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	limit := queryParams.Get("limit")
	if qParms.Limit, ok = utils.CheckIntValue(limit); !ok && limit != "" {
		err := map[string]interface{}{"message": "limit should be a integer value"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	if qParms.Limit, ok = utils.CheckIntValue(limit); !ok && limit != "" {
		err := map[string]interface{}{"message": "limit should be a integer value"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	qParms.SortOrder = queryParams.Get("sort_order")
	if ok = utils.StrInListStatus(qParms.SortOrder, sortOrderList); !ok && qParms.SortOrder != "" {
		err := map[string]interface{}{"message": "enter a valid sort order"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	qParms.SortBy = queryParams.Get("sort_by")
	if ok = utils.StrInListStatus(qParms.SortBy, sortByList); !ok && qParms.SortBy != "" {
		err := map[string]interface{}{"message": "enter a valid sort parameter"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	resp, err := store.GetImages(qParms)
	if err != nil {
		err := map[string]interface{}{"message": "error fetching data from db"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	success := map[string]interface{}{"result": resp}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success)
	return
}

func Stop() (err error) {
	err = store.CloseDbConn()
	if err != nil {
		log.Println("error : ", err.Error())
	}
	return
}
