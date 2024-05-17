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

## 爬楼梯

假设你正在爬楼梯。需要 `n` 阶你才能到达楼顶。

每次你可以爬 `1` 或 `2` 个台阶。你有多少种不同的方法可以爬到楼顶呢？

**函数方式**：

- 到达第n阶的方法数=到达第n-1阶的方法数+到达第n-2阶的方法数

```go
func climbStairs(n int) int {
    if n<2{
        return 1
    }

    dp:=make([]int,n+1)
    dp[0]=1
    dp[1]=1
    for i:=2;i<n+1;i++{
        dp[i]=dp[i-1]+dp[i-2]
    }
    return dp[n]
}
```

**递归方式**：

- 出口：n<2
- 状态转移：`dp[i]=dp[i-1]+dp[i-2]`

```go
func climbStairs(n int) int {
    if (n<2){
        return 1 
    }
    return climbStairs(n-1)+climbStairs(n-2)
}
```

## 打家劫舍

你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，**如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警**。

给定一个代表每个房屋存放金额的非负整数数组，计算你 **不触动警报装置的情况下** ，一夜之内能够偷窃到的最高金额。

- 判断当前index与`index-2`能偷到的最大的值
- 判断`index-1`偷盗的最大的钱
- 选取两个节点最大的钱数

```go
func rob(nums []int) int {
    total:=len(nums)
    if total==1{
        return nums[0]
    }
    before:=nums[0]
    beforeTwo:=0
    for i:=1;i<total;i++{
        tem:=max(before,beforeTwo+nums[i])
        beforeTwo=before
        before=tem
    }
    return before
}
```

