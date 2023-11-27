package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

// a struct to hold the url and the response
type result struct {
	url  string
	data []byte
}

func main() {
	folderList := []string{
		//"_1_6riverdeltafullef__37",
		//"_3_1entrancemouthfull_47",
		//"_3_2entrancejungleful_48",
		//"_3_4secondrivercrossi_49",
		//"_3_5entrancedownlight_50",
		//"_3_6entrancedownmiddl_51",
		//"_4_1handofdogcaverngi_52",
		//"_4_2handofdogtopfulle_53",
		//"_5_1dolineoneascensio_54",
		//"_5_2dolineonestartgig_55",
		"_5_5dolineonedowngiga_57",
		//"_5_7dolineonestalagmi_58",
		//"_6_1dolinetwotunnelgi_62",
		//"_6_2dolinetwoapproach_63",
		//"_6_4dolinetwojunglefu_64",
		//"_6_6dolinetwoapproach_65",
		//"_6_7dolinetwodownfull_66",
		//"_6_9dolinetwoviewpoin_67",
		//"_6_12dolinetwojunglef_60",
		//"_6_13dolinetwoforesto_61",
		//"_7_2passcendaeleendla_68",
		//"_7_3passchendaelestal_69",
		//"_7_4passchendaelestal_70",
	}

	urlList := []string{}

	basePath := "https://www.nationalgeographic.com/news-features/son-doong-cave/2/SonDoong360data"

	l1 := []int{0, 1, 2, 3, 4, 5}

	groupFileList := [][]string{
		{"0_0.jpg", "0_1.jpg", "1_0.jpg", "1_1.jpg"},
		{"0_0.jpg", "0_1.jpg", "0_2.jpg", "1_0.jpg", "1_1.jpg", "1_2.jpg", "2_0.jpg", "2_1.jpg", "2_2.jpg"},
		{"0_0.jpg", "0_1.jpg", "0_2.jpg", "0_3.jpg", "0_4.jpg", "1_0.jpg", "1_1.jpg", "1_2.jpg", "1_3.jpg", "1_4.jpg", "2_0.jpg", "2_1.jpg", "2_2.jpg", "2_3.jpg", "2_4.jpg", "3_0.jpg", "3_1.jpg", "3_2.jpg", "3_3.jpg", "3_4.jpg", "4_0.jpg", "4_1.jpg", "4_2.jpg", "4_3.jpg", "4_4.jpg"},
		{"0_0.jpg", "0_1.jpg", "0_2.jpg", "0_3.jpg", "0_4.jpg", "0_5.jpg", "0_6.jpg", "0_7.jpg", "0_8.jpg", "0_9.jpg", "1_0.jpg", "1_1.jpg", "1_2.jpg", "1_3.jpg", "1_4.jpg", "1_5.jpg", "1_6.jpg", "1_7.jpg", "1_8.jpg", "1_9.jpg", "2_0.jpg", "2_1.jpg", "2_2.jpg", "2_3.jpg", "2_4.jpg", "2_5.jpg", "2_6.jpg", "2_7.jpg", "2_8.jpg", "2_9.jpg", "3_0.jpg", "3_1.jpg", "3_2.jpg", "3_3.jpg", "3_4.jpg", "3_5.jpg", "3_6.jpg", "3_7.jpg", "3_8.jpg", "3_9.jpg", "4_0.jpg", "4_1.jpg", "4_2.jpg", "4_3.jpg", "4_4.jpg", "4_5.jpg", "4_6.jpg", "4_7.jpg", "4_8.jpg", "4_9.jpg", "5_0.jpg", "5_1.jpg", "5_2.jpg", "5_3.jpg", "5_4.jpg", "5_5.jpg", "5_6.jpg", "5_7.jpg", "5_8.jpg", "5_9.jpg", "6_0.jpg", "6_1.jpg", "6_2.jpg", "6_3.jpg", "6_4.jpg", "6_5.jpg", "6_6.jpg", "6_7.jpg", "6_8.jpg", "6_9.jpg", "7_0.jpg", "7_1.jpg", "7_2.jpg", "7_3.jpg", "7_4.jpg", "7_5.jpg", "7_6.jpg", "7_7.jpg", "7_8.jpg", "7_9.jpg", "8_0.jpg", "8_1.jpg", "8_2.jpg", "8_3.jpg", "8_4.jpg", "8_5.jpg", "8_6.jpg", "8_7.jpg", "8_8.jpg", "8_9.jpg", "9_0.jpg", "9_1.jpg", "9_2.jpg", "9_3.jpg", "9_4.jpg", "9_5.jpg", "9_6.jpg", "9_7.jpg", "9_8.jpg", "9_9.jpg"},
		{"0_0.jpg", "0_1.jpg", "0_2.jpg", "0_3.jpg", "0_4.jpg", "0_5.jpg", "0_6.jpg", "0_7.jpg", "0_8.jpg", "0_9.jpg", "0_10.jpg", "0_11.jpg", "0_12.jpg", "0_13.jpg", "0_14.jpg", "0_15.jpg", "0_16.jpg", "0_17.jpg", "0_18.jpg", "0_19.jpg", "1_0.jpg", "1_1.jpg", "1_2.jpg", "1_3.jpg", "1_4.jpg", "1_5.jpg", "1_6.jpg", "1_7.jpg", "1_8.jpg", "1_9.jpg", "1_10.jpg", "1_11.jpg", "1_12.jpg", "1_13.jpg", "1_14.jpg", "1_15.jpg", "1_16.jpg", "1_17.jpg", "1_18.jpg", "1_19.jpg", "2_0.jpg", "2_1.jpg", "2_2.jpg", "2_3.jpg", "2_4.jpg", "2_5.jpg", "2_6.jpg", "2_7.jpg", "2_8.jpg", "2_9.jpg", "2_10.jpg", "2_11.jpg", "2_12.jpg", "2_13.jpg", "2_14.jpg", "2_15.jpg", "2_16.jpg", "2_17.jpg", "2_18.jpg", "2_19.jpg", "3_0.jpg", "3_1.jpg", "3_2.jpg", "3_3.jpg", "3_4.jpg", "3_5.jpg", "3_6.jpg", "3_7.jpg", "3_8.jpg", "3_9.jpg", "3_10.jpg", "3_11.jpg", "3_12.jpg", "3_13.jpg", "3_14.jpg", "3_15.jpg", "3_16.jpg", "3_17.jpg", "3_18.jpg", "3_19.jpg", "4_0.jpg", "4_1.jpg", "4_2.jpg", "4_3.jpg", "4_4.jpg", "4_5.jpg", "4_6.jpg", "4_7.jpg", "4_8.jpg", "4_9.jpg", "4_10.jpg", "4_11.jpg", "4_12.jpg", "4_13.jpg", "4_14.jpg", "4_15.jpg", "4_16.jpg", "4_17.jpg", "4_18.jpg", "4_19.jpg", "5_0.jpg", "5_1.jpg", "5_2.jpg", "5_3.jpg", "5_4.jpg", "5_5.jpg", "5_6.jpg", "5_7.jpg", "5_8.jpg", "5_9.jpg", "5_10.jpg", "5_11.jpg", "5_12.jpg", "5_13.jpg", "5_14.jpg", "5_15.jpg", "5_16.jpg", "5_17.jpg", "5_18.jpg", "5_19.jpg", "6_0.jpg", "6_1.jpg", "6_2.jpg", "6_3.jpg", "6_4.jpg", "6_5.jpg", "6_6.jpg", "6_7.jpg", "6_8.jpg", "6_9.jpg", "6_10.jpg", "6_11.jpg", "6_12.jpg", "6_13.jpg", "6_14.jpg", "6_15.jpg", "6_16.jpg", "6_17.jpg", "6_18.jpg", "6_19.jpg", "7_0.jpg", "7_1.jpg", "7_2.jpg", "7_3.jpg", "7_4.jpg", "7_5.jpg", "7_6.jpg", "7_7.jpg", "7_8.jpg", "7_9.jpg", "7_10.jpg", "7_11.jpg", "7_12.jpg", "7_13.jpg", "7_14.jpg", "7_15.jpg", "7_16.jpg", "7_17.jpg", "7_18.jpg", "7_19.jpg", "8_0.jpg", "8_1.jpg", "8_2.jpg", "8_3.jpg", "8_4.jpg", "8_5.jpg", "8_6.jpg", "8_7.jpg", "8_8.jpg", "8_9.jpg", "8_10.jpg", "8_11.jpg", "8_12.jpg", "8_13.jpg", "8_14.jpg", "8_15.jpg", "8_16.jpg", "8_17.jpg", "8_18.jpg", "8_19.jpg", "9_0.jpg", "9_1.jpg", "9_2.jpg", "9_3.jpg", "9_4.jpg", "9_5.jpg", "9_6.jpg", "9_7.jpg", "9_8.jpg", "9_9.jpg", "9_10.jpg", "9_11.jpg", "9_12.jpg", "9_13.jpg", "9_14.jpg", "9_15.jpg", "9_16.jpg", "9_17.jpg", "9_18.jpg", "9_19.jpg", "10_0.jpg", "10_1.jpg", "10_2.jpg", "10_3.jpg", "10_4.jpg", "10_5.jpg", "10_6.jpg", "10_7.jpg", "10_8.jpg", "10_9.jpg", "10_10.jpg", "10_11.jpg", "10_12.jpg", "10_13.jpg", "10_14.jpg", "10_15.jpg", "10_16.jpg", "10_17.jpg", "10_18.jpg", "10_19.jpg", "11_0.jpg", "11_1.jpg", "11_2.jpg", "11_3.jpg", "11_4.jpg", "11_5.jpg", "11_6.jpg", "11_7.jpg", "11_8.jpg", "11_9.jpg", "11_10.jpg", "11_11.jpg", "11_12.jpg", "11_13.jpg", "11_14.jpg", "11_15.jpg", "11_16.jpg", "11_17.jpg", "11_18.jpg", "11_19.jpg", "12_0.jpg", "12_1.jpg", "12_2.jpg", "12_3.jpg", "12_4.jpg", "12_5.jpg", "12_6.jpg", "12_7.jpg", "12_8.jpg", "12_9.jpg", "12_10.jpg", "12_11.jpg", "12_12.jpg", "12_13.jpg", "12_14.jpg", "12_15.jpg", "12_16.jpg", "12_17.jpg", "12_18.jpg", "12_19.jpg", "13_0.jpg", "13_1.jpg", "13_2.jpg", "13_3.jpg", "13_4.jpg", "13_5.jpg", "13_6.jpg", "13_7.jpg", "13_8.jpg", "13_9.jpg", "13_10.jpg", "13_11.jpg", "13_12.jpg", "13_13.jpg", "13_14.jpg", "13_15.jpg", "13_16.jpg", "13_17.jpg", "13_18.jpg", "13_19.jpg", "14_0.jpg", "14_1.jpg", "14_2.jpg", "14_3.jpg", "14_4.jpg", "14_5.jpg", "14_6.jpg", "14_7.jpg", "14_8.jpg", "14_9.jpg", "14_10.jpg", "14_11.jpg", "14_12.jpg", "14_13.jpg", "14_14.jpg", "14_15.jpg", "14_16.jpg", "14_17.jpg", "14_18.jpg", "14_19.jpg", "15_0.jpg", "15_1.jpg", "15_2.jpg", "15_3.jpg", "15_4.jpg", "15_5.jpg", "15_6.jpg", "15_7.jpg", "15_8.jpg", "15_9.jpg", "15_10.jpg", "15_11.jpg", "15_12.jpg", "15_13.jpg", "15_14.jpg", "15_15.jpg", "15_16.jpg", "15_17.jpg", "15_18.jpg", "15_19.jpg", "16_0.jpg", "16_1.jpg", "16_2.jpg", "16_3.jpg", "16_4.jpg", "16_5.jpg", "16_6.jpg", "16_7.jpg", "16_8.jpg", "16_9.jpg", "16_10.jpg", "16_11.jpg", "16_12.jpg", "16_13.jpg", "16_14.jpg", "16_15.jpg", "16_16.jpg", "16_17.jpg", "16_18.jpg", "16_19.jpg", "17_0.jpg", "17_1.jpg", "17_2.jpg", "17_3.jpg", "17_4.jpg", "17_5.jpg", "17_6.jpg", "17_7.jpg", "17_8.jpg", "17_9.jpg", "17_10.jpg", "17_11.jpg", "17_12.jpg", "17_13.jpg", "17_14.jpg", "17_15.jpg", "17_16.jpg", "17_17.jpg", "17_18.jpg", "17_19.jpg", "18_0.jpg", "18_1.jpg", "18_2.jpg", "18_3.jpg", "18_4.jpg", "18_5.jpg", "18_6.jpg", "18_7.jpg", "18_8.jpg", "18_9.jpg", "18_10.jpg", "18_11.jpg", "18_12.jpg", "18_13.jpg", "18_14.jpg", "18_15.jpg", "18_16.jpg", "18_17.jpg", "18_18.jpg", "18_19.jpg", "19_0.jpg", "19_1.jpg", "19_2.jpg", "19_3.jpg", "19_4.jpg", "19_5.jpg", "19_6.jpg", "19_7.jpg", "19_8.jpg", "19_9.jpg", "19_10.jpg", "19_11.jpg", "19_12.jpg", "19_13.jpg", "19_14.jpg", "19_15.jpg", "19_16.jpg", "19_17.jpg", "19_18.jpg", "19_19.jpg"},
	}

	// loop over the folder_list using a range-based for loop
	for _, folder := range folderList {
		// loop over the l1 using a range-based for loop
		for _, lst := range l1 {
			// loop over the group_file_list using a range-based for loop with an index
			for grpIndex, grp := range groupFileList {
				// loop over the grp using a range-based for loop
				for _, filename := range grp {
					// join the folder name and the file name with "/"
					filePath := "data/sd360/www.nationalgeographic.com/news-features/son-doong-cave/2/SonDoong360data" + "/" + folder + "/" + fmt.Sprintf("%d", lst) + "/" + fmt.Sprintf("%d", grpIndex) + "/" + filename

					if _, err := os.Stat(filePath); os.IsNotExist(err) {
						// construct the download_url using string concatenation and fmt.Sprintf
						downloadUrl := basePath + "/" + folder + "/" + fmt.Sprintf("%d", lst) + "/" + fmt.Sprintf("%d", grpIndex) + "/" + filename
						// append the download_url to the url_list using the built-in append function
						urlList = append(urlList, downloadUrl)
					}
				}
			}
		}
	}

	// create a channel to send and receive the urls
	urlChannel := make(chan string)

	// create a channel to send and receive the results
	resultChannel := make(chan string)

	// create a rate limiter to limit the number of requests per minute
	limiter := time.Tick(1200 * time.Millisecond)

	// create a loop that iterates over the urls and sends them to the url channel
	go func() {
		for _, url := range urlList {
			urlChannel <- url
		}
	}()

	// create a loop that receives from the url channel and launches a goroutine for each url
	go func() {
		for url := range urlChannel {
			// wait for the limiter to send a value
			<-limiter
			// call the download function and send the result to the result channel
			go func(url string) {
				resultChannel <- downloadImage(url)
			}(url)
		}
	}()

	for range urlList {
		fmt.Println(<-resultChannel)
	}
}

