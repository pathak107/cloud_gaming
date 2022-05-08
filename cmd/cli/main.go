package main

import (
	"fmt"

	"github.com/gosimple/slug"
	"github.com/pathak107/cloudesk/pkg/api/helpers"
)

func main() {
	str := "some string"
	strPtr := helpers.StringPtr(str)
	fmt.Println(helpers.ToString(strPtr))
	cat_id := slug.Make(helpers.ToString(strPtr))
	fmt.Println(cat_id)
}
