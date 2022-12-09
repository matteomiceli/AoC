const { getData } = require("../helpers/helpers");
const movements = getData("./data.txt").array.map((move) => move.split(" "));

const tVisited = new Set();

// part 1
const tail = { x: 0, y: 0 };
const head = { x: 0, y: 0 };

function movement(vector, dir) {
  switch (dir) {
    case "U":
      vector.y += 1;
      break;

    case "R":
      vector.x += 1;
      break;

    case "D":
      vector.y -= 1;
      break;

    case "L":
      vector.x -= 1;
      break;
  }
}

function areApart(vHead, vTail) {
  let deltaX = vHead.x - vTail.x;
  let deltaY = vHead.y - vTail.y;

  if (deltaX < 0) deltaX = deltaX * -1;
  if (deltaY < 0) deltaY = deltaY * -1;

  if (deltaX > 1 || deltaY > 1) {
    return true;
  }
  return false;
}

function tailAdjustment(vHead, vTail) {
  const adjustX = vHead.x - vTail.x;
  const adjustY = vHead.y - vTail.y;

  // horizontal adjust
  if (adjustX > 0) {
    movement(tail, "R");
  } else if (adjustX < 0) {
    movement(tail, "L");
  }

  // vertical adjust
  if (adjustY > 0) {
    movement(tail, "U");
  } else if (adjustY < 0) {
    movement(tail, "D");
  }
}

movements.forEach((move, i) => {
  const dir = move[0];
  const length = move[1];

  for (let i = 0; i < length; i++) {
    movement(head, dir);
    if (areApart(head, tail)) {
      tailAdjustment(head, tail);
    }
    tVisited.add(JSON.stringify(tail));
  }
});

console.log(tVisited.size);
