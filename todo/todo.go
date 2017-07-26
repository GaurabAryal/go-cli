package todo
import (
  "io/ioutil"
  "encoding/json"
)

type Item struct {
  Text string
  Priority int
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
  b, err := ioutil.ReadFile(filename)
  if err != nil {
    return []Item{}, err
  }
  var items []Item
  data := json.Unmarshal([]byte(b), &items);
  if  items != nil {
    return items, data
  }
  return []Item{}, nil
}

//why can't we just default assign to pri variable? 
func (i *Item) SetPriority(pri int) {
  switch pri {
  case 1:
    i.Priority = 1
  case 3:
    i.Priority = 3
  default:
    i.Priority = 2
  }
}
