## 贪心

- [45. 跳跃游戏 II](45.go)
- [55. 跳跃游戏](55.go)
- [134. 加油站](134.go)
- [435. 无重叠区间](435.go)
- [452. 用最少数量的箭引爆气球](452.go)

```go
// 动态规划和贪心算法的关系。
// 贪心算法可以理解为一种特殊的动态规划问题。
// 拥有一些更特殊的性质，可以进一步降低动态规划算法的时间复杂度。

/**
贪心是一种在每一步选择中都采取在当前状态下最好或最优（即最有利）的选择，从而达到结果是最优的算法。
什么是贪心选择性质呢，简单说就是：每一步都做出一个局部最优的选择，最终的结果就是全局最优。
这是一种特殊性质，其实只有一部分问题拥有这个性质。
------------------------------------------------------
在所有的 feasible solution 中选择出 optimal solution
feasible solution（可行解）
optimal solution （最优解）
optimization problem （最优化问题）
dynamic programming, greedy algorithm is used for solving optimization problems

A greedy algorithm is an approach for solving a problem
by selecting the best option available at the moment.

Two properties:
1. Greedy Choice Property
If an optimal solution to the problem can be found
by choosing the best choice at each step without reconsidering the previous steps once chosen,
the problem can be solved using a greedy approach.This property is called greedy choice property.

2. Optimal Substructure
If the optimal overall solution to the problem corresponds to the optimal solution to its sub problems,
then the problem can be solved using a greedy approach. This property is called optimal substructure.
*/
```

