package main

import (
	"encoding/json"
)

func main() {
	//a := []int{2}
	//
	//nodes := &algo.ListNode{Val: 1}
	//current := nodes
	//
	//for _, v := range a {
	//	tmp := &algo.ListNode{Val: v}
	//	current.Next = tmp
	//	current = current.Next
	//}
	//
	//fmt.Print(nodes)
	//
	//fmt.Print(algo.ReverseList(nodes))

	//fmt.Println(algo.FirstUniqChar("cc"))

	//fmt.Println(resultToMap("apsidaosnbdapisndapnspodamnpo"))
	//algo.Merge([]int64{2, 3, 1}) //

	//fmt.Println(algo.HasCycle(nodes))

}

func resultToMap(result string) (map[string]int, error) {
	tmp := make(map[string]int)
	err := json.Unmarshal([]byte(result), &tmp)
	if err != nil {
		return nil, err
	}

	return tmp, err
}
