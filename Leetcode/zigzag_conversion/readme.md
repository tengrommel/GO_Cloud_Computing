# ZigZag Conversion

The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

    P   A   H   N
    A P L S I I G
    Y   I   R

And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

    string convert(string text, int nRows);

convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR".

# 解题思路

输入"ABCDEFGHIJKLMNOPQRSTUVWXYZ"和参数5后，得到答案"AGMSYBFHLNRTXZCEIKOQUWDJPV"，按照题目的摆放方法可得：

    A   I   Q   Y
    B  HJ  PR  XZ
    C G K O S W
    DF  LN  TV
    E   M   U
    
可以看到，各行字符在原字符中的索引号为

   1. 0行，0, 8, 16, 24
   2. 1行，1, 7,9, 15,17, 23,25
   3. 2行，2, 6, 10, 14, 18, 22
   4. 3行，3,5, 11,13, 19,21
   5. 4行，4, 12, 20

令p=numRows×2-2，可以总结出以下规律
1. 0行,0×p,1×p, ...
2. r行,r,1×p+r,2×p-r,2×p+r, ...
3. 最后一行， numRow-1, numRow-1+1×p，numRow-1+2×p，...

只需编程依次处理各行即可。
