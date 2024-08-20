import numpy as np
import matplotlib.pyplot as plt
from constantQ.timeseries import TimeSeries
from constantQ.spectrogram import Spectrogram
from scipy import signal
from scipy import signal as scisignal
from astropy.units.quantity import Quantity as q
from astropy.units import Unit
import json
import os

env_path=os.environ['DATART']

#env_srv=os.environ['SRV']



if os.path.exists(env_path):
    print("LOG OK DATA")

    flist=os.scandir(env_path)
    
    listbytm=[]

    flag=0

    for fname in flist:
       
       
       try:
        f=open(env_path+"/"+fname.name,"r")

        gwdata=json.load(f)
        f.close()

        t0=gwdata['t'][0]

        listbytm.append((t0,fname.name))
       
       except:
        print("LOG FS ERROR")
        flag=1
        break
    
    if(flag==0):
          
     listbytm.sort(key=lambda elm:elm[0])

     gwarr=[]

     for item in listbytm:
      

      try:
        f=open(env_path+"/"+item[1],"r")

        gwdata=json.load(f)
        f.close()

        hdata=gwdata['h']
        

        for num in hdata:
           gwarr.append(num)
           
       
      except:
        print("LOG FS ERROR1")
        flag=1
        break
      
     if(flag==0):
        dtd=gwdata['t'][1]
        



        gwdtms = TimeSeries(gwarr,Unit('ad'),0.0,q(dtd,'s'))

        gwdtms.whiten()

        spectre=gwdtms.q_transform(search=None)
        spectre.writeSpec(target="output.hdf5",path='/gwdata',attrs='dict',overwrite=True)
        plt.imshow(spectre.T, origin='lower')				# plot the spectrogram
        plt.colorbar()									# colorbar
        plt.savefig('gwdata.png')


      
      
      
      

     



    
    
    

       






else:
    print('LOGERROR')




