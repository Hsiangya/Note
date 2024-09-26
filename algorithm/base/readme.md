# 基础算法

## 快速排序

### Hore

通常被称为**双指针分区**或**Hoare分区**，在某些情况下，比Lomuto分区方案更高效，因为它减少了交换次数并且更好的处理了重复元素

1. **选择基准**：通常选择第一个或中间或最后一个元素作为基准(pivot)，
2. **初始化指针**：初始化两个指针，left指向数组的开始，right指向数组的结束
3. **双指针遍历**：
   - left指针从左向右移动，直到找到一个不小于pivot的元素
   - right指针从右向左移动，直到找到一个不大于pivot的元素
   - 如果left<=right，交换left与right指向的元素，然后left与right分别移动一位
4. **递归调用**：
   - 一旦完成分区，递归地对基准左边和有便地子数组进行排序

![image-20240429213408227](./assets/image-20240429213408227.png)

```go
func HoareQuickSort(nums []int, start int, end int) {
	// start>end 表示只有一个元素 已经是有序的
	if start >= end {
		return
	}

	// 选择数组中间的数作为pivot
	pivot := nums[(start+end)/2]
	left, right := start, end // 初始化指针
	for left <= right {
		for nums[left] < pivot {
			left++
		}
		for nums[right] > pivot {
			right--
		}
		// 交换元素
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	// 递归对左右区间进行排序
	HoareQuickSort(nums, start, right)
	HoareQuickSort(nums, left, end)
}
```

## Lomuto分区

Lomuto分区算法是快速排序中常用的一种分区方法，，主要用于在数组中选择一个元素作为基准元素(pivot)，然后重新排序，使得所有小于pivot的元素都位于其左侧，所有大于pivot的元素都位于其右侧，最终返回pivot的位置，比Hoare分区方法效率低一些：

1. **选择pivot**：通常选择数组的最后一个元素作为pivot（如果选择其他元素作为pivot,需要将该元素与最后一个元素交换位置）
2. **初始化索引**：设置一个索引i,初始化为`start-1`,其中start是要分区子数组的起始位置
3. **遍历和交换**：
   - 遍历数组的`start`到`end-1`(end是数组的结束位置，不包括pivot)
   - 对于每个元素`a[j]`，如果该元素小于或等于pivot，就增加i的值，并将`a[i]`与`a[j]`交换
4. **放置pivot**：遍历完成后，将pivot与`i+1`位置的元素交换。这样pivot就被放在了其最终位置上，这个位置称为分区索引。
5. **返回分区索引**：返回`i+1`作为分区索引。

## 第K个数

第k个数，在数组中找到第k小的数，方法与快排类似

1. **选择基数pivot**：通常选择第一个元素low或最后一个元素high或最中间的一个元素(low+high)/2
2. **进行分区**：将数组分为大于piovt和小于piovt的区间
3. **判断K所在区间**：k在左侧区间，则将high修改为左侧区间的最大index值，k在右侧区间，则将low修改为右侧区间的起始index
4. **递归调用**：递归执行上述操作，直至k值刚好等于分区边界值

```go
// TheKthNumber 第K个数
func TheKthNumber(nums []int, left, right, k int) int {
    if left == right {
       return nums[left]
    }
    // 查找pivot分区边界的索引
    pivotIndex := TheKthNumberQuickStart(nums, left, right)
    if pivotIndex == k-1 {
       return nums[pivotIndex]
    } else if pivotIndex > k-1 {
       // 第k个数在左边分区，因此结束位置为pivotIndex+1,数组区间为左闭右开
       return TheKthNumber(nums, left, pivotIndex+1, k)
    } else {
       // 第k个数在右边分区，因此起始位置修改为pivotIndex+1即右边区间的起点位置
       return TheKthNumber(nums, pivotIndex+1, right, k)
    }
}

// TheKthNumberQuickStart第K个数快排
func TheKthNumberQuickStart(nums []int, left, right int) int {
    pivot := nums[(left+right)/2]
    for left <= right {
       for nums[left] < pivot {
          left++
       }

       for nums[right] > pivot {
          right--
       }
       nums[left], nums[right] = nums[right], nums[left]
       left++
       right--
    }
    return right
}
```

## 归并排序

归并排序使用了分治法的思想来进行排序，主要包括两个过程，分解和合并：

- **分解**：将原始数组分割成若干个子数组，直到每个子数组只包含一个元素或为空。任何一个单独的元素都可以视为已排序的数组，所以这个步骤珠岙是不断地将数组对半分割，直至分割到不能再分为止
- **合并**：将分解后地子数组重新组合成一个整体有序地数组，这一过程需要额外空间来暂时存储合并后地数组

以下是基本过程:

