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
#rpp1 = pd.read_csv("result_delay_rpp1.csv")
#rpp2 = pd.read_csv("result_delay_rpp2.csv")
#rpp3 = pd.read_csv("result_delay_rpp3.csv")
#rpp4 = pd.read_csv("result_delay_rpp4.csv")

plot(konp, 'b', 'KONP')
plot(nop, 'r', 'NP')
plot(gpp, 'g', 'GPP')
plot(lc, 'y', 'LC')
plot(rpp, 'm', 'RPP')
#plot(rpp1, 'y', "RPP1")
#plot(rpp2, 'k', "RPP2")
#plot(rpp3, 'c', "RPP3")
#plot(rpp4, 'm', "RPP4")

#n2 = pd.read_csv("result_delay_2.csv")
#n8 = pd.read_csv("result_delay_8.csv")
#nUL = pd.read_csv("result_delay_UL.csv")

#plot(n2, 'b', 'NC = 2')
#plot(n8, 'g', 'NC = 8')
#plot(nUL, 'r', 'NC = UL')

plt.xlabel('delays')
plt.ylabel('cdf')
plt.grid()
plt.legend()

plt.show()
