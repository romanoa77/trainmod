import httpx
import Appfconf
from fastapi import FastAPI, Request
from contextlib import asynccontextmanager
from pydantic import BaseModel
from typing import List
import sys
import json


class FrzMsg(BaseModel):
    code: str


class GWData(BaseModel):
      h: List[float] 
      t: List[float]    

@asynccontextmanager
async def lifespan(app: FastAPI):
    app.requests_client = httpx.AsyncClient()
    yield
    await app.requests_client.aclose()

app = FastAPI(lifespan=lifespan)

@app.get("/stats")
async def get_ds_stats(request: Request):
    requests_client = request.app.requests_client
    response = await requests_client.get(Appfconf.env_burl+Appfconf.env_stat)
    return response.json()

@app.post("/train")
async def trainsign(request: Request,msg:FrzMsg):
    requests_client = request.app.requests_client
    
    response = await requests_client.post(Appfconf.env_burl+Appfconf.env_freeze,
                                          data={"code":msg.code})
    
    
    
    
    if(response.status_code==httpx.codes.CREATED):
     buf=response.json()
     n_file=buf['n_itm']
     n_byte=buf['buff_size']
     
     
     output={"resp":'FROZEN',"n_f":n_file,"bt_wrt":n_byte}

     
     #INSERT HERE CODE FOR PREPROCESSING STARTUP
     #AIRFLOW SHOULD FOLLOW THE PIPELINE
     
    else:
     
     output={"resp":'WAITING'}

     
    
    return output


@app.post("/streamdata")
async def streamdat(request: Request,stream:GWData):
    requests_client = request.app.requests_client
    buf={'h':stream.h,'t':stream.t}
    buf=json.dumps(buf)
    

    

    response = await requests_client.post(Appfconf.env_burl+Appfconf.env_send,data=buf)
    
    return buf