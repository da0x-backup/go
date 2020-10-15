package rest
import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "time"
)

// Get calls a url and parses the json output and unmarshals it into a struct. It returns
// the http response code (if one is available) and an optional error.
func Get(url string, response interface{}, headers map[string]string) (int, error) {
    request, err := http.NewRequest("GET", url, bytes.NewReader([]byte{}))
    if err != nil {
        return 0, err
    }
    for key, value := range headers {
        request.Header.Set(key,value)
    }
    request.Header.Set("Content-Type", "application/json")
    request.Header.Set("Connection", "close")
    var client = &http.Client{Timeout: 30 * time.Second}
    resp, err := client.Do(request)
    if err != nil {
        if resp != nil {
            return resp.StatusCode, err
        }
        return 0, err
    }
    read, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return resp.StatusCode, err
    }
    resp.Body.Close()
    err = json.Unmarshal(read,response)
    if err != nil {
        return resp.StatusCode, err
    }
    return resp.StatusCode, nil
}

// Post calls a url with the POST method and sends the provided body. It will
// return a response code (if one is available) and an optional error. This
// function will also unmarshal the reponse in its response struct.
func Post(url string, body, response interface{}, headers map[string]string) (int,error) {
    requestBody, err := json.Marshal(body)
    if err != nil {
        return 0, err
    }
    request, err := http.NewRequest("POST", url, bytes.NewReader(requestBody))
    if err != nil {
        return 0, err
    }
    request.Header.Set("Content-Type", "application/json")
    request.Header.Set("Connection", "close")
    var client = &http.Client{Timeout: 30 * time.Second}
    resp, err := client.Do(request)
    if err != nil {
        if resp!= nil {
            return resp.StatusCode, err
        }
        return 0, err
    }
    read, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return resp.StatusCode, err
    }
    resp.Body.Close()
    err = json.Unmarshal(read,resp)
    if err != nil {
        return resp.StatusCode, err
    }
    return resp.StatusCode, nil
}

// PostAcceptOctetStream will attempt to post the request to the given url and will give back
// the response from the server as a string.
func PostAcceptOctetStream(url string, body, response *string, headers map[string]string) (int,error) {
    requestBody, err := json.Marshal(body)
    if err != nil {
        return 0, err
    }
    request, err := http.NewRequest("POST", url, bytes.NewReader(requestBody))
    if err != nil {
        return 0, err
    }
    request.Header.Set("Content-Type", "application/json")
    request.Header.Set("Connection", "close")
    var client = &http.Client{Timeout: 30 * time.Second}
    resp, err := client.Do(request)
    if err != nil {
        if resp != nil {
            return resp.StatusCode, err
        }
        return 0, err
    }
    read, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return resp.StatusCode, err
    }
    resp.Body.Close()
    *response = string(read)
    return resp.StatusCode, nil
}
