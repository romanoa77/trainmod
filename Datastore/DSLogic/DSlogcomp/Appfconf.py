import os

"""

ENV. TO BE DEFINED INSIDE THE CONTAINER IMG

ENDP_STAT
ENDP_DUMP
ENDP_SEND
DB_BURL
"""

env_stat=os.environ['ENDP_STAT']
env_dump=os.environ['ENDP_DUMP']
env_send=os.environ['ENDP_SEND']
env_burl=os.environ['DB_BURL']
env_max_size=os.environ['MAX_SIZE']




