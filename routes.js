const express = require('express')
const cors = require('cors')
const routes = require('./index')

module.exports = function(app){
    app.use(express.json())
    app.use(cors())
    app.use('/api', routes)
}