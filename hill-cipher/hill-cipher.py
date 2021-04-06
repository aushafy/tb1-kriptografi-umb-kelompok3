import math
import sympy
from sympy import init_printing, pprint
from sympy import Matrix
from sympy.vector import matrix_to_vector, CoordSysCartesian
init_printing()

def decrypt(matrix, words):
    cipher = ''
    M = Matrix(matrix)
    M = M.inv_mod(64)
    length = len(M)
    d = {}
    d2 = {}
    # arr = np.array([d[i] = j; j +=1 for i in "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789_{}"], dtype=int)
    alph = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789_{}"

    for x in range(len(alph)):
        d[alph[x]] = x
        d2[x] = alph[x]
    # print d
    count = 0
    l = []

    for ch in words:
        if (count+1) % (8+1) == 0:
            m = Matrix(l)
            dot_pr_m = M*m
            # print mcd(dot_pr_m)
            # pprint(dot_pr_m.rref()[0][0])
            n = []
            for i in dot_pr_m:
                cipher += d2[i % 64]
            count = 0
            l = []
        l.append(d[ch])
        count += 1
    if (count+1) % (8+1) == 0:
        m = Matrix(l)
        dot_pr_m = M*m
        # print mcd(dot_pr_m)
        # pprint(dot_pr_m.rref()[0][0])
        n = []
        for i in dot_pr_m:
            cipher += d2[i % 64]
    return cipher

if __name__ == '__main__':
    #print mcd([])
    # exit(0)
    secret = [[54, 53, 28, 20, 54, 15, 12, 7],
          [32, 14, 24, 5, 63, 12, 50, 52],
          [63, 59, 40, 18, 55, 33, 17, 3],
          [63, 34, 5, 4, 56, 10, 53, 16],
          [35, 43, 45, 53, 12, 42, 35, 37],
          [20, 59, 42, 10, 46, 56, 12, 61],
          [26, 39, 27, 59, 44, 54, 23, 56],
          [32, 31, 56, 47, 31, 2, 29, 41]]
    ciphertext = "KELOMPOK3_HILLCIPHER"
    # secret = [[7,8], [11,11]]
    # ciphertext = 'APADJTFTWLFJ'.lower()
    print(ciphertext)
    
print(decrypt(secret, ciphertext))