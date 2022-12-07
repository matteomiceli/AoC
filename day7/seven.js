const { getData } = require("../helpers/helpers");

const lines = getData("./data_test.txt").array;

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
    const dir = workingDir[i];
    fileSystem[dir] += size;
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
  if (!fileSystem[subject]) {
    fileSystem[subject] = 0;
  }
}

function findSumOfDirectories(lines, maxSize) {
  lines.forEach((line) => {
    parser(tokenizer(line));
  });

  console.log(fileSystem);
  return Object.values(fileSystem).reduce((acc, size) => {
    if (size <= maxSize) {
      return acc + size;
    }
    return acc + 0;
  }, 0);
}

console.log(findSumOfDirectories(lines, 100000));
