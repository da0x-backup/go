package envie

import (
    "fmt"
    "reflect"
    "os"
)

// Unmarshal reads an entire struct of env variables. Returns an error
// if any of those variables does not exist in the environment.
func Unmarshal(e interface{}) error {
    t := reflect.TypeOf(e).Elem()
    v := reflect.ValueOf(e).Elem()
    errors := []string{}
    for i := 0; i < t.NumField(); i++ {
        tag := t.Field(i).Tag
        env := tag.Get("envie")
        val := os.Getenv(env)
        v.Field(i).SetString(val)
        if len(val) <= 0 {
            errors = append(errors, env)
        }
    }
    if len(errors) != 0 {
        str := "environment variables not found:\n"
        for _, err := range errors {
            str += "\t" + err + "\n"
        }
        return fmt.Errorf(str)
    }
    return nil
}
