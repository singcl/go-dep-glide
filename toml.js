var toml = require('toml');
var fs = require('fs');

try {
    var data = toml.parse(fs.readFileSync('Gopkg.toml', 'utf8'));
    console.dir(data);
} catch (e) {
    console.log(e);
}
