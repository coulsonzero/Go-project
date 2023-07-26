/**
 * 请补全JavaScript函数，要求返回数字参数的阶乘。
 * 注意：参数为大于等于0的整数。
 */

function _factorial(number) {
    // 补全代码
    return number === 1 ? 1 : number * _factorial(number - 1)
}