const fs = require("fs");

function getData(path) {
  const contents = fs.readFileSync(path, "utf-8");
  return {
    raw: contents,
    array: contents.split("\n"),
  };
}

module.exports = { getData };
