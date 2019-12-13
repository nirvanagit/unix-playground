# Autconf

Tutorial - [edwardrosten.com](http://www.edwardrosten.com/code/autoconf/)

## Introduction
Autoconf is a system for generating a script which will automatically determine the presence of things which your project either depends upon, or can make use of. 
Once these have been determined, autoconf can export this information in various formats suitable for your project to use. 
Autoconf scripts themselves are written in a combination of shell and M4 (a macro processing language considerably more capable than the C preprocessor). 
Autoconf essentially runs the preprocessor on your script to produce a portable shell script which will perform all the requisite tests, 
produce handy log files, preprocess template files, for example to generate Makefile from Makefile.in and and take a standard set of command line arguments.

