package sap

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
	"net/textproto"
	"net/url"
	"path"
	"strings"
	"text/template"

	"github.com/pkg/errors"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-sap/" + libraryVersion
	mediaType      = "application/json"
	charset        = "utf-8"
)

var (
	BaseURL = url.URL{
		Scheme: "https",
		Host:   "localhost:50000",
		Path:   "",
	}

///BusinessPartners
)

// NewClient returns a new Exact Globe Client client
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{}

	client.SetHTTPClient(httpClient)
	client.SetBaseURL(BaseURL)
	client.SetDebug(false)
	client.SetUserAgent(userAgent)
	client.SetMediaType(mediaType)
	client.SetCharset(charset)

	return client
}

// Client manages communication with Exact Globe Client
type Client struct {
	// HTTP client used to communicate with the Client.
	http *http.Client

	debug   bool
	baseURL url.URL

	// credentials
	username  string
	password  string
	companyDB string
	cookiejar *cookiejar.Jar

	// User agent for client
	userAgent string

	mediaType             string
	charset               string
	disallowUnknownFields bool

	// Optional function called after every successful request made to the DO Clients
	beforeRequestDo    BeforeRequestDoCallback
	onRequestCompleted RequestCompletionCallback
}

type BeforeRequestDoCallback func(*http.Client, *http.Request, interface{})

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

func (c *Client) SetHTTPClient(client *http.Client) {
	c.http = client
}

func (c Client) Username() string {
	return c.username
}

func (c *Client) SetUsername(username string) {
	c.username = username
}

func (c Client) Password() string {
	return c.password
}

func (c *Client) SetPassword(password string) {
	c.password = password
}

func (c Client) CompanyDB() string {
	return c.companyDB
}

func (c *Client) SetCompanyDB(companyDB string) {
	c.companyDB = companyDB
}

func (c *Client) InitSession(req *http.Request) error {
	if c.cookiejar == nil {
		jar, err := c.FetchLoginCookies()
		if err != nil {
			return errors.WithStack(err)
		}
		c.cookiejar = &jar
	}

	return nil
}

func (c *Client) FetchLoginCookies() (cookiejar.Jar, error) {
	var err error

	req := c.NewLoginPostRequest()
	req.RequestBody().Username = c.Username()
	req.RequestBody().Password = c.Password()
	req.RequestBody().CompanyDB = c.CompanyDB()
	jar, err := cookiejar.New(nil)
	if err != nil {
		return cookiejar.Jar{}, errors.WithStack(err)
	}

	req.client.http.Jar = jar
	_, err = req.Do()
	if err != nil {
		return cookiejar.Jar{}, errors.WithStack(err)
	}
	return *jar, nil
}

func (c *Client) HasCookie() (bool, error) {
	return (c.cookiejar != nil), nil
}

func (c Client) Debug() bool {
	return c.debug
}

func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

func (c Client) BaseURL() url.URL {
	return c.baseURL
}

func (c *Client) SetBaseURL(baseURL url.URL) {
	c.baseURL = baseURL
}

func (c *Client) SetMediaType(mediaType string) {
	c.mediaType = mediaType
}

func (c Client) MediaType() string {
	return mediaType
}

func (c *Client) SetCharset(charset string) {
	c.charset = charset
}

func (c Client) Charset() string {
	return charset
}

func (c *Client) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

func (c Client) UserAgent() string {
	return userAgent
}

func (c *Client) SetDisallowUnknownFields(disallowUnknownFields bool) {
	c.disallowUnknownFields = disallowUnknownFields
}

func (c *Client) SetBeforeRequestDo(fun BeforeRequestDoCallback) {
	c.beforeRequestDo = fun
}

func (c *Client) GetEndpointURL(p string, pathParams PathParams) url.URL {
	clientURL := c.BaseURL()

	parsed, err := url.Parse(p)
	if err != nil {
		log.Fatal(err)
	}
	q := clientURL.Query()
	for k, vv := range parsed.Query() {
		for _, v := range vv {
			q.Add(k, v)
		}
	}
	clientURL.RawQuery = q.Encode()

	clientURL.Path = path.Join(clientURL.Path, parsed.Path)

	tmpl, err := template.New("path").Parse(clientURL.Path)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	params := pathParams.Params()
	// params["administration_id"] = c.Administration()
	err = tmpl.Execute(buf, params)
	if err != nil {
		log.Fatal(err)
	}

	clientURL.Path = buf.String()
	return clientURL
}

func (c *Client) NewRequest(ctx context.Context, req Request) (*http.Request, error) {
	// convert body struct to json
	var body io.Reader
	if req.RequestBodyInterface() != nil {
		if r, ok := req.RequestBodyInterface().(io.Reader); ok {
			body = r
		} else if bb, ok := req.RequestBodyInterface().([]byte); ok {
			body = bytes.NewReader(bb)
		} else {
			buf := new(bytes.Buffer)
			err := json.NewEncoder(buf).Encode(req.RequestBodyInterface())
			if err != nil {
				return nil, err
			}
			body = buf
		}
	}

	// create new http request
	r, err := http.NewRequest(req.Method(), req.URL().String(), body)
	if err != nil {
		return nil, err
	}

	// values := url.Values{}
	// err = AddURLValuesToRequest(values, req, true)
	// if err != nil {
	// 	return nil, err
	// }

	// optionally pass along context
	if ctx != nil {
		r = r.WithContext(ctx)
	}

	// set other headers
	r.Header.Add("Content-Type", fmt.Sprintf("%s; charset=%s", c.MediaType(), c.Charset()))
	r.Header.Add("Accept", c.MediaType())
	r.Header.Add("User-Agent", c.UserAgent())

	return r, nil
}

