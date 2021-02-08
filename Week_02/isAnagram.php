<?php
class Solution {

	/**
	 * @param String $s
	 * @param String $t
	 * @return Boolean
	 */
	function isAnagram($s, $t) {
		if(strlen($s) != strlen($t)){
			return false;
		}
		if($s == '' && $t == ''){
			return true;
		}
		$tmp = [];
		for($i = 0;$i < strlen($s);$i++){
			if(isset($tmp[$s[$i]])){
				$tmp[$s[$i]] += 1;
			}else{
				$tmp[$s[$i]] = 1;
			}
		}
		for($j = 0;$j < strlen($t);$j++){
			$value = $t[$j];
			if(isset($tmp[$value])){
				$tmp[$value] -= 1;
			}
		}
		foreach($tmp as $value){
			if($value !== 0){
				return false;
			}
		}
		return true;
	}
}