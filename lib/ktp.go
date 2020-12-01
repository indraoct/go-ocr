package lib

import (
	"image"
	"os"
	"regexp"
	"strings"
	"errors"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

type Ktp struct {
	Province 		string		`json:"province"`
	District 		string		`json:"district"`
	Nik				string		`json:"nik"`
	Name 			string		`json:"name"`
	PlaceDob 		string		`json:"place_dob"`
	Gender 			string		`json:"gender"`
	Address1 		string		`json:"address_1"`
	Address2 		string		`json:"address_2"`
	Address3 		string		`json:"address_3"`
	Address4 		string		`json:"address_4"`
	Religion 		string		`json:"religion"`
	MarriedStatus 	string		`json:"married_status"`
	Occupation 		string		`json:"occupation"`
	Nationality     string		`json:"nationality"`
	ValidUntil 		string		`json:"valid_until"`
}


func FormatDataKTP(text string)(formatKtp Ktp,err error){

	var(
		datas 		[]string
		reg 		*regexp.Regexp
		countData	int
	)

	//split string
	datas = strings.Split(strings.Replace(text,"\n\n","\n",-1),"\n")

	//regex config
	reg, _ = regexp.Compile("[^A-Za-z0-9 ]+")

	//looping data split
	for _, data := range datas{

		obj := strings.Split(data,":")

		if len(obj) > 1 {
			data = strings.TrimSpace(obj[1])
			data = reg.ReplaceAllString(data,"")
		}else{
			data = reg.ReplaceAllString(data,"")
		}

		if countData == 0{
			formatKtp.Province = data
		}

		if countData == 1{
			formatKtp.District = data
		}

		if countData == 2{
			formatKtp.Nik = data
		}

		if countData == 3{
			formatKtp.Name = data
		}

		if countData == 4{
			formatKtp.PlaceDob = data
		}

		if countData == 5{
			formatKtp.Gender = data
		}

		if countData == 6{
			formatKtp.Address1 = data
		}

		if countData == 7{
			formatKtp.Address2 = data
		}

		if countData == 8{
			formatKtp.Address3 = data
		}

		if countData == 9{
			formatKtp.Address4 = data
		}

		if countData == 10{
			formatKtp.Religion = data
		}

		if countData == 11{
			formatKtp.MarriedStatus = data
		}

		if countData == 12{
			formatKtp.Occupation = data
		}

		if countData == 13{
			formatKtp.Nationality = data
		}

		if countData == 14{
			formatKtp.ValidUntil = data
		}

		countData++
	}

	return formatKtp,nil
}


func ValidateImageKtp(filename string) (err error){
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil{
		return errors.New("Failed to open file!")
	}

	images, _, err := image.DecodeConfig(file)

	if err != nil{
		return err
	}

	if images.Height >= images.Width {
		return errors.New("Images must be landscape!")
	}

	if images.Height  < 350 {
		return errors.New("Minimum height image is 350 px")
	}

	if images.Width  < 550 {
		return errors.New("Minimum width image is 550 px")
	}

	return nil
}
