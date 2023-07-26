package Summary

import (
	"fmt"
	"strconv"
	"strings"
)

// func main() {
// 	s := "1->1->2->3->4->4->5->6"
// 	println(ToLink(s))
// }

func ToStringList(cur *ListNode) (res string) {
	for cur != nil {
		if cur.Next != nil {
			res += strconv.Itoa(cur.Val) + "->"
		} else {
			res += strconv.Itoa(cur.Val)
		}
		cur = cur.Next
	}
	return res
}

func ToLink1(s string) string {
	// s := "1->2->3->4"
	nums := strings.Split(s, "->")
	// fmt.Printf("%v\n", nums)
	var ans string
	ans += "var "
	for i, _ := range nums {
		if i == len(nums)-1 {
			ans += "node" + strconv.Itoa(i+1) + " "
		} else {
			ans += "node" + strconv.Itoa(i+1) + ", "
		}
	}
	ans += "*ListNode\n"

	for i, v := range nums {
		if i == len(nums)-1 {
			ans += fmt.Sprintf("node%d = &ListNode{ %v, nil }\n", i+1, v)
			// fmt.Printf("node%d = &ListNode{ %v, nil }\n", i+1, v)
		} else {
			ans += fmt.Sprintf("node%d = &ListNode{ %v, node%d }\n", i+1, v, i+2)
			// fmt.Printf("node%d = &ListNode{ %v, node%d }\n", i+1, v, i+2)
		}
	}

	return ans
}

func ToLink(s string) string {
	var ans strings.Builder

	ans.WriteString("\n<======================= start =======================>\n")
	ans.WriteString("\nvar ")

	nums := strings.Split(s, "->")
	for i, _ := range nums {
		if i == len(nums)-1 {
			ans.WriteString("node" + strconv.Itoa(i+1) + " ")
		} else {
			ans.WriteString("node" + strconv.Itoa(i+1) + ", ")
		}
	}
	ans.WriteString("*ListNode\n")

	for i, v := range nums {
		if i == len(nums)-1 {
			ans.WriteString(fmt.Sprintf("node%d = &ListNode{ %v, nil }\n", i+1, v))
		} else {
			ans.WriteString(fmt.Sprintf("node%d = &ListNode{ %v, node%d }\n", i+1, v, i+2))
		}
	}

	ans.WriteString("\n<======================== end ========================>\n")
	return ans.String()
}
