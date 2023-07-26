/**
 * 创建正则表达式，一代身份证15位，二代身份证18位
 * 二代身份证18位数字中，最后一位有可能是”x“|”X“
 */

const _isCard = number => {
    // 补全代码
    let reg = /^[\d]{17}[X\d]$/
    return  reg.test(number)
}

const _isCard2 = number => {
    let reg = /^[\d]{17}[\d|X|x]$/
    return  reg.test(number)
}

const _isCard3 = number => {
    // 补全代码
    let reg = /^[\d]{17,18}[X]?$/
    return  reg.test(number)
}