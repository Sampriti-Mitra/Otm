package utils

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"otm/app/constants"
)

var newUUID = uuid.NewV5
var regexCompile = regexp.MustCompile
var currentTime = time.Now

const (
	Method = "method"
)

// GetCurrentTimeStamp will give the current unix timestamp
func GetCurrentTimeStamp() time.Time {
	return currentTime()
}

func GetCurrentTimeStampInMiliSec() int64 {
	return currentTime().UnixNano() / 1000000
}

// Remove a string from an array
func Remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// Find a string "f" in an array and replace it with string "r"
func Replace(s []string, f string, r string) []string {
	for i, n := range s {
		if f == n {
			s[i] = r
		}
	}
	return s
}

// IsEmpty will check for given data is empty as per the go documentation
func IsEmpty(val interface{}) bool {
	if val == nil {
		return true
	}

	reflectVal := reflect.ValueOf(val)

	if val == nil {
		return true
	}

	switch reflectVal.Kind() {
	case reflect.Int:
		return val.(int) == 0

	case reflect.Int64:
		return val.(int64) == 0

	case reflect.String:
		return val.(string) == ""

	case reflect.Map:
	case reflect.Slice:
		return reflectVal.IsNil() || reflectVal.Len() == 0
	}

	return false
}

// GetRequestID fetches the request id from the context
func GetRequestID(ctx *gin.Context) string {
	return ctx.GetString(constants.RequestID)
}

// IsArray : Check if the given value is an array or not
func IsArray(value interface{}) bool {
	rt := reflect.TypeOf(value)
	kind := rt.Kind()

	if kind != reflect.Array && kind != reflect.Slice {
		return false
	}

	return true
}

func InArray(needle interface{}, haystack interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(haystack).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(haystack)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(needle, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

// IsMap : Check if the given value is a map
func IsMap(value interface{}) bool {
	rt := reflect.TypeOf(value)
	kind := rt.Kind()

	if kind != reflect.Map {
		return false
	}

	return true
}

// ConvertToInt : converts a value to int64 type
func ConvertToInt(value interface{}) (int64, error) {
	if IsEmpty(value) {
		return 0, nil
	}

	strValue := fmt.Sprintf("%v", value)

	i, err := strconv.ParseInt(strValue, 10, 32)

	return i, err
}

// StringSliceDiff : returns the elements in `a` that aren't in `b`
func StringSliceDiff(a, b []string) []string {
	mb := map[string]bool{}

	for _, x := range b {
		mb[x] = true
	}

	ab := make([]string, 0)

	for _, x := range a {
		if _, ok := mb[x]; !ok {
			ab = append(ab, x)
		}
	}

	return ab
}

// StringSliceUnique : returns a unique slice of the strings slice input
func StringSliceUnique(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := make([]string, 0)

	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}

func GetMapKeys(mymap map[string]interface{}) []string {
	keys := make([]string, 0, len(mymap))

	for k := range mymap {
		keys = append(keys, k)
	}

	return keys
}

// GenerateReferenceKey : generate reference key used for change log
func GenerateReferenceKey(model interface{}, db *gorm.DB) string {
	var (
		scope         = db.NewScope(model)
		primaryValues []string
	)

	for _, field := range scope.PrimaryFields() {
		primaryValues = append(primaryValues, fmt.Sprint(field.Field.Interface()))
	}

	return strings.Join(primaryValues, "::")
}

func GetTypeName(v interface{}) string {
	ret := ""
	t := reflect.TypeOf(v)
	getType(t, &ret)
	return ret

}

func getType(t reflect.Type, ret *string) {
	if *ret == "" {
		switch t.Kind() {
		case reflect.Array, reflect.Map, reflect.Ptr, reflect.Slice:
			getType(t.Elem(), ret)
		case reflect.Struct:
			*ret = t.Name()
		}
	}
}

func GenerateUUID() string {
	uui, _ := uuid.NewV4()
	return newUUID(uui, "genUUID").String()
}

func IsValidUUID(uuid string) bool {
	r := regexCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}

func Find(list []string, key string) (string, bool) {
	for _, ele := range list {
		ele = strings.ToLower(ele)
		if key == ele {
			return strings.ToUpper(ele), true
		}
	}
	return "", false
}

func MakeCopy(m map[string]interface{}) (map[string]interface{}, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)
	err := enc.Encode(m)
	if err != nil {
		return nil, err
	}
	var copy map[string]interface{}
	err = dec.Decode(&copy)
	if err != nil {
		return nil, err
	}
	return copy, nil
}

func MergeMaps(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

//func GetError(err interface{}) error {
//	var er error
//
//	switch err.(type) {
//	case error:
//		er = err.(error)
//
//	case error_package.IError:
//		er = err.(error_package.IError).GetInternalError()
//
//	case string:
//		er = errors.New(err.(string))
//
//	default:
//		er = errors.New(error_package.ErrorInternalServerError)
//	}
//
//	return er
//}

func QueryArray(query string) []string {
	queryArray := make([]string, 0)
	query = strings.TrimSpace(query)
	if query != "" {
		queryArray = strings.Split(query, ",")
	}
	return queryArray
}

func SplitStringToArray(str string) []string {
	if strings.HasPrefix(str, "(") {
		str = strings.TrimPrefix(str, "(")
	}
	if strings.HasSuffix(str, ")") {
		str = strings.TrimSuffix(str, ")")
	}
	return strings.Split(str, ",")
}

func ConvertToInterfaceArray(strList []string) []interface{} {
	var val []interface{}
	for _, str := range strList {
		val = append(val, str)
	}
	return val
}

// Contains tells whether a contains x.
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// DeserializeJSON will convert the given json into map of string interface
func DeserializeJSON(ctx *gin.Context, in interface{}, out interface{}) error {
	var (
		ierr error
	)

	str, ok := in.(string)

	if !ok {
		return errors.New("error deserializing JSON")
	}

	input := []byte(str)

	err := json.Unmarshal(input, out)

	if err != nil {
		ierr = err
	}

	return ierr
}

func GetTaskId(ctx *gin.Context) string {
	var taskId string
	if traceData, ok := ctx.Get(constants.Request); ok {
		if traceDataMap, ok := traceData.(map[string]interface{}); ok {
			if taskIdValue, ok := traceDataMap[constants.TaskID].(string); ok {
				taskId = taskIdValue
			}
		}
	}
	return taskId
}

func GetError(err interface{}) error {
	var er error

	switch err.(type) {
	case error:
		er = err.(error)

	case string:
		er = errors.New(err.(string))

	default:
		er = errors.New("internal server error")
	}

	return er
}

func ConvertToPointer(str string) *string {
	return &str
}
