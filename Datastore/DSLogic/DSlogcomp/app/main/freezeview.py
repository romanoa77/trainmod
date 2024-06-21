from flask import request,Response,make_response
from app.main.main import main_bp
from . import statview
from .. import AppStatus
from .. import DBConn


@main_bp.route('/freezeds',methods=['POST'])
def freezeApp():
   
   buffer=statview.getBufStat()

   read_size=buffer['buff_size']

   if(AppStatus.isFreeze(read_size)):
      appresp=make_response(buffer,'201')   

      AppStatus.freezeBuf()
   else:
       appresp=Response("Waiting for more incoming data.",status=200,mimetype='text/plain')   
      

   return appresp