func (c *Client) NewFormRequest(ctx context.Context, method string, URL url.URL, form Form) (*http.Request, error) {
	body := &bytes.Buffer{}
	w := multipart.NewWriter(body)

	for k, vv := range form.Values() {
		for _, v := range vv {
			err := w.WriteField(k, v)
			if err != nil {
				return nil, err
			}
		}
	}

	for k, f := range form.Files() {
		part, err := CreateFormFile(w, f.Content, k, f.Filename)
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(part, f.Content)
	}

	err := w.Close()
	if err != nil {
		return nil, err
	}

	// create new http request
	req, err := http.NewRequest(method, URL.String(), body)
	if err != nil {
		return nil, err
	}

	values := url.Values{}
	err = AddURLValuesToRequest(values, req, true)
	if err != nil {
		return nil, err
	}

	// optionally pass along context
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	// set other headers
	req.Header.Add("Content-Type", fmt.Sprintf("%s; charset=%s", w.FormDataContentType(), c.Charset()))
	req.Header.Add("Accept", c.MediaType())
	req.Header.Add("User-Agent", c.UserAgent())

	return req, nil
}

// Do sends an Client request and returns the Client response. The Client response is json decoded and stored in the value
// pointed to by v, or returned as an error if an Client error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, body interface{}) (*http.Response, error) {
	if c.beforeRequestDo != nil {
		c.beforeRequestDo(c.http, req, body)
	}

	if c.debug == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	httpResp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, httpResp)
	}

	// close body io.Reader
	defer func() {
		if rerr := httpResp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if c.debug == true {
		dump, _ := httputil.DumpResponse(httpResp, true)
		log.Println(string(dump))
	}

	// check if the response isn't an error
	err = CheckResponse(httpResp)
	if err != nil {
		return httpResp, err
	}

	// check the provided interface parameter
	if httpResp == nil {
		return httpResp, nil
	}

	if body == nil {
		return httpResp, err
	}

	if httpResp.ContentLength == 0 {
		return httpResp, nil
	}

	errResp := &ErrorResponse{Response: httpResp}
	err = c.Unmarshal(httpResp.Body, body, errResp)
	if err != nil {
		return httpResp, err
	}

	if errResp.Error() != "" {
		return httpResp, errResp
	}

	return httpResp, nil
}

func (c *Client) Unmarshal(r io.Reader, vv ...interface{}) error {
	if len(vv) == 0 {
		return nil
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	errs := []error{}
	for _, v := range vv {
		r := bytes.NewReader(b)
		dec := json.NewDecoder(r)
		if c.disallowUnknownFields {
			dec.DisallowUnknownFields()
		}

		err := dec.Decode(v)
		if err != nil && err != io.EOF {
			errs = append(errs, err)
		}

	}

	if len(errs) == len(vv) {
		// Everything errored
		msgs := make([]string, len(errs))
		for i, e := range errs {
			msgs[i] = fmt.Sprint(e)
		}
		return errors.New(strings.Join(msgs, ", "))
	}

	return nil
}

// CheckResponse checks the Client response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range. Client error responses are expected to have either no response
// body, or a json response body that maps to ErrorResponse. Any other response
// body will be silently ignored.
func CheckResponse(r *http.Response) error {
	errorResponse := &ErrorResponse{Response: r}

	// Don't check content-lenght: a created response, for example, has no body
	// if r.Header.Get("Content-Length") == "0" {
	// 	errorResponse.Errors.Message = r.Status
	// 	return errorResponse
	// }

	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	// read data and copy it back
	data, err := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewReader(data))
	if err != nil {
		return errorResponse
	}

	err = checkContentType(r)
	if err != nil {
		return errors.WithStack(err)
	}

	if r.ContentLength == 0 {
		return errors.New("response body is empty")
	}

	// convert json to struct
	if len(data) != 0 {
		err = json.Unmarshal(data, &errorResponse)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	if errorResponse.Error() != "" {
		return errorResponse
	}

	return nil
}

// {
//   "ErrorType": "AccountViewError",
//   "ErrorNumbers": null,
//   "ErrorMessage": ""
// }

type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response

	Err struct {
		Code    interface{} `json:"code"`
		Message struct {
			Lang  string `json:"lang"`
			Value string `json:"value"`
		} `json:"message"`
	} `json:"error"`
}

func (r ErrorResponse) Error() string {
	if (r.Err.Code == nil || r.Err.Code == "" || r.Err.Code == 0) && r.Err.Message.Value == "" {
		return ""
	}

	return fmt.Sprintf("%v: %s", r.Err.Code, r.Err.Message.Value)
}

func checkContentType(response *http.Response) error {
	header := response.Header.Get("Content-Type")
	contentType := strings.Split(header, ";")[0]
	if contentType != mediaType {
		return fmt.Errorf("Expected Content-Type \"%s\", got \"%s\"", mediaType, contentType)
	}

	return nil
}

func CreateFormFile(w *multipart.Writer, data io.Reader, fieldname, filename string) (io.Writer, error) {
	var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

	escapeQuotes := func(s string) string {
		return quoteEscaper.Replace(s)
	}

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			escapeQuotes(fieldname), escapeQuotes(filename)))

	contentType, err := GetFileContentType(data)
	if err != nil {
		return nil, err
	}
	h.Set("Content-Type", contentType)
	return w.CreatePart(h)
}

func GetFileContentType(file io.Reader) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := file.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
