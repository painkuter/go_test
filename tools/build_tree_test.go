package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

// Format
//id	parent	name	en_name
//10000	0	Земля	Earth
//138	10000	Австралия и Океания	Australia and Oceania

type geoNode struct {
	id       int
	parentID int
	name     string
	enName   string
	parent   *geoNode
	children []*geoNode
}

func TestBuildTree(t *testing.T) {
	file, err := os.ReadFile("geoexport.txt")
	if err != nil {
		panic(err)
	}
	arr := strings.Split(string(file), "\n")
	nodes := make(map[int]*geoNode)

	var parentMinusOneCounter int

	for i, s := range arr {
		s = strings.TrimSpace(s)
		if i == 0 || s == "" { // skipping headers and empty strings
			continue
		}

		parts := strings.Split(s, "\t")
		if len(parts) < 4 {
			//fmt.Printf("Wrong length = %d, string: %s \n", len(parts), s)
			continue
		}

		idStr := parts[0]
		parentIDStr := parts[1]
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		parentID, err := strconv.ParseInt(parentIDStr, 10, 64)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if parentID == -1 {
			parentMinusOneCounter++
			continue
		}
		name := parts[2]
		switch name {
		case "_._._", "Атлантида", "Прочее", "Универсальное", "Арктика и Антарктика", "Яндекс (Москва)", "Яндекс.wifi":
			continue
		}

		node := geoNode{
			id:       int(id),
			parentID: int(parentID),
			name:     name,
			enName:   parts[3],
		}
		nodes[int(id)] = &node
	}
	fmt.Println(len(nodes))
	fmt.Println(parentMinusOneCounter)

	//tree := geoNode{}

	for key, el := range nodes {

		parent, ok := nodes[el.parentID]
		if !ok {
			//fmt.Printf("no parent for %s \n", el.name)
			continue
		}

		if len(parent.children) == 0 {
			parent.children = []*geoNode{}
		}
		parent.children = append(parent.children, el)
		nodes[key].parent = parent
	}
	//fmt.Println(nodes[10000])
	nodes[10000].printChildren(6)
	//INSERT INTO public.cities (id, name, country_id, country_code, latitude, longitude, created_at) VALUES (1, 'Test', 1, 'RR', 0.00000000, 0.00000000, null)
}

func (gn geoNode) printChildren(deep int) bool {
	for i := 5; i > deep; i-- {
		fmt.Print("\t")
	}

	/*	if gn.name == "Россия" {
		deep--
	}*/

	if len(gn.children) == 0 || deep == 0 || gn.name == "Москва" /*|| gn.name == "Москва и Московская область" */ {
		fmt.Println(">", gn.name)
		return true
	}
	if strings.Contains(gn.name, "поселение") {
		return false
	}
	fmt.Println(gn.name)
	if len(gn.children) < 2 {
		return true
	}

	deep--
	var printed int
	for _, child := range gn.children {
		if strings.Contains(child.name, "поселение") {
			return false
		}
		if child.printChildren(deep) {
			printed++
		}
	}
	if printed < 2 {
		fmt.Println(">", gn.name)
	}
	return true
}
