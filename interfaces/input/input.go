package input

import "strings"
import "strconv"
import "log"
import "github.com/gin-gonic/gin"

type Codes struct {
    Uha_id string
	Link_id string
}


func Init(c *gin.Context) *Codes {
	inputInstance := new(Codes)
	
	params := parse(c.Param("codes"))

	// Init code
	inputInstance.Uha_id = params[0]
	inputInstance.Link_id = params[1]
	// ...
	return inputInstance
}

func parse(codes string) []string {
	var params []string
	returnParam := make([]string, 2)
	params = strings.Split(codes, "-")
	log.Println("Params count " + strconv.Itoa(len(params)))
	if len(params) != 2 {
		log.Println("Parameters don't not exist")
		returnParam = []string{params[0], "empty"};
	} else {
		copy(returnParam, params)
	}
	params = nil;
	return returnParam;
}