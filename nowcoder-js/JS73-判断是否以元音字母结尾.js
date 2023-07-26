/*
function endsWithVowel(str) {
    var arr = ['a','e','i','o','u','A','E','I','O','U'];
    return arr.includes(str[str.length-1])
}

function endsWithVowel(str) {
   return ("aeiouAEIOU".indexOf(str[str.length-1]) !== -1)
}
 */

function endsWithVowel(str) {
    var reg = /(a|o|e|i|u)$/gi;
    return reg.test(str);
}