#!/bin/bash

###va lanciato con source

<<comment


GIN CALLS

router.GET("/stat", getStat(&StatDesc))
router.GET("/dumpLogF", getLogF())

router.POST("/sendF", postFile(&StatDesc))   

comment


export DS_STAT=bufstat
export DS_SEND=stream
export DS_FREEZE=freezeds
export DS_BURL=http://localhost:5000/
export MAX_SIZE=3000



