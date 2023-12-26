package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

const (
	PhotosUrl = "https://api.pexels.com/v1"
	VideosUrl = "https://api.pexels.com/vidoes"
)

type Client struct {
	Token          string
	HttpCli        http.Client
	RemainingTimes int32
}

type SearchResult struct {
	PageNum      int32   `json:"page_num"`
	PerPage      int32   `json:"per_page"`
	TotalResults int32   `json:"total_results"`
	NextPage     string  `json:"next_page"`
	Photos       []Photo `json:"photos"`
}

type Photo struct {
	Id              int32       `json:"id"`
	Width           int32       `json:"width"`
	Height          int32       `json:"height"`
	Url             string      `json:"url"`
	Photographer    string      `json:"photographer"`
	PhotographerUrl string      `json:"photographer_url"`
	Source          PhotoSource `json:"source"`
}

type PhotoSource struct {
	Original  string `json:"original"`
	Large     string `json:"large"`
	Large2x   string `json:"large2x"`
	Medium    string `json:"Medium"`
	Small     string `json:"small"`
	Potrait   string `json:"potrait"`
	Square    string `json:"square"`
	Landscape string `json:"landscape"`
	Thumbnail string `json:"thumbnail"`
}

type CuratedResult struct {
	PageNum  int32   `json:"page_num"`
	PerPage  int32   `json:"per_page"`
	NextPage int32   `json:"next_page"`
	Photos   []Photo `json:"photos"`
}

func newClient(token string) *Client {
	cli := http.Client{}
	return &Client{Token: token, HttpCli: cli}
}

type VideoSearchResult struct {
	PageNum      int32   `json:"page_num"`
	PerPage      int32   `json:"per_page"`
	TotalResults int32   `json:"total_results"`
	NextPage     string  `json:"next_page"`
	Videos       []Video `json:"videos"`
}

type Video struct {
	Id             int32           `json:"id"`
	Width          int32           `json:"width"`
	Height         int32           `json:"height"`
	Url            string          `json:"url"`
	Image          string          `json:"image"`
	FullResolution interface{}     `json:"full_resolution"`
	Duration       string          `json:"duration"`
	VideoFiles     []VideoFiles    `json:"video_files"`
	VideoPictures  []VideoPictures `json:"video_pictures"`
}

type PopularVideos struct {
	PageNum      int32   `json:"page_num"`
	PerPage      int32   `json:"per_page"`
	TotalResults int32   `json:"total_results"`
	Url          string  `json:"url"`
	Videos       []Video `json:"videos"`
}

type VideoFiles struct {
	Id       int32  `json:"id"`
	Quality  string `json:"quality"`
	FileType string `json:"file_type"`
	Width    int32  `json:"width"`
	Height   int32  `json:"height"`
	Url      string `json:"url"`
}

type VideoPictures struct {
	Id      int32  `json:"id"`
	Picture string `json:"picture"`
	Number  int32  `json:"number"`
}

func (cli *Client) SearchPhotos(searchQuery string, perPage int32, pageNumber int32) (*SearchResult, error) {
	var result SearchResult

	url := fmt.Sprintf(PhotosUrl+"/search?query=%s&per_page=%d&page=%d", searchQuery, perPage, pageNumber)

	response, err := cli.requestDoWithAuth("GET", url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &result)

	return &result, err

}

func (cli *Client) CuratedPhotos(perPage, pageNumber int32) (*CuratedResult, error) {
	var result CuratedResult

	url := fmt.Sprintf(PhotosUrl+"/curated?per_page=%d&page=%d", perPage, pageNumber)

	response, err := cli.requestDoWithAuth("GET", url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &result)

	return &result, err
}

func (cli *Client) requestDoWithAuth(method, url string) (*http.Response, error) {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", cli.Token)
	response, err := cli.HttpCli.Do(request)
	if err != nil {
		return response, err
	}

	times, err := strconv.Atoi(response.Header.Get("X-Ratelimit-Remaining"))
	if err != nil {
		return response, nil
	} else {
		cli.RemainingTimes = int32(times)
	}

	return response, nil
}

func (cli *Client) GetPhoto(id int32) (*Photo, error) {
	var photo Photo

	url := fmt.Sprintf(PhotosUrl+"/photos/%d", id)

	response, err := cli.requestDoWithAuth("GET", url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &photo)

	return &photo, err
}

func (cli *Client) GetRandomPhoto() (*Photo, error) {
	// rand.Seed(time.Now().Unix()) // Deprecated
	rand.New(rand.NewSource(time.Now().Unix())) // New version

	randNum := rand.Intn(1001)

	result, err := cli.CuratedPhotos(1, int32(randNum))
	if err == nil && len(result.Photos) == 1 {
		return &result.Photos[0], nil
	}
	return nil, err
}

func (cli *Client) searchVideo(searchQuery, perPage, pageNumber int32) (*VideoSearchResult, error) {
	var result VideoSearchResult

	url := fmt.Sprintf(VideosUrl+"/search?query=%s&per_page=%d&page=%d", searchQuery, perPage, pageNumber)

	response, err := cli.requestDoWithAuth("GET", url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &result)

	return &result, err
}

func (cli *Client) popularVideo(perPage, pageNumber int32) (*PopularVideos, error) {
	var result PopularVideos

	url := fmt.Sprintf(VideosUrl+"/popular?per_page=%d&page=%d", perPage, pageNumber)

	response, err := cli.requestDoWithAuth("GET", url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &result)

	return &result, err
}

func (cli *Client) getRandomVideo() (*Video, error) {
	rand.New(rand.NewSource(time.Now().Unix())) // New version

	randNum := rand.Intn(1001)

	result, err := cli.popularVideo(1, int32(randNum))
	if err == nil && len(result.Videos) == 1 {
		return &result.Videos[0], nil
	}
	return nil, err
}

func (cli *Client) getRamainingRequests() int32 {
	// Get Remaining requests left for the month
	return cli.RemainingTimes
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	var (
		token  = os.Getenv("PEXELS_API_KEY")
		client = newClient(token)
	)

	result, err := client.SearchPhotos("waves", 15, 1)
	if err != nil {
		log.Fatalf("Error searching photos: %v", err)
	}
	if result.PageNum == 0 {
		log.Fatalf("Search result incorrect")
	}

	fmt.Println(result)

}
