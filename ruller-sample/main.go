package main

import (
	"fmt"
	"math/rand"

	"github.com/Sirupsen/logrus"
	"github.com/flaviostutz/ruller/ruller"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Infof("====Starting Ruller Sample====")
	ruller.Add("test", "rule1", func(input map[string]interface{}) (map[string]interface{}, error) {
		output := make(map[string]interface{})
		output["opt1"] = "Some tests"
		output["opt2"] = 129.99
		rnd := fmt.Sprintln("v", rand.Int())
		if input["children"] == true {
			child := make(map[string]interface{})
			child["c1"] = "123"
			child["c2"] = rnd
			output["children"] = child
		}
		return output, nil
	})
	ruller.Add("test", "rule2", func(input map[string]interface{}) (map[string]interface{}, error) {
		output := make(map[string]interface{})
		output["opt1"] = "Lots of tests"
		age, ok := input["age"].(float64)
		if !ok {
			return nil, fmt.Errorf("Invalid 'age' detected. age=%s", input["age"])
		}
		if age > 60 {
			output["category"] = "elder"
		} else {
			output["category"] = "young"
		}
		return output, nil
	})

	ruller.StartServer()

	// input := make(map[string]interface{})
	// output, err := Process("test", input)
	// if err != nil {
	// 	logrus.Errorf("Error evaluating rules. err=%s", err)
	// 	os.Exit(1)
	// }
	// logrus.Infof("Rules evaluated. output=%s", output)
}