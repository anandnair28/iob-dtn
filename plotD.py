import matplotlib.pyplot as plt
import numpy as np
import pandas as pd

def plot(df, color, policyName):
	delays = df['deliverytime'] - df['creationtime']
	part = 16
	
	delay_thresholds = np.linspace(0.0, delays.max()+1, part)
	cdf = np.array([0.0]*part)
	total = len(delays)

	ind = 0
	for threshold in delay_thresholds:
		res = np.sum(delays <= threshold)/total
		cdf[ind] = res
		ind += 1
	
	plt.plot(delay_thresholds, cdf, '--.'+color, label=policyName)

konp = pd.read_csv("result_delay_konp.csv")
nop = pd.read_csv("result_delay_np.csv")
gpp = pd.read_csv("result_delay_gpp.csv")
lc = pd.read_csv("result_delay_lc.csv")
rpp = pd.read_csv("result_delay_rpp.csv")

plot(konp, 'b', 'KONP')
plot(nop, 'r', 'NP')
plot(gpp, 'g', 'GPP')
plot(lc, 'y', 'LC')
plot(rpp, 'm', "RPP")

plt.xlabel('delays')
plt.ylabel('cdf')
plt.grid()
plt.legend()

plt.show()
