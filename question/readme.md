## 跳跃游戏

给你一个非负整数数组 `nums` ，你最初位于数组的 **第一个下标** 。数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个下标，如果可以，返回 `true` ；否则，返回 `false` 。

题解：

- 判断当前元素能到达的最远距离
- 如果最远距离大于目标，则返回true
- 如果当前索引位置大于之前元素最远距离，返回false

```go
func canJump(nums []int) bool {
	target := len(nums) - 1
	var maxIndex int

	for i := 0; i < target+1; i++ {
		if i > maxIndex {
			return false
		}
		maxIndex = max(i+nums[i], maxIndex)
		if maxIndex >= target {
			return true
		}
	}
	return true
}
```

