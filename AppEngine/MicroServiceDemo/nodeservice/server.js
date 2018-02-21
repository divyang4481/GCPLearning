"use strict";

const express = require('express'); 

const app = express(); 

app.get('/', (req, res) => {    
   res.status(200).send('This session 4 demo for nodejs appengine it is changed in demo');
});

app.listen(process.env.PORT);