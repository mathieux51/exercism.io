package tree

import (
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID         int
	IDChildren []*Node
	Children   []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	// Sort records
	// Could be optimize by swaping
	sort.SliceStable(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	nodes := make([]Node, len(records))

	for i, r := range records {

		nodes[i].ID = i
		// if i != 0 {
		// 	p := &nodes[r.Parent]
		// 	p.Children = append(p.Children, &nodes[i])
		// 	fmt.Printf("%+v\n", p)
		// }
	}

	return &nodes[0], nil
}
