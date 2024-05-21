#!/bin/bash

###va lanciato con source

<<comment


GIN CALLS

router.GET("/stat", getStat(&StatDesc))
router.GET("/dumpLogF", getLogF())

router.POST("/sendF", postFile(&StatDesc))   

comment


export ENDP_STAT=stat
export ENDP_DUMP=dumpLogF
export ENDP_SEND=sendF
export DB_BURL=http://localhost:8081/
export MAX_SIZE=3000



