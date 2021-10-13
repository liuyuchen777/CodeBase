import random

def fun(a, b, c, d, e, f, g, h):
    output = 0.25 * a + 0.15 * b + 0.1 * c + 0.12 * e + 0.08 * f + 0.17 *g + 0.13 * h
    return output

files = open('data2.txt', 'w')

for i in range(1, 50 + 1):
    a = random.random()
    b = random.random()
    c = random.random()
    d = random.random()
    e = random.random()
    f = random.random()
    g = random.random()
    h = random.random()
    if fun(a, b, c, d, e, f, g, h) > 0.5:
        label = 1
    else:
        label = 0
    files.write(str(i) + '\t' + str(a) + '\t' + str(b) + '\t' + str(c) + '\t' + \
            str(d) + '\t' + str(e) + '\t' + str(f) + '\t' + str(g) + \
            '\t' + str(h) +\
            '\t' + str(label) + '\n')
    i += 1

files.close()