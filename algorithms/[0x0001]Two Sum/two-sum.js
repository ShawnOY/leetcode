const twoSum = function(nums, target) {
    const len = nums.length;

    for(let x = 0; x < len; x++) {
        for(let y =0; y < len; y++) {
            if (x === y) continue;
            if (nums[x] + nums[y] === target) return [x, y];
        }
    }
};
