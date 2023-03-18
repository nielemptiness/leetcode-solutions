var maxArea = function(height) {
    let start = 0;
    let end = height.length - 1;
    let current, result = 0;
    const getArea = (arr, a, b) => {
        const height = Math.min(arr[a], arr[b]);
        const length = b-a;
        console.log('Height + length: ', height, length);
        return height * length;
    }
    while (start < end) {
        current = getArea(height, start, end);
        result = Math.max(current, result);
        console.log('area: ', result);
        if (height[start] <= height[end]) {
            start++
        } else end--;
    }
    return result;
};

console.log(maxArea([1,8,6,2,5,4,8,3,7]));