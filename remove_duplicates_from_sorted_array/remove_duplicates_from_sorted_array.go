// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
package remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Menggunakan dua pointer untuk melacak elemen unik
	uniqueIndex := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[uniqueIndex] {
			uniqueIndex++
			nums[uniqueIndex] = nums[i]
		}
	}

	// Panjang slice setelah menghapus elemen duplikat
	return uniqueIndex + 1
}
