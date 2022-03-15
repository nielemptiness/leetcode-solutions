var singleNumber = function(nums) {
    const set = new Set();
    
    let element;
    for (let index = 0; index < nums.length; index ++) {
        
        element = nums[index];

        if (set.has(element)) {
            console.log('removed')
            set.delete(element);
            continue;
        }

        console.log('added')
        set.add(element);
        
    }

    for (const iterator of set) {
        return iterator;
    }
};

console.log(singleNumber([1,1,2,2,3]));