const acorn = require("acorn");
const express = require('express');
var bodyParser = require('body-parser');

const app = express();
const port = 3000;

app.use(bodyParser.text());

app.get('/', (req, res) => {
    res.send('Hello World!');
});
app.post('/', (req, res) => {
    res.send(req.body);
});
app.post('/parse', (req, res) => {
    const result = acorn.parse(req.body);
    res.send(JSON.stringify(result));
});



app.listen(port, () => {
    console.log(`Example app listening at http://localhost:${port}`);
});

