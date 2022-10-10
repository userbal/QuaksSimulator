package main

func setupPot() playerBoard {
	pot := []potLocation{
		{
			money:         1,
			victoryPoints: 0,
			ruby:          false,
		},
		{
			money:         2,
			victoryPoints: 0,
			ruby:          false,
		},
		{
			money:         3,
			victoryPoints: 0,
			ruby:          false,
		}, {
			money:         4,
			victoryPoints: 0,
			ruby:          false,
		},
		{
			money:         5,
			victoryPoints: 0,
			ruby:          true,
		},
		{
			money:         6,
			victoryPoints: 1,
			ruby:          false,
		},
		{
			money:         7,
			victoryPoints: 1,
			ruby:          false,
		},
		{
			money:         8,
			victoryPoints: 1,
			ruby:          false,
		},
		{
			money:         9,
			victoryPoints: 1,
			ruby:          true,
		},
		{
			money:         10,
			victoryPoints: 2,
			ruby:          false,
		},
		{
			money:         11,
			victoryPoints: 2,
			ruby:          false,
		},
		{
			money:         12,
			victoryPoints: 2,
			ruby:          false,
		},
		{
			money:         13,
			victoryPoints: 2,
			ruby:          true,
		},
		{
			money:         14,
			victoryPoints: 3,
			ruby:          false,
		},
		{
			money:         15,
			victoryPoints: 3,
			ruby:          false,
		},
		{
			money:         15,
			victoryPoints: 3,
			ruby:          true,
		},
		{
			money:         16,
			victoryPoints: 3,
			ruby:          false,
		},
		{
			money:         16,
			victoryPoints: 4,
			ruby:          false,
		},
		{
			money:         17,
			victoryPoints: 4,
			ruby:          false,
		},
		{
			money:         17,
			victoryPoints: 4,
			ruby:          true,
		},
		{
			money:         18,
			victoryPoints: 4,
			ruby:          false,
		},
		{
			money:         18,
			victoryPoints: 5,
			ruby:          false,
		},
		{
			money:         19,
			victoryPoints: 5,
			ruby:          false,
		},
		{
			money:         19,
			victoryPoints: 5,
			ruby:          true,
		},
		{
			money:         20,
			victoryPoints: 5,
			ruby:          false,
		},
		{
			money:         20,
			victoryPoints: 6,
			ruby:          false,
		},
		{
			money:         21,
			victoryPoints: 6,
			ruby:          false,
		},
		{
			money:         21,
			victoryPoints: 6,
			ruby:          true,
		},
		{
			money:         22,
			victoryPoints: 7,
			ruby:          false,
		},
		{
			money:         22,
			victoryPoints: 7,
			ruby:          true,
		},
		{
			money:         23,
			victoryPoints: 7,
			ruby:          false,
		},
		{
			money:         23,
			victoryPoints: 8,
			ruby:          false,
		},
		{
			money:         24,
			victoryPoints: 8,
			ruby:          false,
		},
		{
			money:         24,
			victoryPoints: 8,
			ruby:          true,
		},
		{
			money:         25,
			victoryPoints: 9,
			ruby:          false,
		},
		{
			money:         25,
			victoryPoints: 9,
			ruby:          true,
		},
		{
			money:         26,
			victoryPoints: 9,
			ruby:          false,
		},
		{
			money:         26,
			victoryPoints: 10,
			ruby:          false,
		},
		{
			money:         27,
			victoryPoints: 10,
			ruby:          false,
		},
		{
			money:         27,
			victoryPoints: 10,
			ruby:          true,
		},
		{
			money:         28,
			victoryPoints: 11,
			ruby:          false,
		},
		{
			money:         28,
			victoryPoints: 11,
			ruby:          true,
		},
		{
			money:         29,
			victoryPoints: 11,
			ruby:          false,
		},
		{
			money:         29,
			victoryPoints: 12,
			ruby:          false,
		},
		{
			money:         30,
			victoryPoints: 12,
			ruby:          false,
		},
		{
			money:         30,
			victoryPoints: 12,
			ruby:          true,
		},
		{
			money:         31,
			victoryPoints: 12,
			ruby:          false,
		},
		{
			money:         31,
			victoryPoints: 13,
			ruby:          false,
		},
		{
			money:         32,
			victoryPoints: 13,
			ruby:          false,
		},
		{
			money:         32,
			victoryPoints: 13,
			ruby:          true,
		},
		{
			money:         33,
			victoryPoints: 14,
			ruby:          false,
		},
		{
			money:         33,
			victoryPoints: 14,
			ruby:          true,
		},
	}

	pb := playerBoard{
		pot:         pot,
		ingredients: []chip{},
	}

	return pb
}

func setupBag() []chip {
	m := getChipSet1Functions()

	bag := []chip{
		// {
		// 	color:    "orange",
		// 	value:    1,
		// 	function: m["orange"],
		// },
		// {
		// 	color:    "green",
		// 	value:    1,
		// 	function: m["green"],
		// },
		// {
		// 	color:    "white",
		// 	value:    1,
		// 	function: m["white"],
		// },
		// {
		// 	color:    "white",
		// 	value:    1,
		// 	function: m["white"],
		// },
		// {
		// 	color:    "white",
		// 	value:    1,
		// 	function: m["white"],
		// },
		// {
		// 	color:    "white",
		// 	value:    1,
		// 	function: m["white"],
		// },
		// {
		// 	color:    "white",
		// 	value:    2,
		// 	function: m["white"],
		// },
		// {
		// 	color:    "white",
		// 	value:    2,
		// 	function: m["white"],
		// },
		// {
		// 	color:    "white",
		// 	value:    3,
		// 	function: m["white"],
		// },
		// testing blue
		{
			color:    "blue",
			value:    2,
			function: m["blue"],
		},
	}

	return bag
}

func setupBoard() []boardLocation {
	scoreBoard := []boardLocation{}
	for i := 1; i <= 50; i++ {
		scoreBoard = append(scoreBoard, boardLocation{
			value:   i,
			ratTail: false,
		})
	}
	return scoreBoard
}
