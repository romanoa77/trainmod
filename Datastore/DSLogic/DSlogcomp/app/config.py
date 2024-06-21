import os
basedir=os.path.abspath(os.path.dirname(__file__))

"""

ENV. TO BE DEFINED INSIDE THE CONTAINER IMG

ENDP_STAT
ENDP_DUMP
ENDP_SEND
DB_BURL
"""
class Config:
   env_stat=os.environ['ENDP_STAT']
   env_dump=os.environ['ENDP_DUMP']
   env_send=os.environ['ENDP_SEND']
   env_burl=os.environ['DB_BURL']
   env_max_size=os.environ['MAX_SIZE']

   @staticmethod
   def init_app(app):
      pass
   
class DevConfig(Config):
   DEBUG=False

#class ProdConfig(Config):


config={'dev':DevConfig,'prod':None,'default':DevConfig}




