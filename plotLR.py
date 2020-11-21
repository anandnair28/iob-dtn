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

plot(konp, 'b', "KONP")
plot(nop, 'r', "NP")
plot(gpp, 'g', "GPP")
plot(lc, 'y', "LC")
plot(rpp, 'm', "RPP")

plt.xlabel('loss rate')
plt.ylabel('cdf')
plt.grid()
plt.legend()

plt.show()
