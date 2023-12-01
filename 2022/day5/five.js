const { getData } = require("../helpers/helpers");

const raw = getData("./data.txt").raw;
const [_, rawMovements] = raw.split("\n\n");

const stacks = {
  1: ["Q", "W", "P", "S", "Z", "R", "H", "D"],
  2: ["V", "B", "R", "W", "Q", "H", "F"],
  3: ["C", "V", "S", "H"],
  4: ["H", "F", "G"],
  5: ["P", "G", "J", "B", "Z"],
  6: ["Q", "T", "J", "H", "W", "F", "L"],
  7: ["Z", "T", "W", "D", "L", "V", "J", "N"],
  8: ["D", "T", "Z", "C", "J", "G", "H", "F"],
  9: ["W", "P", "V", "M", "B", "H"],
};

const parsedMovements = rawMovements.split("\n").map((move) => {
  return move.split(/[(a-z)]+\s+/).filter((el) => el);
});

function executeMovement(movement) {
  const numMoved = movement[0];

  const slicePosition = stacks[parseInt(movement[1])].length - numMoved;
  const crane = stacks[parseInt(movement[1])].splice(slicePosition, numMoved);
  // disabled for part 2
  // crane.reverse();
  stacks[parseInt(movement[2])] = stacks[parseInt(movement[2])].concat(crane);
}

parsedMovements.forEach((movement) => {
  executeMovement(movement);
});

console.log(stacks);
