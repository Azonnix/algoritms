package main

func main() {
	// DELETE GRAITEST VALUE IN EACH ROW
	// var a = [][]int{{1, 2, 4}, {3, 3, 1}}
	// fmt.Println(deleteGraitestValue(a))
	//==============================================================

	// MAX VALUE OF A STRING IN AN ARRAY
	// fmt.Println(maximumValue([]string{"5232", "yv", "22", "c", "yawgs", "928", "4003", "2"}))
	//===============================================================

	// SEARCH INSERT INDEX
	//fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
	//================================================================

	// PLUSE ONE
	//fmt.Println(plusOne([]int{9}))
	//===============================================================

	// MERGE TWO SORTED ARRAYS
	// nums1 := []int{1, 2, 3, 0, 0, 0}
	// merge(nums1, 3, []int{2, 5, 6}, 3)
	// fmt.Println(nums1)
	//=================================================================

	// PASCAL`S TRIANGLE
	// fmt.Println(generate(5))
	//===============================================================

	// PASCAL`S TRIANGLE 2
	// fmt.Println(getRow(5))
	//==================================================================

	// BEST TIME TO BUY AND SEND STOCK
	// fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	//================================================================

	// SINGLE NUMBER
	// fmt.Println(singleNumber([]int{2, 1, 4, 1, 2}))
	//==================================================================

	// MAJORITY ELEMENT
	// fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
	//==================================================================

	// CONTAINS DUPLICATE
	// fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	//====================================================================

	// CONTAINS DUPLICATE 2
	// fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	//==================================================================

	// SUMMARY RANGES
	// fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
	//====================================================================

	// MISSING NUBER
	// fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	//===================================================================

	// MOVE ZEROES
	// moveZeroes([]int{1, 0, 0, 2, 0, 3, 4, 5})
	//===================================================================

	// RANGE SUM QUERY - IMMUTABLE
	//===================================================================

	// INSPECTION OF TWO ARRAYS
	// fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
	//====================================================================

	// INSPECTION OF TWO ARRAYS 2
	// fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	// fmt.Println(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4, 7}))
	//====================================================================

	// THIRD MAXIMUM NUMBER
	// fmt.Println(thirdMax([]int{3, 2, 1}))
	// fmt.Println(thirdMax([]int{1, 1, 2}))
	//=======================================================================

	// FIND ALL NUMBERS DISAPPEARED IN AN ARRAY
	// fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	//======================================================================

	// ISLAND PERIMETER
	// fmt.Println(islandPerimeter([][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}))
	//=========================================================================

	// MAX CONSECUNTIVE ONES
	// fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
	//===========================================================================

	// TEEMO ATTACKING
	// fmt.Println(findPoisonedDuration([]int{1, 4}, 2))
	//============================================================================

	// NEXT GREATER ELEMENT
	// fmt.Println(nextGreaterElement([]int{1, 3, 5, 2, 4}, []int{6, 5, 4, 3, 2, 1, 7}))
	//======================================================================================

	// KEYBOARD ROW
	// fmt.Println(findWords([]string{"asdfghjklASDFGHJKL", "qwertyuiopQWERTYUIOP", "zxcvbnmZXCVBNM"}))
	//====================================================================================================

	// BINARY SEARCH
	// fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))

	// RELETIVE RANKS
	// fmt.Println(findRelativeRanks([]int{5, 4, 3, 2, 1}))

	// ARRAY PARTITION
	// fmt.Println(arrayPairSum([]int{6, 2, 6, 5, 1, 2}))

	// RESHAPE THE MATRIX
	// fmt.Println(matrixReshape([][]int{{1, 2}, {3, 4}}, 1, 4))

	// DISTRIBUTE CANDIES
	// fmt.Println(distributeCandies([]int{1, 1, 2, 2, 3, 3}))

	// LONGEST HARMONIUS SUBSEQUENCE

	// MERGE TWO SORTED LISTS

}

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// func inorderTraversal(root *TreeNode) []int {

// }

