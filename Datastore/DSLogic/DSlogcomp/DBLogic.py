from flask import Flask, render_template,Response,request,make_response
from flask import jsonify

import sys
from DSDesc import AppDesc
from DBhooks import DBhook
import Appfconf


#import json
app=Flask(__name__)

AppStatus=AppDesc(Appfconf.env_max_size)
DBConn=DBhook(
    Appfconf.env_burl,Appfconf.env_stat,Appfconf.env_dump,Appfconf.env_send
)





@app.route('/')
def index():
    
    output=DBConn.getBufStat()

     #print(output['resp_msg'],file=sys.stderr)

    if (output['code']==200): 
       

       appresp=render_template(
          'index.html',size=output['resp_msg']['buff_size']
          ,state=AppStatus.getState(),
          nitm=output['resp_msg']['n_itm'])
    else:
       appresp=Response("Internal error!",status=500,mimetype='text/plain')
    
    return appresp


@app.route('/stream',methods=['POST'])
def pushStream():
    
     if(AppStatus.isOp()):

      data=request.data

      #print('Incoming data',file=sys.stderr)
      #print(request.data,file=sys.stderr)

    

      output=DBConn.postDataStream(data)

      if (output['code']==201): 
        appresp=Response(output['resp_msg'],status=201,mimetype='text/plain') 
      else:
       appresp=Response("Internal error!",status=500,mimetype='text/plain')   

     else: 
      appresp=Response("Datastore frozen.",status=503,mimetype='text/plain')   

     return appresp

@app.route('/freezeds',methods=['POST'])
def freezeApp():
   
   buffer=getBufStat()

   read_size=buffer['buff_size']

   if(AppStatus.isFreeze(read_size)):
      appresp=make_response(buffer,'201')   

      AppStatus.freezeBuf()
   else:
       appresp=Response("Waiting for more incoming data.",status=200,mimetype='text/plain')   
      

   return appresp


#GET ROUTES

@app.route('/bufstat',methods=['GET'])
def getBufStat():
     
     
     output=DBConn.getBufStat()

     #print(output['resp_msg'],file=sys.stderr)

     

     if (output['code']==200): 
       appresp=output['resp_msg']
     else:
       appresp=Response("Internal error!",status=500,mimetype='text/plain')   
     
     

     return appresp


@app.route('/dumpflog',methods=['GET'])
def getFlog():
     
     
     output=DBConn.getLogDump()

     #print(output['resp_msg'],file=sys.stderr)

     

     if (output['code']==200): 
       appresp=output['resp_msg']
     else:
       appresp=Response("Internal error!",status=500,mimetype='text/plain')   
     
     

     return appresp



@app.errorhandler(500)
def handle_bad_request(e):
    return 'Internal error!', 500

     

if __name__=='__main__':
    app.run(debug=False)