package parser

type (
	actionTable [numStates]actionRow
	actionRow   struct {
		canRecover bool
		actions    [numSymbols]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
		canRecover: false,
		actions: [numSymbols]action{
			nil,      /* INVALID */
			nil,      /* $ */
			shift(2), /* if */
			shift(3), /* id */
			nil,      /* then */
			nil,      /* else */

		},
	},
	actionRow{ // S1
		canRecover: false,
		actions: [numSymbols]action{
			nil,          /* INVALID */
			accept(true), /* $ */
			nil,          /* if */
			nil,          /* id */
			nil,          /* then */
			nil,          /* else */

		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: [numSymbols]action{
			nil,      /* INVALID */
			nil,      /* $ */
			nil,      /* if */
			shift(4), /* id */
			nil,      /* then */
			nil,      /* else */

		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(3), /* $, reduce: Stmt */
			nil,       /* if */
			nil,       /* id */
			nil,       /* then */
			nil,       /* else */

		},
	},
	actionRow{ // S4
		canRecover: false,
		actions: [numSymbols]action{
			nil,      /* INVALID */
			nil,      /* $ */
			nil,      /* if */
			nil,      /* id */
			shift(5), /* then */
			nil,      /* else */

		},
	},
	actionRow{ // S5
		canRecover: false,
		actions: [numSymbols]action{
			nil,      /* INVALID */
			nil,      /* $ */
			shift(7), /* if */
			shift(8), /* id */
			nil,      /* then */
			nil,      /* else */

		},
	},
	actionRow{ // S6
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(1), /* $, reduce: Stmt */
			nil,       /* if */
			nil,       /* id */
			nil,       /* then */
			shift(9),  /* else */

		},
	},
	actionRow{ // S7
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* if */
			shift(10), /* id */
			nil,       /* then */
			nil,       /* else */

		},
	},
	actionRow{ // S8
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(3), /* $, reduce: Stmt */
			nil,       /* if */
			nil,       /* id */
			nil,       /* then */
			reduce(3), /* else, reduce: Stmt */

		},
	},
	actionRow{ // S9
		canRecover: false,
		actions: [numSymbols]action{
			nil,      /* INVALID */
			nil,      /* $ */
			shift(2), /* if */
			shift(3), /* id */
			nil,      /* then */
			nil,      /* else */

		},
	},
	actionRow{ // S10
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* if */
			nil,       /* id */
			shift(12), /* then */
			nil,       /* else */

		},
	},
	actionRow{ // S11
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(2), /* $, reduce: Stmt */
			nil,       /* if */
			nil,       /* id */
			nil,       /* then */
			nil,       /* else */

		},
	},
	actionRow{ // S12
		canRecover: false,
		actions: [numSymbols]action{
			nil,      /* INVALID */
			nil,      /* $ */
			shift(7), /* if */
			shift(8), /* id */
			nil,      /* then */
			nil,      /* else */

		},
	},
	actionRow{ // S13
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(1), /* $, reduce: Stmt */
			nil,       /* if */
			nil,       /* id */
			nil,       /* then */
			shift(14), /* else */

		},
	},
	actionRow{ // S14
		canRecover: false,
		actions: [numSymbols]action{
			nil,      /* INVALID */
			nil,      /* $ */
			shift(7), /* if */
			shift(8), /* id */
			nil,      /* then */
			nil,      /* else */

		},
	},
	actionRow{ // S15
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(2), /* $, reduce: Stmt */
			nil,       /* if */
			nil,       /* id */
			nil,       /* then */
			reduce(2), /* else, reduce: Stmt */

		},
	},
}
