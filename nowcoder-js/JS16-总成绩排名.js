/**
 * JS16 总成绩排名
 * 要求将数组参数中的对象以总成绩(包括属性"chinese"、"math"、"english")从高到低进行排序并返回
 */

const _rank = array => {
    // 补全代码
    array.sort((item1,item2) =>{
        return (item2.math + item2.chinese + item2.english) - (item1.math + item1.chinese + item1.english)
    })
    return array
}