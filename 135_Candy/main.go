package main

import "fmt"

/*
There are n children standing in a line. Each child is assigned a rating value given in the integer array ratings.

You are giving candies to these children subjected to the following requirements:

Each child must have at least one candy.
Children with a higher rating get more candies than their neighbors.
Return the minimum number of candies you need to have to distribute the candies to the children.


Example 1:

Input: ratings = [1,0,2]
Output: 5
Explanation: You can allocate to the first, second and third child with 2, 1, 2 candies respectively.


Example 2:

Input: ratings = [1,2,2]
Output: 4
Explanation: You can allocate to the first, second and third child with 1, 2, 1 candies respectively.
The third child gets 1 candy because it satisfies the above two conditions.


2, 1, 2, 4, 1
*/

func candy(ratings []int) int {
	// Get left neig
	// Get right neig

	// Check minimun in left neig
	// Select current and right neig value

	candies := make([]int, len(ratings))
	candies_sum := 0
	for i := 0; i < len(ratings); i++ {
		var l_rating, r_rating, c_rating int
		var l_candies
		if i != 0 {
			l_rating = ratings[i-1]
			l_candies = candies[i-1]
		} else {
			l_rating = ratings[i]
			l_candies = 1
		}
		c_rating = ratings[i]
		if i != len(ratings)-1 {
			r_rating = ratings[i+1]
		} else {
			r_rating = ratings[i]
		}

		// Choose candies
		if l_rating == c_rating {
			candies[i] = l_candies
		}
	}

	fmt.Println(candies)
	return 0
}

func main() {
	ratings := []int{1, 0, 2}
	fmt.Println(candy(ratings))
}
