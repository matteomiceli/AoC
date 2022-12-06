const { getData } = require("../helpers/helpers");

const raw = getData("./data.txt").raw;

// value of 4 for part 1
const distinctChars = 14;

function noDuplicates(str) {
  const collection = {};
  const potentialMarker = str.split("");
  potentialMarker.forEach((char) => {
    collection[char] = true;
  });
  return Object.keys(collection).length === distinctChars ? true : false;
}

let markerEnd = 0;
for (let i = 0; i < raw.length; i++) {
  const elements = raw.slice(i, i + distinctChars);
  const noDupes = noDuplicates(elements);
  if (noDupes) {
    markerEnd = i + distinctChars;
    console.log(markerEnd, elements);
    break;
  }
}
