# Design
*NULL is a platformer where you have to reach the top of the world as fast as possible.

Sounds simple? Well, memory management isn't always as it seems! The world is tile based, and some tiles are dereferencers. You can use these to dereference a pointer to create a new tile.

But be careful! You might hit a segfault an end your game!

There will also be nulls floating around, which if are in a tile you dereference, will cause a null pointer dereference, and crash the program.
