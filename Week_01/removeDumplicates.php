<?php
/**
 * Created by PhpStorm.
 * User: xiaomu
 * Date: 2021/1/31
 * Time: 20:07
 */
class Solution {

	/**
	 * @param Integer[] $nums
	 * @return Integer
	 */
	function removeDuplicates(&$nums) {
		if(count($nums) <= 1){
			return count($nums);
		}
		$slow = 0;
		for($fast=1;$fast < count($nums);$fast++){
			if($nums[$slow] != $nums[$fast]){
				$nums[$slow+1] = $nums[$fast];
				$slow++;
			}
		}
		return $slow + 1;
	}
}