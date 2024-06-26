from . DBCalls import RqCalls


"""
GIN CALLS

router.GET("/stat", getStat(&StatDesc))
router.GET("/dumpLogF", getLogF())

router.POST("/sendF", postFile(&StatDesc))   

"""

"""

Response struct

response={'code':req.status_code,'resp_msg':req.json()}
response={'code':req.status_code,'resp_msg':req.text}


"""

class DBhook:
    """
    def __init__(self,burl,stat,dump,send):
     self.Caller=RqCalls(burl)

     self.serviceurl=burl
     self.status=stat  
     self.log_dump=dump
     self.send_data=send
    """
    def getBufStat(self):
        
        buf_desc=self.Caller.getReq(self.status)

        return buf_desc
    
    def getLogDump(self):
       
       log_cont=self.Caller.getReq(self.log_dump)

       return log_cont
    
    def getDSDsc(self):
       
       buf_desc=buf_desc=self.Caller.getReq(self.desc_stat)

       return buf_desc

    def postDataStream(self,data_struct):


        #postJson expects a dict rapresenting a json object
        post_resp=self.Caller.postJson(self.send_data,data_struct)

        return post_resp
    
    def postDsc(self,data_struct):


        #postJson expects a dict rapresenting a json object
        post_resp=self.Caller.postJson(self.updd_stat,data_struct)

        return post_resp
    
    def postFlush(self,data_struct):


        #postJson expects a dict rapresenting a json object
        post_resp=self.Caller.postJson(self.flush,data_struct)

        return post_resp
    

    

    def connect(self,burl,stat,dump,send,desc,upddesc,clean):
     self.Caller=RqCalls(burl)

     self.serviceurl=burl
     self.status=stat  
     self.log_dump=dump
     self.send_data=send
     self.desc_stat=desc
     self.updd_stat=upddesc
     self.flush=clean
    

    

"""
DBConn=DBhook(
   Appfconf.env_burl,Appfconf.env_stat,Appfconf.env_dump,Appfconf.env_send) 

gwbuf={'h':[-1.114753991083138831e-18,-1.114753991083138831e-18],
       't':[1.126259452401367188e+09,1.126259452401367188e+09]}

output=DBConn.postDataStream(gwbuf)

print(output)


output=DBConn.getBufStat()

print (output)
"""




     







     