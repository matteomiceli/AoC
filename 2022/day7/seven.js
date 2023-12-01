const { getData } = require("../helpers/helpers");

const lines = getData("./data.txt").array;

const workingDir = [];
const fileSystem = {};

function tokenizer(line) {
  return line.split(" ");
}

function parser(tokens) {
  switch (tokens[0]) {
    case "$":
      // command
      command(tokens);
      break;

    case "dir":
      break;

    default:
      // file
      resolveFile(tokens);
      break;
  }
}

function resolveFile(tokens) {
  const size = parseInt(tokens[0]);

  workingDir.forEach((_, i) => {
    const path = workingDir.slice(0, i + 1).join(" ");
    fileSystem[path] += size;
  });
}

function command(tokens) {
  const cmd = tokens[1];
  const subject = tokens[2];

  if (cmd === "ls") {
    return;
  }

  if (subject === "..") {
    workingDir.pop();
    return;
  }
  workingDir.push(subject);
  const path = workingDir.join(" ");
  if (!fileSystem[path]) {
    fileSystem[path] = 0;
  }
}

function findSumOfDirectories(lines, maxSize) {
  lines.forEach((line) => {
    parser(tokenizer(line));
  });

  return Object.values(fileSystem).reduce((acc, size) => {
    if (size <= maxSize) {
      return acc + size;
    }
    return acc + 0;
  }, 0);
}

findSumOfDirectories(lines, 100000);

const fsSize = 70000000;
const desiredUnused = 30000000;
function deleteSpace(fileSystem) {
  const toDelete = fileSystem["/"] - (fsSize - desiredUnused);
  const sorted = Object.values(fileSystem)
    .sort((a, b) => a - b)
    .filter((val) => val > toDelete);
  return sorted[0];
}

console.log(deleteSpace(fileSystem));
