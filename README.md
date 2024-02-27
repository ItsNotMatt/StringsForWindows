# StringsForWindows
Unix Strings but for windows but with added functionality 

Usage
Basic command
<sub>
strings filename.exe
</sub>

To remove duplicated lines in place of a count
<sub>
strings filename.exe -f
</sub>

Example of using -f: 
PSHUFB 1
PIECES 4
CONV 1
CMPXCHG 2
tlsthrd 1
writable 1
used 1
environ 3
WITH 2
CHAINS 2

To dump into a file and format
<sub>
strings filename.exe > dump.txt -f
</sub>