func downloadImage(url string) string {
	// split the url by "/" and filter out the empty parts
	urlParts := []string{}
	for _, part := range strings.Split(url, "/") {
		if part != "" {
			urlParts = append(urlParts, part)
		}
	}

	// check if the url_parts has more than 5 elements
	if len(urlParts) > 5 {
		// delete the first 5 elements of the url_parts
		urlParts = urlParts[5:]

		// pop the last element of the url_parts as the file name
		fileName := urlParts[len(urlParts)-1]
		urlParts = urlParts[:len(urlParts)-1]

		// join the remaining elements of the url_parts as the folder name
		folderName := strings.Join(urlParts, "/")

		// create the directory if it does not exist
		os.MkdirAll("data/sd360/www.nationalgeographic.com/news-features/son-doong-cave/2"+"/"+folderName, os.ModePerm)

		// check if the file name is empty
		if fileName == "" {
			return ""
		}

		// join the folder name and the file name with "/"
		filePath := "data/sd360/www.nationalgeographic.com/news-features/son-doong-cave/2" + "/" + folderName + "/" + fileName

		// get the response from the url
		response, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		defer response.Body.Close()

		// check the status code of the response
		if response.StatusCode == 200 {
			// read the response body as bytes
			data, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Println(err)
				return ""
			}

			// write the data to the file
			err = ioutil.WriteFile(filePath, data, 0644)
			if err != nil {
				fmt.Println(err)
				return ""
			}
			return url
		} else if response.StatusCode == 404 {
			// create an empty file
			_, err := os.Create(filePath)
			if err != nil {
				fmt.Println(err)
				return ""
			}
			return "File not found"
		} else {
			return "Download failed"
		}
	} else {
		return ""
	}
}
