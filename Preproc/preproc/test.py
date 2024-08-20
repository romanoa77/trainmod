import numpy as np
import matplotlib.pyplot as plt
from constantQ.timeseries import TimeSeries
from constantQ.spectrogram import Spectrogram
from scipy import signal
from scipy import signal as scisignal
from astropy.units.quantity import Quantity as q
from astropy.units import Unit
import json


# Generate np.array chirp signal
dt = 0.001
t = np.arange(0,10,dt)
f0 = 50
f1 = 250
t1 = 2
x = np.cos(2*np.pi*t*(f0 + (f1 - f0)*np.power(t, 2)/(3*t1**2)))
fs = 1/dt

plt.plot(x)				# plot the chirp signal
plt.savefig('a.png')				# display

print(len(x))



series = TimeSeries(x, dt = 0.001, unit='m', name='test', t0=0)     #np.array --> constantQ.timeseries





hdata = series
sq = hdata.q_transform(search=None)

sq.writeSpec(target="output.hdf5",path='/foo',attrs='dict',overwrite=True)







plt.imshow(sq.T, origin='lower')				# plot the spectrogram
plt.colorbar()									# colorbar
plt.savefig('b.png')

freq, ts, Sxx = scisignal.spectrogram(x)		# scipy spectrogram

plt.pcolor(ts, freq, Sxx, shading='auto')		# plot the spectrogram
plt.colorbar()									# colorbar
plt.savefig('c.png')


"""
rds=Spectrogram.readSpec(source="output.hdf5",path='/foo')



plt.imshow(rds.T, origin='lower')				# plot the spectrogram
plt.colorbar()									# colorbar
plt.savefig('output.png')
"""

f=open("target.json","r")

gwdata=json.load(f)

f.close()




gwarr=gwdata['h']

for i in x:
    gwarr.append(-i)


dtd=gwdata['t'][1]



gwdtms = TimeSeries(gwarr,Unit('ad'),0.0,q(dtd,'s'))

print(gwdtms)



spectre=gwdtms.q_transform(search=None)

plt.imshow(spectre.T, origin='lower')				# plot the spectrogram
plt.colorbar()									# colorbar
plt.savefig('gwdata.png')












