package main

import(
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main(){


	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord\n")

	for scanner.Scan(){
		checkdomain(scanner.Text())
	}

	if err:= scanner.Err() ; err!=nil{
		log.Fatalf("error: could not read from input: %v\n", err)
	}
}

func checkdomain(domain string){

	var hasMX, hasSPF, hasDMARC bool
	var spfrecord, dmarcRecord string

	mxrecordss, err :=net.LookupMX(domain)

	if err!= nil{
		log.Printf("error :%v\n", err)

	}
	if len(mxrecordss)>0{
		hasMX=true
	}


	txtrecords, err :=net.LookupTXT(domain)
	if err!=nil{
		log.Printf("error: %v\n",err)
	}

	for _, record := range txtrecords{
		if strings.HasPrefix(record, "v=spf1"){
			hasSPF=true
		    spfrecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err!= nil{
		log.Printf("error %v\n", err)

	}

	for _, record := range dmarcRecords{
		if strings.HasPrefix(record, "v=DMARC1"){
		hasDMARC = true
		dmarcRecord = record
		break
	}
}

	fmt.Printf("%v,%v,%v,%v,%v,%v\n", domain, hasMX, hasSPF, spfrecord, hasDMARC, dmarcRecord)

} 