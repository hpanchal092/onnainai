# onnainai

## description

do you suck at 2048? fear no more, because onnainai plays 2048 and wins for
you!

## how it works

(update this later to be more accurate)

 1. a python script uses opencv to see the current state of the 2048 board and
    uses ocr to figure out each tile and its value

 2. this data is sent to an algorithm written in go to calculate the most 
    optimal move to make

 3. the move is sent back to the python script to play the move onto the board

 4. repeat

## faq

there are no frequently asked questions. this is a random github repo that no
one knows about. although if by some chance you are looking at this repo you
might wonder: 

**why choose to write the algorithm in a differenent language to the script?
why not just have them be the same language to make things easier?**

well first of all i can do whatever i want. to answer the question i wanted to
write the program in a language that would be fast. i could have chosen a
language like c or c++, but who really wants to write code in c or c++ in 2022?
i could have chosen a language like python or javascript, but those are slow.
that left me with rust or go. rust is still mildly challenging to write code in
so i figured i would pick go because it is a nice balance of being fast and
easy.

**why choose the name onnainai? how does that have anything to do with 2048?**

the name onnainai definitely has no meaning whatsoever and was definitely not 
chosen by my friend who wanted to make fun of me.

## how to run

idk fill this in once i figure that part out
