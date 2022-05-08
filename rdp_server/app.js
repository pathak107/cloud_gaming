const GuacamoleLite = require('guacamole-lite');
const express = require('express');
const http = require('http');

const app = express();

const server = http.createServer(app);

const websocketOptions = {
    port: 8080 // we will accept connections to this port
};


const guacdOptions = {
    port: 4822 // port of guacd
};

const clientOptions = {
    crypt: {
        cypher: 'AES-256-CBC',
        key: 'MySuperSecretKeyForParamsToken12'
    }
};

const guacServer = new GuacamoleLite({server}, guacdOptions, clientOptions);

const val = {
    "connection": {
        "type": "rdp",
        "settings": {
            "hostname": "43.204.116.247",
            "username": "Administrator",
            "password": "$jKu-aExwAKlqQEenz?;;4?aSkUv4uZ?",
            "enable-drive": true,
            "create-drive-path": true,
            "security": "any",
            "ignore-cert": true,
            "enable-wallpaper": false
        }
    }
}
const crypto = require('crypto');

const encrypt = (value) => {
    const iv = crypto.randomBytes(16);
    const cipher = crypto.createCipheriv(clientOptions.crypt.cypher, clientOptions.crypt.key, iv);

    let crypted = cipher.update(JSON.stringify(value), 'utf8', 'base64');
    crypted += cipher.final('base64');

    const data = {
        iv: iv.toString('base64'),
        value: crypted
    };

    return new Buffer(JSON.stringify(data)).toString('base64');
};
// console.log(encrypt(val))

server.listen(8080);
