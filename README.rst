Gifexplode
==========

Gifexploder takes a GIF image file and creates new images from each of the frames.  This is a
fun little toy app to help me get more acquainted with the Golang standard image library.

This will take the input file, open it up, decode the image, and for each frame in the GIF
will create a new PNG file of the format "<input filename>.<image index>.png" in the current
working directory.

Usage
------

.. code::

    go build gifexplode.go
    ./gifexplode.go sample.gif
