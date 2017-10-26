/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
const twoSum = function (nums, target) {
  const map = new Map()
  let complement

  for (let index in nums) {
    map.set(nums[index], index)
  }

  for (let index in nums) {
    complement = target - nums[index]

    if (map.has(complement) && map.get(complement) !== index) {
      return [index, map.get(complement)].map(Number)
    }
  }
}
