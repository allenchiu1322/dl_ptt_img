package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "os"
    "net/http"
    "path"
    "regexp"
    "strconv"
)

func main() {

    // get parameters
    if len(os.Args) > 1 {
        file_url := os.Args[1]
        filename_part := path.Base(file_url)

        // download file
        if err := DownloadFile(filename_part, file_url); err != nil {
            panic(err)
        }

        // load fie

        content, err := ioutil.ReadFile(filename_part)
        if err != nil {
            panic(err)
        }
        text := string(content)

        // find match url
        r, _ := regexp.Compile("https://i.imgur.com/[0-9A-Za-z]{5,8}.jpg")
        matches := r.FindAllString(text, -1)
        matches = RemoveDuplicatesFromSlice(matches)
        match_count := strconv.Itoa(len(matches))

        // download each img
        for index, element := range matches {
            index_str := strconv.Itoa(index + 1)
            fmt.Print(index_str + "/" + match_count + " ")
            if err := DownloadFile(path.Base(element), element); err != nil {
                panic(err)
            }
        }

    } else {
        fmt.Println("Usage: dl_ptt_img url")
    }


}

func DownloadFile(filepath string, url string) error {
    fmt.Println("Downloading from " + url)
    // Get the data
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // Create the file
    out, err := os.Create(filepath)
    if err != nil {
        return err
    }
    defer out.Close()

    // Write the body to file
    _, err = io.Copy(out, resp.Body)
    return err
}

func RemoveDuplicatesFromSlice(s []string) []string {
      m := make(map[string]bool)
      for _, item := range s {
              if _, ok := m[item]; ok {
                      // duplicate item
//                      fmt.Println(item, "is a duplicate")
              } else {
                      m[item] = true
              }
      }

      var result []string
      for item, _ := range m {
              result = append(result, item)
      }
      return result
}

