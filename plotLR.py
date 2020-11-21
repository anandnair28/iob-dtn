import matplotlib.pyplot as plt
import numpy as np
import pandas as pd


def plot(df, color, name):
    loss_rate = (df['packetCount'] - df['success']) / df['packetCount']

    part = 16

    ls = np.linspace(0.0, 0.85, part)
    cdf = np.array([0.0] * part)
    total = len(loss_rate)

    ind = 0
    for rate in ls:
        t = np.sum(loss_rate <= rate) / total
        cdf[ind] = t
        ind += 1

    plt.plot(ls, cdf, '--.'+color, label=name)

konp = pd.read_csv("result_lr_konp.csv")
nop = pd.read_csv("result_lr_np.csv")
gpp = pd.read_csv("result_lr_gpp.csv")
lc = pd.read_csv("result_lr_lc.csv")
rpp = pd.read_csv("result_lr_rpp.csv")
#rpp1 = pd.read_csv("result_lr_rpp1.csv")
#rpp2 = pd.read_csv("result_lr_rpp2.csv")
#rpp3 = pd.read_csv("result_lr_rpp3.csv")
#rpp4 = pd.read_csv("result_lr_rpp4.csv")

plot(konp, 'b', "KONP")
plot(nop, 'r', "NP")
plot(gpp, 'g', "GPP")
plot(lc, 'y', "LC")
plot(rpp, 'm', 'RPP')
#plot(rpp1, 'y', "RPP1")
#plot(rpp2, 'k', "RPP2")
#plot(rpp3, 'c', "RPP3")
#plot(rpp4, 'm', "RPP4")

#n2 = pd.read_csv("result_lr_2.csv")
#n8 = pd.read_csv("result_lr_8.csv")
#nUL = pd.read_csv("result_lr_UL.csv")

#plot(n2, 'b', 'NC = 2')
#plot(n8, 'g', 'NC = 8')
#plot(nUL, 'r', 'NC = UL')

plt.xlabel('loss rate')
plt.ylabel('cdf')
plt.grid()
plt.legend()

plt.show()