// REVERSE LINKED LIST
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func initList(arr []int) *ListNode {
// 	resultNode := new(ListNode)
// 	r := resultNode

// 	for _, val := range arr {
// 		r.Next = &ListNode{Val: val, Next: nil}
// 		r = r.Next
// 	}

// 	return resultNode.Next
// }

// func printList(head *ListNode) {
// 	for head != nil {
// 		fmt.Print(head, " ")
// 		head = head.Next
// 	}
// 	fmt.Println()
// }

// func reverseList(head *ListNode) *ListNode {
// 	if head == nil {
// 		return nil
// 	}

// 	currentNode := head
// 	var resultList *ListNode

// 	for currentNode != nil {
// 		tempNode := currentNode.Next
// 		currentNode.Next = resultList
// 		resultList = currentNode
// 		currentNode = tempNode
// 	}

// 	return resultList
// }

// REMOVE LINKED LIST ELEMENTS
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func initList(arr []int) *ListNode {
// 	resultNode := new(ListNode)
// 	r := resultNode

// 	for _, val := range arr {
// 		r.Next = &ListNode{Val: val, Next: nil}
// 		r = r.Next
// 	}

// 	return resultNode.Next
// }

// func printList(head *ListNode) {
// 	for head != nil {
// 		fmt.Print(head, " ")
// 		head = head.Next
// 	}
// 	fmt.Println()
// }

// func removeElements(head *ListNode, val int) *ListNode {
// 	var previousNode *ListNode
// 	currentNode := head

// 	for currentNode != nil {
// 		if currentNode.Val == val {
// 			if previousNode != nil {
// 				previousNode.Next = currentNode.Next
// 			} else {
// 				head = currentNode.Next
// 			}
// 		} else {
// 			previousNode = currentNode
// 		}
// 		currentNode = currentNode.Next
// 	}

// 	return head
// }

// func removeElements(head *ListNode, val int) *ListNode {
// 	if head == nil {
// 		return nil
// 	}
// 	result := &ListNode{Val: 0, Next: nil}
// 	var r = result
// 	for head != nil {
// 		fmt.Println(r, result)
// 		if head.Val != val {
// 			r.Next = &ListNode{Val: head.Val, Next: nil}
// 			r = r.Next
// 		}
// 		head = head.Next
// 	}

// 	return result.Next
// }

// func removeElements(head *ListNode, val int) *ListNode {
// 	if head == nil {
// 		return nil
// 	}
// 	// h := &ListNode{Val: 0, Next: head}
// 	h := new(ListNode)
// 	h.Next = head
// 	for h.Next != nil {
// 		fmt.Println(h, h.Next, h.Next.Next)
// 		if h.Next.Val == val {
// 			h.Next = h.Next.Next
// 		} else {
// 			h = h.Next
// 		}
// 	}
// 	return head
// }

// INTERSECTION OF TWO LINKED LISTS
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func initList(arr []int) *ListNode {
// 	resultNode := new(ListNode)
// 	r := resultNode

// 	for _, val := range arr {
// 		r.Next = &ListNode{Val: val, Next: nil}
// 		r = r.Next
// 	}

// 	return resultNode.Next
// }

// func printList(head *ListNode) {
// 	for head != nil {
// 		fmt.Print(head.Val, " ")
// 		head = head.Next
// 	}
// 	fmt.Println()
// }

// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	check := make(map[*ListNode]int)

// 	for headA != nil {
// 		check[headA] = 1
// 		headA = headA.Next
// 	}

// 	for headB != nil {
// 		if check[headB] == 1 {
// 			return headB
// 		}
// 		headB = headB.Next
// 	}

// 	return nil
// }

// LINKED LIST CYCLE
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func initList(arr []int) *ListNode {
// 	resultNode := new(ListNode)
// 	r := resultNode

// 	for _, val := range arr {
// 		r.Next = &ListNode{Val: val, Next: nil}
// 		r = r.Next
// 	}

// 	return resultNode.Next
// }

// func printList(head *ListNode) {
// 	for head != nil {
// 		fmt.Print(head.Val, " ")
// 		head = head.Next
// 	}
// 	fmt.Println()
// }

// func hasCycle(head *ListNode) bool {
// 	check := make(map[*ListNode]int)

// 	for head != nil {
// 		if check[head] == 0 {
// 			check[head] = 1
// 		} else {
// 			return true
// 		}
// 		head = head.Next
// 	}

// 	return false
// }

// REMOVE DUPLICATES FROM SORTED LIST
// Definition for singly-linked list.
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func initList(arr []int) *ListNode {
// 	resultNode := new(ListNode)
// 	r := resultNode

// 	for _, val := range arr {
// 		r.Next = &ListNode{Val: val, Next: nil}
// 		r = r.Next
// 	}

// 	return resultNode.Next
// }

// func printList(head *ListNode) {
// 	for head != nil {
// 		fmt.Print(head.Val, " ")
// 		head = head.Next
// 	}
// 	fmt.Println()
// }

// func deleteDuplicates(head *ListNode) *ListNode {
// 	if head == nil {
// 		return nil
// 	}

// 	tempNode := head

// 	for tempNode.Next != nil {
// 		if tempNode.Val == tempNode.Next.Val {
// 			new := tempNode.Next.Next
// 			tempNode.Next = nil
// 			tempNode.Next = new
// 		} else {
// 			tempNode = tempNode.Next
// 		}
// 	}

// 	return head
// }

// func deleteDuplicates(head *ListNode) *ListNode {
// 	if head == nil {
// 		return nil
// 	}

// 	resultNode := &ListNode{Val: head.Val, Next: nil}
// 	r := resultNode

// 	for head != nil {
// 		if r.Val != head.Val {
// 			r.Next = &ListNode{Val: head.Val, Next: nil}
// 			r = r.Next
// 		}
// 		head = head.Next
// 	}

// 	return resultNode
// }

// MERGE TWO SORTED LISTS
// Definition for singly-linked list.
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	var resultNode *ListNode
// 	r := resultNode

// 	for list1 != nil || list2 != nil {
// 		if list1 == nil {
// 			resultNode = list2
// 			list2 = list2.Next
// 			continue
// 		}
// 		if list2 == nil {
// 			resultNode = list1
// 			list1 = list1.Next
// 			continue
// 		}

// 		if list1.Val < list2.Val {
// 			resultNode = list1
// 			list1 = list1.Next
// 		} else {
// 			resultNode = list2
// 			list2 = list2.Next
// 		}
// 	}

// 	return r
// }

// LONGEST HARMONIUS SUBSEQUENCE
// func findLHS(nums []int) int {

// }

// DISTRIBUTE CANDIES
// func distributeCandies(candyType []int) int {
// 	setCandyType := make(map[int]bool)

// 	for _, ct := range candyType {
// 		setCandyType[ct] = true
// 	}

// 	if len(candyType)/2 < len(setCandyType) {
// 		return len(candyType) / 2
// 	}

// 	return len(setCandyType)
// }

// RESHAPE THE MATRIX
// func matrixReshape(mat [][]int, r int, c int) [][]int {
// 	l0 := len(mat)
// 	l1 := len(mat[0])
// 	l := l0 * l1

// 	if r*c != l {
// 		return mat
// 	}

// 	result := make([][]int, r)
// 	for i := 0; i < r; i++ {
// 		result[i] = make([]int, c)
// 	}

// 	for i := 0; i < l; i++ {
// 		result[i/c][i%c] = mat[i/l1][i%l1]
// 	}

// 	return result
// }

// func matrixReshape(mat [][]int, r int, c int) [][]int {
// 	if len(mat)*len(mat[0]) != r*c {
// 		return mat
// 	}

// 	l0 := len(mat)
// 	l1 := len(mat[0])

// 	result := make([][]int, r)
// 	for i := 0; i < r; i++ {
// 		result[i] = make([]int, c)
// 	}

// 	// count := 0

// 	// if r == 1 || c == 1 {
// 	// 	count = r * c
// 	// } else {
// 	// 	count = r + c
// 	// }

// 	for i := 0; i < r*c; i++ {
// 		it := 0
// 		jt := 0
// 		if i < l0 {
// 			if l1 == 1 {
// 				it = i
// 			} else {
// 				it = 0
// 			}
// 		} else {
// 			if l0 == 1 {
// 				it = 0
// 			} else {
// 				it = 1
// 			}
// 		}
// 		if i < l0 {
// 			if l1 == 1 {
// 				jt = 0
// 			} else {
// 				jt = i
// 			}
// 		} else {
// 			if l0 == 1 {
// 				jt = i - l0 + 1
// 			} else {
// 				jt = i - l0
// 			}
// 		}

// 		rit := 0
// 		rjt := 0
// 		if i < r {
// 			if c == 1 {
// 				rit = i
// 			} else {
// 				rit = 0
// 			}
// 		} else {
// 			if r == 1 {
// 				rit = 0
// 			} else {
// 				rit = 1
// 			}
// 		}
// 		if i < r {
// 			if c == 1 {
// 				rjt = 0
// 			} else {
// 				rjt = i
// 			}
// 		} else {
// 			if r == 1 {
// 				rjt = i - r + 1
// 			} else {
// 				rjt = i - r
// 			}
// 		}
// 		fmt.Println(it, jt, "    ", rit, rjt)
// 		// result[rit][rjt] = mat[it][jt]
// 	}

// 	return result
// }

// ARRAY PARTITION
// func arrayPairSum(nums []int) int {
// 	sort.Slice(nums, func(i, j int) bool {
// 		return nums[i] < nums[j]
// 	})

// 	sum := 0

// 	for i := 0; i < len(nums); i += 2 {
// 		sum += nums[i]
// 	}

// 	return sum
// }

// RELETIVE RANKS
// func findRelativeRanks(score []int) []string {
// 	relativeRanks := make([]string, len(score))
// 	newScore := make([]int, len(score))

// 	copy(newScore, score)

// 	sort.Slice(newScore, func(i, j int) bool {
// 		return newScore[i] > newScore[j]
// 	})

// 	linkScoreRank := make(map[int]int)

// 	for i, val := range newScore {
// 		linkScoreRank[val] = i
// 	}

// 	for i, val := range score {
// 		if linkScoreRank[val] == 0 {
// 			relativeRanks[i] = "Gold Medal"
// 		} else if linkScoreRank[val] == 1 {
// 			relativeRanks[i] = "Silver Medal"
// 		} else if linkScoreRank[val] == 2 {
// 			relativeRanks[i] = "Bronze Medal"
// 		} else {
// 			relativeRanks[i] = strconv.Itoa(linkScoreRank[val] + 1)
// 		}
// 	}

// 	return relativeRanks
// }

// BINARY SEARCH
// func search(nums []int, target int) int {
// 	leftBorder := 0
// 	rightBorder := len(nums) - 1
// 	var middle int

// 	for leftBorder <= rightBorder {
// 		middle = (leftBorder + rightBorder) / 2

// 		if nums[middle] == target {
// 			return middle
// 		} else if nums[middle] > target {
// 			rightBorder = middle - 1
// 		} else if nums[middle] < target {
// 			leftBorder = middle + 1
// 		}
// 	}

// 	return -1
// }

// KEYBOARD ROW
// func findWords(words []string) []string {
// 	availableWards := []string{}

// 	if len(words) == 0 {
// 		return availableWards
// 	}

// 	oneS := []rune{'Q', 'W', 'E', 'R', 'T', 'Y', 'U', 'I', 'U', 'P', 'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p'}

// 	one := make(map[rune]int)
// 	for _, s := range oneS {
// 		one[s] = 1
// 	}

// 	twoS := []rune{'A', 'S', 'D', 'F', 'G', 'H', 'J', 'K', 'L', 'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l'}

// 	two := make(map[rune]int)
// 	for _, s := range twoS {
// 		two[s] = 2
// 	}

// 	threeS := []rune{'Z', 'X', 'C', 'V', 'B', 'N', 'M', 'z', 'x', 'c', 'v', 'b', 'n', 'm'}

// 	three := make(map[rune]int)
// 	for _, s := range threeS {
// 		three[s] = 3
// 	}

// 	for _, word := range words {
// 		countOne := 0
// 		countTwo := 0
// 		countThree := 0
// 		for _, letter := range word {
// 			countOne += one[letter]
// 			countTwo += two[letter]
// 			countThree += three[letter]
// 		}
// 		fmt.Println(len(word), countOne, countTwo, countThree)
// 		if countOne+countTwo == 0 || countOne+countThree == 0 || countTwo+countThree == 0 {
// 			availableWards = append(availableWards, word)
// 		}
// 	}
//=================================================================================================

// 	return availableWards
// }

// NEXT GREATER ELEMENT
// func nextGreaterElement(nums1 []int, nums2 []int) []int {
// 	result := []int{}
// 	for _, v := range nums1 {
// 		isV := false
// 		isExist := false
// 		for _, val := range nums2 {
// 			if v == val {
// 				isV = true
// 			}
// 			if isV {
// 				if val > v {
// 					result = append(result, val)
// 					isExist = true
// 					break
// 				}
// 			}
// 		}
// 		if !isExist {
// 			result = append(result, -1)
// 		}
// 	}
// 	return result
// }
//==============================================================================

// if j < len(nums2)-1 {
// 	if v < nums2[j+1] {
// 		result = append(result, nums2[j+1])
// 	} else {
// 		result = append(result, -1)
// 	}
// } else {
// 	result = append(result, -1)
// }

// TEEMO ATTACKING
// func findPoisonedDuration(timeSeries []int, duration int) int {
// 	if len(timeSeries) == 0 {
// 		return 0
// 	}

// 	secondsDuration := 0

// 	for i := 0; i < len(timeSeries)-1; i++ {
// 		if timeSeries[i+1]-timeSeries[i] >= duration {
// 			secondsDuration += duration
// 		} else {
// 			secondsDuration += timeSeries[i+1] - timeSeries[i]
// 		}
// 	}

// 	return secondsDuration + duration
// }
//======================================================================

// MAX CONSECUNTIVE ONES
// func findMaxConsecutiveOnes(nums []int) int {
// 	maxOnes := 0
// 	maxTempOnes := 0

// 	for _, v := range nums {
// 		if v == 1 {
// 			maxTempOnes++
// 		} else {
// 			if maxTempOnes > maxOnes {
// 				maxOnes = maxTempOnes
// 			}
// 			maxTempOnes = 0
// 		}
// 	}

// 	if maxTempOnes > maxOnes {
// 		maxOnes = maxTempOnes
// 	}

// 	return maxOnes
// }
//====================================================================

// ISLAND PERIMETER
// func islandPerimeter(grid [][]int) int {
// 	countFacet := 0
// 	countIntersection := 0

// 	for i, v := range grid {
// 		for j, val := range v {
// 			if val == 1 {
// 				countFacet += 4
// 				if j < len(v)-1 {
// 					if grid[i][j+1] == 1 {
// 						countIntersection += 2
// 					}
// 				}
// 				if i < len(grid)-1 {
// 					if grid[i+1][j] == 1 {
// 						countIntersection += 2
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return countFacet - countIntersection
// }
//==================================================================

// FIND ALL NUMBERS DISAPPEARED IN AN ARRAY
// func findDisappearedNumbers(nums []int) []int {
// m := make(map[int]int)

// for i, v := range nums {
// 	if m[i+1] != 1 {
// 		m[i+1] = 0
// 	}
// 	m[v] = 1
// }

// fmt.Println(m)

// disappearedNumers := []int{}

// for k, v := range m {
// 	if v == 0 {
// 		disappearedNumers = append(disappearedNumers, k)
// 	}
// }

// return disappearedNumers
// }
//===========================================================================

// THIRD MAXIUM NUMBER
// func thirdMax(nums []int) int {
// 	firstMaximum := math.MinInt32
// 	secondMaximum := math.MinInt32
// 	thirdMaximum := math.MinInt32
// 	isThirdMaximum := false

// 	for _, num := range nums {
// 		if num > firstMaximum {
// 			firstMaximum = num
// 		}
// 	}

// 	for _, num := range nums {
// 		if num > secondMaximum && num < firstMaximum {
// 			secondMaximum = num
// 		}
// 	}

// 	for _, num := range nums {
// 		if num > thirdMaximum && num < secondMaximum {
// 			thirdMaximum = num
// 			isThirdMaximum = true
// 		}
// 	}

// 	if !isThirdMaximum {
// 		return firstMaximum
// 	}

// 	return thirdMaximum
// }
//================================================================

// INSPECTION OF TWO ARRAYS 2
// func intersect(nums1 []int, nums2 []int) []int {
// 	m1 := make(map[int]int)
// 	inter := []int{}

// 	for _, num := range nums1 {
// 		m1[num] += 1
// 	}

// 	for _, num := range nums2 {
// 		if m1[num] != 0 {
// 			m1[num] -= 1
// 			inter = append(inter, num)
// 		}
// 	}

// for k := range m1 {
// 	if m1[k] != 0 && m2[k] != 0 {
// 		count := 0
// 		if m1[k] > m2[k] {
// 			count = m2[k]
// 		} else {
// 			count = m1[k]
// 		}

// 		for i := 0; i < count; i++ {
// 			inter = append(inter, k)
// 		}
// 	}
// }

// 	return inter
// }
//==================================================================

// INSPECTION OF TWO ARRAYS
// func intersection(nums1 []int, nums2 []int) []int {
// 	m := make(map[int]int)
// 	result := []int{}

// 	for _, num := range nums1 {
// 		m[num] = 1
// 	}

// 	for _, num := range nums2 {
// 		if m[num] == 1 {
// 			m[num] = 2
// 		}
// 	}

// 	for k := range m {
// 		if m[k] == 2 {
// 			result = append(result, k)
// 		}
// 	}

// 	return result
// }
//=======================================================================

// RANGE SUM QUEERY - IMMUTABLE
// type NumArray struct {
// 	arr []int
// }

// func Constructor(nums []int) NumArray {
// 	return NumArray{arr: nums}
// }

// func (this *NumArray) SumRange(left int, right int) int {
// 	sum := 0

// 	for i := left; i <= right; i++ {
// 		sum += this.arr[i]
// 	}

// 	return sum
// }
//==========================================================================

// MOVE ZEROES
// func moveZeroes(nums []int) {
// 	counter_zeroes := 0

// 	for iterator := range nums {
// 		if nums[iterator] == 0 {
// 			counter_zeroes++
// 		} else {
// 			nums[iterator], nums[iterator-counter_zeroes] = nums[iterator-counter_zeroes], nums[iterator]
// 		}
// 	}
// }

// MISSING NUMBER
// func missingNumber(nums []int) int {
// 	target := 0
// 	current := 0

// 	for iterator, num := range nums {
// 		target += iterator
// 		current += num
// 	}

// 	return target + len(nums) - current
// }
//===================================================================================

// SUMMARY RANGES
// func summaryRanges(nums []int) []string {
// 	if len(nums) == 0 {
// 		return []string{}
// 	}

// 	var s []string
// 	start, end := nums[0], nums[0]

// 	for i := 1; i < len(nums); i++ {
// 		if nums[i]-end == 1 {
// 			end = nums[i]
// 		} else {
// 			if start == end {
// 				s = append(s, strconv.Itoa(start))
// 			} else {
// 				s = append(s, strconv.Itoa(start)+"->"+strconv.Itoa(end))
// 			}
// 			start, end = nums[i], nums[i]
// 		}
// 	}

// 	if start == end {
// 		s = append(s, strconv.Itoa(start))
// 	} else {
// 		s = append(s, strconv.Itoa(start)+"->"+strconv.Itoa(end))
// 	}

// 	return s
// }
//==============================================================================

// CONTAINS DUPLICATE 2
// func containsNearbyDuplicate(nums []int, k int) bool {
// 	m := make(map[int]int)

// 	for iterator, num := range nums {
// 		if m[num] == 0 {
// 			m[num] = iterator + 1
// 		} else {
// 			temp := (m[num] - 1) - iterator

// 			if temp < 0 {
// 				temp = -temp
// 			}

// 			if temp <= k {
// 				return true
// 			} else {
// 				m[num] = iterator + 1
// 			}
// 		}
// 	}

// 	return false
// }
//========================================================================

// CONTAINS DUPLICATE
// func containsDuplicate(nums []int) bool {
// 	m := make(map[int]int)

// 	for _, value := range nums {
// 		if m[value] == 0 {
// 			m[value] = 1
// 		} else if m[value] == 1 {
// 			return true
// 		}
// 	}

// 	return false
// }
//======================================================================

// MAJORITY ELEMENT
// func majorityElement(nums []int) int {
// 	major := nums[0]
// 	counter := 1

// 	for i := 1; i < len(nums); i++ {
// 		if counter == 0 {
// 			major = nums[i]
// 			counter = 1
// 		} else if major == nums[i] {
// 			counter++
// 		} else {
// 			counter--
// 		}
// 	}

// 	return major

// m := make(map[int]int)

// for _, value := range nums {
// 	if isIn(value, m) {
// 		m[value] += 1
// 		if m[value] > len(nums)/2 {
// 			return value
// 		}
// 	} else {
// 		m[value] = 1
// 	}
// }

// return 0
// }

// func isIn(num int, m map[int]int) bool {
// 	for k := range m {
// 		if num == k {
// 			return true
// 		}
// 	}
// 	return false
// }

// func getMax(m map[int]int) int {
// 	maxVal := 0
// 	maxKey := 0
// 	for k, v := range m {
// 		if v > maxVal {
// 			maxVal = v
// 			maxKey = k
// 		}
// 	}

// 	return maxKey
// }
//===========================================================

// SINGLE NUMBER
// func singleNumber(nums []int) int {
// 	result := 0

// 	for _, value := range nums {
// 		result = result ^ value
// 	}

// 	return result

// sort.Slice(nums, func(i, j int) bool {
// 	return nums[i] < nums[j]
// })

// for i := 0; i < len(nums)-1; i += 2 {
// 	if nums[i] != nums[i+1] {
// 		return nums[i]
// 	}
// }

// return nums[len(nums)-1]
// }
//================================================================

// BEST TIME TO BUY AND SEND STOCK
// func maxProfit(prices []int) int {
// 	minValue := 2147483647
// 	maxValue := 0
// 	minIndex := 0
// 	maxIndex := 0

// 	maxDiff := []int{}

// 	for index, value := range prices {
// 		if value < minValue {
// 			minValue = value
// 			maxValue = 0
// 			minIndex = index
// 		}

// 		if value > maxValue {
// 			maxValue = value
// 			maxIndex = index
// 			if minIndex != maxIndex {
// 				maxDiff = append(maxDiff, maxValue-minValue)
// 			}
// 		}
// 	}

// 	maxResult := 0

// 	for _, val := range maxDiff {
// 		if val > maxResult {
// 			maxResult = val
// 		}
// 	}

// 	return maxResult
// }
//=================================================================

// PASCAL'S TRIANGLE 2
// func getRow(rowIndex int) []int {
// 	tempSlice := make([]int, rowIndex+1)
// 	tempSlice[0] = 1

// 	for i := 0; i <= rowIndex; i++ {
// 		for j := i; j > 0; j-- {
// 			tempSlice[j] += tempSlice[j-1]
// 		}
// 	}

// 	return tempSlice

// if rowIndex == 0 {
// 	return []int{1}
// }

// pascalsTriangle := [][]int{{1}}

// for i := 1; i <= rowIndex; i++ {
// 	var rowPascalsTriangle []int

// 	for j := 0; j <= i; j++ {
// 		if j == 0 {
// 			rowPascalsTriangle = append(rowPascalsTriangle, 1)
// 			continue
// 		}

// 		if j == i {
// 			rowPascalsTriangle = append(rowPascalsTriangle, 1)
// 			continue
// 		}

// 		rowPascalsTriangle = append(rowPascalsTriangle, pascalsTriangle[i-1][j-1]+pascalsTriangle[i-1][j])
// 	}

// 	pascalsTriangle = append(pascalsTriangle, rowPascalsTriangle)
// }

// return pascalsTriangle[rowIndex]
// }
//=================================================================================

// PASCAL`S TRIANGLE
// func generate(numRows int) [][]int {
// 	if numRows == 0 {
// 		return [][]int{}
// 	}

// 	pascalsTriangle := [][]int{{1}}

// 	for i := 1; i < numRows; i++ {
// 		var tempSlice []int

// 		for j := 0; j <= i; j++ {
// 			fmt.Println("j = ", j)

// 			if j == 0 {
// 				tempSlice = append(tempSlice, pascalsTriangle[i-1][j])
// 				continue
// 			}

// 			if j == i {
// 				tempSlice = append(tempSlice, pascalsTriangle[i-1][j-1])
// 				break
// 			}

// 			tempSlice = append(tempSlice, pascalsTriangle[i-1][j-1]+pascalsTriangle[i-1][j])
// 		}

// 		pascalsTriangle = append(pascalsTriangle, tempSlice)
// 	}

// 	return pascalsTriangle
// }
//=========================================================

// MERGE TWO SORTED ARRAYS
// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	nums1Iterator := m - 1
// 	nums2Iterator := n - 1
// 	resultIterator := n + m - 1

// 	for nums1Iterator > -1 && nums2Iterator > -1 {
// 		if nums1[nums1Iterator] > nums2[nums2Iterator] {
// 			nums1[resultIterator] = nums1[nums1Iterator]
// 			resultIterator--
// 			nums1Iterator--
// 		} else {
// 			nums1[resultIterator] = nums2[nums2Iterator]
// 			resultIterator--
// 			nums2Iterator--
// 		}
// 	}

// 	for nums2Iterator > -1 {
// 		nums1[resultIterator] = nums2[nums2Iterator]
// 		resultIterator--
// 		nums2Iterator--
// 	}

// 	// var tempNums1 []int

// 	// for i := 0; i < m; i++ {
// 	// 	tempNums1 = append(tempNums1, nums1[i])
// 	// }

// 	// nums1Iterator := 0
// 	// nums2Iterator := 0
// 	// tempNums1Iterator := 0

// 	// for tempNums1Iterator < m && nums2Iterator < n {
// 	// 	if tempNums1[tempNums1Iterator] <= nums2[nums2Iterator] {
// 	// 		nums1[nums1Iterator] = tempNums1[tempNums1Iterator]
// 	// 		tempNums1Iterator++
// 	// 		nums1Iterator++
// 	// 	} else {
// 	// 		nums1[nums1Iterator] = nums2[nums2Iterator]
// 	// 		nums2Iterator++
// 	// 		nums1Iterator++
// 	// 	}
// 	// }

// 	// if tempNums1Iterator == m {
// 	// 	for ; nums2Iterator < n; nums2Iterator++ {
// 	// 		nums1[nums1Iterator] = nums2[nums2Iterator]
// 	// 		nums1Iterator++
// 	// 	}
// 	// } else if nums2Iterator == n {
// 	// 	for ; tempNums1Iterator < m; tempNums1Iterator++ {
// 	// 		nums1[nums1Iterator] = tempNums1[tempNums1Iterator]
// 	// 		nums1Iterator++
// 	// 	}
// 	// }
// }
//=================================================================

// PLUSE ONE
// func plusOne(digits []int) []int {
// 	for i := len(digits) - 1; i >= 0; i-- {
// 		temp := digits[i] + 1
// 		if temp != 10 {
// 			digits[i] += 1
// 			break
// 		}
// 		digits[i] = 0
// 	}

// 	if digits[0] == 0 {
// 		digits = append([]int{1}, digits...)
// 	}

// 	return digits
// }
//==============================================================

// SEARCH INSERT INDEX
// func searchInsert(nums []int, target int) int {
// 	leftBorder := 0
// 	rightBorder := len(nums) - 1
// 	var middle int

// 	for leftBorder <= rightBorder {
// 		middle = (leftBorder + rightBorder) / 2

// 		if nums[middle] == target {
// 			return middle
// 		} else if target < nums[middle] {
// 			rightBorder = middle - 1
// 		} else if target > nums[middle] {
// 			leftBorder = middle + 1
// 		}
// 	}

// 	if nums[middle] > target {
// 		return middle
// 	} else {
// 		return middle + 1
// 	}
// }
//======================================================================

// MAX VALUE OF A STRING IN AN ARRAY
// func maximumValue(strs []string) int {
// 	maxValue := 0
// 	for _, str := range strs {
// 		value := getStringValue(str)
// 		if value > maxValue {
// 			maxValue = value
// 		}
// 	}

// 	return maxValue
// }

// func getStringValue(str string) int {
// 	for _, s := range str {
// 		if !unicode.IsDigit(s) {
// 			return utf8.RuneCountInString(str)
// 		}
// 	}
// 	maxDigit, _ := strconv.Atoi(str)
// 	return maxDigit
// }
//=======================================================

// DELETE GRAITEST VALUE IN EACH ROW
// func deleteGraitestValue(grid [][]int) int {
// 	sumMaxVal := 0
// 	for len(grid[0]) > 0 {
// 		maxValCol := 0
// 		for i, val := range grid {
// 			maxValRow, iterator := getMax(val)
// 			grid[i] = append(grid[i][:iterator], grid[i][iterator+1:]...)
// 			if maxValRow > maxValCol {
// 				maxValCol = maxValRow
// 			}
// 		}
// 		sumMaxVal += maxValCol
// 	}

// 	return sum

// func getMax(s []int) (int, int) {
// 	result := 0
// 	it := 0

// 	for i, val := range s {
// 		if val > result {
// 			result = val
// 			it = i
// 		}
// 	}

// 	return result, it
// }
//====================================================================
