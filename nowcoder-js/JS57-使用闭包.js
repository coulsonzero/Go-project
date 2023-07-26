/**
 * JS57 使用闭包
 * 实现函数 makeClosures，调用之后满足如下条件：
 * 1、返回一个函数数组 result，长度与 arr 相同
 * 2、运行 result 中第 i 个函数，即 result[i]()，结果与 fn(arr[i]) 相同
 * 
 * var arr = [1,2,3];
 * var fn = function (x) {
 *     return x * x;
 * }
 * var result = makeClosures(arr,fn);
 * (result[1]() === 4) === (fn(arr[1]) === 4) === true
 */


function makeClosures(arr, fn) {
    let result = []
    for (let i = 0; i < arr.length; i++) {
        result[i] = fn.bind(null, arr[i])
    }
    return result

}

function makeClosures2(arr, fn) {
    let result = []
    for (let i = 0; i < arr.length; i++) {
        result[i] = function() {
            return fn(arr[i])
        }
    }
    return result
}

function makeClosures3(arr, fn) {
    let result = [];
    for (let i = 0; i < arr.length; i++) {
        result.push(function(){
            return fn(arr[i]);
        });
    }
    return result;
}
