package input

import "errors"
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

	params, e := parse(c.Param("codes"))

	if e != nil {
		return nil
	}

	// Init code
	inputInstance.Uha_id = params[0]
	inputInstance.Link_id = params[1]

	return inputInstance
}



func parse(codes string) ([]string, error) {
	var params []string

	params = strings.Split(codes, "-")

	log.Println("Params count " + strconv.Itoa(len(params)))

	if len(params) != 2 {
		return nil, errors.New("Parameters don't not exist");
	}

	return params, nil;
}