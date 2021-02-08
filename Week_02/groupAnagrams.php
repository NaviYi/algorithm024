<?php
class Solution {

	/**
	 * @param String[] $strs
	 * @return String[][]
	 */
	function groupAnagrams($strs) {
		if(empty($strs)){
			return [];
		}
		$j = 0;
		$tmp = [];
		for($i=0;$i<count($strs);$i++){
			$s_arr = str_split($strs[$i]);
			sort($s_arr);
			$t = implode('',$s_arr);
			$tmp[$t][] = $strs[$i];
		}
		return array_values($tmp);
	}
}
