package homework01

import (
	"fmt"
	"sort"
	"strconv"
)

// 1. 只出现一次的数字
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func SingleNumber(nums []int) int {
	//count:统计数字个数
	count := 0
	//target:目标数字
	target := 0
	map1 := make(map[int]int)
	for _, v := range nums {
		_, ok := map1[v]
		if ok {
			map1[v] = v + 1
		} else {
			map1[v] = count + 1
		}
	}
	for k, v := range map1 {
		if v == 1 {
			target = k
			break
		}
	}
	return target
}

// 2. 回文数
// 判断一个整数是否是回文数
func IsPalindrome(x int) bool {
	//reverseNumber:反转数
	reverseNumber := 0
	//oriNumber:原始数据
	oriNumber := x
	//mod:取余
	mod := 0
	if x < 0 {
		return false
	} else {
		for {
			if x > 0 {
				//计算反转后的数字
				mod = x % 10
				x /= 10
				reverseNumber = reverseNumber*10 + mod
			} else {
				break
			}
		}
		fmt.Println("原始数字为:", oriNumber)
		fmt.Println("反转后的数字为:", reverseNumber)
		return reverseNumber == oriNumber
	}
}

// 3. 有效的括号
// 给定一个只包括 '(', ')', '{', '}', '[', ']' 的字符串，判断字符串是否有效
func IsValid(s string) bool {
	//stack:存放左括号('(','{','[')的切片
	stack := make([]rune, 0)
	//pairs:存放左括号与右括号的映射
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	if len(s)%2 != 0 {
		//若字符长度不是偶数,说明不匹配
		return false
	} else {
		for _, char := range s {
			//若字符是左括号('(','{','[')任意之一,加入到切片中
			if char == '(' || char == '[' || char == '{' {
				stack = append(stack, char)
			} else {
				//若字符是右括号(')','}',']')任意之一
				if len(stack) == 0 {
					//若此时stack切片为空,说明还没有左括号,说明不匹配
					return false
				}
				//取出stack切片中最后一位字符
				top := stack[len(stack)-1]
				//更新stack切片
				stack = stack[:len(stack)-1]
				//判断stack切片中最后一位字符(左括号)是否等于映射的字符(左括号)
				if top != pairs[char] {
					//若不相等,说明不匹配
					return false
				}
			}
		}
		//若len(stack)==0,说明匹配
		//若len(stack)!=0,说明不匹配
		return len(stack) == 0
	}
}

// 4. 最长公共前缀
// 查找字符串数组中的最长公共前缀
func LongestCommonPrefix(strs []string) string {
	//minLen 最短字符串的长度
	minLen := len(strs[0])
	//minIndex 最短字符串的索引
	minIndex := 0
	//找出最短的字符串
	for k, v := range strs {
		if len(v) < minLen {
			minLen = len(v)
			minIndex = k
		}
	}
	//prefix:以最短字符串作为前缀
	prefix := strs[minIndex]
	for i := 0; i < len(strs); i++ {
		for j := 0; j < minLen; j++ {
			if len(prefix) == 0 {
				return ""
			}
			//若某个字符串中的某个字符与最短字符串的对应字符不匹配
			//取该字符之前的字符串作为最长公共前缀并退出循环
			if strs[i][j] != strs[minIndex][j] {
				prefix = prefix[:j]
				break
			}
		}
	}
	return prefix
}

// 5. 加一
// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func PlusOne(digits []int) []int {
	str := ""
	//将数组转为string
	for _, v := range digits {
		str += strconv.Itoa(v)
	}
	//将string转为int
	i, err := strconv.Atoi(str)
	if err == nil {
		i += 1
	}
	//将int转为string
	str = strconv.Itoa(i)
	var nums = make([]int, len(str))
	//string再转为[]int
	for k, v := range str {
		i, _ := strconv.Atoi(fmt.Sprintf("%c", v))
		nums[k] = i
	}
	return nums
}

// 6. 删除有序数组中的重复项
// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
func RemoveDuplicates(nums []int) int {
	lenNums := len(nums)
	switch lenNums {
	case 0:
		return 0
	case 1:
		return 1
	default:
		slow := 0
		for fast := 1; fast < lenNums; fast++ {
			if nums[fast] != nums[slow] {
				slow++
				nums[slow] = nums[fast]
			}
		}
		return slow + 1
	}
}

// 7. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
func Merge(intervals [][]int) [][]int {
	//先按照起始位置将区间从小到大排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	//result:结果切片
	result := make([][]int, 0)
	//current:当前区间
	current := intervals[0]
	for i := 1; i < len(intervals); i++ {
		//如果下一个区间的起始位置小于等于当前区间的结束位置,说明两个区间有重叠
		if intervals[i][0] <= current[1] {
			//如果下一个区间的结束位置大于等于当前区间的起始位置
			if intervals[i][1] >= current[1] {
				//合并区间:将下一个区间的结束位置赋给当前区间的结束位置
				current[1] = intervals[i][1]
			}
		} else {
			//没有重叠,将区间直接加入结果切片
			result = append(result, current)
			//遍历下一个区间
			current = intervals[i]
		}
	}
	//将最后一个区间加入结果数组
	result = append(result, current)
	return result
}

// 8. 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func TwoSum(nums []int, target int) []int {
	numsIndex1 := 0
	numsIndex2 := 0
	flag := false
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				fmt.Println("i=", i, "j=", j)
				numsIndex1 = i
				numsIndex2 = j
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}
	if numsIndex1 > numsIndex2 {
		return []int{numsIndex2, numsIndex1}
	} else {
		return []int{numsIndex1, numsIndex2}
	}
}
