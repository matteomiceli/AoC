const { getData } = require("../helpers/helpers");
const movements = getData("./data.txt").array.map((move) => move.split(" "));

const tVisited = new Set();
const rope = {
  0: { x: 0, y: 0 },
  1: { x: 0, y: 0 },
  2: { x: 0, y: 0 },
  3: { x: 0, y: 0 },
  4: { x: 0, y: 0 },
  5: { x: 0, y: 0 },
  6: { x: 0, y: 0 },
  7: { x: 0, y: 0 },
  8: { x: 0, y: 0 },
  9: { x: 0, y: 0 },
};

// part 1
// const tail = { x: 0, y: 0 };
// const head = { x: 0, y: 0 };

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
    movement(vTail, "R");
  } else if (adjustX < 0) {
    movement(vTail, "L");
  }

  // vertical adjust
  if (adjustY > 0) {
    movement(vTail, "U");
  } else if (adjustY < 0) {
    movement(vTail, "D");
  }
}

movements.forEach((move, i) => {
  const dir = move[0];
  const length = move[1];

  for (let i = 0; i < length; i++) {
    movement(rope[0], dir);
    for (let j = 0; j < Object.values(rope).length; j++) {
      if (j === Object.values(rope).length - 1) break;

      const segment = rope[j];
      const trailing = rope[j + 1];

      if (areApart(segment, trailing)) {
        tailAdjustment(segment, trailing);
      }
      tVisited.add(JSON.stringify(rope[9]));
    }
  }
});

console.log(tVisited.size);
