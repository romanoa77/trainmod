import requests
import json


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

class RqCalls:

  def __init__(self,url):
    self.baseurl=url
    #self.max_time=1

  def getReq(self,dest):


    
    req=requests.get(self.baseurl+dest)
    

    response={'code':req.status_code,'resp_msg':req.json()}

    return response


  def postJson(self,dest,data):
    #input=json.loads(data)
    input=data
    

    req=requests.post(self.baseurl+dest,input)

    response={'code':req.status_code,'resp_msg':req.text}

    return response





