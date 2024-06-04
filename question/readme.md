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

## 找出缺失的观测数据

现有一份 `n + m` 次投掷单个 **六面** 骰子的观测数据，骰子的每个面从 `1` 到 `6` 编号。观测数据中缺失了 `n` 份，你手上只拿到剩余 `m` 次投掷的数据。幸好你有之前计算过的这 `n + m` 次投掷数据的 **平均值** 。

给你一个长度为 `m` 的整数数组 `rolls` ，其中 `rolls[i]` 是第 `i` 次观测的值。同时给你两个整数 `mean` 和 `n` 。

返回一个长度为 `n` 的数组，包含所有缺失的观测数据，且满足这 `n + m` 次投掷的 **平均值** 是 `mean` 。如果存在多组符合要求的答案，只需要返回其中任意一组即可。如果不存在答案，返回一个空数组。

`k` 个数字的 **平均值** 为这些数字求和后再除以 `k` 。

注意 `mean` 是一个整数，所以 `n + m` 次投掷的总和需要被 `n + m` 整除。

- 计算与总值的差值
- 插值大于6*n或者小于n直接无法得到
- 利用插值对当前剩余数组个数取整，并进行赋值

```go
func missingRolls(rolls []int, mean int, n int) []int {
    var hasNums int
    m:=len(rolls)
    for i:=0;i<m;i++{
        hasNums+=rolls[i]
    }

    missNums:=mean*(n+m)-hasNums
    if missNums>6*n || missNums<n{
        return  []int{}
    }

    result:=make([]int,n)
    for i:=0;i<n;i++{
        // 可优化点:如果curNum是整数，则后续的值直接赋值 不需要再进行比较
        curNum:=missNums/(n-i)
        result[i]=curNum
        missNums-=curNum
    }
    return result 
}
```

# 栈与队列

## 最小栈

设计一个支持 `push` ，`pop` ，`top` 操作，并能在常数时间内检索到最小元素的栈。

实现 `MinStack` 类:

- `MinStack()` 初始化堆栈对象。
- `void push(int val)` 将元素val推入堆栈。
- `void pop()` 删除堆栈顶部的元素。
- `int top()` 获取堆栈顶部的元素。
- `int getMin()` 获取堆栈中的最小元素。

题解：

- 实现两个stack，一个用于存放数据，一个用户存放当前数据时的最小值
- push以及pop时同时新增或删除栈顶数据

```go
type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minStack: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	nowLen := len(this.minStack)
	if nowLen == 0 {
		this.minStack = append(this.minStack, val)
	} else {
		this.minStack = append(this.minStack, min(this.minStack[nowLen-1], val))
	}
}

func (this *MinStack) Pop() {

	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {

	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
```

