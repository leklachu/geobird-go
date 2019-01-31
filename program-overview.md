----------------------
-- G .. O .. Bird ! --
----------------------

Overview of the GeoBird Program
-------------------------------

Aim: To download sets of satellite images from earthdata.nasa.gov

The program:
1. Takes inputs (basic: command line options)
2. works out from those what images are wanted; in turn what URLs
3. queries those URLs and downloads the relevant files
4. (or 2b) works out from the options where to put the files
5. puts the files in the right place
6. (3b) writes output as it's working (and when finished?)

A few notes:
- Might want to check if the relevant images already exist (e.g. for updating a
  folder of all images till 'today')


