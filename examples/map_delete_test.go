package examples

import (
	"testing"
)

type node struct {
	ID     int
	parent *node
}

func mockData() []node {
	var result []node
	for i := 0; i < size; i++ {
		result = append(result, node{ID: i})
	}
	return result
}

func TestDeleteFromMap(t *testing.T) {
	tree := mockData()
	for _, elem := range tree {
		for _, elem2 := range tree {
			if elem.ID%elem2.ID == 0 {
				// elem.parent
			}
		}
	}
}

/*

func (c Controller) GetTree(ctx context.Context, request *GetTreeRequest) (*GetTreeResponse, error) {

	commercialCategories, err := c.AttributeService.GetCommercialCategories(ctx)
	if err != nil {
		return nil, err
	}

	commercialToDescription, err := c.AttributeService.GetCommercialCategoriesMap(ctx)
	if err != nil {
		return nil, err
	}

	result := make(map[string]*Node)

	for _, elem := range descriptionCategories {
		// skip unmapped
		if _, ok := unmapCatsMap[elem.ID]; ok {
			continue
		}

		result[strconv.Itoa(int(elem.ID))] = &Node{
			ID:       elem.ID,
			Disabled: elem.Disabled,
			Name:     elem.LongName, // from PHP
			ParentID: elem.ParentID,
			Nodes:    make(map[string]*Node),
		}
	}

	commTypeFilter := make(map[int64]bool)
	for _, elem := range commercialCategories {

		if commercialToDescription[elem.ID] == 0 ||
			commTypeFilter[elem.CommercialTypeID] ||
			result[strconv.Itoa(int(commercialToDescription[elem.ID]))] == nil {

			continue
		}

		commTypeFilter[elem.CommercialTypeID] = true

		result[strconv.Itoa(int(elem.ID))] = &Node{
			ID:       elem.ID,
			Name:     elem.Name,
			ParentID: commercialToDescription[elem.ID],
			Nodes:    make(map[string]*Node),
		}
	}

	for i, elem := range result {
		if elem.ParentID != -1 && elem.ParentID != 0 {
			delete(result, i) // remove not root elements
			if result[strconv.Itoa(int(elem.ParentID))] == nil {
				continue
			}
			result[strconv.Itoa(int(elem.ParentID))].Nodes[strconv.Itoa(int(elem.ID))] = elem
		}
	}

	return &GetTreeResponse{
		Result: result,
	}, nil
}
*/
