<?php
class Solution {

	/**
	 * @param Integer[] $nums
	 * @param Integer $target
	 * @return Integer[]
	 */
	function twoSum($nums, $target) {
		$tmp = [];
		foreach($nums as $k => $v){
			$tmp[$v] = $k;
		}
		foreach($nums as $key => $value){
			$diff = $target - $value;
			if(isset($tmp[$diff]) && $tmp[$diff] != $key){
				return [$key,$tmp[$diff]];
			}
		}
		return [];
	}
}