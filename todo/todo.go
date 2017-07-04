package todo
import (
  "io/ioutil"
  "fmt"
  "encoding/json"
)

type Item struct {
  Text string
}

func SaveItems(filename string, items []Item) error {
  b, err := json.Marshal(items)
  if err != nil {
    return err
  }
  err = ioutil.WriteFile(filename, b, 0644)
  if err != nil {
    return err
  }
  return nil
}

//([]Item, error) just specify the return types
func ReadItems(filename string) ([]Item, error) {
  fmt.Println(filename)
  b, err := ioutil.ReadFile(filename)
  if err != nil {
    return []Item{}, err
  }
  var items []Item
  data := json.Unmarshal(b, &items);
  fmt.Println(err)
  if  data != nil {
    return []Item{}, data
  }
  return []Item{}, nil
}
