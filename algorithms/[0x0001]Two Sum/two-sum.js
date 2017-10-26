/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
const twoSum = function (nums, target) {
  const map = new Map()

  for (let index in nums) {
    map.set(nums[index], index)
  }

  for (let index in nums) {
    let complement = target - nums[index]

    if (map.has(complement) && map.get(complement) !== index) {
      return [parseInt(index), parseInt(map.get(complement))]
    }
  }
}
