class AppDesc:
    def __init__(self,size):
        self.buf_state='OPERATIONAL'
        self.max_buf_size=int(size)

    def freezeBuf(self):
         self.buf_state='FROZEN'

    def toOpState(self):
        self.buf_state='OPERATIONAL' 

    def getState(self):
        return self.buf_state
    
    def isFreeze(self,curr_size):

        code=0

        if(curr_size>=self.max_buf_size):
            code=1

        return code  

    def isOp(self): 
        code=0

        if(self.buf_state=='OPERATIONAL'):
            code=1

        return code    

              
        