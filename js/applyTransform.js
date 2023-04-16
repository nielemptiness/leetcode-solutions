var map = function(arr, fn) {
    for (const i in arr) {
        arr[i] = fn(arr[i], Number(i));
    }
    return arr;
};
const fn = function plusone(n) { return n + 1; }; 
const arr = [1,2,3];
console.log(map(arr, fn))
const arr1 = [1,2,3]; const fn1 = function plusI(n, i) { return n + i; }
console.log(map(arr1, fn1));