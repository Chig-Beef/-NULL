# *NULL
This game was released on itch and now here is the source code. I've removed as
much dead code as I could, cleaned it all up, and hopefully it all still works.
If it doesn't, fix it! This git repo also includes the level editor, which I
have not looked over, so it may be full of even worse coding practices than
usual. The level editor, however, should be more that usable to create your own
levels, and a starting point for a more enhanced editor.

## Go Original Version
I originally wrote the game in Golang, and it's over 2k lines. It's not that
complicated, but here's a walkthrough of the code.

First, try to keep away from the audio, image, and font code. It's not nice, so
only work with it if you are sure. You may want to add images and sounds, which
isn't too hard. First, add the image you want into the correct folder (e.g. 
`assets/images`. Then, in `images.json` add an entry for this image. This will
load the image into the game, and you can use it anywhere with the name you
gave it. This means you can also switch out images for tiles by simply editing
`images.json`. Sounds is a bit more complicated. You'll have to jump to around
line 200 of `Audio.go`, and add a line similar to those you find. It works in
a similar fashion, but sounds are hard coded (which is why I don't use this
code anymore!).

The game is split into phases, as most games are, and this is all handled in
`Game.go`, which splits drawing and updating. Most game logic will be in
`Game.go`, `Level.go`, and `Player.go`.

Last thing I want to talk about is `UI.go`, which predefines all UI elements
using `stagui`. This is an old version of `stagui` and `stagerror`, and they're
more private, but I left them in for completion-sake. It might be good to have
a look through `stagui` to understand how the UI works if you want to change
something.

## C Port
After writing the code, I started learning to use C and Raylib. I thought it
would be very useful as an exercise to port an existing game, so I did. The
code is mostly complete, however, it's missing sounds. I wasn't very
comfortable with C when I started porting (and I'm still not fully). There's
going to be *heaps* of bad practices and errors, but it works (mostly), which
is the main point.

## Missing Features
In the original game (that I think is on itch), there are little guys going
around in circles and numbers on the tiles. These features were useless, so I
removed them. They were interesting ideas, but they weren't fun, nor did they
contribute to a good speedrun.

## To-Do
Looking for some project ideas? You could:

Add sound effects to CNULL. You could also add in the soundtrack.

Add back in nulls and dereferencing.

Add either same device, local, or remote multiplayer.

Port the game to another language.

Make a new level that is more interesting.

Allow for more horizontal movement, allowing for larger levels and more sophisticated paths.

Add ghost replay, so that players can race themselves.

Add horizontal boosts, which could be used to create flow tubes.

Add obstacles that could kill you.

Add grip tiles, that could be used for wall climbing.

Add gravity switching mechanics.

Add complex levels, such as using keys to open doors, or teleporters.