1. **递归分解**：从中间位置将数组分解成两个子数组，递归地对这两个数组进行归并排序
2. **合并排序**：将两个排序好地数组合并成一个有序的数组。合并时，通常需要一个辅助的数组来存储合并后的元素，确保在合并过程中数组元素的顺序正确

复杂度:

- 时间复杂度：

  - 最佳情况O(n log n)

  - 平均情况O(n log n)

  - 最差情况O(n log n)

- 空间复杂度：
  - 需要额外的空间存储临时数组，空间复杂度为O(n)

优点与缺点:

- 优点：
  - 稳定排序，相同元素的相对顺序不会改变
  - 对于大数据量效率很高
- 缺点：
  - 需要额外的存储空间
  - 对于小规模数组，其他排序算法如快读排序可能更优

```go
func MergeSortMerge(left, right []int) []int {
	var result []int
	leftIndex, rightIndex := 0, 0
	for leftIndex < len(left) && rightIndex < len(right) {

		if left[leftIndex] < right[rightIndex] {
			result = append(result, left[leftIndex])
			leftIndex++
		} else {
			result = append(result, right[rightIndex])
			rightIndex++
		}
	}

	result = append(result, left[leftIndex:]...)
	result = append(result, right[rightIndex:]...)

	return result
}

// MergeSort 合并排序
func MergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	midIndex := len(nums) / 2
	leftNums := MergeSort(nums[:midIndex])
	rightNums := MergeSort(nums[midIndex:])
	return MergeSortMerge(leftNums, rightNums)
}
```



## 二分 

二分查找算法，也称为二分搜索，是一种在有序数组中查找某一特定元素的搜索算法：

1. **确定边界**：定义搜索范围，通常是整个数组的最低索引(low)和最高索引(high)

2. **中点计算**：计算重点mid的索引，通常是`low+(high-low)/2`,避免某些语言中的整数溢出问题

3. **比较**：将中点元素与目标值比较:
   - 中点值等于目标元素，返回结果
   - 中点值大于目标元素，调整high为mid-1
   - 中点值小于目标元素，调整low为mid+1
4. **重复搜索**：根据上一步比较结果调整边界，重复执行2与3，直到low>high,表示搜索失败

```go
func BinarySearch(num []int, target int) int {
	low := 0
	high := len(num) - 1

	// low <high的时候 一直循环处理
	for low < high {
		// 取整
		mid := low + (high-low)/2
		if num[mid] == target {
			return mid
		} else if num[mid] > target {
			// 避免重复比较mid 这里mid索引的值已经大于target了
			high = mid - 1
		} else {
			// 同上，避免重复比较mid
			low = mid + 1
		}
	}
	return -1
}
```

## 双指针算法

### 最长不重复连续子序列

指在从给定字符穿中找到最长的没有重复的字串。通常通过滑动窗口术来解决，主要思路是使用两个指针（left 和right）来定义当前考虑的子串的边界。

1. **初始化**：定义两个指针left和right，初始都位于字符串的起始位置。使用哈希表来记录当前窗口内字符的出现情况
2. **扩展窗口**：移动right指针向右扩展窗口，每移动一次，九江新加入窗口的字符加入哈希表，并检查是否重复
3. **调整窗口**：如果发现right指针指向的字符在哈希表中已存在，意味着窗口内有重复字符。此时移动left向右缩小串钩，直到重复的字符被移除窗口。同时更新哈希表
4. **重复步骤2到4**：直到right指针到达字符串末尾。
5. **输出结果**：最长不重复子串的长度

```go
// LongestNonRepeatingSubsequence 最长子串
func LongestNonRepeatingSubsequence(str string) int {
	strMap := make(map[rune]int)
	left := 0
	maxLength := 0
	for right, s := range str {
		// 判断hash表中是否有该字符串。如果存在，更新最长字串的起始index为之前保存的索引位置+1
		if lastStr, ok := strMap[s]; ok && lastStr > left {
			left = lastStr + 1
		}

		// 更新hash表中字串的位置
		strMap[s] = right

		// 获取当前子串的长度并于maxLength比较 取最长的字符串
		if currentLength := right - left + 1; currentLength > maxLength {
			maxLength = currentLength
		}	
	}

	return maxLength
}
```

### 两数之和

给定一个数组和目标数值，找出数组中是否存在一组元素的和等于目标元素

```go
// ToSum 两数之和
func ToSum(arr []int, target int) (int, int) {
	intMap := make(map[int]int)
	for index, value := range arr {
		if existIndex, ok := intMap[target-value]; ok {
			return index, existIndex
		} else {
			intMap[value] = index
		}
	}
	return -1, -1
}
```

# 搜索与图论

## DFS（深度优先）

### 排列数字

给定一个整数n,将数字1~n排成一排，将会有很多种排列方法，现在，请你按照字典序将所有的排列方法输出。

- 输入格式：共一行，包含一个整数n
- 输出格式，按字典序输出所有排列方案，每个方案占一行
