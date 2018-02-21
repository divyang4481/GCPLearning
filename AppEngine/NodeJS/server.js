"use strict"; 

const express = require('express'); 

const app = express(); 

app.get('/', (req, res) => {    
   res.status(200).send('Hello Team! <br\> we are in session 4.');
});

app.listen(process.env.PORT);