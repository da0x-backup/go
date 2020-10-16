package envie

import (
    "testing"
    "os"
)

type entity struct {
    V1 string `envie:"951f83e8-c682-405f-9b49-49b498a41613"`
    V2 string `envie:"99217878-c9c3-4eaf-90ec-2b54d4da396b"`
}

func TestUnmarshal(t * testing.T) {
    hello := "hello"
    world := "world"
    os.Setenv("951f83e8-c682-405f-9b49-49b498a41613", hello)
    os.Setenv("99217878-c9c3-4eaf-90ec-2b54d4da396b", world)
    var e entity
    Unmarshal(&e)
    if e.V1 != hello || e.V2 != world {
        t.Errorf("Incorrect environment variables")
    }
}
