/**
 * 统计数组 arr 中值等于 item 的元素出现的次数
 * input: [1, 2, 4, 4, 3, 4, 3], 4
 * output: 3
 */

function count(arr, item) {
    return arr.filter(elem => elem === item).length
}

function count2(arr, item) {
    let count = 0
    for(let i = 0; i < arr.length; i++) {
        if(arr[i] === item) count++
    }
    return count
}