<?php
/**
 * Created by PhpStorm.
 * User: xiaomu
 * Date: 2021/2/19
 * Time: 22:33
 */
class Solution {

	/**
	 * @param Integer $n
	 * @return String[]
	 */
	public $arr = [];
	function generateParenthesis($n) {
		$this->gen(0,0,$n,'');
		return $this->arr;
	}

	function gen($left,$right,$n,$s){
		//此处第一次写错了,判断条件写成了$left == $right,必须保证左边的括号与右边的括号都是n
		if($left == $n && $right == $n){
			$this->arr[] = $s;
			return;
		}
		if($left < $n){
			$this->gen($left+1,$right,$n,$s.'(');
		}
		if($left > $right){
			$this->gen($left,$right+1,$n,$s.')');
		}
	}
}