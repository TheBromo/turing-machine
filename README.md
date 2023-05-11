# turing-machine


q0,1,q1,X,R
q1,1,q1,1,R
q1,x,q2,x,R
q2,1,q4,Y,R
q4,1,q4,1,R
q4,=,q5,=,R
q5,1,q5,1,R
q5,B,q6,1,L
q6,1,q6,1,L
q6,=,q7,=,L
q7,1,q7,1,L
q7,Y,q8,Y,R
q8,1,q4,Y,R
q8,=,q3,=,L
q2,=,q3,=,L
q3,Y,q3,1,L
q3,x,q9,x,L
q9,1,q9,1,L
q9,X,q0,X,R
q0,x,q10,x,L
q10,X,q10,1,L
q10,B,q11,B,R


1 = 0
x = 00
= = 000
X = 0000
Y = 00000
B = 000000
