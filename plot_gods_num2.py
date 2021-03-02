# library
from mpl_toolkits.mplot3d import Axes3D
import matplotlib.pyplot as plt
import pandas as pd
# import seaborn as sns

single_use_data = [
    (1,1,1),
    (2,1,2),
    (2,2,2),
    (3,1,3),
    (3,2,3),
    (3,3,4),
    (4,1,4),
    (4,2,3),
    (4,3,4),
    (4,4,5),
    (5,1,5),
    (5,2,4),
    (5,3,4),
    (5,4,5),
    (5,5,6),
    (6,1,6),
    (6,2,4),
    (6,3,5),
    (6,4,5),
    (6,5,6),
    (6,5,7),
]

multi_use_data = [
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

datadict = {"X": [], "Y": [], "Z": []}
for x,y,z in single_use_data:
    datadict["X"].append(x)
    datadict["Y"].append(y)
    datadict["Z"].append(z)

print(datadict)

df = pd.DataFrame(datadict, columns=["X", "Y", "Z"])

print(df.head())

# Get the data (csv file is hosted on the web)
# url = 'https://python-graph-gallery.com/wp-content/uploads/volcano.csv'
# data = pd.read_csv(url)
 
# Transform it to a long format
# df=data.unstack().reset_index()
# df.columns=["X","Y","Z"]
 
# And transform the old column name in something numeric
# df['X']=pd.Categorical(df['X'])
# df['X']=df['X'].cat.codes
 
# Make the plot
fig = plt.figure()
ax = fig.gca(projection='3d')
ax.plot_trisurf(df['X'],df['Y'], df['Z'], cmap=plt.cm.viridis, linewidth=0.2)
ax.set_title('Mastermind Gods Number trend')
ax.set_xlabel('Total pins')
ax.set_ylabel('Total Selectable pins')
ax.set_zlabel('Gods Number')
plt.show()
 
# to Add a color bar which maps values to colors.
# surf=ax.plot_trisurf(df['Y'], df['X'], df['Z'], cmap=plt.cm.viridis, linewidth=0.2)
# fig.colorbar( surf, shrink=0.5, aspect=5)
# plt.show()
 
# Rotate it
# ax.view_init(30, 45)
# plt.show()
 
# Other palette
# ax.plot_trisurf(df['Y'], df['X'], df['Z'], cmap=plt.cm.jet, linewidth=0.01)
# plt.show()
