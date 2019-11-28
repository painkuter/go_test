package examples

// https://play.golang.org/p/tocM2Qq6SrH
import (
	"fmt"
	"sort"
	"testing"
)

// CategoryAttribute ...
type CategoryAttribute struct {
	ID         int64
	IsRequired bool
	SortOrder  int64 `json:"-"`
}

type AttributeList []CategoryAttribute

// Len is method to implement sort.Interface
func (al AttributeList) Len() int {
	return len(al)
}

// Less is method to implement sort.Interface
func (al AttributeList) Less(i, j int) bool {
	if al[i].IsRequired {
		return !al[j].IsRequired || al[i].SortOrder < al[j].SortOrder
	}

	if al[j].IsRequired {
		return false
	}

	return al[i].SortOrder < al[j].SortOrder
}

// Swap is method to implement sort.Interface
func (al AttributeList) Swap(i, j int) {
	al[i], al[j] = al[j], al[i]
}

func TestSort(t *testing.T) {
	list := AttributeList{
		{
			ID:         1,
			IsRequired: false,
			SortOrder:  4,
		}, {
			ID:         2,
			IsRequired: false,
			SortOrder:  5,
		}, {
			ID:         3,
			IsRequired: true,
			SortOrder:  2,
		}, {
			ID:         4,
			IsRequired: true,
			SortOrder:  1,
		}}

	sort.Sort(list)
	fmt.Println(list[0].ID == int64(4))
	fmt.Println(list[1].ID == int64(3))
	fmt.Println(list[2].ID == int64(1))
	fmt.Println(list[3].ID == int64(2))
}
