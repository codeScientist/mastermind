import matplotlib.pyplot as plt
import numpy as np

v = [
    (1,1,1),
    (1,2,1),
    (1,3,1),
    (1,4,1),
    (1,5,1),
    (1,6,1),
    (2,1,2),
    (2,2,3),
    (2,3,4),
    (2,4,4),
    (2,5,5),
    (2,6,6),
    (3,1,3),
    (3,2,4),
    (3,3,4),
    (3,4,4),
    (3,5,4),
    (3,6,5),
    (4,1,4),
    (4,2,4),
    (4,3,4),
    (4,4,4),
    (4,5,5),
    (4,6,6),
    (5,1,5),
    (5,2,5),
    (5,3,5),
    (5,4,5),
    (5,5,6),
    (6,1,6),
    (6,2,5),
    (6,3,5),
    (6,4,5),
    (6,5,6),
]

def f(x, y):
    for xt,yt,z in v:
        if xt == x and yt == y:
            return z
    return 0


x = np.arange(1, 7, 1)
y = np.arange(1, 7, 1)

# X, Y = np.meshgrid(x, y)
# print(X)
# print(Y)

X = [
    [1, 2, 3, 4, 5, 6],
    [1, 2, 3, 4, 5, 6],
    [1, 2, 3, 4, 5, 6],
    [1, 2, 3, 4, 5, 6],
    [1, 2, 3, 4, 5, 6],
    [1, 2, 3, 4, 5, 6]
]
X = np.array(X)

Y = [
    [1, 1, 1, 1, 1, 1],
    [2, 2, 2, 2, 2, 2],
    [3, 3, 3, 3, 3, 3],
    [4, 4, 4, 4, 4, 4],
    [5, 5, 5, 5, 5, 5],
    [6, 6, 6, 6, 6, 6]
]
Y = np.array(Y)

Z = []
for yi in y:
    Z.append([])
    for xi in x:
        Z[yi-1].append(f(xi, yi))
Z = np.array(Z)

# print(Z)


ax = plt.axes(projection='3d')
ax.plot_surface(X, Y, Z, rstride=1, cstride=1, cmap='viridis', edgecolor='none')
ax.set_title('surface')
ax.set_xlabel('X')
ax.set_ylabel('Y')
ax.set_zlabel('Z')
plt.show()
