package tools

import (
	"strconv"
	"time"
	"fmt"
)

func GetStoreImageCover(storeId int, suffix string) string {
	name := strconv.Itoa(storeId) + "_" + strconv.Itoa(int(time.Now().UnixNano() / 1e3)) +"." + suffix
	return name
}

func GetStoreImageSavePath(storeId int, name string) string {
	//path := "/Users/kenn/www/webcb/images/store/" + strconv.Itoa(storeId)
	path := "/alidata/dockerdata/nginx/www/webcb-common/images/store/" + strconv.Itoa(storeId)
	err := MkDir(path, 0777)
	if err == nil {
		return path + "/" + name
	} else {
		fmt.Println(err)
		return ""
	}

}