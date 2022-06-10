package main

/**
 * question: GP38 团队闯关
 * description: 有一个团队闯关游戏，通过规则是团队中的某一个人的分数大于标准分数，这个团队就算通关。给出该团队所有人的分数，判断该团队是否能通关，能通关返回true,不能返回false。
 * input : [1,2,3,4,8],7
 * output: true
 */
func canPass(score []int, target int) bool {
	// write code here
	for _, v := range score {
		if v > target {
			return true
		}
	}
	return false
}
