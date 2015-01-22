var express = require('express');
var session = require('express-session');
var logger = require('express-logger');
var favicon = require('serve-favicon');
var servestatic = require('serve-static')

var config = require('./config.json');
var routes = require('./lib/routes');

var app = express();

app.use(logger({ path: "./log/app.log" }));
app.use(favicon(__dirname + '/public/img/favicon.png'));
app.use(servestatic(__dirname + '/public'));

app.set('view engine', 'jade');
app.set('views', __dirname + '/src/views');

app.use(session({
    secret: 'keyboard cat',
    resave: false,
    saveUninitialized: true
}));

app.get('/', routes.index);
app.get('/about', routes.about);

app.get('/home', routes.home);

app.get('/logout', routes.authLogout);
app.get('/auth/twitter', routes.authTwitter);
app.get('/auth/twitter/callback', routes.authTwitterCallback);

var server = app.listen(config.port, function () {
    var host = server.address().address;
    var port = server.address().port;
    console.log('moview.io front listening at http://%s:%s', host, port);
});
