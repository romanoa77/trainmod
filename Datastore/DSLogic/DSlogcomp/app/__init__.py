from flask import Flask, render_template


from . import DSDesc
from . import DBhooks

from .config import config

AppStatus=DSDesc.AppDesc()
DBConn=DBhooks.DBhook()



def create_app(config_name):

    app=Flask(__name__)
    app.config.from_object(config[config_name])
    #config[config_name].init_app(app)

    
  

    AppStatus.initDesc(config[config_name].env_max_size)
    

    DBConn.connect(config[config_name].env_burl,config[config_name].env_stat,
                 config[config_name].env_dump,config[config_name].env_send,
                 config[config_name].env_desc,config[config_name].env_upd_desc,
                 config[config_name].env_flush)
    
    from app.main.main import main_bp as mainbp
    main_blueprint=mainbp
    app.register_blueprint(main_blueprint)
    

    return app
















