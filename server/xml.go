package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/nodauf/ReconFramwork/server/models"
)

func main2() {
	xmlFile, err := os.Open("/tmp/nmap.dtd")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// we initialize our Users array
	var nmap models.Nmaprun
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &nmap)
	fmt.Printf("%#v \n", nmap)
	empJSON, err := json.MarshalIndent(nmap, "", "  ")
	fmt.Println(string(empJSON))
}
