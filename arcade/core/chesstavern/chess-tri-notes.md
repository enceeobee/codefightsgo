# Chess Triangle

Let's use numbers 1-8 to describe the 2x3 permutations in the problem's description.

<!-- ## Notes

Each triangle works flipped inverted both horizontally and vertically.

In 2x3, knight will never be in the center column

`n` = rows; `m` = columns

Ultimately, there's only 'knight takes rook', and 'knight takes bishop'. After that, there are only inversions and extensions.

So maybe we begin identifying how many 'k->r' and 'k->b' there are (that's our baseline). Then find the extensions on that. Then check if we can invert them vertically and/or horizontally. Then check if we can slide them around the board

Thus, for 2x3, we:

* Identify that there are two (2) essential triangles:
    * k->r
    * k->b
* Find that we can extend neither (0) of them
* Find that we can invert each (2) horizontally
* Find that we can invert each (2) vertically
* Total = 2 essential + 2 h invert + 2 v invert = 6

----

Let's try 2x4

* Start with our two (2) essential triangles
* Check sliding: (1)
    * Can slide right one col
* Check k->r extensions: (1)
    *  Bishop can extend
* Check k->b extensions: (0)
    * Nope
* Check inverts:
    * Basic config, both x & y (2)
    * Slide, both x & y (2)
    * k->r extension, x (1)

Total = 2+1+1+2+2+1 = 9 (?)

No idea if that's right..

---

Let's try 3x3 (which should equal 48 aka 12 configs)

`k->r`:


|  k | .  |  . |
|---|---|---|
| .  | b  | r |
| . | . | . |

**Summary**

**Note: these counts are off**

* Basic config (1)
* Slide (1)
	* +1 for each additional column past 3: (0)
	* +1 for each row past 2: (1)
* Extend (1)
	* +1 for each bishop [-1,-1] square:
		* bishop going under rook (1)
* Invert (6)
	* +1 for each previous config per axis:
		* basic config (2)
		* slide config (2)
		* extend config(2)

**Total for `k->r`:** 9

`k->b`:

| .  | .  | b  |
|---|---|---|
| k | r | . |
|  . | .  |  . |

*fig 2:*
| .  | b  | . |
|---|---|---|
| r | . | . |
| k | . |  . |

* Basic config (1)
* Slide (1)
	* +1 for each additional column past 3 (0)
	* +1 for each row past 2 (1)
		* This config can be extended to put the rook over the knight (see fig 2)
		* I'm sure we missed this one
		* So it's kinda recursive. Every new config we make, we have to check all the steps
* Extend (1)
	* Rook goes under Knight (1)
* Invert (6)
	* basic config (2)
	* slide config (2)
	* extend config (2)

**Total for `k->b`:** 9

Total total = 9 + 9 = 18

Um, fuck.

Wait. Each config is equivalent to 4 because:

* Initial config
* Invert over x axis
* Invert *that* over the y axis
* Finally, invert over x axis
	* i.e. invert x -> y -> x
	* If we inverted this last one over the y, we'd return to initial config

So that would mean for 3x3:

`k->r`:

3 configs (init, 1 slide, 1 extend) =

4 + 4 + 4 = 12

`k->b`

3 configs (init, 1 slide, 1 extend) =

4 + 4 + 4 = 12

= 24 ... still short

Ok, we need to account for when we create a new config via sliding or extending. Can we slide or extend from that config? It's recursive. We must do this, however, without double counting.

Once we have a total count of configs, we can return `configsCount*4`

So! Let's try 3x3 again

So maybe it's:

* Determine all the slides
* Determine all the shifts
* x by 2
* x by 4

---

With knight in top right:

In 2x3, there is only one way to have the knight take the bishop or rook. Everything else is derivative.

So in 2x3 the knight has a single move from the top right, resulting in one config per second piece (i.e. k->b and k->r)

In 3x3, the knight has two moves from the top right, resulting in two configs per second piece. It has one move from center right, resulting in one config per second piece.

configs = movesFromTopRight + numRowsMoreThan2 (for slides) + numColsMoreThan3 (for slides) + shifts/extensions

return configs * 4


### 3x3

Knight top right:

* Rook is down 2, over 1; Bishop over rook
	* Can slide right 1
* Rook is down 2, over 1; Bishop right of rook
* Rook is down 1, over 2: Bishop left of rook
	* Can slide down 1
* Rook is down 1, over 2: Bishop under rook

= 6 configs

6 configs * 2 enemies (rook and bishop) * 4 configs each = 48 👍

4 2x3s = 32; 32/48 = 2/3
6/9 = 2/3

is it really as easy as that?

2x3 total squares = 6
3x3 total squares = 9

ratio = 2/3

divide 32 by (2/3) = 48

1. Ok, first, determine how many 2x3s can fit in our board.
	* in a 3x3, we can fit 4 2x3s. n = 4
	* TODO - figure out how to calculate this
3. Multiply that number by 8
	* 4 * 8 = 32
4. Get ratio of total squares
	* 2x3 = 6; 3x3 = 9
	* 6/9 = 2/3
5. Divide 32 by ratio
	32 / (2/3) = 48


### 3x4

Knight top right:

* Rook is down 2, over 1
	* bishop over rook (total: 1)
		* Can slide right 2 (total: 3)
	* bishop right of rook (total: 4)
		* Can slide right 1 (total: 5)
* Rook is down 1, over 2
	* bishop left of rook (total: 6)
		* Can slide right 1 (total: 7)
		* Can slide down 1 (total: 8)
		* Can slide right AND down 1 (total: 9)
	* bishop under rook (total: 10)
		* Can slide right 1 (total: 11)

= 11 configs

11 * 2 * 4 = 88 💩 (should be 112)

7 2x3 = 56; 56 x 2 = 112

**Trying the new method**

1. How many 2x3s fit in our board?
	* 7
3. 7 * 8 = 56
4. Ratio = 1/2
	* 2x3 squares = 6
	* 3x4 squares = 12
	* 6/12 = 1/2
4. Divide 56 by ratio = 56/.5 = 112

💥💥💥

### 5x2

Should be 40

1. How many 2x3s fit in our board? 3
2. 3*8 = 24
3. Ratio =
	* 2x3 squares = 6
	* 5x2 squares = 10
	* 6/10 = 3/5
4. Divide 24 by ratio = 24*5/3 = 40 oh shit


**FUUUUUCK, THIS DIDN'T WORK.**

Let's try again -->

---

## Trying again!

I still like the idea of finding all the 2x3s, and multiplying that by 8. Now let's find all the additional configs outside of 2x3.

### 2x4

Number of 2x3s = 2 = 16 triangles from that

Number of additional triangles (by config):

1. 1 (bishop moves right)
2. 1 (bishop moves right)
3. 0
4. 0
5. 1 (rook moves right)
6. 0
7. 0
8. 1 (rook moves right)

Total: 4

**Grand total: 20**

### 3x3

Number of 2x3s = 4 = 32 tris from that

Number of additional tris:

1. 1 (bishop moves down and left)
2. 1 (bishop moves up and left)

OK GODDAMMIT I HATE THIS APPROACH I'M JUST GONNA BRUTE FORCE IT
