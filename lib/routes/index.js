var oauth = require('oauth');
var util = require('util');

/**
 *
 * @type {string}
 * @private
 */
var _twitterConsumerKey = "P1LBfEJNOHDaTSOQrmAb5Gvil";

/**
 *
 * @type {string}
 * @private
 */
var _twitterConsumerSecret = "8B6wkuXDloxrmx1uUZlQEmaSnUzVU3n48XYEcHfzFe33sNu8KK";

/**
 *
 * @type {exports.OAuth}
 */
var consumer = new oauth.OAuth(
    "https://twitter.com/oauth/request_token",
    "https://twitter.com/oauth/access_token",
    _twitterConsumerKey,
    _twitterConsumerSecret,
    "1.0A",
    "http://localhost:8081/auth/twitter/callback",
    "HMAC-SHA1");


/**
 *
 * @param req
 * @param res
 */
exports.index = function (req, res) {
    var context = {};
    res.render('index', context);
};

/**
 *
 * @param req
 * @param res
 */
exports.home = function (req, res) {
    var context = {};
    res.render('home', context);
};

/**
 *
 * @param req
 * @param res
 */
exports.authLogout = function (req, res) {
    res.session = null;
    res.redirect('/');
};

/**
 *
 * @param req
 * @param res
 */
exports.authTwitter = function (req, res) {
    consumer.getOAuthRequestToken(function (error, oauthToken, oauthTokenSecret, results) {
        if (error) {
            res.send("Error getting OAuth request token : " + util.inspect(error), 500);
        } else {
            req.session.oauthRequestToken = oauthToken;
            req.session.oauthRequestTokenSecret = oauthTokenSecret;
            res.redirect("https://twitter.com/oauth/authorize?oauth_token=" + req.session.oauthRequestToken);
        }
    });
};

/**
 *
 * @param req
 * @param res
 */
exports.authTwitterCallback = function (req, res) {
    // util.puts(">> Request Token : " + req.session.oauthRequestToken);
    // util.puts(">> Request Token Secret : " + req.session.oauthRequestTokenSecret);
    // util.puts(">> OAuth Verifier : " + req.query.oauth_verifier);
    consumer.getOAuthAccessToken(
        req.session.oauthRequestToken,
        req.session.oauthRequestTokenSecret,
        req.query.oauth_verifier, function (error, oauthAccessToken, oauthAccessTokenSecret, results) {
            if (error) {
                var msg = "Error getting OAuth access token : ";
                msg += util.inspect(error);
                msg += "[ oauthAccessToken : " + oauthAccessToken + "]";
                msg += "[ oauthAccessTokenSecret : " + oauthAccessTokenSecret + "]";
                msg += "[ results : " + util.inspect(results) + "]";
                res.send(msg, 500);
            } else {
                req.session.oauthAccessToken = oauthAccessToken;
                req.session.oauthAccessTokenSecret = oauthAccessTokenSecret;
                res.redirect('/home');
            }
        }
    );

    /*
    consumer.get(
        "http://twitter.com/account/verify_credentials.json",
        req.session.oauthAccessToken,
        req.session.oauthAccessTokenSecret, function (error, data, response) {
            if (error) {
                res.redirect('/sessions/connect');
                // res.send("Error getting twitter screen name : " + util.inspect(error), 500);
            } else {
                var parsedData = JSON.parse(data);
                // req.session.twitterScreenName = response.screen_name;
                res.send('You are signed in: ' + parsedData.screen_name);
            }
        }
    );
    */
};