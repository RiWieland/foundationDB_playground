playing around with foundationDB

# Description


# Design
The functions to manipulate our Key-Value store are build in different layers.

The "lowest" layer describe functions with simple interaction to the key-value store. These are given by query, update, drop, etc. These functions build the basic interactions with the key-value store and used for the different subs.

On top of these there are the functions that build more complex interactions and need to call different functions (also for different subs). 
Here the transaction properties of foundationDB comes into play as we are normally want a all-or-nothing write (atomicity).



# Data Model:

The datamodel follows access patterns. 

# Entities:

Image Files:
The path of the image files represents the key of the file.

Rectangles:
Rectangles are modeled with index for identification and the positions of the 4 corners of the rectangle.
The index is used as the same rectangle position could be used for different images. 



