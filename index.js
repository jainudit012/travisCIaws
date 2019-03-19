const express = require('express')
const router = express.Router()

router.get('/index', (req,res)=>{
    res.send({error: false, msg: 'Success Ok. INDEX PAGE.'})
})

router.get('/', (req,res)=>{
    res.send({error: false, msg: 'Success Ok. HOME PAGE.'})
})

module.exports = router