from flask import render_template
from app.main.main import main_bp
from .. import AppStatus
from .. import DBConn





@main_bp.route('/')
def index():
    
    output=DBConn.getBufStat()

     #print(output['resp_msg'],file=sys.stderr)

    if (output['code']==200): 
       

       appresp=render_template(
          'index.html',size=output['resp_msg']['buff_size']
          ,state=AppStatus.getState(),
          nitm=output['resp_msg']['n_itm'])
    
    
    return appresp