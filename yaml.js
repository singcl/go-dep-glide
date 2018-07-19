var fs = require('fs');
var yaml = require('js-yaml');

// Get document, or throw exception on error
try {
  var doc = yaml.safeLoad(fs.readFileSync('glide.yaml', 'utf8'));
  console.log(doc);
} catch (e) {
  console.log(e);
}